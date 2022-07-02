// Package reports provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package reports

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"gopkg.me/selling-partner-api-sdk-makrorof/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetReportDocument request
	GetReportDocument(ctx context.Context, reportDocumentId string) (*http.Response, error)

	// GetReports request
	GetReports(ctx context.Context, params *GetReportsParams) (*http.Response, error)

	// CreateReport request  with any body
	CreateReportWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	CreateReport(ctx context.Context, body CreateReportJSONRequestBody) (*http.Response, error)

	// CancelReport request
	CancelReport(ctx context.Context, reportId string) (*http.Response, error)

	// GetReport request
	GetReport(ctx context.Context, reportId string) (*http.Response, error)

	// GetReportSchedules request
	GetReportSchedules(ctx context.Context, params *GetReportSchedulesParams) (*http.Response, error)

	// CreateReportSchedule request  with any body
	CreateReportScheduleWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	CreateReportSchedule(ctx context.Context, body CreateReportScheduleJSONRequestBody) (*http.Response, error)

	// CancelReportSchedule request
	CancelReportSchedule(ctx context.Context, reportScheduleId string) (*http.Response, error)

	// GetReportSchedule request
	GetReportSchedule(ctx context.Context, reportScheduleId string) (*http.Response, error)
}

func (c *Client) GetReportDocument(ctx context.Context, reportDocumentId string) (*http.Response, error) {
	req, err := NewGetReportDocumentRequest(c.Endpoint, reportDocumentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetReports(ctx context.Context, params *GetReportsParams) (*http.Response, error) {
	req, err := NewGetReportsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateReportWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateReportRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateReport(ctx context.Context, body CreateReportJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateReportRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CancelReport(ctx context.Context, reportId string) (*http.Response, error) {
	req, err := NewCancelReportRequest(c.Endpoint, reportId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetReport(ctx context.Context, reportId string) (*http.Response, error) {
	req, err := NewGetReportRequest(c.Endpoint, reportId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetReportSchedules(ctx context.Context, params *GetReportSchedulesParams) (*http.Response, error) {
	req, err := NewGetReportSchedulesRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateReportScheduleWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateReportScheduleRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateReportSchedule(ctx context.Context, body CreateReportScheduleJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateReportScheduleRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CancelReportSchedule(ctx context.Context, reportScheduleId string) (*http.Response, error) {
	req, err := NewCancelReportScheduleRequest(c.Endpoint, reportScheduleId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetReportSchedule(ctx context.Context, reportScheduleId string) (*http.Response, error) {
	req, err := NewGetReportScheduleRequest(c.Endpoint, reportScheduleId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetReportDocumentRequest generates requests for GetReportDocument
func NewGetReportDocumentRequest(endpoint string, reportDocumentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "reportDocumentId", reportDocumentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/documents/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetReportsRequest generates requests for GetReports
func NewGetReportsRequest(endpoint string, params *GetReportsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/reports")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.ReportTypes != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "reportTypes", *params.ReportTypes); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ProcessingStatuses != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "processingStatuses", *params.ProcessingStatuses); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.MarketplaceIds != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", *params.MarketplaceIds); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageSize != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "pageSize", *params.PageSize); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedSince != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdSince", *params.CreatedSince); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedUntil != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdUntil", *params.CreatedUntil); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "nextToken", *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateReportRequest calls the generic CreateReport builder with application/json body
func NewCreateReportRequest(endpoint string, body CreateReportJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateReportRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewCreateReportRequestWithBody generates requests for CreateReport with any type of body
func NewCreateReportRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/reports")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCancelReportRequest generates requests for CancelReport
func NewCancelReportRequest(endpoint string, reportId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "reportId", reportId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/reports/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetReportRequest generates requests for GetReport
func NewGetReportRequest(endpoint string, reportId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "reportId", reportId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/reports/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetReportSchedulesRequest generates requests for GetReportSchedules
func NewGetReportSchedulesRequest(endpoint string, params *GetReportSchedulesParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/schedules")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "reportTypes", params.ReportTypes); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateReportScheduleRequest calls the generic CreateReportSchedule builder with application/json body
func NewCreateReportScheduleRequest(endpoint string, body CreateReportScheduleJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateReportScheduleRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewCreateReportScheduleRequestWithBody generates requests for CreateReportSchedule with any type of body
func NewCreateReportScheduleRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/schedules")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCancelReportScheduleRequest generates requests for CancelReportSchedule
func NewCancelReportScheduleRequest(endpoint string, reportScheduleId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "reportScheduleId", reportScheduleId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/schedules/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetReportScheduleRequest generates requests for GetReportSchedule
func NewGetReportScheduleRequest(endpoint string, reportScheduleId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "reportScheduleId", reportScheduleId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/reports/2020-09-04/schedules/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetReportDocument request
	GetReportDocumentWithResponse(ctx context.Context, reportDocumentId string) (*GetReportDocumentResp, error)

	// GetReports request
	GetReportsWithResponse(ctx context.Context, params *GetReportsParams) (*GetReportsResp, error)

	// CreateReport request  with any body
	CreateReportWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateReportResp, error)

	CreateReportWithResponse(ctx context.Context, body CreateReportJSONRequestBody) (*CreateReportResp, error)

	// CancelReport request
	CancelReportWithResponse(ctx context.Context, reportId string) (*CancelReportResp, error)

	// GetReport request
	GetReportWithResponse(ctx context.Context, reportId string) (*GetReportResp, error)

	// GetReportSchedules request
	GetReportSchedulesWithResponse(ctx context.Context, params *GetReportSchedulesParams) (*GetReportSchedulesResp, error)

	// CreateReportSchedule request  with any body
	CreateReportScheduleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateReportScheduleResp, error)

	CreateReportScheduleWithResponse(ctx context.Context, body CreateReportScheduleJSONRequestBody) (*CreateReportScheduleResp, error)

	// CancelReportSchedule request
	CancelReportScheduleWithResponse(ctx context.Context, reportScheduleId string) (*CancelReportScheduleResp, error)

	// GetReportSchedule request
	GetReportScheduleWithResponse(ctx context.Context, reportScheduleId string) (*GetReportScheduleResp, error)
}

type GetReportDocumentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetReportDocumentResponse
}

// Status returns HTTPResponse.Status
func (r GetReportDocumentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetReportDocumentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetReportsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetReportsResponse
}

// Status returns HTTPResponse.Status
func (r GetReportsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetReportsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateReportResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CreateReportResponse
}

// Status returns HTTPResponse.Status
func (r CreateReportResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateReportResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelReportResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelReportResponse
}

// Status returns HTTPResponse.Status
func (r CancelReportResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelReportResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetReportResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetReportResponse
}

// Status returns HTTPResponse.Status
func (r GetReportResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetReportResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetReportSchedulesResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetReportSchedulesResponse
}

// Status returns HTTPResponse.Status
func (r GetReportSchedulesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetReportSchedulesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateReportScheduleResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CreateReportScheduleResponse
}

// Status returns HTTPResponse.Status
func (r CreateReportScheduleResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateReportScheduleResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelReportScheduleResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelReportScheduleResponse
}

// Status returns HTTPResponse.Status
func (r CancelReportScheduleResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelReportScheduleResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetReportScheduleResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetReportScheduleResponse
}

// Status returns HTTPResponse.Status
func (r GetReportScheduleResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetReportScheduleResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetReportDocumentWithResponse request returning *GetReportDocumentResponse
func (c *ClientWithResponses) GetReportDocumentWithResponse(ctx context.Context, reportDocumentId string) (*GetReportDocumentResp, error) {
	rsp, err := c.GetReportDocument(ctx, reportDocumentId)
	if err != nil {
		return nil, err
	}
	return ParseGetReportDocumentResp(rsp)
}

// GetReportsWithResponse request returning *GetReportsResponse
func (c *ClientWithResponses) GetReportsWithResponse(ctx context.Context, params *GetReportsParams) (*GetReportsResp, error) {
	rsp, err := c.GetReports(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetReportsResp(rsp)
}

// CreateReportWithBodyWithResponse request with arbitrary body returning *CreateReportResponse
func (c *ClientWithResponses) CreateReportWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateReportResp, error) {
	rsp, err := c.CreateReportWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateReportResp(rsp)
}

func (c *ClientWithResponses) CreateReportWithResponse(ctx context.Context, body CreateReportJSONRequestBody) (*CreateReportResp, error) {
	rsp, err := c.CreateReport(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateReportResp(rsp)
}

// CancelReportWithResponse request returning *CancelReportResponse
func (c *ClientWithResponses) CancelReportWithResponse(ctx context.Context, reportId string) (*CancelReportResp, error) {
	rsp, err := c.CancelReport(ctx, reportId)
	if err != nil {
		return nil, err
	}
	return ParseCancelReportResp(rsp)
}

// GetReportWithResponse request returning *GetReportResponse
func (c *ClientWithResponses) GetReportWithResponse(ctx context.Context, reportId string) (*GetReportResp, error) {
	rsp, err := c.GetReport(ctx, reportId)
	if err != nil {
		return nil, err
	}
	return ParseGetReportResp(rsp)
}

// GetReportSchedulesWithResponse request returning *GetReportSchedulesResponse
func (c *ClientWithResponses) GetReportSchedulesWithResponse(ctx context.Context, params *GetReportSchedulesParams) (*GetReportSchedulesResp, error) {
	rsp, err := c.GetReportSchedules(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetReportSchedulesResp(rsp)
}

// CreateReportScheduleWithBodyWithResponse request with arbitrary body returning *CreateReportScheduleResponse
func (c *ClientWithResponses) CreateReportScheduleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateReportScheduleResp, error) {
	rsp, err := c.CreateReportScheduleWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateReportScheduleResp(rsp)
}

func (c *ClientWithResponses) CreateReportScheduleWithResponse(ctx context.Context, body CreateReportScheduleJSONRequestBody) (*CreateReportScheduleResp, error) {
	rsp, err := c.CreateReportSchedule(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateReportScheduleResp(rsp)
}

// CancelReportScheduleWithResponse request returning *CancelReportScheduleResponse
func (c *ClientWithResponses) CancelReportScheduleWithResponse(ctx context.Context, reportScheduleId string) (*CancelReportScheduleResp, error) {
	rsp, err := c.CancelReportSchedule(ctx, reportScheduleId)
	if err != nil {
		return nil, err
	}
	return ParseCancelReportScheduleResp(rsp)
}

// GetReportScheduleWithResponse request returning *GetReportScheduleResponse
func (c *ClientWithResponses) GetReportScheduleWithResponse(ctx context.Context, reportScheduleId string) (*GetReportScheduleResp, error) {
	rsp, err := c.GetReportSchedule(ctx, reportScheduleId)
	if err != nil {
		return nil, err
	}
	return ParseGetReportScheduleResp(rsp)
}

// ParseGetReportDocumentResp parses an HTTP response from a GetReportDocumentWithResponse call
func ParseGetReportDocumentResp(rsp *http.Response) (*GetReportDocumentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetReportDocumentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetReportDocumentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetReportsResp parses an HTTP response from a GetReportsWithResponse call
func ParseGetReportsResp(rsp *http.Response) (*GetReportsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetReportsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetReportsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCreateReportResp parses an HTTP response from a CreateReportWithResponse call
func ParseCreateReportResp(rsp *http.Response) (*CreateReportResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateReportResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CreateReportResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelReportResp parses an HTTP response from a CancelReportWithResponse call
func ParseCancelReportResp(rsp *http.Response) (*CancelReportResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelReportResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelReportResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetReportResp parses an HTTP response from a GetReportWithResponse call
func ParseGetReportResp(rsp *http.Response) (*GetReportResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetReportResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetReportResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetReportSchedulesResp parses an HTTP response from a GetReportSchedulesWithResponse call
func ParseGetReportSchedulesResp(rsp *http.Response) (*GetReportSchedulesResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetReportSchedulesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetReportSchedulesResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCreateReportScheduleResp parses an HTTP response from a CreateReportScheduleWithResponse call
func ParseCreateReportScheduleResp(rsp *http.Response) (*CreateReportScheduleResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateReportScheduleResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CreateReportScheduleResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelReportScheduleResp parses an HTTP response from a CancelReportScheduleWithResponse call
func ParseCancelReportScheduleResp(rsp *http.Response) (*CancelReportScheduleResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelReportScheduleResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelReportScheduleResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetReportScheduleResp parses an HTTP response from a GetReportScheduleWithResponse call
func ParseGetReportScheduleResp(rsp *http.Response) (*GetReportScheduleResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetReportScheduleResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetReportScheduleResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
