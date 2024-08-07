// This file was auto-generated by Fern from our API Definition.

package engine

import (
	stix "github.com/synqly/go-sdk/client/engine/stix"
)

type DeleteIocsRequest struct {
	// list of ids to delete
	Ids *string `json:"-"`
}

type QueryAlertsRequest struct {
	// Number of threats to return. Defaults to 50.
	Limit *int `json:"-"`
	// Start search from cursor position.
	Cursor *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order.
	// The ordering defaults to `asc` if not specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
	// Include the raw data from the EDR in the response. Defaults to `false`.
	IncludeRawData *bool `json:"-"`
}

type QueryApplicationsRequest struct {
	// Number of applications to return. Defaults to 50.
	Limit *int `json:"-"`
	// Start search from cursor position.
	Cursor *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order.
	// The ordering defaults to `asc` if not specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
}

type QueryEndpointsRequest struct {
	// Number of endpoint assets to return. Defaults to 50.
	Limit *int `json:"-"`
	// Start search from cursor position.
	Cursor *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order.
	// The ordering defaults to `asc` if not specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
}

type QueryIocsRequest struct {
	// Number of threats to return. Defaults to 50.
	Limit *int `json:"-"`
	// Start search from cursor position.
	Cursor *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order.
	// The ordering defaults to `asc` if not specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
	// Include the raw data from the EDR in the response. Defaults to `false`.
	IncludeRawData *bool `json:"-"`
}

type QueryThreatsRequest struct {
	// Number of threats to return. Defaults to 50.
	Limit *int `json:"-"`
	// Start search from cursor position.
	Cursor *string `json:"-"`
	// Select a field to order the results by. Defaults to `name`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[asc]` will sort the results by `name` in ascending order.
	// The ordering defaults to `asc` if not specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
	// Include the raw data from the EDR in the response. Defaults to `false`.
	IncludeRawData *bool `json:"-"`
}

type CreateIocsRequest struct {
	// The list of iocs to create
	Indicators []*stix.Indicator `json:"indicators,omitempty"`
}

type CreateIocsResponse struct {
	// A list of the indicators that were created
	Result []*stix.Indicator `json:"result,omitempty"`
}

type DeleteIocsResponse struct {
	// A list of ids of the iocs that were deleted.
	Result []string `json:"result,omitempty"`
}

type NetworkQuarantineRequest struct {
	// The connection state (Connect or Disconnect) to enforce for the provided endpoint IDs.
	State ConnectionState `json:"state,omitempty"`
	// The list of endpoint IDs to enforce the connection state on.
	EndpointIds []string `json:"endpoint_ids,omitempty"`
}

type QueryAlertsResponse struct {
	// List of alerts that match the query.
	Result []*Event `json:"result,omitempty"`
	// Cursor to use to retrieve the next page of results.
	Cursor string `json:"cursor"`
}

type QueryApplicationsResponse struct {
	// List of applications that match the query.
	Result []Application `json:"result,omitempty"`
	// Cursor to use to retrieve the next page of results.
	Cursor string `json:"cursor"`
}

type QueryEndpointsResponse struct {
	// List of endpoint assets that match the query.
	Result []Device `json:"result,omitempty"`
	// Cursor to use to retrieve the next page of results.
	Cursor string `json:"cursor"`
}

type QueryIocsResponse struct {
	// List of iocs that match the query.
	Result []*stix.Indicator `json:"result,omitempty"`
	// Cursor to use to retrieve the next page of results.
	Cursor string `json:"cursor"`
}

type QueryThreatsResponse struct {
	// List of threats that match the query.
	Result []ThreatEvent `json:"result,omitempty"`
	// Cursor to use to retrieve the next page of results.
	Cursor string `json:"cursor"`
}
