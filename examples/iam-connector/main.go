package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	k        = koanf.New(".")
	oktaConf = oktaConfig{}
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

type provider interface {
	ProviderConfig() *mgmt.ProviderConfig
}

type oktaConfig struct {
	URL      string `koanf:"url"`
	Token    string `koanf:"token"`
	ClientId string `koanf:"clientid"`
}

func (c *oktaConfig) ProviderConfig() *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		IdentityOkta: &mgmt.IdentityOkta{
			Url: c.URL,
			Credential: &mgmt.OktaCredential{
				OAuthClient: &mgmt.OAuthClientCredential{
					ClientId:     c.ClientId,
					ClientSecret: c.Token,
				},
			},
		},
	}
}

func (c *oktaConfig) Validate() error {
	if c.URL == "" {
		return errors.New("okta.url is required")
	}
	if c.Token == "" {
		return errors.New("okta.token is required")
	}
	if c.ClientId == "" {
		return errors.New("okta.client.id is required")
	}
	return nil
}

type entraConfig struct {
	TenantId     string `koanf:"tenant_id"`
	ClientId     string `koanf:"client_id"`
	ClientSecret string `koanf:"client_secret"`
}

func (c *entraConfig) ProviderConfig() *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		IdentityEntraId: &mgmt.IdentityEntraId{
			TenantId: c.TenantId,
			Credential: &mgmt.EntraIdCredential{
				OAuthClient: &mgmt.OAuthClientCredential{
					ClientId:     c.ClientId,
					ClientSecret: c.ClientSecret,
				},
			},
		},
	}
}

func waitForAuditLogResult(ctx context.Context, client *engineClient.Client, req *engine.QueryIdentityAuditLogRequest) (*engine.QueryIdentityAuditLogResponse, error) {
	var resp *engine.QueryIdentityAuditLogResponse = nil
	fmt.Print("Waiting for the audit log to update ")
	for attempt := 1; attempt < 60 && resp == nil; attempt++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		if attempt%10 != 0 {
			continue
		}

		resp, err := client.Identity.QueryAuditLog(ctx, req)
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
func demoActions(userEmail, orgToken string, p provider) error {
	ctx := context.Background()

	t, err := common.NewTenant(ctx, "Secure Identities Corp", "tenant_store-okta.yaml", orgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdIdentity: {
			Name:           engine.String("iam"),
			ProviderConfig: p.ProviderConfig(),
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create tenant: %w", err)
	}

	consoleLogger.Printf("Looking up user ID for email %s", userEmail)

	user, err := t.Synqly.EngineClients["identity"].Identity.QueryUsers(ctx, &engine.QueryUserRequest{
		Filter: []*string{
			engine.String("entity.user.name[eq]" + userEmail),
		},
	})
	if err != nil {
		return err
	}
	if len(user.Result) == 0 {
		return fmt.Errorf("user with email %s not found", userEmail)
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
	consoleLogger.Printf("Forcing password reset for user %s", userEmail)

	err = t.Synqly.EngineClients["identity"].Identity.ForceUserPasswordReset(ctx, userID)
	if err != nil {
		return fmt.Errorf("error forcing password reset: %w", err)
	} else {
		consoleLogger.Printf("Forced password reset")
	}

	resp, err := waitForAuditLogResult(ctx, t.Synqly.EngineClients["identity"], &engine.QueryIdentityAuditLogRequest{
		Filter: []*string{
			engine.String("user.email_addr[eq]" + userEmail),
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
	consoleLogger.Printf("Audit logs for user %s:\n%s\n", userEmail, logs)

	// Disable user

	now = time.Now().Add(-1 * time.Second).UTC().Format(time.RFC3339)

	fmt.Printf("\n\n\n\n")
	consoleLogger.Printf("Disabling user %s", userEmail)

	err = t.Synqly.EngineClients["identity"].Identity.DisableUser(ctx, userID)
	if err != nil {
		consoleLogger.Printf("Error disabling user: %s", err)
	} else {
		consoleLogger.Printf("Disabled user")
	}

	resp, err = waitForAuditLogResult(ctx, t.Synqly.EngineClients["identity"], &engine.QueryIdentityAuditLogRequest{
		Filter: []*string{
			engine.String("user.email_addr[eq]" + userEmail),
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
	consoleLogger.Printf("Audit logs for user %s:\n%s\n", userEmail, logs)

	time.Sleep(1 * time.Second)

	// Re-enable user

	fmt.Printf("\n\n\n\n")
	consoleLogger.Printf("Enabling user %s", userEmail)

	now = time.Now().Add(-1 * time.Second).UTC().Format(time.RFC3339)

	err = t.Synqly.EngineClients["identity"].Identity.EnableUser(ctx, userID)
	if err != nil {
		consoleLogger.Printf("Error enabling user: %s", err)
	} else {
		consoleLogger.Printf("Enabled user")
	}

	resp, err = waitForAuditLogResult(ctx, t.Synqly.EngineClients["identity"], &engine.QueryIdentityAuditLogRequest{
		Filter: []*string{
			engine.String("user.email_addr[eq]" + userEmail),
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
	consoleLogger.Printf("Audit logs for user %s:\n%s\n", userEmail, logs)

	return nil
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	// load config -- try from a config file, and if that is not present, then try env vars
	parser := koanfyaml.Parser()
	if err := k.Load(file.Provider("config-okta.yaml"), parser); err != nil {
		k.Load(env.Provider("SYNQLY_", "_", func(s string) string {
			return strings.ToLower(s)
		}), nil)
	}

	orgToken := k.String("synqly.org.token")
	if orgToken == "" {
		log.Fatal("synqly.org.token is required. Set SYNQLY_ORG_TOKEN or add it to config.yaml.")
	}

	oktaConf := oktaConfig{}

	err := k.Unmarshal("synqly.okta", &oktaConf)
	if err != nil {
		log.Fatal(err)
	}
	if err := oktaConf.Validate(); err != nil {
		log.Fatalf("invalid okta config: %s. Set Okta configuration through env vars or add it to config.yaml", err)
	}

	// entraConf := entraConfig{}
	// err := k.Unmarshal("synqly.entra_id", &entraConf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	userEmail := k.String("synqly.user.email")
	if userEmail == "" {
		log.Fatal("synqly.okta.user.email is required. Set SYNQLY_OKTA_USER_EMAIL or add it to config.yaml.")
	}

	if err := demoActions(userEmail, orgToken, &oktaConf); err != nil {
		log.Fatal(err)
	}
}
