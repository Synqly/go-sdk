// This file was auto-generated by Fern from our API Definition.

package management

import (
	fmt "fmt"
)

type GetIntegrationTimeseriesRequest struct {
	// [minute|hour] provide most recent 60 minute or 24 hour timeseries. default: minute
	Interval *string `json:"-"`
}

type ListStatusRequest struct {
	// Number of `Status` objects to return in this page. Defaults to 100.
	Limit *int `json:"-"`
	// Return `Status` objects starting after this `account_id,integration_id`.
	StartAfter *string `json:"-"`
	// Select a field to order the results by. Defaults to `account_id,integration_id`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
	// The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
	// ordering is applied in the order the fields are specified.
	Order []*string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
	// Expand the status result with the related integration and/or account information.
	Expand []*ListStatusOptions `json:"-"`
}

type ListStatusEventsRequest struct {
	// Number of `StatusEvent` objects to return in this page. Defaults to 100.
	Limit *int `json:"-"`
	// Return `StatusEvent` objects starting after this `created_at`.
	StartAfter *string `json:"-"`
	// The order defaults to created_at[asc] and can changed to descending order by specifying created_at[desc].
	Order *string `json:"-"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-"`
}

// Get integration timeseries
type GetIntegrationTimeseries struct {
	Result GetIntegrationTimeseriesResult `json:"result,omitempty"`
}

type GetStatusResponse struct {
	Result *Status `json:"result,omitempty"`
}

// Get status last hour timeseries
type GetStatusTimeseries struct {
	Result *GetStatusTimeseriesResult `json:"result,omitempty"`
}

type ListStatusEventsResponse struct {
	Result []*StatusEvent `json:"result,omitempty"`
}

type ListStatusOptions string

const (
	ListStatusOptionsAccount     ListStatusOptions = "account"
	ListStatusOptionsIntegration ListStatusOptions = "integration"
	ListStatusOptionsAll         ListStatusOptions = "all"
)

func NewListStatusOptionsFromString(s string) (ListStatusOptions, error) {
	switch s {
	case "account":
		return ListStatusOptionsAccount, nil
	case "integration":
		return ListStatusOptionsIntegration, nil
	case "all":
		return ListStatusOptionsAll, nil
	}
	var t ListStatusOptions
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListStatusOptions) Ptr() *ListStatusOptions {
	return &l
}

type ListStatusResponse struct {
	Result []*Status `json:"result,omitempty"`
}
