// This file was auto-generated by Fern from our API Definition.

package management

type ListOrganizationWebhooksRequest struct {
	// Number of `Webhook` objects to return in this page. Defaults to 100.
	Limit *int `json:"-"`
	// Return `Webhook` objects starting after this `name`.
	StartAfter *string `json:"-"`
	// Return `Webhook` objects ending before this `name`.
	EndBefore *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
	// The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
	// ordering is applied in the order the fields are specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
}

type CreateOrganizationWebhookRequest struct {
	// Human friendly slug for this webhook
	Name *string `json:"name,omitempty"`
	// Fullname for this webhook
	Fullname *string `json:"fullname,omitempty"`
	// Environment that the webhook is configured for. Only events for accounts associated with this environment will trigger the webhook.
	Environment *Environment `json:"environment,omitempty"`
	// Specifies which Webhooks to send.
	Filters []WebhookFilter `json:"filters,omitempty"`
	// URL that webhooks will be sent to.
	Url string `json:"url"`
	// Secret used for signing webhooks. This value is used to verify the authenticity of the webhook payload.
	Secret *OrganizationWebhookSecret `json:"secret,omitempty"`
}

type CreateOrganizationWebhookResponse struct {
	// The created webhook.
	Result *OrganizationWebhook `json:"result,omitempty"`
}

type GetOrganizationWebhookResponse struct {
	Result *OrganizationWebhook `json:"result,omitempty"`
}

type ListOrganizationWebhooksResponse struct {
	// List of webhooks for the organization.
	Result []*OrganizationWebhook `json:"result,omitempty"`
}

type PatchOrganizationWebhookResponse struct {
	Result *OrganizationWebhook `json:"result,omitempty"`
}

type RefreshOrganizationWebhookRequest struct {
	// Secret used for signing webhooks. This value is used to verify the authenticity of the webhook payload.
	Secret *OrganizationWebhookSecret `json:"secret,omitempty"`
}

type UpdateOrganizationWebhookRequest = *OrganizationWebhook

type UpdateOrganizationWebhookResponse struct {
	Result *OrganizationWebhook `json:"result,omitempty"`
}

// Unique identifier for a Webhook
type WebhookId = Id