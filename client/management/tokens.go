// This file was auto-generated by Fern from our API Definition.

package management

type ListTokensRequest struct {
	// Number of `Token` objects to return in this page. Defaults to 100.
	Limit *int `json:"-"`
	// Return `Token` objects starting after this `name`.
	StartAfter *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
	// The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
	// ordering is applied in the order the fields are specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
}

// Unique identifier for this Token
type TokenId = Id

type CreateIntegrationTokenRequest struct {
	// Unique name token. If not provided, defaults to generated newly created refresh token id.
	Name *string `json:"name,omitempty"`
	// Token time-to-live. If not provided, defaults to the TTL of the token used to call this API. Use the format "1h", "1m", "1s" for hours, minutes, and seconds respectively.
	TokenTtl *string `json:"token_ttl,omitempty"`
}

type CreateIntegrationTokenResponse struct {
	Result *Token `json:"result,omitempty"`
}

type CreateTokenRequest struct {
	// Unique name token. If not provided, defaults to generated newly created refresh token id.
	Name *string `json:"name,omitempty"`
	// Limit access to supplied resources
	Resources *Resources `json:"resources,omitempty"`
	// Limit access to supplied permissions
	PermissionSet Permissions `json:"permission_set,omitempty"`
	// Token time-to-live. If not provided, defaults to the TTL of the token used to call this API. Use the format "1h", "1m", "1s" for hours, minutes, and seconds respectively.
	TokenTtl *string `json:"token_ttl,omitempty"`
}

type CreateTokenResponse struct {
	Result *RefreshToken `json:"result,omitempty"`
}

type GetTokenResponse struct {
	Result *RefreshToken `json:"result,omitempty"`
}

type ListTokensResponse struct {
	Result []*RefreshToken `json:"result,omitempty"`
}

type RefreshTokenResponse struct {
	Result *RefreshToken `json:"result,omitempty"`
}

type ResetTokenResponse struct {
	Result *RefreshToken `json:"result,omitempty"`
}
