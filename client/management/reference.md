# Reference
## Accounts
<details><summary><code>client.Accounts.List() -> *management.ListAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Account` objects. For more information on
Organizations and Accounts, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListAccountsRequest{}
client.Accounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Account` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Account` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*bool` — Return total number of accounts in the system, respecting all applied filters. This is expensive, use sparingly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Get(AccountId) -> *management.GetAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Account` object matching `{accountId}`. For more information on
Organizations and Accounts, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.Get(
        context.TODO(),
        "accountId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Create(request) -> *management.CreateAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Account` object. For more information on Organizations and
Accounts, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateAccountRequest{}
client.Accounts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateAccountRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Update(AccountId, request) -> *management.UpdateAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Account` object matching `{accountId}`. For more information on
Organizations and Accounts, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Account{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Fullname: "fullname",
        OrganizationId: "organization_id",
        Environment: management.EnvironmentTest,
    }
client.Accounts.Update(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateAccountRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Patch(AccountId, request) -> *management.PatchAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `Account` object matching `{accountId}`. For more information on
Organizations and Accounts, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Accounts.Patch(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Delete(AccountId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the `Account` matching `{accountId}`. Deleting an `Account` also deletea
all `Tokens` and `Credentials` belonging to the `Account`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.Delete(
        context.TODO(),
        "accountId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Audit
<details><summary><code>client.Audit.List() -> *management.ListAuditEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all Synqly `Audit` events for the `Organization`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListAuditEventsRequest{}
client.Audit.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Audit` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Audit` objects starting after this `created_at`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — The order defaults to created_at[asc] and can changed to descending order by specifying created_at[desc].
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Auth
<details><summary><code>client.Auth.Logon(OrganizationId, request) -> *management.LogonResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Initiates a new session for the given member in specified Synqly organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.LogonRequest{
        Name: "name",
        Secret: "secret",
    }
client.Auth.Logon(
        context.TODO(),
        "organizationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationId:** `management.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.LogonRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.ChangePassword(request) -> *management.ChangePasswordResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Change member password.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ChangePasswordRequest{
        OldSecret: "old_secret",
        NewSecret: "new_secret",
    }
