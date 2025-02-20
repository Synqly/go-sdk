package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	"github.com/synqly/go-sdk/examples/common"
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

func ProviderConfig(c Config) *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		IdentityOkta: &mgmt.IdentityOkta{
			Url: c.url,
			Credential: &mgmt.OktaCredential{
				OAuthClient: &mgmt.OAuthClientCredential{
					ClientId:     c.clientId,
					ClientSecret: c.token,
				},
			},
		},
	}
}

func waitForAuditLogResult(ctx context.Context, client *engineClient.Client, req *engine.QueryIdentityAuditLogRequest) (*engine.QueryIdentityAuditLogResponse, error) {
	var resp *engine.QueryIdentityAuditLogResponse = nil
	var err error
	fmt.Print("Waiting for the audit log to update ")
	for attempt := 1; attempt < 60 && resp == nil; attempt++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		if attempt%10 != 0 {
			continue
		}

		resp, err = client.Identity.QueryAuditLog(ctx, req)
		if err != nil {
			return nil, err
		}

		if len(resp.Result) > 0 {
			fmt.Println()
			return resp, nil
		}
	}

	return nil, errors.New("no audit log found")
}

// demoActions demonstrates how to use the Synqly Identity API to perform actions on a user
// First, it looks up the user's ID by email address because actions are performed on a user ID.
// Next, it forces a password reset for the user, and waits for the audit log to update with the
// result of this action. Then, it disables the user, and again waits for the audit log to update.
// Finally, it re-enables the user, and waits for the audit log to update.
func demoActions(config Config) error {
	ctx := context.Background()

	t, err := common.NewTenant(ctx, "Secure Identities Corp", "tenant_store-okta.yaml", config.synqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdIdentity: {
			Name:           engine.String("iam"),
			ProviderConfig: ProviderConfig(config),
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create tenant: %w", err)
	}

	consoleLogger.Printf("Looking up user ID for email %s", config.userEmail)

	user, err := t.Synqly.EngineClients["identity"].Identity.QueryUsers(ctx, &engine.QueryUserRequest{
		Filter: []*string{
			engine.String("entity.user.name[eq]" + config.userEmail),
		},
	})
	if err != nil {
		return err
	}
	if len(user.Result) == 0 {
		return fmt.Errorf("user with email %s not found", config.userEmail)
	}

	b, _ := json.MarshalIndent(user.Result, "", "  ")
	consoleLogger.Printf("User details:\n%s\n\n\n\n\n\n", b)

	groupName := "Example Group"
	consoleLogger.Printf("Looking up group %s", groupName)

	groups, err := t.Synqly.EngineClients["identity"].Identity.QueryGroups(ctx, &engine.QueryGroupRequest{})
	if err != nil {
		return err
	}

	b, _ = json.MarshalIndent(groups.Result, "", "  ")
	consoleLogger.Printf("Groups:\n%s\n\n\n\n\n\n", b)

	userID := *user.Result[0].Entity.Uid

	// Reset user's password

	now := time.Now().Add(-1 * time.Second).UTC().Format(time.RFC3339)

	fmt.Printf("\n\n\n\n")
	consoleLogger.Printf("Forcing password reset for user %s", config.userEmail)

	err = t.Synqly.EngineClients["identity"].Identity.ForceUserPasswordReset(ctx, userID)
	if err != nil {
		return fmt.Errorf("error forcing password reset: %w", err)
	} else {
		consoleLogger.Printf("Forced password reset")
	}

	resp, err := waitForAuditLogResult(ctx, t.Synqly.EngineClients["identity"], &engine.QueryIdentityAuditLogRequest{
		Filter: []*string{
			engine.String("user.email_addr[eq]" + config.userEmail),
			engine.String("time[gte]" + now),
		},
		Order: []*string{
			engine.String("time[desc]"),
		},
	})
	if err != nil {
		log.Print(err)
	}

	logs, _ := json.MarshalIndent(resp.Result, "", "  ")
	fmt.Println()
	consoleLogger.Printf("Audit logs for user %s:\n%s\n", config.userEmail, logs)

	// Disable user

	now = time.Now().Add(-1 * time.Second).UTC().Format(time.RFC3339)

	fmt.Printf("\n\n\n\n")
	consoleLogger.Printf("Disabling user %s", config.userEmail)

	err = t.Synqly.EngineClients["identity"].Identity.DisableUser(ctx, userID)
	if err != nil {
		consoleLogger.Printf("Error disabling user: %s", err)
	} else {
		consoleLogger.Printf("Disabled user")
	}

	resp, err = waitForAuditLogResult(ctx, t.Synqly.EngineClients["identity"], &engine.QueryIdentityAuditLogRequest{
		Filter: []*string{
			engine.String("user.email_addr[eq]" + config.userEmail),
			engine.String("time[gte]" + now),
		},
		Order: []*string{
			engine.String("time[desc]"),
		},
		Cursor: &resp.Cursor,
	})
	if err != nil {
		log.Print(err)
	}

	logs, _ = json.MarshalIndent(resp.Result, "", "  ")
	fmt.Println()
	consoleLogger.Printf("Audit logs for user %s:\n%s\n", config.userEmail, logs)

	time.Sleep(1 * time.Second)

	// Re-enable user

	fmt.Printf("\n\n\n\n")
	consoleLogger.Printf("Enabling user %s", config.userEmail)

	now = time.Now().Add(-1 * time.Second).UTC().Format(time.RFC3339)

	err = t.Synqly.EngineClients["identity"].Identity.EnableUser(ctx, userID)
	if err != nil {
		consoleLogger.Printf("Error enabling user: %s", err)
	} else {
		consoleLogger.Printf("Enabled user")
	}

	resp, err = waitForAuditLogResult(ctx, t.Synqly.EngineClients["identity"], &engine.QueryIdentityAuditLogRequest{
		Filter: []*string{
			engine.String("user.email_addr[eq]" + config.userEmail),
			engine.String("time[gte]" + now),
		},
		Order: []*string{
			engine.String("time[desc]"),
		},
		Cursor: &resp.Cursor,
	})
	if err != nil {
		log.Print(err)
	}

	logs, _ = json.MarshalIndent(resp.Result, "", "  ")
	fmt.Println()
	consoleLogger.Printf("Audit logs for user %s:\n%s\n", config.userEmail, logs)

	return nil
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	// Load config variables from the env file
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	if config.synqlyOrgToken == "" || config.clientId == "" || config.token == "" || config.userEmail == "" || config.url == "" {
		log.Fatal("Must set following environment variables: SYNQLY_ORG_TOKEN CLIENT_ID TOKEN USER_EMAIL URL")
	}

	if err := demoActions(config); err != nil {
		log.Fatal(err)
	}
}
