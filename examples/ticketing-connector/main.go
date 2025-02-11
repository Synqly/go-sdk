package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	koanfyaml "github.com/knadh/koanf/parsers/yaml"

	"github.com/google/uuid"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var (
	k              = koanf.New(".")
	synqlyOrgToken = os.Getenv("SYNQLY_ORG_TOKEN")
	jiraURL        = os.Getenv("SYNQLY_JIRA_URL")
	jiraUser       = os.Getenv("SYNQLY_JIRA_USER")
	jiraToken      = os.Getenv("SYNQLY_JIRA_TOKEN")
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

type App struct {
	Tenant *Tenant
}

// NewApp instantiates a new App
func NewApp() *App {
	return &App{}
}

// A Tenant object represents one of your customers, as well as the state you
// would maintain for them in your application.
type Tenant struct {
	// ID: A unique identifier for one of your customers.
	ID string
	// SynqlyAccountId: The ID of the Synqly Account in which Integrations for
	// this tenant will be created. This example creates a new Account for every
	// Tenant, but it would also be possible to use the same Account for all Tenants.
	SynqlyAccountId string
	// SynqlyClient: A cached management client, used to manage Integrations.
	SynqlyClient *mgmtClient.Client
	// TicketClient: A cached engine client, used to log events to a third-party
	// logging Provider by way of an Integration.
	TicketClient *engineClient.Client
}

// NewTenant initializes a Synqly Account for a given Tenant. This example
// creates a new Account for every Tenant to keep their credentials isolated.
//
// Returns an error if a tenant with the same ID already exists or if a Synqly
// Account cannot be created for the Tenant.
func (a *App) NewTenant(ctx context.Context, id string) error {
	// Do not allow duplicate tenant names

	if a.Tenant != nil && a.Tenant.ID == id {
		return fmt.Errorf("duplicate tenant name %v", id)
	}

	// Create a Synqly Client that can be used to interact with the tenant
	// We will use this client to create an Account and set up an Integration with an event logging provider
	client := mgmtClient.NewClient(
		mgmtClient.WithToken(synqlyOrgToken),
	)

	// Create a Synqly Account for this tenant
	account, err := client.Accounts.Create(ctx, &mgmt.CreateAccountRequest{
		Fullname: &id,
	})
	if err != nil {
		return fmt.Errorf("unable to create account: %w", err)
	}

	// Store the Tenant and associated Synqly objects in an in-memory cache.
	a.Tenant = &Tenant{
		ID:              id,
		SynqlyAccountId: account.Result.Account.Id,
		SynqlyClient:    client,
		TicketClient:    nil,
	}
	return nil
}

func (a *App) inmemConfig() *mgmt.CreateIntegrationRequest {
	return &mgmt.CreateIntegrationRequest{
		Fullname:       engine.String("In-Memory Connector"),
		ProviderConfig: &mgmt.ProviderConfig{SiemMockSiem: &mgmt.SiemMock{}},
	}
}

func (app *App) cleanup() {
	// Clean up the Accounts created in Synqly. In your application, you would
	// persist the Synqly Account id and Integration tokens to handle process
	// restarts. We are not doing that in this example, so we need to clean up
	// the accounts we created in Synqly.
	consoleLogger.Println("Cleaning up Synqly Account")
	ctx := context.Background()

	// Deleting the account will delete all credentials and integrations associated with the account.
	if err := app.Tenant.SynqlyClient.Accounts.Delete(ctx, app.Tenant.SynqlyAccountId); err != nil {
		consoleLogger.Printf("Error deleting account %s: %s\n", app.Tenant.SynqlyAccountId, err)
	}

	os.Exit(0)
}

func (a *App) configureTicketing(ctx context.Context, ticketProviderType string) error {
	var integrationReq *mgmt.CreateIntegrationRequest
	switch ticketProviderType {
	case "jira":
		integrationReq = &mgmt.CreateIntegrationRequest{
			Fullname: engine.String("Ticketing Connector"),
			ProviderConfig: &mgmt.ProviderConfig{
				TicketingJira: &mgmt.TicketingJira{
					Url: jiraURL,
					Credential: &mgmt.JiraCredential{
						Basic: &mgmt.BasicCredential{
							Username: jiraUser,
							Secret:   jiraToken,
						},
					},
				},
			},
		}

	case "inmem":
		integrationReq = a.inmemConfig()

	default:
		return fmt.Errorf("invalid siem provider type: %s", ticketProviderType)
	}

	integration, err := a.Tenant.SynqlyClient.Integrations.Create(ctx, a.Tenant.SynqlyAccountId, integrationReq)
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	// Create an Event Logger for the tenant
	// We will use this client to post events to the provider
	a.Tenant.TicketClient = engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)

	return nil
}

