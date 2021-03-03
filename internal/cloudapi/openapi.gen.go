// Package cloudapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package cloudapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	Ec2    AWSUploadRequestOptionsEc2 `json:"ec2"`
	Region string                     `json:"region"`
	S3     AWSUploadRequestOptionsS3  `json:"s3"`
}

// AWSUploadRequestOptionsEc2 defines model for AWSUploadRequestOptionsEc2.
type AWSUploadRequestOptionsEc2 struct {
	AccessKeyId       string    `json:"access_key_id"`
	SecretAccessKey   string    `json:"secret_access_key"`
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
	SnapshotName      *string   `json:"snapshot_name,omitempty"`
}

// AWSUploadRequestOptionsS3 defines model for AWSUploadRequestOptionsS3.
type AWSUploadRequestOptionsS3 struct {
	AccessKeyId     string `json:"access_key_id"`
	Bucket          string `json:"bucket"`
	SecretAccessKey string `json:"secret_access_key"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the uploaded image. It must be unique in the given resource group.
	// If name is omitted from the request, a random one based on a UUID is
	// generated.
	ImageName *string `json:"image_name,omitempty"`

	// Location where the image should be uploaded and registered. This link explain
	// how to list all locations:
	// https://docs.microsoft.com/en-us/cli/azure/account?view=azure-cli-latest#az_account_list_locations'
	Location string `json:"location"`

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResult defines model for ComposeResult.
type ComposeResult struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// Name of an existing STANDARD Storage class Bucket.
	Bucket string `json:"bucket"`

	// The name to use for the imported and shared Compute Node image.
	// The image name must be unique within the GCP project, which is used
	// for the OS image upload and import. If not specified a random
	// 'composer-api-<uuid>' string is used as the image name.
	ImageName *string `json:"image_name,omitempty"`

	// The GCP region where the OS image will be imported to and shared from.
	// The value must be a valid GCP location. See https://cloud.google.com/storage/docs/locations.
	// If not specified, the multi-region location closest to the source
	// (source Storage Bucket location) is chosen automatically.
	Region *string `json:"region,omitempty"`

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	// If not specified, the imported Compute Node image is not shared with any
	// account.
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     string        `json:"image_type"`
	Repositories  []Repository  `json:"repositories"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl    *string `json:"baseurl,omitempty"`
	Metalink   *string `json:"metalink,omitempty"`
	Mirrorlist *string `json:"mirrorlist,omitempty"`
	Rhsm       bool    `json:"rhsm"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws   UploadTypes = "aws"
	UploadTypes_azure UploadTypes = "azure"
	UploadTypes_gcp   UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeJSONBody defines parameters for Compose.
type ComposeJSONBody ComposeRequest

// ComposeRequestBody defines body for Compose for application/json ContentType.
type ComposeJSONRequestBody ComposeJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
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

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Compose request  with any body
	ComposeWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	Compose(ctx context.Context, body ComposeJSONRequestBody) (*http.Response, error)

	// ComposeStatus request
	ComposeStatus(ctx context.Context, id string) (*http.Response, error)

	// GetOpenapiJson request
	GetOpenapiJson(ctx context.Context) (*http.Response, error)

	// GetVersion request
	GetVersion(ctx context.Context) (*http.Response, error)
}

func (c *Client) ComposeWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewComposeRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Compose(ctx context.Context, body ComposeJSONRequestBody) (*http.Response, error) {
	req, err := NewComposeRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ComposeStatus(ctx context.Context, id string) (*http.Response, error) {
	req, err := NewComposeStatusRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetOpenapiJson(ctx context.Context) (*http.Response, error) {
	req, err := NewGetOpenapiJsonRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetVersion(ctx context.Context) (*http.Response, error) {
	req, err := NewGetVersionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewComposeRequest calls the generic Compose builder with application/json body
func NewComposeRequest(server string, body ComposeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewComposeRequestWithBody(server, "application/json", bodyReader)
}

// NewComposeRequestWithBody generates requests for Compose with any type of body
func NewComposeRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/compose")
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

// NewComposeStatusRequest generates requests for ComposeStatus
func NewComposeStatusRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/compose/%s", pathParam0)
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

// NewGetOpenapiJsonRequest generates requests for GetOpenapiJson
func NewGetOpenapiJsonRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/openapi.json")
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

// NewGetVersionRequest generates requests for GetVersion
func NewGetVersionRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/version")
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
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
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
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// Compose request  with any body
	ComposeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ComposeResponse, error)

	ComposeWithResponse(ctx context.Context, body ComposeJSONRequestBody) (*ComposeResponse, error)

	// ComposeStatus request
	ComposeStatusWithResponse(ctx context.Context, id string) (*ComposeStatusResponse, error)

	// GetOpenapiJson request
	GetOpenapiJsonWithResponse(ctx context.Context) (*GetOpenapiJsonResponse, error)

	// GetVersion request
	GetVersionWithResponse(ctx context.Context) (*GetVersionResponse, error)
}

type ComposeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ComposeResult
}

// Status returns HTTPResponse.Status
func (r ComposeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ComposeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ComposeStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ComposeStatus
}

// Status returns HTTPResponse.Status
func (r ComposeStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ComposeStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOpenapiJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetOpenapiJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOpenapiJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetVersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Version
}

// Status returns HTTPResponse.Status
func (r GetVersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ComposeWithBodyWithResponse request with arbitrary body returning *ComposeResponse
func (c *ClientWithResponses) ComposeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ComposeResponse, error) {
	rsp, err := c.ComposeWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseComposeResponse(rsp)
}

func (c *ClientWithResponses) ComposeWithResponse(ctx context.Context, body ComposeJSONRequestBody) (*ComposeResponse, error) {
	rsp, err := c.Compose(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseComposeResponse(rsp)
}

// ComposeStatusWithResponse request returning *ComposeStatusResponse
func (c *ClientWithResponses) ComposeStatusWithResponse(ctx context.Context, id string) (*ComposeStatusResponse, error) {
	rsp, err := c.ComposeStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseComposeStatusResponse(rsp)
}

// GetOpenapiJsonWithResponse request returning *GetOpenapiJsonResponse
func (c *ClientWithResponses) GetOpenapiJsonWithResponse(ctx context.Context) (*GetOpenapiJsonResponse, error) {
	rsp, err := c.GetOpenapiJson(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetOpenapiJsonResponse(rsp)
}

// GetVersionWithResponse request returning *GetVersionResponse
func (c *ClientWithResponses) GetVersionWithResponse(ctx context.Context) (*GetVersionResponse, error) {
	rsp, err := c.GetVersion(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetVersionResponse(rsp)
}

// ParseComposeResponse parses an HTTP response from a ComposeWithResponse call
func ParseComposeResponse(rsp *http.Response) (*ComposeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ComposeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ComposeResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseComposeStatusResponse parses an HTTP response from a ComposeStatusWithResponse call
func ParseComposeStatusResponse(rsp *http.Response) (*ComposeStatusResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ComposeStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ComposeStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetOpenapiJsonResponse parses an HTTP response from a GetOpenapiJsonWithResponse call
func ParseGetOpenapiJsonResponse(rsp *http.Response) (*GetOpenapiJsonResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetOpenapiJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetVersionResponse parses an HTTP response from a GetVersionWithResponse call
func ParseGetVersionResponse(rsp *http.Response) (*GetVersionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetVersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Version
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create compose
	// (POST /compose)
	Compose(w http.ResponseWriter, r *http.Request)
	// The status of a compose
	// (GET /compose/{id})
	ComposeStatus(w http.ResponseWriter, r *http.Request, id string)
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(w http.ResponseWriter, r *http.Request)
	// get the service version
	// (GET /version)
	GetVersion(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Compose operation middleware
func (siw *ServerInterfaceWrapper) Compose(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.Compose(w, r.WithContext(ctx))
}

// ComposeStatus operation middleware
func (siw *ServerInterfaceWrapper) ComposeStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ComposeStatus(w, r.WithContext(ctx), id)
}

// GetOpenapiJson operation middleware
func (siw *ServerInterfaceWrapper) GetOpenapiJson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.GetOpenapiJson(w, r.WithContext(ctx))
}

// GetVersion operation middleware
func (siw *ServerInterfaceWrapper) GetVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.GetVersion(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Group(func(r chi.Router) {
		r.Post("/compose", wrapper.Compose)
	})
	r.Group(func(r chi.Router) {
		r.Get("/compose/{id}", wrapper.ComposeStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get("/openapi.json", wrapper.GetOpenapiJson)
	})
	r.Group(func(r chi.Router) {
		r.Get("/version", wrapper.GetVersion)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