client.Auth.ChangePassword(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.ChangePasswordRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.Logoff() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminates the session identified by the given logon token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Auth.Logoff(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.CreateSso(request) -> *management.CreateSsoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Configure an identity provider for Single Sign-On (SSO) with the organization.
Returns the created resource with its assigned SSO configuration ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateSsoRequest{
        Fullname: "fullname",
        Config: &management.CreateSsoConfiguration{
            Oidc: &management.CreateOidcSsoConfiguration{
                IssuerUrl: "issuer_url",
                ClientId: "client_id",
                ClientSecret: "client_secret",
            },
        },
    }
client.Auth.CreateSso(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateSsoRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.ListSso() -> *management.ListSsoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all Single Sign-On (SSO) configurations for the organization.
Client secrets are not included.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Auth.ListSso(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.GetSso(SsoId) -> *management.GetSsoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific Single Sign-On (SSO) configuration. Client
secrets are not included.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Auth.GetSso(
        context.TODO(),
        "ssoId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ssoId:** `management.SsoConfigurationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.GetSsoMetadata() -> *management.GetSsoMetadataResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns configuration that must be used when configuring Identity
Providers, such as the OIDC redirect URI or SAML ACS URL and SP
entity ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Auth.GetSsoMetadata(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.UpdateSso(SsoId, request) -> *management.UpdateSsoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the SSO configuration by ID. This is a full replacement of the
configuration. The `updated_at` field is used for optimistic locking
to prevent concurrent updates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSsoRequest{
        Fullname: "fullname",
        Config: &management.UpdateSsoConfiguration{
            Oidc: &management.UpdateOidcSsoConfiguration{
                IssuerUrl: "issuer_url",
                ClientId: "client_id",
            },
        },
    }
client.Auth.UpdateSso(
        context.TODO(),
        "ssoId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ssoId:** `management.SsoConfigurationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.UpdateSsoRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Auth.DeleteSso(SsoId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a specific Single Sign-On (SSO) configuration. This may disable signing on
with the identity provider defined in the configuration, and can result in members
linked to that identity provider no longer being able to access the Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Auth.DeleteSso(
        context.TODO(),
        "ssoId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ssoId:** `management.SsoConfigurationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Billing
<details><summary><code>client.Billing.List() -> *management.ListBillingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of billing reports. Use a RootOrganization token for multiple organizations. For more information on Organizations and Billings, refer to our [Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListBillingRequest{}
client.Billing.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of billing reports to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return billing reports starting after this `organization_id`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `organization_id`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `organization_id[desc]` will sort the results by `organization_id` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.Get(OrganizationId, Month) -> *management.GetBillingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the billing reports matching `{organizationId} and {month}`. For more information on
Organizations and Billings, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Billing.Get(
        context.TODO(),
        "organizationId",
        management.BillingMonthPartial.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationId:** `management.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**month:** `*management.BillingMonth` — Specify month for a previously generated monthly report or 'partial' to generate the current month-to-date partial billing report.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bridges
<details><summary><code>client.Bridges.List(AccountId) -> *management.ListBridgesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Bridge Group` objects that match the query params.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListBridgesRequest{}
client.Bridges.List(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Bridge` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Bridge` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bridges.Get(AccountId, BridgeId) -> *management.GetBridgeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the Bridge Group object matching `{bridgeId}`. For more information on Bridges, refer to the
[Bridge Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bridges.Get(
        context.TODO(),
        "accountId",
        "bridgeId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**bridgeId:** `management.BridgeGroupId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bridges.GetStatus(AccountId, BridgeId) -> *management.GetBridgeStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the status and local configuration of running Bridges Agents in the Bridge Group `{bridgeId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bridges.GetStatus(
        context.TODO(),
        "accountId",
        "bridgeId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**bridgeId:** `management.BridgeGroupId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bridges.Create(AccountId, request) -> *management.CreateBridgeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Bridge Group` with a unique identifier and authentication
credentials. This allows for Bridge Agents to connect to Synqly. For
more information on Bridges, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateBridgeRequest{}
client.Bridges.Create(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.CreateBridgeRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bridges.Update(AccountId, BridgeId, request) -> *management.UpdateBridgeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Bridge Group` object matching `{bridgeId}`. For more information on Bridges, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.BridgeGroup{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Fullname: "fullname",
    }
client.Bridges.Update(
        context.TODO(),
        "accountId",
        "bridgeId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**bridgeId:** `management.BridgeGroupId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateBridgeRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bridges.Patch(AccountId, BridgeId, request) -> *management.PatchBridgeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `Bridge Group` object matching `{bridgeId}`. For more information on Bridges, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Bridges.Patch(
        context.TODO(),
        "accountId",
        "bridgeId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**bridgeId:** `management.BridgeGroupId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bridges.Delete(AccountId, BridgeId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the `Bridge Group` matching `{bridgeId}`. Deleting an `Bridge Group` also deletea
all `Tokens` and `Credentials` belonging to the `Bridge Group`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Bridges.Delete(
        context.TODO(),
        "accountId",
        "bridgeId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**bridgeId:** `management.BridgeGroupId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Capabilities
<details><summary><code>client.Capabilities.ListConnectors() -> *management.ListConnectorsCapabilitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Connectors`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListConnectorsCapabilitiesRequest{}
client.Capabilities.ListConnectors(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**expand:** `*management.ListConnectorCapabilitiesExpandOptions` 

Expand the capabilities result fields that are otherwise
omitted or returned as references to OpenAPI spec components.
NOTE: This can yield very big response objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Capabilities.ListProviders() -> *management.ListProvidersCapabilitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all Provider capabilities and their configurations.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListProvidersCapabilitiesRequest{}
client.Capabilities.ListProviders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**expand:** `*management.ListProviderCapabilitiesExpandOptions` 

Expand the capabilities result fields that are otherwise
omitted or returned as references to OpenAPI spec components.
NOTE: This can yield very big response objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Capabilities.GetProvider(ProviderId) -> *management.ProviderCapabilitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the capabilities and configurations for a specific Provider type
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ProviderCapabilitiesRequest{}
client.Capabilities.GetProvider(
        context.TODO(),
        management.ProviderConfigIdAppsecAmazonInspector.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**providerId:** `*management.ProviderConfigId` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*management.GetProviderCapabilitiesExpandOptions` 

Expand the capabilities result fields that are otherwise
omitted or returned as references to OpenAPI spec components.
NOTE: This can yield very big response objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Credentials
<details><summary><code>client.Credentials.List(OwnerId) -> *management.ListCredentialsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Credential` objects belonging to the `Account`, `Integration`, or `IntegrationPoint`, or `OrganizationWebhook` matching
`{ownerId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListCredentialsRequest{}
client.Credentials.List(
        context.TODO(),
        "ownerId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Credential` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Credential` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.Get(OwnerId, CredentialId) -> *management.GetCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Credential` object matching `{credentialId}` where the
`Credential` belongs to the `Account`, `Integration`, `IntegrationPoint` or `OrganizationWebhook` matching `{ownerId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credentials.Get(
        context.TODO(),
        "ownerId",
        "credentialId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**credentialId:** `management.CredentialId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.Lookup(CredentialId) -> *management.LookupCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Credential` object matching `{credentialId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credentials.Lookup(
        context.TODO(),
        "credentialId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**credentialId:** `management.CredentialId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.Create(OwnerId, request) -> *management.CreateCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Credential` object in the `Account`, `IntegrationPoint`, or `OrganizationWebhook` matching matching
`{ownerId}`. A `Credential` may only by used by a single `Account`, `Integration`, `IntegrationPoint` or `OrganizationWebhook`;
however, `Credential` objects can be shared by multiple `Integrations` within an `Account`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateCredentialRequest{}
client.Credentials.Create(
        context.TODO(),
        "ownerId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.CreateCredentialRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.Update(OwnerId, CredentialId, request) -> *management.UpdateCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Credential` object matching `{credentialId}`, where the
`Credential` belongs to the `Account`, `Integration`, `IntegrationPoint` or `OrganizationWebhook` matching `{ownerId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Credential{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        OwnerType: management.OwnerTypeAccount,
        Fullname: "fullname",
        Managed: management.ManagedTypeManaged,
    }
client.Credentials.Update(
        context.TODO(),
        "ownerId",
        "credentialId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**credentialId:** `management.CredentialId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateCredentialRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.Patch(OwnerId, CredentialId, request) -> *management.PatchCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `Credential` object matching `{credentialId}`, where the
`Credential` belongs to the `Account`, `Integration`, `IntegrationPoint` or `OrganizationWebhook` matching `{ownerId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Credentials.Patch(
        context.TODO(),
        "ownerId",
        "credentialId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**credentialId:** `management.CredentialId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.Delete(OwnerId, CredentialId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the `Credential` object matching `{credentialId}`, where the
`Credential` belongs to the `Account`, `Integration`, `IntegrationPoint` or `OrganizationWebhook` matching `{ownerId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credentials.Delete(
        context.TODO(),
        "ownerId",
        "credentialId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**credentialId:** `management.CredentialId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Customs
<details><summary><code>client.Customs.List() -> *management.ListCustomsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Custom` provider objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListCustomsRequest{}
client.Customs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of custom provider objects to return in this page. Defaults to 100. Items listed will not include the `data` field.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return custom provider objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*bool` — Return total number of custom providers in the system, respecting all applied filters. This is expensive, use sparingly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customs.Get(CustomId) -> *management.GetCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Custom` provider object matching `{customId}`. This response will include the full `data` field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Customs.Get(
        context.TODO(),
        "customId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customId:** `management.CustomId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customs.Create(request) -> *management.CreateCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an custom provider object.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateCustomRequest{
        Data: "data",
    }
client.Customs.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateCustomRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customs.Update(CustomId, request) -> *management.UpdateCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the custom provider object matching `{customId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Custom{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Fullname: "fullname",
        OrganizationId: "organization_id",
        Data: "data",
    }
client.Customs.Update(
        context.TODO(),
        "customId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customId:** `management.CustomId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateCustomRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customs.Patch(CustomId, request) -> *management.PatchCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the custom provider object matching `{customId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Customs.Patch(
        context.TODO(),
        "customId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customId:** `management.CustomId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Customs.Delete(CustomId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the custom provider object matching `{customId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Customs.Delete(
        context.TODO(),
        "customId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customId:** `management.CustomId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## IntegrationPoints
<details><summary><code>client.IntegrationPoints.List() -> *management.ListIntegrationPointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `IntegrationPoint` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListIntegrationPointsRequest{}
client.IntegrationPoints.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `IntegrationPoint` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `IntegrationPoint` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*bool` — Return total number of integration points in the system, respecting all applied filters. This is expensive, use sparingly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IntegrationPoints.Get(IntegrationPointId) -> *management.GetIntegrationPointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `IntegrationPoint` object matching `{integrationPointId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.IntegrationPoints.Get(
        context.TODO(),
        "integrationPointId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**integrationPointId:** `management.IntegrationPointId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IntegrationPoints.Create(request) -> *management.CreateIntegrationPointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a `IntegrationPoint` object.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateIntegrationPointRequest{
        Connector: management.CategoryIdAppsec,
        Environments: &management.IntegrationEnvironments{},
    }
client.IntegrationPoints.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateIntegrationPointRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IntegrationPoints.Update(IntegrationPointId, request) -> *management.UpdateIntegrationPointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `IntegrationPoint` object matching `{integrationPointId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.IntegrationPoint{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Connector: management.CategoryIdAppsec,
        Environments: &management.IntegrationEnvironments{},
    }
client.IntegrationPoints.Update(
        context.TODO(),
        "integrationPointId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**integrationPointId:** `management.IntegrationPointId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateIntegrationPointRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IntegrationPoints.Patch(IntegrationPointId, request) -> *management.PatchIntegrationPointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `IntegrationPoint` object matching `{integrationPointId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.IntegrationPoints.Patch(
        context.TODO(),
        "integrationPointId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**integrationPointId:** `management.IntegrationPointId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IntegrationPoints.Delete(IntegrationPointId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the `IntegrationPoint` object matching `{integrationPointId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.IntegrationPoints.Delete(
        context.TODO(),
        "integrationPointId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**integrationPointId:** `management.IntegrationPointId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Integrations
<details><summary><code>client.Integrations.List() -> *management.ListIntegrationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Integration` objects that match the query params.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListIntegrationsRequest{}
client.Integrations.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Integration` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Integration` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*management.ListIntegrationOptions` — Expand the integration result with the related integration point and/or account information.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*bool` — Return total number of all integrations in the system, respecting all applied filters. This is expensive, use sparingly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.ListAccount(AccountId) -> *management.ListAccountIntegrationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Integration` objects belonging to the
`Account` matching `{accountId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListAccountIntegrationsRequest{}
client.Integrations.ListAccount(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Integration` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Integration` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*management.ListIntegrationOptions` — Expand the integration result with the related integration point and/or account information.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*bool` — Return total number of integrations for a particular account. This is expensive, use sparingly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.Get(AccountId, IntegrationId) -> *management.GetIntegrationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Integration` object matching `{integrationId}` where
the `Integration` belongs to the `Account` matching `{accountId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.Get(
        context.TODO(),
        "accountId",
        "integrationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.Create(AccountId, request) -> *management.CreateIntegrationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Integration` object belonging to the `Account` matching
`{accountId}`. Configures the `Integration` with the Provider specified
in the request. Returns an `Integration` token for use with `Integration` APIs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateIntegrationRequest{
        ProviderConfig: &management.ProviderConfig{
            AppsecAmazonInspector: &management.AppsecAmazonInspector{
                Credential: &management.AwsProviderCredential{
                    Aws: &management.AwsCredential{
                        AccessKeyId: "access_key_id",
                        SecretAccessKey: "secret_access_key",
                    },
                },
                Region: management.AwsRegionUsEast1,
            },
        },
    }
client.Integrations.Create(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.CreateIntegrationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.Verify(AccountId, request) -> *management.VerifyIntegrationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verifies an ephemeral `Integration` and provider configuration and tests the authentication and provider connectivity.
The provider config credential IDs can utilize persistent IDs or use "#/n" reference IDs;
where (n) is the zero based offset in the optional credentials list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.VerifyIntegrationRequest{
        Integration: &management.CreateIntegrationRequest{
            ProviderConfig: &management.ProviderConfig{
                AppsecAmazonInspector: &management.AppsecAmazonInspector{
                    Credential: &management.AwsProviderCredential{
                        Aws: &management.AwsCredential{
                            AccessKeyId: "access_key_id",
                            SecretAccessKey: "secret_access_key",
                        },
                    },
                    Region: management.AwsRegionUsEast1,
                },
            },
        },
    }
client.Integrations.Verify(
        context.TODO(),
        "accountId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.VerifyIntegrationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.VerifyExisting(AccountId, IntegrationId, request) -> *management.VerifyIntegrationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verifies the `Integration` matching `{integrationId}` and tests authentication
and provider connectivity.

When called without a request body, the stored integration configuration is
verified. When a request body is provided, it uses the same shape as the
Update Integration API; the integration is verified as if that update were
applied, without persisting changes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.VerifyExisting(
        context.TODO(),
        "accountId",
        "integrationId",
        nil,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.UpdateIntegrationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.Update(AccountId, IntegrationId, request) -> *management.UpdateIntegrationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Integration` object matching `{integrationId}`, where the
`Integration` belongs to the `Account` matching `{accountId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Integration{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Fullname: "fullname",
        RefreshTokenId: "refresh_token_id",
        AccountId: "account_id",
        Category: management.CategoryIdAppsec,
        ProviderConfig: &management.ProviderConfig{
            AppsecAmazonInspector: &management.AppsecAmazonInspector{
                Credential: &management.AwsProviderCredential{
                    Aws: &management.AwsCredential{
                        AccessKeyId: "access_key_id",
                        SecretAccessKey: "secret_access_key",
                    },
                },
                Region: management.AwsRegionUsEast1,
            },
        },
        ProviderFullname: "provider_fullname",
        ProviderType: "provider_type",
    }
client.Integrations.Update(
        context.TODO(),
        "accountId",
        "integrationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateIntegrationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.Patch(AccountId, IntegrationId, request) -> *management.PatchIntegrationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `Integration` object matching `{integrationId}`, where the
`Integration` belongs to the `Account` matching `{accountId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Integrations.Patch(
        context.TODO(),
        "accountId",
        "integrationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.Delete(AccountId, IntegrationId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the `Integration` object matching `{integrationId}, where the
`Integration` belongs to the `Account` matching `{accountId}`. Deleting
an `Integration` also deletes any tokens that belong to it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.Delete(
        context.TODO(),
        "accountId",
        "integrationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Mappings
<details><summary><code>client.Mappings.List() -> *management.ListMappingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Mapping` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListMappingsRequest{}
client.Mappings.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of mapping objects to return in this page. Defaults to 100. Items listed will not include the `data` field.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return mapping objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**total:** `*bool` — Return total number of mappings in the system, respecting all applied filters. This is expensive, use sparingly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mappings.Get(MappingId) -> *management.GetMappingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Mapping` object matching `{mappingId}`. This response will include the full `data` field.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Mappings.Get(
        context.TODO(),
        "mappingId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mappingId:** `management.MappingId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mappings.Create(request) -> *management.CreateMappingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an mapping object.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateMappingRequest{
        Data: "data",
    }
client.Mappings.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateMappingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mappings.Update(MappingId, request) -> *management.UpdateMappingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the mapping object matching `{mappingId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Mapping{
        Id: "id",
        Fullname: "fullname",
        OrganizationId: "organization_id",
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Data: "data",
    }
client.Mappings.Update(
        context.TODO(),
        "mappingId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mappingId:** `management.MappingId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateMappingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mappings.Patch(MappingId, request) -> *management.PatchMappingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the mapping object matching `{mappingId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Mappings.Patch(
        context.TODO(),
        "mappingId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mappingId:** `management.MappingId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mappings.Delete(MappingId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the mapping object matching `{mappingId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Mappings.Delete(
        context.TODO(),
        "mappingId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mappingId:** `management.MappingId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mappings.Apply(request) -> *management.ApplyMappingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Apply a list of mapping transforms against the JSON input.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ApplyMappingRequest{
        Body: &management.ApplyMappingRequestBody{
            Mappings: []string{
                "mappings",
                "mappings",
            },
            Data: map[string]any{
                "data": map[string]any{
                    "key": "value",
                },
            },
        },
    }
client.Mappings.Apply(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the input in the response
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.ApplyMappingRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Members
<details><summary><code>client.Members.List() -> *management.ListMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all members
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListMembersRequest{}
client.Members.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Member` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Member` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Members.Get(MemberId) -> *management.GetMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a Member by ID
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Members.Get(
        context.TODO(),
        "memberId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberId:** `management.MemberId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Members.Create(request) -> *management.CreateMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new member for this Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateMemberRequest{
        Name: "name",
        Secret: "secret",
    }
client.Members.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateMemberRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Members.Update(MemberId, request) -> *management.UpdateMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a Member by ID
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Member{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        State: management.StateDisabled,
        LastLogon: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Fullname: "fullname",
        Ttl: "ttl",
        TokenTtl: "token_ttl",
        Expires: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        PinExpires: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        RoleBinding: []management.RoleName{
            "role_binding",
            "role_binding",
        },
        Type: management.MemberTypePersonal,
    }
client.Members.Update(
        context.TODO(),
        "memberId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberId:** `management.MemberId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateMemberRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Members.Patch(MemberId, request) -> *management.PatchMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a Member by ID
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Members.Patch(
        context.TODO(),
        "memberId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberId:** `management.MemberId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Members.Delete(MemberId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a Member by ID. Also deletes all Tokens for the Member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Members.Delete(
        context.TODO(),
        "memberId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberId:** `management.MemberId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Meta
<details><summary><code>client.Meta.ManagementOpenapiSpec() -> management.GetOpenApiSpecResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the OpenAPI spec for the Management API
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Meta.ManagementOpenapiSpec(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Meta.EngineOpenapiSpec() -> management.GetOpenApiSpecResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the OpenAPI spec for the Engine APIs
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Meta.EngineOpenapiSpec(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Operations
<details><summary><code>client.Operations.List() -> *management.ListOperationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Asynchronous Operations` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOperationsRequest{}
client.Operations.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Asynchronous Operations` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Asynchronous Operations` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `id`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field id. For example, `id[desc]` will sort the results by `id` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Operations.ListExecutionHistory() -> *management.ListExecutionHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the execution history for all operations.
History is stored for the configured retention period (default 4 weeks).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListExecutionHistoryRequestParams{}
client.Operations.ListExecutionHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**filter:** `*string` 

Filter operation execution history by this query.
Available filters:
- `id[eq]` - Filter by specific operation ID
- `integration_id[eq]` - Filter by specific integration ID
- `account_id[eq]` - Filter by specific account ID
- `execution_id[eq]` - Filter by specific execution ID
- `started_at[gte]` - Filter executions started at or after a specific datetime (RFC3339 format)
- `started_at[lte]` - Filter executions started at or before a specific datetime (RFC3339 format)
- `started_at[gt]` - Filter executions started after a specific datetime
- `started_at[lt]` - Filter executions started before a specific datetime
- `operation_id[eq]` - Filter by operation name (e.g., "assets_query_devices")
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `started_at[desc]`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field id. For example, `started_at[asc]` will sort the results by `started_at` in ascending order.
The ordering defaults to `desc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return (default 100, max 500)
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return execution history starting after this cursor.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## OrganizationWebhooks
<details><summary><code>client.OrganizationWebhooks.List() -> *management.ListOrganizationWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List webhooks for the organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationWebhooksRequest{}
client.OrganizationWebhooks.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Webhook` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Webhook` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrganizationWebhooks.Get(WebhookId) -> *management.GetOrganizationWebhookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Webhook` object matching `{webhookId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OrganizationWebhooks.Get(
        context.TODO(),
        "webhookId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookId:** `management.WebhookId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrganizationWebhooks.Create(request) -> *management.CreateOrganizationWebhookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new webhook for the organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateOrganizationWebhookRequest{
        Filters: []management.WebhookFilter{
            management.WebhookFilterAll,
            management.WebhookFilterAll,
        },
        Url: "url",
        Secret: &management.OrganizationWebhookSecret{
            Value: "value",
        },
    }
client.OrganizationWebhooks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateOrganizationWebhookRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrganizationWebhooks.Update(WebhookId, request) -> *management.UpdateOrganizationWebhookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Webhook` object matching `{webhookId}`. For more information on
Organizations and Webhooks, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.OrganizationWebhook{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Fullname: "fullname",
        Environment: management.EnvironmentTest,
        Filters: []management.WebhookFilter{
            management.WebhookFilterAll,
            management.WebhookFilterAll,
        },
        Url: "url",
        CredentialId: "credential_id",
    }
client.OrganizationWebhooks.Update(
        context.TODO(),
        "webhookId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookId:** `management.WebhookId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateOrganizationWebhookRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrganizationWebhooks.Patch(WebhookId, request) -> *management.PatchOrganizationWebhookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `Webhook` object matching `{webhookId}`. For more information on
Organizations and Webhooks, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.OrganizationWebhooks.Patch(
        context.TODO(),
        "webhookId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookId:** `management.WebhookId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OrganizationWebhooks.Delete(WebhookId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a webhook for the organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OrganizationWebhooks.Delete(
        context.TODO(),
        "webhookId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookId:** `management.WebhookId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization
<details><summary><code>client.Organization.Get() -> *management.GetOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve Organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Update(request) -> *management.UpdateOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update Organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.Organization{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        RefreshTokenId: "refresh_token_id",
        OrganizationType: management.OrganizationTypeRoot,
        Fullname: "fullname",
    }
client.Organization.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `management.UpdateOrganizationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Patch(request) -> *management.PatchOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patch Organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Organization.Patch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Permissionset
<details><summary><code>client.Permissionset.List() -> *management.ListPermissionSetsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `PermissionSets` objects that match the query params.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissionset.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissionset.Get(PermissionsetId) -> *management.GetPermissionSetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `PermissionSet` object matching `{permissionsetId}`. For more information on PermissionSets, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissionset.Get(
        context.TODO(),
        management.PermissionsAdministrator.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**permissionsetId:** `*management.Permissions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Roles
<details><summary><code>client.Roles.List() -> *management.ListRolesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `Roles` objects that match the query params.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListRolesRequest{}
client.Roles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Role` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Role` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Get(RoleId) -> *management.GetRoleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Role` object matching `{roleId}`. For more information on Roles, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Roles.Get(
        context.TODO(),
        "roleId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleId:** `management.RoleId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Create(request) -> *management.CreateRoleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Role` object. For more information on Roles, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateRoleRequest{
        PermissionSet: management.PermissionsAdministrator,
    }
client.Roles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateRoleRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Update(RoleId, request) -> *management.UpdateRoleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Role` object matching `{roleId}`. For more information on Roles, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.RoleDefinition{
        Name: "name",
        CreatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        UpdatedAt: management.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        Id: "id",
        Fullname: "fullname",
        PermissionSet: management.PermissionsAdministrator,
    }
client.Roles.Update(
        context.TODO(),
        "roleId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleId:** `management.RoleId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateRoleRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Patch(RoleId, request) -> *management.PatchRoleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Patches the `Role` object matching `{roleId}`. For more information on Roles, refer to our
[Synqly Overview](https://docs.synqly.com/docs/synqly-overview).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.PatchOperation{
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
        &management.PatchOperation{
            Op: management.PatchOpAdd,
            Path: "path",
        },
    }
client.Roles.Patch(
        context.TODO(),
        "roleId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleId:** `management.RoleId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*management.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Delete(RoleId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the `Role` matching `{roleId}`. Deleting an `Role` also deletea
all `Tokens` and `Credentials` belonging to the `Role`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Roles.Delete(
        context.TODO(),
        "roleId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleId:** `management.RoleId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Status
<details><summary><code>client.Status.List() -> *management.ListStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all matching `Status` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListStatusRequest{}
client.Status.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Status` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Status` objects starting after this `account_id,integration_id`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `account_id,integration_id`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*management.ListStatusOptions` — Expand the status result with the related integration and/or account information.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Status.Get(AccountId, IntegrationId) -> *management.GetStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the integration `Status` object.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Status.Get(
        context.TODO(),
        "accountId",
        "integrationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Status.Reset(AccountId, IntegrationId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resets the integration `Status` object.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Status.Reset(
        context.TODO(),
        "accountId",
        "integrationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Status.ListEvents(AccountId, IntegrationId) -> *management.ListStatusEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns integration `Status` object list of `StatusEvent` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListStatusEventsRequest{}
client.Status.ListEvents(
        context.TODO(),
        "accountId",
        "integrationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `StatusEvent` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `StatusEvent` objects starting after this `created_at`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — The order defaults to created_at[asc] and can changed to descending order by specifying created_at[desc].
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Status.GetTimeseries() -> *management.GetStatusTimeseries</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns organization last hour usage timeseries.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Status.GetTimeseries(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Status.GetIntegrationTimeseries(AccountId, IntegrationId) -> *management.GetIntegrationTimeseries</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns organization last hour usage timeseries.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetIntegrationTimeseriesRequest{}
client.Status.GetIntegrationTimeseries(
        context.TODO(),
        "accountId",
        "integrationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*string` — [hour] provide most recent 24 hour timeseries. default: hour
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SubOrgs
<details><summary><code>client.SubOrgs.List() -> *management.ListOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all organizations
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SubOrgs.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubOrgs.Get(OrganizationId) -> *management.GetOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve an Organization by ID
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SubOrgs.Get(
        context.TODO(),
        "organizationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationId:** `management.OrganizationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubOrgs.Create(request) -> *management.CreateOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateOrganizationRequest{}
client.SubOrgs.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateOrganizationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubOrgs.Delete(OrganizationId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a Organization by ID. Also deletes
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SubOrgs.Delete(
        context.TODO(),
        "organizationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationId:** `management.OrganizationId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tokens
<details><summary><code>client.Tokens.CreateToken(request) -> *management.CreateTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an adhoc organization token restricted to specified resources and permission set.
Tokens can only be reduced in scope, never expanded.
Permissions are inherited from the token used to call this API.
Permissions assigned to the new token will not be persisted, this is not a way to create roles.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateTokenRequest{
        Resources: &management.Resources{},
        PermissionSet: management.PermissionsAdministrator,
    }
client.Tokens.CreateToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateTokenRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.CreateMcpToken(request) -> *management.CreateMcpTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a token for MCP authentication. This token is soley for MCP authentication and cannot
authenticate with any other API.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateMcpTokenRequest{
        Scope: &management.McpTokenScope{
            Management: &management.McpManagementScopeOptions{},
        },
    }
client.Tokens.CreateMcpToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateMcpTokenRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.CreateIntegrationToken(AccountId, IntegrationId, request) -> *management.CreateIntegrationTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an adhoc integration token restricted to a single integration. The token used to call
this API must have the necessary permissions to create tokens and have access to the account
and integration IDs. Permissions may not be escalated, so any operation that the invocation
token does not have access to cannot be granted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateIntegrationTokenRequest{}
client.Tokens.CreateIntegrationToken(
        context.TODO(),
        "accountId",
        "integrationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `management.AccountId` 
    
</dd>
</dl>

<dl>
<dd>

**integrationId:** `management.IntegrationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.CreateIntegrationTokenRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.CreateSynqlyIntegrationsToken(request) -> *management.CreateSynqlyIntegrationsTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a token for managing Synqly-specific integrations. This token can be used with the integration APIs to manage
integrations for Synqly-specific integrations, such as status events exports and async operations. See the
[Synqly Integrations](https://docs.synqly.com/guides/synqly-integrations) documentation for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateSynqlyIntegrationsTokenRequest{}
client.Tokens.CreateSynqlyIntegrationsToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateSynqlyIntegrationsTokenRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Delete(RefreshTokenId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the Refresh Token with id `{id}`. This immediately
invalidates both the primary and secondary token pairs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Delete(
        context.TODO(),
        "refreshTokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refreshTokenId:** `management.TokenId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.List() -> *management.ListTokensResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all `RefreshToken` objects belonging to the Authorization Bearer
token. For more infromation on Tokens, refer to
[Authentication](/api-reference/authentication).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListTokensRequest{}
client.Tokens.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of `Token` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**startAfter:** `*string` — Return `Token` objects starting after this `name`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` 

Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/getting-started/management-api-filtering). Defaults to no filter.
If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Get(RefreshTokenId) -> *management.GetTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `RefreshToken` object matching `{tokenId}`. For more information on
Tokens, refer to
[Authentication](/api-reference/authentication).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Get(
        context.TODO(),
        "refreshTokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refreshTokenId:** `management.TokenId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Reset(OwnerId, RefreshTokenId) -> *management.ResetTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API can be used to reset `Organization` or `Integration` `RefreshTokens`.
Resets the specified `RefreshToken` and expiration time, removes the secondary, and resets access and refresh tokens for the
`RefreshToken` object matching `{ownerId}/{refreshTokenId}` where `ownerId` is an `organizationId` or `integrationId`.
An `Organization` token with `administrator` permissions can be used to perform this operation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Reset(
        context.TODO(),
        "ownerId",
        "refreshTokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**refreshTokenId:** `management.TokenId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Rotate(OwnerId, RefreshTokenId) -> *management.RotateTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API can be used to rotate an `Organization` or `Integration` `RefreshTokens`.
Rotate deletes the existing `Secondary` TokenPair, moves the `Primary` TokenPair to the `Secondary` TokenPair and creates a new `Primary` TokenPair for the
`RefreshToken` object matching `{ownerId}/{refreshTokenId}` where `ownerId` is an `organizationId` or `integrationId`.
An `Organization` token with `administrator` permissions can be used to perform this operation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Rotate(
        context.TODO(),
        "ownerId",
        "refreshTokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `management.Id` 
    
</dd>
</dl>

<dl>
<dd>

**refreshTokenId:** `management.TokenId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Refresh(RefreshTokenId) -> *management.RefreshTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new primary `TokenPair` object, setting the secondary `TokenPair`
to the previous primary value. Call `/v1/removeSecondaryToken` to remove
this secondary backup once the new primary `TokenPair` has been deployed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Refresh(
        context.TODO(),
        "refreshTokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refreshTokenId:** `management.TokenId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.RemoveSecondary(RefreshTokenId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the secondary `TokenPair` for the `RefreshToken` object
matching `{refreshTokenId}`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.RemoveSecondary(
        context.TODO(),
        "refreshTokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refreshTokenId:** `management.TokenId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

