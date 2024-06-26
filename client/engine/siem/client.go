// This file was auto-generated by Fern from our API Definition.

package siem

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	engine "github.com/synqly/go-sdk/client/engine"
	core "github.com/synqly/go-sdk/client/engine/core"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

// Queries investigations
func (c *Client) QueryInvestigations(ctx context.Context, request *engine.QueryInvestigationsRequest) (*engine.QueryInvestigationResponse, error) {
	baseURL := "https://api.synqly.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/siem/investigations"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	for _, value := range request.Order {
		queryParams.Add("order", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Filter {
		queryParams.Add("filter", fmt.Sprintf("%v", *value))
	}
	if request.IncludeRawData != nil {
		queryParams.Add("include_raw_data", fmt.Sprintf("%v", *request.IncludeRawData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(engine.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(engine.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(engine.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(engine.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(engine.MethodNotAllowedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 409:
			value := new(engine.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 415:
			value := new(engine.UnsupportedMediaTypeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(engine.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(engine.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(engine.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(engine.BadGatewayError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(engine.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(engine.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *engine.QueryInvestigationResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Retrieves an investigation by ID.
//
// ID of the investigation to retrieve.
func (c *Client) GetInvestigation(ctx context.Context, id string, request *engine.GetInvestigationRequest) (*engine.GetInvestigationResponse, error) {
	baseURL := "https://api.synqly.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/siem/investigations/%v", id)

	queryParams := make(url.Values)
	if request.IncludeRawData != nil {
		queryParams.Add("include_raw_data", fmt.Sprintf("%v", *request.IncludeRawData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(engine.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(engine.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(engine.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(engine.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(engine.MethodNotAllowedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 409:
			value := new(engine.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 415:
			value := new(engine.UnsupportedMediaTypeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(engine.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(engine.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(engine.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(engine.BadGatewayError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(engine.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(engine.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *engine.GetInvestigationResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Updates an investigation by ID.
//
// ID of the investigation to update.
func (c *Client) PatchInvestigation(ctx context.Context, id string, request engine.PatchInvestigationRequest) (*engine.GetInvestigationResponse, error) {
	baseURL := "https://api.synqly.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/siem/investigations/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(engine.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(engine.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(engine.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(engine.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(engine.MethodNotAllowedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 409:
			value := new(engine.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 415:
			value := new(engine.UnsupportedMediaTypeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(engine.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(engine.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(engine.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(engine.BadGatewayError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(engine.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(engine.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *engine.GetInvestigationResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Retrieves the evidence for an investigation.
//
// ID of the investigation to retrieve evidence for.
func (c *Client) GetEvidence(ctx context.Context, id string, request *engine.GetInvestigationEvidenceRequest) (*engine.GetEvidenceResponse, error) {
	baseURL := "https://api.synqly.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/siem/investigations/%v/evidence", id)

	queryParams := make(url.Values)
	if request.IncludeRawData != nil {
		queryParams.Add("include_raw_data", fmt.Sprintf("%v", *request.IncludeRawData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(engine.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(engine.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(engine.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(engine.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(engine.MethodNotAllowedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 409:
			value := new(engine.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 415:
			value := new(engine.UnsupportedMediaTypeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(engine.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(engine.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(engine.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(engine.BadGatewayError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(engine.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(engine.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *engine.GetEvidenceResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Writes a batch of `Event` objects to the SIEM configured with the token used for authentication.
func (c *Client) PostEvents(ctx context.Context, request []*engine.Event) error {
	baseURL := "https://api.synqly.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/siem/events"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(engine.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(engine.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(engine.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(engine.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(engine.MethodNotAllowedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 409:
			value := new(engine.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 415:
			value := new(engine.UnsupportedMediaTypeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(engine.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(engine.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(engine.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(engine.BadGatewayError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(engine.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(engine.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		nil,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

// Queries events from the SIEM configured with the token used for authentication.
func (c *Client) QueryEvents(ctx context.Context, request *engine.QuerySiemEventsRequest) (*engine.QuerySiemEventsResponse, error) {
	baseURL := "https://api.synqly.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/siem/events"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	for _, value := range request.Order {
		queryParams.Add("order", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Filter {
		queryParams.Add("filter", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Meta {
		queryParams.Add("meta", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.PassthroughParam {
		queryParams.Add("passthrough-param", fmt.Sprintf("%v", *value))
	}
	if request.IncludeRawData != nil {
		queryParams.Add("include_raw_data", fmt.Sprintf("%v", *request.IncludeRawData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(engine.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(engine.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(engine.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(engine.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(engine.MethodNotAllowedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 409:
			value := new(engine.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 415:
			value := new(engine.UnsupportedMediaTypeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(engine.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(engine.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(engine.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(engine.BadGatewayError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(engine.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(engine.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *engine.QuerySiemEventsResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
