// This file was auto-generated by Fern from our API Definition.

package management

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/synqly/go-sdk/client/management/core"
)

type GetIntegrationTimeseriesRequest struct {
	// [hour] provide most recent 24 hour timeseries. default: hour
	Interval *string `json:"-" url:"interval,omitempty"`
}

type ListStatusRequest struct {
	// Number of `Status` objects to return in this page. Defaults to 100.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Return `Status` objects starting after this `account_id,integration_id`.
	StartAfter *string `json:"-" url:"start_after,omitempty"`
	// Select a field to order the results by. Defaults to `account_id,integration_id`. To control the direction of the sorting, append
	// `[asc]` or `[desc]` to the field name. For example, `name[desc]` will sort the results by `name` in descending order.
	// The ordering defaults to `asc` if not specified. May be used multiple times to order by multiple fields, and the
	// ordering is applied in the order the fields are specified.
	Order []*string `json:"-" url:"order,omitempty"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-" url:"filter,omitempty"`
	// Expand the status result with the related integration and/or account information.
	Expand []*ListStatusOptions `json:"-" url:"expand,omitempty"`
}

type ListStatusEventsRequest struct {
	// Number of `StatusEvent` objects to return in this page. Defaults to 100.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Return `StatusEvent` objects starting after this `created_at`.
	StartAfter *string `json:"-" url:"start_after,omitempty"`
	// The order defaults to created_at[asc] and can changed to descending order by specifying created_at[desc].
	Order *string `json:"-" url:"order,omitempty"`
	// Filter results by this query. For more information on filtering, refer to our Filtering Guide. Defaults to no filter.
	// If used more than once, the queries are ANDed together.
	Filter []*string `json:"-" url:"filter,omitempty"`
}

// Get integration timeseries
type GetIntegrationTimeseries struct {
	Result GetIntegrationTimeseriesResult `json:"result" url:"result"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetIntegrationTimeseries) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetIntegrationTimeseries) UnmarshalJSON(data []byte) error {
	type unmarshaler GetIntegrationTimeseries
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetIntegrationTimeseries(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = nil
	return nil
}

func (g *GetIntegrationTimeseries) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GetStatusResponse struct {
	Result *Status `json:"result" url:"result"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetStatusResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetStatusResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetStatusResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetStatusResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = nil
	return nil
}

func (g *GetStatusResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Get status last hour timeseries
type GetStatusTimeseries struct {
	Result *GetStatusTimeseriesResult `json:"result" url:"result"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetStatusTimeseries) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetStatusTimeseries) UnmarshalJSON(data []byte) error {
	type unmarshaler GetStatusTimeseries
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetStatusTimeseries(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = nil
	return nil
}

func (g *GetStatusTimeseries) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type ListStatusEventsResponse struct {
	Result []*StatusEvent `json:"result" url:"result"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListStatusEventsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListStatusEventsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListStatusEventsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListStatusEventsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = nil
	return nil
}

func (l *ListStatusEventsResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListStatusOptions string

const (
	ListStatusOptionsAccount      ListStatusOptions = "account"
	ListStatusOptionsAccounts     ListStatusOptions = "accounts"
	ListStatusOptionsIntegration  ListStatusOptions = "integration"
	ListStatusOptionsIntegrations ListStatusOptions = "integrations"
	ListStatusOptionsAll          ListStatusOptions = "all"
)

func NewListStatusOptionsFromString(s string) (ListStatusOptions, error) {
	switch s {
	case "account":
		return ListStatusOptionsAccount, nil
	case "accounts":
		return ListStatusOptionsAccounts, nil
	case "integration":
		return ListStatusOptionsIntegration, nil
	case "integrations":
		return ListStatusOptionsIntegrations, nil
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
	Result []*Status `json:"result" url:"result"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListStatusResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListStatusResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListStatusResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListStatusResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = nil
	return nil
}

func (l *ListStatusResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}
