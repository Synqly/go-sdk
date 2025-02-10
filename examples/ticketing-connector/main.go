package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	koanfyaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	"github.com/synqly/go-sdk/examples/common"
)

var (
	k = koanf.New(".")
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

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
		Project:     engine.String("SDK Examples"),
		Tags:        tags,
	})
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Created a ticket with id %s\n", newTicket.Result.Id)

	return newTicket, nil
}

// deleteComment deletes the comment with the given ID in the ticket with the given ID
func deleteComment(ctx context.Context, client *engineClient.Client, ticketID, commentID string) error {
	if err := client.Ticketing.DeleteComment(ctx, ticketID, commentID); err != nil {
		return err
	}

	consoleLogger.Printf("Deleted comment with id %s\n", commentID)

	return nil
}

// getAttachmentsMetadata retrieves the attachments metadata in the ticket with the given ID
func getAttachmentsMetadata(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.ListAttachmentsMetadataResponse, error) {
	attachments, err := client.Ticketing.ListAttachmentsMetadata(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Retrieved attachments metadata in the ticket with id %s\n", ticketID)

	return attachments, nil
}

// getComments retrieves the comments in the ticket with the given ID
func getComments(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.ListCommentsResponse, error) {
	comments, err := client.Ticketing.ListComments(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Retrieved comments in the ticket with id %s\n", ticketID)

	return comments, nil
}

// getProject retrieves all the projects in the tenant's ticketing integreation
func getProjects(ctx context.Context, client *engineClient.Client) (*engine.ListProjectsResponse, error) {
	projects, err := client.Ticketing.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Retrieved list of projects\n")

	return projects, nil
}

// getRemoteFields retrieves all the remote fields for all the projects in the tenant's ticketing integration
func getRemoteFields(ctx context.Context, client *engineClient.Client) (*engine.ListRemoteFieldsResponse, error) {
	listRemoteFields, err := client.Ticketing.ListRemoteFields(ctx)
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Retrieved list of remote fields\n")

	return listRemoteFields, nil
}

// searchTicket searches for an existing ticket with the given tags
func searchTicket(ctx context.Context, client *engineClient.Client, tags []string) (*engine.QueryTicketsResponse, error) {
	res, err := client.Ticketing.QueryTickets(ctx, &engine.QueryTicketsRequest{
		Filter: []*string{
			engine.String(fmt.Sprintf("tags[in]%s", strings.Join(tags, ","))),
		},
	})
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Found %d existing tickets\n", len(res.Result))

	return res, nil
}

// getTicket retrieves the ticket with the given ID
func getTicket(ctx context.Context, client *engineClient.Client, ticketID string) (*engine.GetTicketResponse, error) {
	ticket, err := client.Ticketing.GetTicket(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Retrieved ticket with id %s\n", ticket.Result.Id)

	return ticket, nil
}

// patchTicket updates the ticket with the given ID
func patchTicket(ctx context.Context, client *engineClient.Client, ticketID string, ticket *Ticket) (*engine.PatchTicketResponse, error) {
	patchRequest := []*engine.PatchOperation{
		{Op: "replace", Path: "/description", Value: ticket.Description},
		{Op: "replace", Path: "/summary", Value: ticket.Summary},
	}

	ticketPatched, err := client.Ticketing.PatchTicket(ctx, ticketID, patchRequest)
	if err != nil {
		return nil, err
	}

	consoleLogger.Printf("Patched ticket with id %s\n", ticketPatched.Result.Id)

	return ticketPatched, nil
}

func ticketingProviderConfig(jiraURL, jiraUser, jiraToken string) (*mgmt.CreateIntegrationRequest, error) {
	if jiraUser == "" || jiraToken == "" {
		return nil, fmt.Errorf("missing ticketing provider config")
	}

	integrationReq := &mgmt.CreateIntegrationRequest{
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

	return integrationReq, nil
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	parser := koanfyaml.Parser()
	if err := k.Load(file.Provider("config.yaml"), parser); err != nil {
		k.Load(env.Provider("SYNQLY_", "_", func(s string) string {
			return strings.ToLower(s)
		}), nil)
	}

	synqlyOrgToken := k.String("synqly.org.token")
	jiraURL := k.String("synqly.jira.url")
	jiraUser := k.String("synqly.jira.user")
	jiraToken := k.String("synqly.jira.token")

	if synqlyOrgToken == "" {
		log.Fatal("Missing Synqly Org token")
	}

	ticketingProvider, err := ticketingProviderConfig(jiraURL, jiraUser, jiraToken)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	t, err := common.NewTenant(ctx, "Zenith Systems", "tenant_store_qualys.yaml", synqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdTicketing: ticketingProvider,
	})
	if err != nil {
		log.Fatal(err)
	}

	ticket := &Ticket{
		Uid:         "123",
		Title:       "Title of the vulnerability ticket",
		Summary:     "Summary of the vulnerability",
		Description: "Description of the vulnerability",
	}

	ticketingClient := t.Synqly.EngineClients["ticketing"]

	// get the list of projects
	listProjectsResponse, err := getProjects(ctx, ticketingClient)
	if err != nil {
		log.Fatal(err)
	}

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

	consoleLogger.Printf("List of remote fields:\n")
	for key, spec := range specsMap {
		consoleLogger.Printf("Field Name: %s, Field Id: %s\n", key, spec.FieldId)
	}

	// create a new ticket
	newTicket, err := createTicket(ctx, ticketingClient, ticket)
	if err != nil {
		log.Fatal(err)
	}

	// create a comment in the ticket
	createComment(ctx, ticketingClient, newTicket.Result.Id, "This is a comment")

	// get the comments in the ticket
	comments, err := getComments(ctx, ticketingClient, newTicket.Result.Id)
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger.Printf("Comments in the ticket:\n")
	for _, comment := range comments.Result {
		consoleLogger.Printf("Comment: %s\n", comment.Content)
	}

	// update the ticket
	ticket.Description = "Updated description of the vulnerability"
	ticket.Summary = "Updated summary of the vulnerability"
	patchedTicket, err := patchTicket(ctx, ticketingClient, newTicket.Result.Id, ticket)
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger.Printf("Updated description %s\n", *patchedTicket.Result.Description)
	consoleLogger.Printf("Updated summary %s\n", patchedTicket.Result.Summary)

}