type Ticket struct {
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

// createAttachment creates an attachment in the ticket with the given ID
func createAttachment(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.CreateAttachmentResponse, error) {
	file, err := os.ReadFile("attachment.png")
	if err != nil {
		return nil, err
	}

	attachmentResponse, err := client.Ticketing.CreateAttachment(ctx, ticketID, &engine.CreateAttachmentRequest{
		FileName: "image.png",
		Content:  file,
	})
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Created attachment with id %s\n", attachmentResponse.Result.Id)

	return attachmentResponse, nil
}

// createComment creates a comment in the ticket with the given ID
func createComment(ctx context.Context, client *engineClient.Client, ticketID, comment string) (*engine.CreateCommentResponse, error) {
	res, err := client.Ticketing.CreateComment(ctx, ticketID, &engine.CreateCommentRequest{
		Content: comment,
	})
	if err != nil {
		return nil, err

		consoleLogger.Println()
	}

	consoleLogger.Printf("Created a comment with id %s\n", res.Result.Id)

	return res, nil
}

// createTicket creates a ticket in the tenant's ticketing connector
func createTicket(ctx context.Context, client *engineClient.Client, ticket *Ticket) (*engine.CreateTicketResponse, error) {
	tags := []string{fmt.Sprintf("asset:%s", ticket.Uid)}

	priority := engine.PriorityLow
	switch ticket.Severity {
	case "Critical":
		priority = engine.PriorityUrgent
	case "High":
		priority = engine.PriorityHigh
	case "Medium":
		priority = engine.PriorityMedium
	}

	newTicket, err := client.Ticketing.CreateTicket(ctx, &engine.CreateTicketRequest{
		Name:        ticket.Title,
		Summary:     ticket.Summary,
		Description: engine.String(ticket.Description),
		IssueType:   engine.String("Bug"),
		Priority:    priority.Ptr(),
		Project:     engine.String("Test"),
		Tags:        tags,
	})
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Created a ticket with id %s\n", newTicket.Result.Id)

	return newTicket, nil
}

// deleteAttachment deletes the attachment with the given ID in the ticket with the given ID
func deleteAttachment(ctx context.Context, client *engineClient.Client, ticketID, attachmentID string) error {
	if err := client.Ticketing.DeleteAttachment(ctx, ticketID, attachmentID); err != nil {
		return err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Deleted attachment with id %s\n", attachmentID)

	return nil
}

// deleteComment deletes the comment with the given ID in the ticket with the given ID
func deleteComment(ctx context.Context, client *engineClient.Client, ticketID, commentID string) error {
	if err := client.Ticketing.DeleteComment(ctx, ticketID, commentID); err != nil {
		return err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Deleted comment with id %s\n", commentID)

	return nil
}

// getAttachment retrieves the attachment with the given ID in the ticket with the given ID
func getAttachment(ctx context.Context, client *engineClient.Client, ticketID, attachmentID string) (*engine.DownloadAttachmentResponse, error) {
	attachment, err := client.Ticketing.DownloadAttachment(ctx, ticketID, attachmentID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Downloaded attachment %s with id %s\n", attachment.Result.FileName, attachmentID)

	return attachment, nil
}

// getAttachmentsMetadata retrieves the attachments metadata in the ticket with the given ID
func getAttachmentsMetadata(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.ListAttachmentsMetadataResponse, error) {
	attachments, err := client.Ticketing.ListAttachmentsMetadata(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Retrieved attachments metadata in the ticket with id %s\n", ticketID)

	return attachments, nil
}

// getComments retrieves the comments in the ticket with the given ID
func getComments(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.ListCommentsResponse, error) {
	comments, err := client.Ticketing.ListComments(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Retrieved comments in the ticket with id %s\n", ticketID)

	return comments, nil
}

// getProject retrieves all the projects in the tenant's ticketing integreation
func getProjects(ctx context.Context, client *engineClient.Client) (*engine.ListProjectsResponse, error) {
	projects, err := client.Ticketing.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Retrieved list of projects\n")

	return projects, nil
}

// getRemoteFields retrieves all the remote fields for all the projects in the tenant's ticketing integration
func getRemoteFields(ctx context.Context, client *engineClient.Client) (*engine.ListRemoteFieldsResponse, error) {
	listRemoteFields, err := client.Ticketing.ListRemoteFields(ctx)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Retrieved list of remote fields\n")

	return listRemoteFields, nil
}

// getTicket retrieves the ticket with the given ID
func getTicket(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.GetTicketResponse, error) {
	ticket, err := client.Ticketing.GetTicket(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Retrieved ticket with id %s\n", ticket.Result.Id)

	return ticket, nil
}

// searchTickets searches for an existing ticket with the given tags
func searchTickets(ctx context.Context, client *engineClient.Client, tags []string) (*engine.QueryTicketsResponse, error) {
	res, err := client.Ticketing.QueryTickets(ctx, &engine.QueryTicketsRequest{
		Filter: []*string{
			engine.String(fmt.Sprintf("tags[in]%s", strings.Join(tags, ","))),
		},
	})
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Found %d existing tickets\n", len(res.Result))

	return res, nil
}

// updateTicket updates the ticket with the given ID
func updateTicket(ctx context.Context, client *engineClient.Client, ticketID string, ticket *Ticket) (*engine.PatchTicketResponse, error) {
	patchRequest := []*engine.PatchOperation{
		{Op: "replace", Path: "/description", Value: ticket.Description},
		{Op: "replace", Path: "/summary", Value: ticket.Summary},
	}

	ticketPatched, err := client.Ticketing.PatchTicket(ctx, ticketID, patchRequest)
	if err != nil {
		return nil, err
	}

	consoleLogger.Println()

	consoleLogger.Printf("Patched ticket with id %s\n", ticketPatched.Result.Id)

	return ticketPatched, nil
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	parser := koanfyaml.Parser()
	if err := k.Load(file.Provider("config.yaml"), parser); err != nil {
		k.Load(env.Provider("SYNQLY_", "_", func(s string) string {
			return strings.ToLower(s)
		}), nil)
	}

	synqlyOrgToken = k.String("synqly.org.token")
	jiraURL = k.String("synqly.jira.url")
	jiraUser = k.String("synqly.jira.user")
	jiraToken = k.String("synqly.jira.token")

	ctx := context.Background()

	if synqlyOrgToken == "" {
		log.Fatal("Missing Synqly Org token")
	}
	if jiraToken == "" || jiraURL == "" || jiraUser == "" {
		log.Fatal("Must set Jira URL, Jira User, and Jira Token")
	}

	// Instantiate App object
	app := NewApp()

	// Create an interrupt handler to clean up tenants if the program is shut down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		// Intercept all ^C
		for range c {
			app.cleanup()
		}
	}()
	// Also be sure to run clean up if the program exits gracefully
	defer app.cleanup()

	uid := uuid.New().String()

	ticket := &Ticket{
		Uid:         uid,
		Title:       "Title of the vulnerability ticket",
		Summary:     "Summary of the vulnerability",
		Description: "Description of the vulnerability",
	}

	consoleLogger.Print("Creating Tenant ABC tenant\n")
	if err := app.NewTenant(ctx, "Tenant ABC"); err != nil {
		log.Fatal(err)
	}

	err := app.configureTicketing(ctx, "jira")
	if err != nil {
		log.Fatal(err)
	}

	ticketingClient := app.Tenant.TicketClient

	// get the list of projects
	listProjectsResponse, err := getProjects(ctx, ticketingClient)
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger.Println()

	consoleLogger.Printf("List of projects:\n")
	for _, project := range listProjectsResponse.Result {
		consoleLogger.Printf("Project Name: %s, Project Id: %s\n", project.Name, project.Id)
	}

	// get the list of remote fields
	listRemoteResponse, err := getRemoteFields(ctx, ticketingClient)
	if err != nil {
		log.Fatal(err)
	}

	specsMap := make(map[string]*engine.RemoteField, len(listRemoteResponse.Result))
	for _, spec := range listRemoteResponse.Result {
		if spec == nil {
			continue
		}
		specsMap[spec.ProviderFieldName] = spec
	}

	consoleLogger.Println()

	consoleLogger.Printf("List of 10 remote fields:\n")
	i := 0
	for key, spec := range specsMap {
		consoleLogger.Printf("Field Name: %s, Field Id: %s\n", key, spec.FieldId)
		if i == 10 {
			break
		}
		i++
	}

	// create a new ticket
	newTicket, err := createTicket(ctx, ticketingClient, ticket)
	if err != nil {
		log.Fatal(err)
	}

	// create an attachment in the ticket
	createAttachmentResponse, err := createAttachment(ctx, ticketingClient, newTicket.Result.Id)
	if err != nil {
		log.Fatal(err)
	}

	// get the attachment
	_, err = getAttachment(ctx, ticketingClient, newTicket.Result.Id, createAttachmentResponse.Result.Id)
	if err != nil {
		log.Fatal(err)
	}

	// delete the attachment
	err = deleteAttachment(ctx, ticketingClient, newTicket.Result.Id, createAttachmentResponse.Result.Id)
	if err != nil {
		log.Fatal(err)
	}

	// create a comment in the ticket
	createCommentResponse, err := createComment(ctx, ticketingClient, newTicket.Result.Id, "This is a comment")
	if err != nil {
		log.Fatal(err)
	}

	// get the comments in the ticket
	comments, err := getComments(ctx, ticketingClient, newTicket.Result.Id)
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger.Println()

	consoleLogger.Printf("Comments in the ticket:\n")
	for _, comment := range comments.Result {
		consoleLogger.Printf("Comment: %s\n", comment.Content)
	}

	// delete the comment
	err = deleteComment(ctx, ticketingClient, newTicket.Result.Id, createCommentResponse.Result.Id)
	if err != nil {
		log.Fatal(err)
	}

	// update the ticket
	ticket.Description = "Updated description of the vulnerability"
	ticket.Summary = "Updated summary of the vulnerability"
	_, err = updateTicket(ctx, ticketingClient, newTicket.Result.Id, ticket)
	if err != nil {
		log.Fatal(err)
	}
}
