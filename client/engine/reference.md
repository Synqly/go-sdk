# Reference
## Appsec
<details><summary><code>client.Appsec.QueryApplications() -> *engine.AppSecQueryApplicationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of applications matching the query from a the token-linked application security integration.
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
request := &engine.AppSecQueryApplicationsRequest{}
client.Appsec.QueryApplications(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of applications to return. Defaults to 100 with a maximum of 5000. If a provider has a maximum limit lower than 5000, the provider's maximum limit will be used instead.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/appsec/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Appsec.QueryApplicationFindings(ApplicationId) -> *engine.AppSecQueryApplicationFindingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of an application's findings matching `{applictionId}` and the query from a the token-linked application security integration.
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
request := &engine.AppSecQueryApplicationFindingsRequest{}
client.Appsec.QueryApplicationFindings(
        context.TODO(),
        "applicationId",
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

**applicationId:** `engine.ApplicationId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of findings to return. Defaults to 100 with a maximum of 5000. If a provider has a maximum limit lower than 5000, the provider's maximum limit will be used instead.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/appsec/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Appsec.QueryFindings() -> *engine.AppSecQueryFindingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of each findings details combined with the application details for all applications in the token-linked application security integration. This API may perform multiple provider API calls per executation so can be slower to respond.
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
request := &engine.AppSecQueryFindingsRequest{}
client.Appsec.QueryFindings(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of findings to return. Defaults to 100 with a maximum of 5000. If a provider has a maximum limit lower than 5000, the provider's maximum limit will be used instead.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/appsec/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Appsec.GetApplicationFindingDetails(ApplicationId, FindingId) -> *engine.AppSecGetApplicationFindingDetailsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the details of the finding matching `{findingId}` where the finding belongs to the application matching `{applicationId}` from the token-linked application security integration.
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
request := &engine.AppSecGetApplicationFindingDetailsRequest{}
client.Appsec.GetApplicationFindingDetails(
        context.TODO(),
        "applicationId",
        "findingId",
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

**applicationId:** `engine.ApplicationId` 
    
</dd>
</dl>

<dl>
<dd>

**findingId:** `engine.FindingId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Assets
<details><summary><code>client.Assets.QueryDevices() -> *engine.QueryDevicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query devices from an asset inventory system
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
request := &engine.QueryDevicesRequest{}
client.Assets.QueryDevices(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of finding reports to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/assets/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `time[asc]` will sort the results by `time` in ascending order.
The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.CreateAsset(request) -> *engine.CreateDeviceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Device` object in the token-linked Integration.
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
request := &engine.CreateDeviceRequestInput{
        Body: &engine.CreateDeviceRequest{
            Device: &inventoryinfo.InventoryInfo{
                ActivityId: 1,
                CategoryUid: 1,
                ClassUid: 1,
                Device: &v130.Device{
                    TypeId: 1,
                },
                Metadata: &v130.Metadata{
                    Product: &v130.Product{},
                    Version: "version",
                },
                SeverityId: 1,
                Time: 1,
                TypeUid: 1,
            },
        },
    }
client.Assets.CreateAsset(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateDeviceRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.CreateDevices(request) -> *engine.CreateDevicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Bulk create devices in an device inventory system
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
request := &engine.CreateDevicesRequestInput{
        Body: &engine.CreateDevicesRequest{
            Devices: []engine.Device{
                &inventoryinfo.InventoryInfo{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Device: &v130.Device{
                        TypeId: 1,
                    },
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                },
                &inventoryinfo.InventoryInfo{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Device: &v130.Device{
                        TypeId: 1,
                    },
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                },
            },
            SourceName: "source_name",
        },
    }
client.Assets.CreateDevices(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateDevicesRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.GetLabels() -> *engine.GetLabelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get labels from an asset inventory system
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
request := &engine.GetLabelsRequest{}
client.Assets.GetLabels(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.QuerySoftware() -> *engine.QuerySoftwareInventoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query software inventory records from an asset management system
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
request := &engine.QuerySoftwareInventoryRequest{}
client.Assets.QuerySoftware(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of software inventory records to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to the [Assets Filtering Guide](https://docs.synqly.com/guides/connectors/assets/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. Defaults to `package.name`. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `package.name[asc]` will sort the results by package name in ascending order.
The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.CreateSoftware(request) -> *engine.CreateSoftwareInventoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a software inventory record in the token-linked Integration.
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
request := &engine.CreateSoftwareInventoryRequestInput{
        Body: &engine.CreateSoftwareInventoryRequest{
            SoftwareInventory: &softwareinventoryinfo.SoftwareInfo{
                ActivityId: 1,
                CategoryUid: 1,
                ClassUid: 1,
                Device: &v180.Device{
                    TypeId: 1,
                },
                Metadata: &v180.Metadata{
                    Product: &v180.Product{},
                    Version: "version",
                },
                SeverityId: 1,
                Time: 1,
                TypeUid: 1,
            },
        },
    }
client.Assets.CreateSoftware(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateSoftwareInventoryRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Chat
<details><summary><code>client.Chat.QueryUsers() -> *engine.ChatQueryUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of users in the connected workspace or tenant. Used to discover whose conversations can be queried.
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
request := &engine.QueryChatUsers{}
client.Chat.QueryUsers(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of users to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/chat/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Chat.QueryUserConversations(UserId) -> *engine.ChatQueryConversationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of conversations scoped to a specific user — DMs, group chats, or (for AI providers) synthesised sessions.
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
request := &engine.QueryUserConversations{}
client.Chat.QueryUserConversations(
        context.TODO(),
        "userId",
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

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of conversations to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Chat.QueryUserConversationMembers(UserId, ConversationId) -> *engine.ChatQueryConversationMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the members of a conversation scoped to a specific user. Required for providers where conversations are only addressable via the user (e.g. Teams DMs, Copilot AI sessions).
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
request := &engine.QueryUserConversationMembers{}
client.Chat.QueryUserConversationMembers(
        context.TODO(),
        "userId",
        "conversationId",
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

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**conversationId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of members to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Chat.QueryUserConversationMessages(UserId, ConversationId) -> *engine.ChatQueryMessagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns messages from a conversation scoped to a specific user. Required for providers where conversations are only addressable via the user (e.g. Teams DMs, Copilot AI sessions).
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
request := &engine.QueryUserConversationMessages{}
client.Chat.QueryUserConversationMessages(
        context.TODO(),
        "userId",
        "conversationId",
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

**userId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**conversationId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of messages to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. Supported filter fields vary by provider. Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Chat.QueryConversations() -> *engine.ChatQueryConversationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of conversations (channels or workspaces) not scoped to a specific user. Providers that can only enumerate conversations per-user return NotImplemented; use query_user_conversations instead.
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
request := &engine.QueryConversations{}
client.Chat.QueryConversations(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of conversations to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Chat.QueryConversationMembers(ConversationId) -> *engine.ChatQueryConversationMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the members of a specific conversation. Use the user-scoped variant (query_user_conversation_members) for chats that only exist in a user's mailbox (e.g. Teams 1:1 chats).
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
request := &engine.QueryConversationMembers{}
client.Chat.QueryConversationMembers(
        context.TODO(),
        "conversationId",
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

**conversationId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of members to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Chat.QueryConversationMessages(ConversationId) -> *engine.ChatQueryMessagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns messages from a specific conversation. The conversationId must be obtained from query_conversations or query_user_conversations. Use the user-scoped variant (query_user_conversation_messages) for chats that only exist in a user's mailbox.
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
request := &engine.QueryConversationMessages{}
client.Chat.QueryConversationMessages(
        context.TODO(),
        "conversationId",
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

**conversationId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of messages to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. Supported filter fields vary by provider. Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Cloudsecurity
<details><summary><code>client.Cloudsecurity.QueryEvents() -> *engine.QueryEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of events that match the query from the cloud security provider.
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
request := &engine.QueryEvents{}
client.Cloudsecurity.QueryEvents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of events to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/cloudsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cloudsecurity.QueryIoms() -> *engine.QueryIomsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of Indicators of Misconfiguration (IOM) findings that match the query from the cloud security provider.
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
request := &engine.QueryIoms{}
client.Cloudsecurity.QueryIoms(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of cloud resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/cloudsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the CloudSecurity in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cloudsecurity.QueryCloudResourceInventory() -> *engine.QueryCloudResourceInventoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of cloud resources that match the query from the cloud security provider.
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
request := &engine.QueryCloudResourceInventory{}
client.Cloudsecurity.QueryCloudResourceInventory(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.  
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of cloud resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/cloudsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the CloudSecurity in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cloudsecurity.QueryComplianceFindings() -> *engine.QueryComplianceFindingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of compliance findings matching the query from the cloud security provider.
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
request := &engine.QueryComplianceFindings{}
client.Cloudsecurity.QueryComplianceFindings(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of compliance findings to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/cloudsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the CloudSecurity in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cloudsecurity.QueryThreats() -> *engine.QueryCloudSecurityThreatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of threat detections that match the query from the cloud security provider.
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
request := &engine.QueryCloudSecurityThreats{}
client.Cloudsecurity.QueryThreats(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of threats to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/cloudsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the CloudSecurity in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Custom
<details><summary><code>client.Custom.Query(Operation) -> *engine.QueryCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queries operation
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
request := &engine.QueryCustomRequest{}
client.Custom.Query(
        context.TODO(),
        engine.OperationIdAppsecGetApplicationFindingDetails.Ptr(),
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

**operation:** `*engine.OperationId` — what `operation` to query
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `operation` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Custom.Get(Operation, Id) -> *engine.GetCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an Topic by ID.
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
request := &engine.GetCustomRequest{}
client.Custom.Get(
        context.TODO(),
        engine.OperationIdAppsecGetApplicationFindingDetails.Ptr(),
        "id",
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

**operation:** `*engine.OperationId` — what `operation` to query
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — ID of the `operation` to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Custom.Patch(Operation, Id, request) -> *engine.GetCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an `operation` by ID.
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
request := &engine.PatchCustomRequestInput{
        Body: []*engine.PatchOperation{
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
        },
    }
client.Custom.Patch(
        context.TODO(),
        engine.OperationIdAppsecGetApplicationFindingDetails.Ptr(),
        "id",
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

**operation:** `*engine.OperationId` — what `operation` to query
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — ID of the `operation` to update.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `engine.PatchCustomRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Custom.Delete(Operation, Id) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the object matching `id` for the operation matching `operation`.
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
request := &engine.DeleteRequest{}
client.Custom.Delete(
        context.TODO(),
        engine.OperationIdAppsecGetApplicationFindingDetails.Ptr(),
        "id",
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

**operation:** `*engine.OperationId` — what `operation` to delete.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — ID of the `operation` to delete.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Custom.Post(Operation, request) -> *engine.CreateCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Writes an `operation` object to the Custom provider configured with the token used for authentication.
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
request := &engine.PostCustomRequest{
        Body: map[string]any{
            "string": map[string]any{
                "key": "value",
            },
        },
    }
client.Custom.Post(
        context.TODO(),
        engine.OperationIdAppsecGetApplicationFindingDetails.Ptr(),
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

**operation:** `*engine.OperationId` — what `operation` to post
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `map[string]any` — The `operation` object to be sent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Custom.PostBatch(Operation, request) -> *engine.CreateBatchCustomResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Writes a batch of `operation` objects to the Custom provider configured with the token used for authentication.
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
request := &engine.PostBatchCustomRequest{
        Body: []map[string]any{
            map[string]any{
                "string": map[string]any{
                    "key": "value",
                },
            },
            map[string]any{
                "string": map[string]any{
                    "key": "value",
                },
            },
        },
    }
client.Custom.PostBatch(
        context.TODO(),
        engine.OperationIdAppsecGetApplicationFindingDetails.Ptr(),
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

**operation:** `*engine.OperationId` — what `operation` to post
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]map[string]any` — The `operation` object to be sent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Edr
<details><summary><code>client.Edr.QueryEndpoints() -> *engine.QueryEndpointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of endpoint assets matching the query from the token-linked EDR source.
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
request := &engine.QueryEndpointsRequest{}
client.Edr.QueryEndpoints(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of endpoint assets to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/edr/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.GetEndpoint(Id) -> *engine.GetEndpointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a single endpoint assets matching the UID from the token-linked EDR source.
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
request := &engine.GetEndpointRequest{}
client.Edr.GetEndpoint(
        context.TODO(),
        "id",
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

**id:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.ExecuteCommand(Uid, request) -> *engine.ExecuteCommandResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Runs a provider-backed command on the endpoint identified by `{uid}` and returns normalized stdout and stderr without exposing provider session details.
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
request := &engine.ExecuteCommandRequestInput{
        Body: &engine.ExecuteCommandRequest{
            Command: "command",
        },
    }
client.Edr.ExecuteCommand(
        context.TODO(),
        "uid",
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

**uid:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.ExecuteCommandRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.RetrieveFile(Uid, request) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a file from the endpoint identified by `{uid}` and returns the provider artifact as a binary file response.
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
request := &engine.RetrieveFileRequest{
        Path: "path",
    }
client.Edr.RetrieveFile(
        context.TODO(),
        "uid",
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

**uid:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.RetrieveFileRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.QueryApplications() -> *engine.QueryApplicationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of applications matching the query from the token-linked EDR source.
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
request := &engine.QueryApplicationsRequest{}
client.Edr.QueryApplications(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of applications to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/edr/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.NetworkQuarantine(request) -> *engine.NetworkQuarantineResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Connect or disconnect one or more endpoints assets to the network, allowing or disallowing connections.
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
request := &engine.NetworkQuarantineRequestInput{
        Body: &engine.NetworkQuarantineRequest{
            State: engine.ConnectionStateConnect,
            EndpointIds: []string{
                "endpoint_ids",
                "endpoint_ids",
            },
        },
    }
client.Edr.NetworkQuarantine(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.NetworkQuarantineRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.QueryThreatevents() -> *engine.QueryThreatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of threats that match the query from the token-linked EDR source.
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
request := &engine.QueryThreatsRequest{}
client.Edr.QueryThreatevents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of threats to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/edr/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.QueryAlerts() -> *engine.QueryAlertsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of alerts that match the query from the token-linked EDR source.
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
request := &engine.QueryAlertsRequest{}
client.Edr.QueryAlerts(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of threats to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/edr/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.QueryIocs() -> *engine.QueryIocsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of iocs that match the query from the token-linked EDR source.
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
request := &engine.QueryIocsRequest{}
client.Edr.QueryIocs(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of threats to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/edr/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.CreateIocs(request) -> *engine.CreateIocsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a list of iocs that match the stix input for the EDR source.
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
request := &engine.CreateIocsRequestInput{
        Body: &engine.CreateIocsRequest{
            Indicators: []*stix.Indicator{
                &stix.Indicator{
                    Id: "id",
                    Created: engine.MustParseDateTime(
                        "2024-01-15T09:30:00Z",
                    ),
                    Modified: engine.MustParseDateTime(
                        "2024-01-15T09:30:00Z",
                    ),
                    ValidFrom: engine.MustParseDateTime(
                        "2024-01-15T09:30:00Z",
                    ),
                },
                &stix.Indicator{
                    Id: "id",
                    Created: engine.MustParseDateTime(
                        "2024-01-15T09:30:00Z",
                    ),
                    Modified: engine.MustParseDateTime(
                        "2024-01-15T09:30:00Z",
                    ),
                    ValidFrom: engine.MustParseDateTime(
                        "2024-01-15T09:30:00Z",
                    ),
                },
            },
        },
    }
client.Edr.CreateIocs(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateIocsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.DeleteIocs() -> *engine.DeleteIocsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a list of iocs that match the input of ids in the query param
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
request := &engine.DeleteIocsRequest{}
client.Edr.DeleteIocs(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — list of ids to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.QueryPostureScore() -> *engine.QueryPostureScoreResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the posture score of the endpoint assets that match the query from the token-linked EDR source.
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
request := &engine.QueryPostureScoreRequest{}
client.Edr.QueryPostureScore(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of scores for endpoints to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/edr/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.QueryEdrEvents() -> *engine.QueryEdrEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of EDR events that match the query from the token-linked EDR source. 
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
request := &engine.QueryEdrEventsRequest{}
client.Edr.QueryEdrEvents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**passthroughParam:** `*string` — Provider-specific query to pass through to the EDR. This is useful for advanced queries that require additional filtering.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter events. Multiple filters can be provided.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the EDR in the response. Defaults to `false`.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of events to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` 

Select a field to order the results by. To control the direction of the sorting, append
`[asc]` or `[desc]` to the field name. For example, `timestamp[asc]` will sort the results by `timestamp` in ascending order.
The ordering defaults to `asc` if not specified.            
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.GetThreatNotes(ThreatId) -> *engine.GetThreatNotesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of notes for a threat.
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
request := &engine.GetThreatNotesRequest{}
client.Edr.GetThreatNotes(
        context.TODO(),
        "threatId",
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

**threatId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of notes to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor from a previous response to retrieve the next page of notes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Edr.CreateThreatNote(ThreatId, request) -> *engine.CreateThreatNoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a note for a threat.
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
request := &engine.CreateThreatNoteRequestInput{
        Body: &engine.CreateThreatNoteRequest{
            Note: &noteactivity.NoteActivity{
                ActivityId: 1,
                CategoryUid: 1,
                ClassUid: 1,
                FindingInfo: &v180.FindingInfo{
                    Uid: "uid",
                },
                Metadata: &v180.Metadata{
                    Product: &v180.Product{},
                    Version: "version",
                },
                Note: &v180.Annotation{},
                SeverityId: 1,
                Time: 1,
                TypeUid: 1,
            },
        },
    }
client.Edr.CreateThreatNote(
        context.TODO(),
        "threatId",
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

**threatId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateThreatNoteRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Emailsecurity
<details><summary><code>client.Emailsecurity.QueryThreats() -> *engine.EmailSecurityQueryThreatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of threats matching the query from the token-linked Email Security provider. Defaults to the last 30 days of threats. This can be overridden by using the `time` filter. Note that some providers may have a maximum time range limit. A threat is an automated detection that was deemed to be a threat by the Email Security provider.
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
request := &engine.EmailSecurityQueryThreatsRequest{}
client.Emailsecurity.QueryThreats(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of threats to return. Defaults to 1000 with a maximum of 5000. If a provider has a maximum limit lower than 5000, the provider's maximum limit will be used instead.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/emailsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Emailsecurity.GetThreatDetails(ThreatId) -> *engine.EmailSecurityGetThreatDetailsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the details of the threat matching `{threatId}` from the token-linked Email Security provider. If a provider allows for the gathering of more detailed information about a threat, the response will include the additional information. Otherwise, the response will only include the basic information about the threat returned by the query_threats endpoint.
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
request := &engine.EmailSecurityGetThreatDetailsRequest{}
client.Emailsecurity.GetThreatDetails(
        context.TODO(),
        "threatId",
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

**threatId:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Emailsecurity.QueryEmailEvents() -> *engine.EmailSecurityQueryEmailEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of email events matching the query from the token-linked Email Security provider. Defaults to the last 30 days of email events. This can be overridden by using the `time` filter. Note that some providers may have a maximum time range limit.
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
request := &engine.EmailSecurityQueryEmailEventsRequest{}
client.Emailsecurity.QueryEmailEvents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of email events to return. Defaults to 1000 with a maximum of 5000. If a provider has a maximum limit lower than 5000, the provider's maximum limit will be used instead.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/emailsecurity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Endpointmanagement
<details><summary><code>client.Endpointmanagement.QueryDevices() -> *engine.QueryEndpointManagementDevicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of managed devices matching the query from the token-linked endpoint management source.
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
request := &engine.QueryEndpointManagementDevicesRequest{}
client.Endpointmanagement.QueryDevices(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of devices to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `device.name[asc]` will sort the results by `device.name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/endpointmanagement/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.GetDevice(Id) -> *engine.GetEndpointManagementDeviceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a single managed device matching the ID from the token-linked endpoint management source.
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
request := &engine.GetDeviceRequest{}
client.Endpointmanagement.GetDevice(
        context.TODO(),
        "id",
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

**id:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.QueryComplianceFindings() -> *engine.QueryDeviceComplianceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of device compliance findings matching the query from the token-linked endpoint management source.
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
request := &engine.QueryDeviceComplianceRequest{}
client.Endpointmanagement.QueryComplianceFindings(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of compliance findings to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/endpointmanagement/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.RemediateDevice(request) -> *engine.RemediateDeviceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Performs a remediation action on a managed device from the token-linked endpoint management source.
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
request := &engine.RemediationRequestInput{
        Body: &engine.RemediationRequest{},
    }
client.Endpointmanagement.RemediateDevice(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.RemediationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.UpdateDevice(request) -> *engine.DeviceActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Triggers the device to check in and enforce all assigned policies, profiles, and pending updates.
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
request := &engine.DeviceActionRequest{
        DeviceId: "device_id",
    }
client.Endpointmanagement.UpdateDevice(
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

**request:** `*engine.DeviceActionRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.LockDevice(request) -> *engine.DeviceActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remotely locks the device screen.
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
request := &engine.DeviceActionRequest{
        DeviceId: "device_id",
    }
client.Endpointmanagement.LockDevice(
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

**request:** `*engine.DeviceActionRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.RestartDevice(request) -> *engine.DeviceActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remotely reboots the device.
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
request := &engine.DeviceActionRequest{
        DeviceId: "device_id",
    }
client.Endpointmanagement.RestartDevice(
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

**request:** `*engine.DeviceActionRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.WipeDevice(request) -> *engine.DeviceActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Factory resets or erases all data from the device.
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
request := &engine.DeviceActionRequest{
        DeviceId: "device_id",
    }
client.Endpointmanagement.WipeDevice(
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

**request:** `*engine.DeviceActionRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpointmanagement.QueryDeviceApplications(Id) -> *engine.QueryDeviceApplicationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of applications installed on a managed device from the token-linked endpoint management source.
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
request := &engine.QueryDeviceApplicationsRequest{}
client.Endpointmanagement.QueryDeviceApplications(
        context.TODO(),
        "id",
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

**id:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of applications to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `product.name[asc]` will sort the results by `product.name` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/endpointmanagement/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hooks
<details><summary><code>client.Hooks.Proxy(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Proxy webhook messages from webhook providers to webhook recievers. For exact webhook implementations please refer to providers e.g. Ticketing. This is just an API call used in that context, not a standalone implementation.
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
request := &engine.PostProxyHook{
        Token: "token",
        Body: map[string]any{
            "key": "value",
        },
    }
client.Hooks.Proxy(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**token:** `string` — Optional: if you can't use the HTTP Authorization Bearer, specify integration access token here.
    
</dd>
</dl>

<dl>
<dd>

**request:** `any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Identity
<details><summary><code>client.Identity.QueryAuditLog() -> *engine.QueryIdentityAuditLogResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Event` objects from the token-linked audit log.
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
request := &engine.QueryIdentityAuditLogRequest{}
client.Identity.QueryAuditLog(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of events to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `time[asc]` will sort the results by `time` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/identity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.QueryRiskEvents() -> *engine.QueryIdentityRiskEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns identity threat / risk events (for example Microsoft Entra Identity Protection risk detections for users), normalized to OCSF.
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
request := &engine.QueryIdentityRiskEventsRequest{}
client.Identity.QueryRiskEvents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of events to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `time[asc]` will sort the results by `time` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/identity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the identity provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.QueryRiskyUsers() -> *engine.QueryIdentityRiskyUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns rolled-up risky user records (for example Microsoft Entra Identity Protection riskyUsers), each normalized to an OCSF Entity Management Read event.
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
request := &engine.QueryIdentityRiskyUsersRequest{}
client.Identity.QueryRiskyUsers(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of users to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `time[asc]` will sort the results by `time` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/identity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the identity provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.QueryUsers() -> *engine.QueryUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `User` objects wrapped in the OCSF Entity Management event of type Read from the token-linked identity provider.
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
request := &engine.QueryUserRequest{}
client.Identity.QueryUsers(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of users to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `uid`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `email_addr[asc]` will sort the results by `email_addr` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/identity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.GetUser(UserId) -> *engine.GetUserResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `User` object wrapped in an OCSF Entity Management event of type Read from the token-linked identity provider. Depending 
on the providers offerings, this may include additional user information, such as the user's current groups, roles, and Identity Protection risk state (when the provider exposes it and the user has a risk record).
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
request := &engine.GetUserRequest{}
client.Identity.GetUser(
        context.TODO(),
        "userId",
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

**userId:** `engine.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.QueryGroups() -> *engine.QueryGroupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Group` objects wrapped in the OCSF Entity Management event of type Read from the token-linked identity provider.
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
request := &engine.QueryGroupRequest{}
client.Identity.QueryGroups(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of users to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `uid`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `email_addr[asc]` will sort the results by `email_addr` in ascending order. The ordering defaults to `asc` if not specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/identity/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.GetGroup(GroupId) -> *engine.GetGroupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Group` object wrapped in an OCSF Entity Management event of type Read from the token-linked identity provider. Depending 
on the providers offerings, this may include additional group information, such as the roles assigned.
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
request := &engine.GetGroupRequest{}
client.Identity.GetGroup(
        context.TODO(),
        "groupId",
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

**groupId:** `engine.GroupId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.GetGroupMembers(GroupId) -> *engine.GetGroupMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns list of `User` objects wrapped in an OCSF Entity Management event of type Read from the token-linked identity provider that are members in the group referenced by ID.
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
request := &engine.GetGroupMembersRequest{}
client.Identity.GetGroupMembers(
        context.TODO(),
        "groupId",
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

**groupId:** `engine.GroupId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of users to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.EnableUser(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reenables a disabled user in the identity system based on user ID.
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
request := &engine.EnableUserRequest{}
client.Identity.EnableUser(
        context.TODO(),
        "userId",
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

**userId:** `engine.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.DisableUser(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disables a user in the identity system based on user ID.
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
request := &engine.DisableUserRequest{}
client.Identity.DisableUser(
        context.TODO(),
        "userId",
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

**userId:** `engine.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.ForceUserPasswordReset(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Forces a user to reset their password before they can log in again.
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
request := &engine.ForceUserPasswordResetRequest{}
client.Identity.ForceUserPasswordReset(
        context.TODO(),
        "userId",
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

**userId:** `engine.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.ExpireAllUserSessions(UserId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Logs a user out of all current sessions so they must log in again.
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
request := &engine.ExpireAllUserSessionsRequest{}
client.Identity.ExpireAllUserSessions(
        context.TODO(),
        "userId",
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

**userId:** `engine.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.GetUserPicture(UserId) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the profile picture for a user as binary image data. The Content-Type header indicates the image format (e.g., image/jpeg, image/png).
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
client.Identity.GetUserPicture(
        context.TODO(),
        "userId",
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

**userId:** `engine.UserId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Incidentresponse
<details><summary><code>client.Incidentresponse.QueryEscalationPolicies() -> *engine.IncidentResponseQueryEscalationPoliciesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of incident response escalation policies.
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
request := &engine.IncidentResponseQueryEscalationPoliciesRequest{}
client.Incidentresponse.QueryEscalationPolicies(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of escalation policies to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Incidentresponse.QueryEscalationPolicyUsersOnCall(EscalationPolicyId) -> *engine.IncidentResponseQueryEscalationPolicyUsersOnCallResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of users that are currently on call for an incident response escalation policy.
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
request := &engine.IncidentResponseQueryEscalationPolicyUsersOnCallRequest{}
client.Incidentresponse.QueryEscalationPolicyUsersOnCall(
        context.TODO(),
        "escalationPolicyId",
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

**escalationPolicyId:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## IntegrationWebhooks
<details><summary><code>client.IntegrationWebhooks.CreateWebhook(request) -> *engine.CreateIntegrationWebHookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a WebHook for the token-linked Integration.
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
request := &engine.CreateIntegrationWebHookRequest{
        Project: "project",
    }
client.IntegrationWebhooks.CreateWebhook(
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

**request:** `*engine.CreateIntegrationWebHookRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IntegrationWebhooks.DeleteWebhook() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the WebHook matching the token-linked Integration.
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
client.IntegrationWebhooks.DeleteWebhook(
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

<details><summary><code>client.IntegrationWebhooks.ListWebhooks() -> *engine.ListIntegrationWebHooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all WebHooks from the token-linked Integration.
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
client.IntegrationWebhooks.ListWebhooks(
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

## Notifications
<details><summary><code>client.Notifications.GetMessage(NotificationId) -> *engine.GetNotificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Notification` object matching `{notificationId}` from the token-linked `Integration`.
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
request := &engine.GetMessageRequest{}
client.Notifications.GetMessage(
        context.TODO(),
        "notificationId",
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

**notificationId:** `engine.NotificationId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.CreateMessage(request) -> *engine.CreateNotificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Notification` object in the token-linked `Integration`.
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
request := &engine.CreateNotificationRequestInput{
        Body: &engine.CreateNotificationRequest{
            Summary: "summary",
        },
    }
client.Notifications.CreateMessage(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateNotificationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.ClearMessage(NotificationId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resolves a `Notification` object in the token-linked `Integration`.
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
request := &engine.ClearMessageRequest{}
client.Notifications.ClearMessage(
        context.TODO(),
        "notificationId",
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

**notificationId:** `engine.NotificationId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Operations
<details><summary><code>client.Operations.Get(OperationId) -> *engine.GetOperationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the `Asynchronous Operation` object matching `{operationId}`.
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
client.Operations.Get(
        context.TODO(),
        "operationId",
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

**operationId:** `engine.AsyncOperationRequestId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Operations.Create(request) -> *engine.CreateOperationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Asynchronous Operation` object.
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
request := &engine.CreateOperationRequest{
        Operation: "operation",
        Input: &engine.OperationInput{},
    }
client.Operations.Create(
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

**request:** `*engine.CreateOperationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Operations.Cancel(OperationId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels the `Asynchronous Operation` matching `{operationId}`.
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
client.Operations.Cancel(
        context.TODO(),
        "operationId",
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

**operationId:** `engine.AsyncOperationRequestId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Siem
<details><summary><code>client.Siem.QueryInvestigations() -> *engine.QueryInvestigationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queries investigations
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
request := &engine.QueryInvestigationsRequest{}
client.Siem.QueryInvestigations(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Investigation` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/siem/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.GetInvestigation(Id) -> *engine.GetInvestigationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an investigation by ID.
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
request := &engine.GetInvestigationRequest{}
client.Siem.GetInvestigation(
        context.TODO(),
        "id",
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

**id:** `string` — ID of the investigation to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.PatchInvestigation(Id, request) -> *engine.GetInvestigationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an investigation by ID.
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
request := &engine.PatchInvestigationRequestInput{
        Body: []*engine.PatchOperation{
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
        },
    }
client.Siem.PatchInvestigation(
        context.TODO(),
        "id",
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

**id:** `string` — ID of the investigation to update.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `engine.PatchInvestigationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.GetEvidence(Id) -> *engine.GetEvidenceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the evidence for an investigation.
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
request := &engine.GetInvestigationEvidenceRequest{}
client.Siem.GetEvidence(
        context.TODO(),
        "id",
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

**id:** `string` — ID of the investigation to retrieve evidence for.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.QueryLogProviders() -> *engine.QueryLogProvidersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queries available log providers in the source SIEM
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
request := &engine.QueryLogProvidersRequest{}
client.Siem.QueryLogProviders(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of log provider objects to return in this page. Defaults to 100.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.PostEvents(request) -> *engine.CreateSiemEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Writes a batch of `Event` objects to the SIEM configured with the token used for authentication.
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
request := &engine.PostSiemEventRequest{
        Body: []*engine.Event{
            &engine.Event{
                AccountChange: &accountchange.AccountChange{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                    User: &v130.User{},
                },
            },
            &engine.Event{
                AccountChange: &accountchange.AccountChange{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                    User: &v130.User{},
                },
            },
        },
    }
client.Siem.PostEvents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*engine.Event` — The `Event` object to be sent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.QueryEvents() -> *engine.QuerySiemEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queries events from the SIEM configured with the token used for authentication.
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
request := &engine.QuerySiemEventsRequest{}
client.Siem.QueryEvents(
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

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Account` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order. The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/siem/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**passthroughParam:** `*string` 

Provider-specific query to pass through to the SIEM. This is useful for advanced queries that are not
supported by the API. The keys and values are provider-specific. For example, to perform a specific
query in Rapid7 IDR, you can use the `query: "{advanced query}"` key-value pair.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. This is useful for debugging and troubleshooting. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.QueryAlerts() -> *engine.QuerySiemAlertsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queries alerts from the SIEM configured with the token used for authentication.
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
request := &engine.QuerySiemAlertsRequest{}
client.Siem.QueryAlerts(
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

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Account` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order. The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/siem/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. This is useful for debugging and troubleshooting. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Siem.GetAlert(Id) -> *engine.GetAlertResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an alert by ID.
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
request := &engine.GetAlertRequest{}
client.Siem.GetAlert(
        context.TODO(),
        "id",
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

**id:** `string` — ID of the alert to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the SIEM in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sink
<details><summary><code>client.Sink.PostEvents(request) -> *engine.CreateSinkEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Writes a batch of `Event` objects to the Sink configured with the token used for authentication.
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
request := &engine.PostSinkEventRequest{
        Body: []*engine.Event{
            &engine.Event{
                AccountChange: &accountchange.AccountChange{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                    User: &v130.User{},
                },
            },
            &engine.Event{
                AccountChange: &accountchange.AccountChange{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                    User: &v130.User{},
                },
            },
        },
    }
client.Sink.PostEvents(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — Updates the path or queue name for the sink. If not provided, the default path or queue name is used. When provided, the path or queue name is appended to the default path or queue name.
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*engine.Event` — The `Event` object to be sent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Storage
<details><summary><code>client.Storage.ListFiles(Path) -> *engine.ListStorageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of contents from the token-linked `Integration`.
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
request := &engine.ListStorageRequest{}
client.Storage.ListFiles(
        context.TODO(),
        "path",
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

**path:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to fetch the next set of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return. Default is 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Storage.DownloadFile(Path) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Downloads a file from the provided `{path}` in the token-linked `Integration`.
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
client.Storage.DownloadFile(
        context.TODO(),
        "path",
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

**path:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Storage.DeleteFile(Path) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a file from the provided `{path}` in the token-linked `Integration`.
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
request := &engine.DeleteFileRequest{}
client.Storage.DeleteFile(
        context.TODO(),
        "path",
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

**path:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing
<details><summary><code>client.Ticketing.ListRemoteFields() -> *engine.ListRemoteFieldsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all remote fields for all Projects in a ticketing integration. The response will include a list of fields for each issue type in the ticketing provider.
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
request := &engine.ListRemoteFieldsRequest{}
client.Ticketing.ListRemoteFields(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.ListProjects() -> *engine.ListProjectsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Projects` from the token-linked `Integration`. Tickets must be created and retrieved within the context of a specific Project.
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
request := &engine.ListProjectsRequest{}
client.Ticketing.ListProjects(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Projects` objects to return in this page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.QueryTickets() -> *engine.QueryTicketsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Ticket` objects from the token-linked `Integration`.
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
request := &engine.QueryTicketsRequest{}
client.Ticketing.QueryTickets(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor to use to retrieve the next page of results.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of `Account` objects to return in this page. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — Select a field to order the results by. Defaults to `time`. To control the direction of the sorting, append `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order. The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the ordering is applied in the order the fields are specified.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/ticketing/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.GetTicket(TicketId) -> *engine.GetTicketResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Ticket` object matching `{ticketId}` from the token-linked `Integration`.
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
request := &engine.GetTicketRequest{}
client.Ticketing.GetTicket(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.CreateTicket(request) -> *engine.CreateTicketResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Ticket` object in the token-linked Integration.
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
request := &engine.CreateTicketRequestInput{
        Body: &engine.CreateTicketRequest{
            Name: "name",
        },
    }
client.Ticketing.CreateTicket(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateTicketRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.PatchTicket(TicketId, request) -> *engine.PatchTicketResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `Ticket` object matching `{ticketId}` in the token-linked `Integration`.
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
request := &engine.PatchTicketRequestInput{
        Body: []*engine.PatchOperation{
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
        },
    }
client.Ticketing.PatchTicket(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*engine.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.CreateAttachment(TicketId, request) -> *engine.CreateAttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

[beta: currently supported by Jira] Creates an `Attachment` for the ticket with id `{ticketId}` in the token-linked `Integration`.
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
request := &engine.CreateAttachmentRequestInput{
        Body: &engine.CreateAttachmentRequest{
            FileName: "file_name",
            Content: []byte("SGVsbG8gd29ybGQh"),
        },
    }
client.Ticketing.CreateAttachment(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateAttachmentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.ListAttachmentsMetadata(TicketId) -> *engine.ListAttachmentsMetadataResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

[beta: currently supported by Jira] Returns metadata for all Attachments for a `Ticket` object matching `{ticketId}` from the token-linked `Integration`.
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
request := &engine.ListAttachmentsMetadataRequest{}
client.Ticketing.ListAttachmentsMetadata(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.DownloadAttachment(TicketId, AttachmentId) -> *engine.DownloadAttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

[beta: currently supported by Jira] Downloads the Attachment object matching {attachmentId} for the Ticket matching {tickedId} from the token-linked Integration.
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
request := &engine.DownloadAttachmentRequest{}
client.Ticketing.DownloadAttachment(
        context.TODO(),
        "ticketId",
        "attachmentId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**attachmentId:** `engine.AttachmentId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.DeleteAttachment(TicketId, AttachmentId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

[beta: currently supported by Jira] Deletes the Attachment object matching {attachmentId} for the Ticket matching {tickedId} from the token-linked Integration.
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
request := &engine.DeleteAttachmentRequest{}
client.Ticketing.DeleteAttachment(
        context.TODO(),
        "ticketId",
        "attachmentId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**attachmentId:** `engine.AttachmentId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.ListComments(TicketId) -> *engine.ListCommentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all comments for the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.ListCommentsRequest{}
client.Ticketing.ListComments(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.CreateComment(TicketId, request) -> *engine.CreateCommentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a comment on the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.CreateCommentRequestInput{
        Body: &engine.CreateCommentRequest{
            Content: "content",
        },
    }
client.Ticketing.CreateComment(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateCommentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.DeleteComment(TicketId, CommentId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the comment matching {commentId} form the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.DeleteCommentRequest{}
client.Ticketing.DeleteComment(
        context.TODO(),
        "ticketId",
        "commentId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**commentId:** `engine.CommentId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.CreateNote(TicketId, request) -> *engine.CreateNoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a note on the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.CreateNoteRequestInput{
        Body: &engine.CreateNoteRequest{
            Content: "content",
            Title: "title",
        },
    }
client.Ticketing.CreateNote(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateNoteRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.DeleteNote(TicketId, NoteId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the note matching {noteId} form the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.DeleteNoteRequest{}
client.Ticketing.DeleteNote(
        context.TODO(),
        "ticketId",
        "noteId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**noteId:** `engine.NoteId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.ListNotes(TicketId) -> *engine.ListNotesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists all notes for the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.ListNotesRequest{}
client.Ticketing.ListNotes(
        context.TODO(),
        "ticketId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.PatchNote(TicketId, NoteId, request) -> *engine.PatchNoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a note matching {noteId} title and/or content on the ticket matching {ticketId} from the token-linked Integration.
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
request := &engine.PatchNoteRequestInput{
        Body: []*engine.PatchOperation{
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
            &engine.PatchOperation{
                Op: engine.PatchOpAdd,
                Path: "path",
            },
        },
    }
client.Ticketing.PatchNote(
        context.TODO(),
        "ticketId",
        "noteId",
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

**ticketId:** `engine.TicketId` 
    
</dd>
</dl>

<dl>
<dd>

**noteId:** `engine.NoteId` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*engine.PatchOperation` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.QueryEscalationPolicies() -> *engine.QueryEscalationPoliciesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of escalation policies.
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
request := &engine.QueryEscalationPoliciesRequest{}
client.Ticketing.QueryEscalationPolicies(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of escalation policies to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/ticketing/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.ListOnCall(EscalationPolicyId) -> *engine.ListOnCallResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of all on-call agents for an escalation policy.
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
request := &engine.ListOnCallRequest{}
client.Ticketing.ListOnCall(
        context.TODO(),
        "escalationPolicyId",
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

**escalationPolicyId:** `engine.Id` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Vulnerabilities
<details><summary><code>client.Vulnerabilities.QueryFindings() -> *engine.QueryFindingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query vulnerability findings
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
request := &engine.QueryFindingsRequest{}
client.Vulnerabilities.QueryFindings(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of finding reports to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/vulnerabilities/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.CreateFindings(request) -> *engine.CreateFindingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create findings (bulk) in a vulnerability scanning system
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
request := &engine.CreateFindingsRequestInput{
        Body: &engine.CreateFindingsRequest{
            Findings: []engine.SecurityFinding{
                &securityfinding.SecurityFinding{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Finding: &v130.Finding{
                        Title: "title",
                        Uid: "uid",
                    },
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    StateId: 1,
                    Time: 1,
                    TypeUid: 1,
                },
                &securityfinding.SecurityFinding{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Finding: &v130.Finding{
                        Title: "title",
                        Uid: "uid",
                    },
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    StateId: 1,
                    Time: 1,
                    TypeUid: 1,
                },
            },
        },
    }
client.Vulnerabilities.CreateFindings(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateFindingsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.UpdateFinding(FindingId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

update a finding in a vulnerability scanning system
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
request := &engine.UpdateFindingRequestInput{
        Body: &engine.UpdateFindingRequest{
            SeverityId: 1,
        },
    }
client.Vulnerabilities.UpdateFinding(
        context.TODO(),
        "findingId",
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

**findingId:** `string` — Uid of the Finding (URL encoded). This will be `finding.uid` in the OCSF model.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.UpdateFindingRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.QueryAssets() -> *engine.QueryAssetsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query assets in a vulnerability scanning system
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
request := &engine.QueryAssetsRequest{}
client.Vulnerabilities.QueryAssets(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of assets to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/vulnerabilities/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.CreateAsset(request) -> *engine.CreateAssetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create assets in a vulnerability scanning system
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
request := &engine.CreateAssetRequestInput{
        Body: &engine.CreateAssetRequest{
            Asset: &inventoryinfo.InventoryInfo{
                ActivityId: 1,
                CategoryUid: 1,
                ClassUid: 1,
                Device: &v130.Device{
                    TypeId: 1,
                },
                Metadata: &v130.Metadata{
                    Product: &v130.Product{},
                    Version: "version",
                },
                SeverityId: 1,
                Time: 1,
                TypeUid: 1,
            },
            SourceName: "source_name",
        },
    }
client.Vulnerabilities.CreateAsset(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateAssetRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.UpdateAsset(AssetId, request) -> *engine.Asset</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

update an asset in a vulnerability scanning system
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
request := &engine.UpdateAssetRequestInput{
        Body: &engine.CreateAssetRequest{
            Asset: &inventoryinfo.InventoryInfo{
                ActivityId: 1,
                CategoryUid: 1,
                ClassUid: 1,
                Device: &v130.Device{
                    TypeId: 1,
                },
                Metadata: &v130.Metadata{
                    Product: &v130.Product{},
                    Version: "version",
                },
                SeverityId: 1,
                Time: 1,
                TypeUid: 1,
            },
            SourceName: "source_name",
        },
    }
client.Vulnerabilities.UpdateAsset(
        context.TODO(),
        "assetId",
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

**assetId:** `string` — Uid of the Asset. This will be `devices.uid` in the OCSF model.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.CreateAssetRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.QueryScans() -> *engine.QueryScansResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query scans in a vulnerability scanning system
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
request := &engine.QueryScansRequest{}
client.Vulnerabilities.QueryScans(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of scans to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/vulnerabilities/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.GetScanActivity(ScanId) -> *engine.GetScanActivityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deprecated: use GET /scans/{scan_id}/activities instead, which returns the full execution history for a scan definition.
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
request := &engine.VulnerabilitiesGetScanActivityRequest{}
client.Vulnerabilities.GetScanActivity(
        context.TODO(),
        "scan_id",
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

**scanId:** `string` — ID of the scan to get activity for (provider-specific; often scan definition or template id).
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.GetScanActivities(ScanId) -> *engine.GetScanActivitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get scan execution history for a configured scan. Returns one OCSF Scan Activity per execution (e.g. each Tenable.sc scan result for the given scan definition id).
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
request := &engine.VulnerabilitiesGetScanActivitiesRequest{}
client.Vulnerabilities.GetScanActivities(
        context.TODO(),
        "scan_id",
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

**scanId:** `string` — ID of the scan definition whose execution history is requested.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.UploadScan(request) -> *engine.UploadScanResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload a scan in a vulnerability scanning system
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
request := &engine.UploadScanRequestInput{
        Body: &engine.UploadScanRequest{
            Assets: []engine.Asset{
                &inventoryinfo.InventoryInfo{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Device: &v130.Device{
                        TypeId: 1,
                    },
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                },
                &inventoryinfo.InventoryInfo{
                    ActivityId: 1,
                    CategoryUid: 1,
                    ClassUid: 1,
                    Device: &v130.Device{
                        TypeId: 1,
                    },
                    Metadata: &v130.Metadata{
                        Product: &v130.Product{},
                        Version: "version",
                    },
                    SeverityId: 1,
                    Time: 1,
                    TypeUid: 1,
                },
            },
        },
    }
client.Vulnerabilities.UploadScan(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*engine.UploadScanRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.GetScanStatus(ScanId) -> *engine.GetScanStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the status of a upload scan
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
request := &engine.GetScanStatusRequest{}
client.Vulnerabilities.GetScanStatus(
        context.TODO(),
        "scan_id",
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

**scanId:** `string` — ID of the scan to get the status of.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.GetLabels() -> *engine.GetLabelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get labels from an asset inventory system
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
request := &engine.VulnerabilitiesGetLabelsRequest{}
client.Vulnerabilities.GetLabels(
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

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Vulnerabilities.QueryScanFindings(ScanId) -> *engine.VulnerabilitiesQueryScanFindingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query findings for a scan
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
request := &engine.VulnerabilitiesQueryScanFindingsRequest{}
client.Vulnerabilities.QueryScanFindings(
        context.TODO(),
        "scanId",
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

**scanId:** `string` — ID of the scan to query findings for. Use the `query_scans` API to find available scan IDs.
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*string` — Add metadata to the response by invoking meta functions. Documentation for [meta functions](https://docs.synqly.com/api-reference/meta-functions) is available. Not all meta functions are available at every endpoint.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of findings to return. Defaults to 50 with a maximum of 5000. If a provider has a maximum limit lower than 5000, the provider's maximum limit will be used instead.
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Filter results by this query. For more information on filtering, refer to our [Filtering Guide](https://docs.synqly.com/guides/connectors/vulnerabilities/query-filters). Defaults to no filter. If used more than once, the queries are ANDed together.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Start search from cursor position.
    
</dd>
</dl>

<dl>
<dd>

**includeRawData:** `*bool` — Include the raw data from the provider in the response. Defaults to `false`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

