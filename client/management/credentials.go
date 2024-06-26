// This file was auto-generated by Fern from our API Definition.

package management

import (
	time "time"
)

type ListCredentialsRequest struct {
	// Number of `Credential` objects to return in this page. Defaults to 100.
	Limit *int `json:"-"`
	// Return `Credential` objects starting after this `name`.
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

type CreateCredentialRequest struct {
	// Unique short name for this Credential (lowercase [a-z0-9_-], can be used in URLs). Also used for case insensitive duplicate name detection and default sort order. Defaults to CredentialId if both name and fullname are not specified.
	Name *string `json:"name,omitempty"`
	// Human friendly display name for this Credential, will auto-generate 'name' field (if 'name' is not specified). Defaults to the same value as the 'name' field if not specified.
	Fullname *string `json:"fullname,omitempty"`
	// Credential configuration
	Config *CredentialConfig `json:"config,omitempty"`
	// One of `account` or `integration_point`; defaults to `account` if not specified.
	OwnerType *OwnerType `json:"owner_type,omitempty"`
	// Time when this credential expires and can no longer be used again.
	Expires *time.Time `json:"expires,omitempty"`
}

type CreateCredentialResponse struct {
	Result *CredentialResponse `json:"result,omitempty"`
}

// Unique identifier for this Credential
type CredentialId = Id

type GetCredentialResponse struct {
	Result *CredentialResponse `json:"result,omitempty"`
}

type ListCredentialsResponse struct {
	Result []*CredentialResponse `json:"result,omitempty"`
}

type LookupCredentialResponse struct {
	Result *CredentialResponse `json:"result,omitempty"`
}

type PatchCredentialResponse struct {
	Result *CredentialResponse `json:"result,omitempty"`
}

type UpdateCredentialRequest = *Credential

type UpdateCredentialResponse struct {
	Result *CredentialResponse `json:"result,omitempty"`
}
