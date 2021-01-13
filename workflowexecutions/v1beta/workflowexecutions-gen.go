// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package workflowexecutions provides access to the Workflow Executions API.
//
// For product documentation, see: https://cloud.google.com/workflows
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/workflowexecutions/v1beta"
//   ...
//   ctx := context.Background()
//   workflowexecutionsService, err := workflowexecutions.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   workflowexecutionsService, err := workflowexecutions.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   workflowexecutionsService, err := workflowexecutions.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package workflowexecutions // import "google.golang.org/api/workflowexecutions/v1beta"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "workflowexecutions:v1beta"
const apiName = "workflowexecutions"
const apiVersion = "v1beta"
const basePath = "https://workflowexecutions.googleapis.com/"
const mtlsBasePath = "https://workflowexecutions.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.Workflows = NewProjectsLocationsWorkflowsService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	Workflows *ProjectsLocationsWorkflowsService
}

func NewProjectsLocationsWorkflowsService(s *Service) *ProjectsLocationsWorkflowsService {
	rs := &ProjectsLocationsWorkflowsService{s: s}
	rs.Executions = NewProjectsLocationsWorkflowsExecutionsService(s)
	return rs
}

type ProjectsLocationsWorkflowsService struct {
	s *Service

	Executions *ProjectsLocationsWorkflowsExecutionsService
}

func NewProjectsLocationsWorkflowsExecutionsService(s *Service) *ProjectsLocationsWorkflowsExecutionsService {
	rs := &ProjectsLocationsWorkflowsExecutionsService{s: s}
	return rs
}

type ProjectsLocationsWorkflowsExecutionsService struct {
	s *Service
}

// CancelExecutionRequest: Request for the CancelExecution method.
type CancelExecutionRequest struct {
}

// Error: Error describes why the execution was abnormally terminated.
type Error struct {
	// Context: Human readable error context, helpful for debugging
	// purposes.
	Context string `json:"context,omitempty"`

	// Payload: Error payload returned by the execution, represented as a
	// JSON string.
	Payload string `json:"payload,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Context") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Context") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Error) MarshalJSON() ([]byte, error) {
	type NoMethod Error
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Execution: A running instance of a Workflow
// (/workflows/docs/reference/rest/v1beta/projects.locations.workflows).
type Execution struct {
	// Argument: Input parameters of the execution represented as a JSON
	// string. The size limit is 32KB. *Note*: If you are using the REST API
	// directly to run your workflow, you must escape any JSON string value
	// of `argument`. Example:
	// `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument string `json:"argument,omitempty"`

	// EndTime: Output only. Marks the end of execution, successful or not.
	EndTime string `json:"endTime,omitempty"`

	// Error: Output only. The error which caused the execution to finish
	// prematurely. The value is only present if the execution's state is
	// `FAILED` or `CANCELLED`.
	Error *Error `json:"error,omitempty"`

	// Name: Output only. The resource name of the execution. Format:
	// projects/{project}/locations/{location}/workflows/{workflow}/execution
	// s/{execution}
	Name string `json:"name,omitempty"`

	// Result: Output only. Output of the execution represented as a JSON
	// string. The value can only be present if the execution's state is
	// `SUCCEEDED`.
	Result string `json:"result,omitempty"`

	// StartTime: Output only. Marks the beginning of execution.
	StartTime string `json:"startTime,omitempty"`

	// State: Output only. Current state of the execution.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Invalid state.
	//   "ACTIVE" - The execution is in progress.
	//   "SUCCEEDED" - The execution finished successfully.
	//   "FAILED" - The execution failed with an error.
	//   "CANCELLED" - The execution was stopped intentionally.
	State string `json:"state,omitempty"`

	// WorkflowRevisionId: Output only. Revision of the workflow this
	// execution is using.
	WorkflowRevisionId string `json:"workflowRevisionId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Argument") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Argument") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Execution) MarshalJSON() ([]byte, error) {
	type NoMethod Execution
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListExecutionsResponse: Response for the ListExecutions method.
type ListExecutionsResponse struct {
	// Executions: The executions which match the request.
	Executions []*Execution `json:"executions,omitempty"`

	// NextPageToken: A token, which can be sent as `page_token` to retrieve
	// the next page. If this field is omitted, there are no subsequent
	// pages.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Executions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Executions") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListExecutionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListExecutionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "workflowexecutions.projects.locations.workflows.executions.cancel":

type ProjectsLocationsWorkflowsExecutionsCancelCall struct {
	s                      *Service
	name                   string
	cancelexecutionrequest *CancelExecutionRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Cancel: Cancels an execution of the given name.
func (r *ProjectsLocationsWorkflowsExecutionsService) Cancel(name string, cancelexecutionrequest *CancelExecutionRequest) *ProjectsLocationsWorkflowsExecutionsCancelCall {
	c := &ProjectsLocationsWorkflowsExecutionsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.cancelexecutionrequest = cancelexecutionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210112")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cancelexecutionrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.cancel" call.
// Exactly one of *Execution or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Execution.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Do(opts ...googleapi.CallOption) (*Execution, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Execution{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels an execution of the given name.",
	//   "flatPath": "v1beta/projects/{projectsId}/locations/{locationsId}/workflows/{workflowsId}/executions/{executionsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "workflowexecutions.projects.locations.workflows.executions.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Name of the execution to be cancelled. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflows/[^/]+/executions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}:cancel",
	//   "request": {
	//     "$ref": "CancelExecutionRequest"
	//   },
	//   "response": {
	//     "$ref": "Execution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "workflowexecutions.projects.locations.workflows.executions.create":

type ProjectsLocationsWorkflowsExecutionsCreateCall struct {
	s          *Service
	parent     string
	execution  *Execution
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new execution using the latest revision of the
// given workflow.
func (r *ProjectsLocationsWorkflowsExecutionsService) Create(parent string, execution *Execution) *ProjectsLocationsWorkflowsExecutionsCreateCall {
	c := &ProjectsLocationsWorkflowsExecutionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.execution = execution
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210112")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.execution)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/executions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.create" call.
// Exactly one of *Execution or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Execution.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Do(opts ...googleapi.CallOption) (*Execution, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Execution{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new execution using the latest revision of the given workflow.",
	//   "flatPath": "v1beta/projects/{projectsId}/locations/{locationsId}/workflows/{workflowsId}/executions",
	//   "httpMethod": "POST",
	//   "id": "workflowexecutions.projects.locations.workflows.executions.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Name of the workflow for which an execution should be created. Format: projects/{project}/locations/{location}/workflows/{workflow} The latest revision of the workflow will be used.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflows/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+parent}/executions",
	//   "request": {
	//     "$ref": "Execution"
	//   },
	//   "response": {
	//     "$ref": "Execution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "workflowexecutions.projects.locations.workflows.executions.get":

type ProjectsLocationsWorkflowsExecutionsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns an execution of the given name.
func (r *ProjectsLocationsWorkflowsExecutionsService) Get(name string) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c := &ProjectsLocationsWorkflowsExecutionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// View sets the optional parameter "view": A view defining which fields
// should be filled in the returned execution. The API will default to
// the FULL view.
//
// Possible values:
//   "EXECUTION_VIEW_UNSPECIFIED" - The default / unset value.
//   "BASIC" - Includes only basic metadata about the execution.
// Following fields are returned: name, start_time, end_time, state and
// workflow_revision_id.
//   "FULL" - Includes all data.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) View(view string) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210112")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.get" call.
// Exactly one of *Execution or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Execution.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Do(opts ...googleapi.CallOption) (*Execution, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Execution{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns an execution of the given name.",
	//   "flatPath": "v1beta/projects/{projectsId}/locations/{locationsId}/workflows/{workflowsId}/executions/{executionsId}",
	//   "httpMethod": "GET",
	//   "id": "workflowexecutions.projects.locations.workflows.executions.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Name of the execution to be retrieved. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflows/[^/]+/executions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Optional. A view defining which fields should be filled in the returned execution. The API will default to the FULL view.",
	//       "enum": [
	//         "EXECUTION_VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "The default / unset value.",
	//         "Includes only basic metadata about the execution. Following fields are returned: name, start_time, end_time, state and workflow_revision_id.",
	//         "Includes all data."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+name}",
	//   "response": {
	//     "$ref": "Execution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "workflowexecutions.projects.locations.workflows.executions.list":

type ProjectsLocationsWorkflowsExecutionsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns a list of executions which belong to the workflow with
// the given name. The method returns executions of all workflow
// revisions. Returned executions are ordered by their start time
// (newest first).
func (r *ProjectsLocationsWorkflowsExecutionsService) List(parent string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c := &ProjectsLocationsWorkflowsExecutionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// executions to return per call. Max supported value depends on the
// selected Execution view: it's 10000 for BASIC and 100 for FULL. The
// default value used if the field is not specified is 100, regardless
// of the selected view. Values greater than the max value will be
// coerced down to it.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) PageSize(pageSize int64) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token,
// received from a previous `ListExecutions` call. Provide this to
// retrieve the subsequent page. When paginating, all other parameters
// provided to `ListExecutions` must match the call that provided the
// page token.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) PageToken(pageToken string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// View sets the optional parameter "view": A view defining which fields
// should be filled in the returned executions. The API will default to
// the BASIC view.
//
// Possible values:
//   "EXECUTION_VIEW_UNSPECIFIED" - The default / unset value.
//   "BASIC" - Includes only basic metadata about the execution.
// Following fields are returned: name, start_time, end_time, state and
// workflow_revision_id.
//   "FULL" - Includes all data.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) View(view string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210112")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/executions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.list" call.
// Exactly one of *ListExecutionsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListExecutionsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Do(opts ...googleapi.CallOption) (*ListExecutionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListExecutionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of executions which belong to the workflow with the given name. The method returns executions of all workflow revisions. Returned executions are ordered by their start time (newest first).",
	//   "flatPath": "v1beta/projects/{projectsId}/locations/{locationsId}/workflows/{workflowsId}/executions",
	//   "httpMethod": "GET",
	//   "id": "workflowexecutions.projects.locations.workflows.executions.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Maximum number of executions to return per call. Max supported value depends on the selected Execution view: it's 10000 for BASIC and 100 for FULL. The default value used if the field is not specified is 100, regardless of the selected view. Values greater than the max value will be coerced down to it.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A page token, received from a previous `ListExecutions` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListExecutions` must match the call that provided the page token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. Name of the workflow for which the executions should be listed. Format: projects/{project}/locations/{location}/workflows/{workflow}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflows/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Optional. A view defining which fields should be filled in the returned executions. The API will default to the BASIC view.",
	//       "enum": [
	//         "EXECUTION_VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "The default / unset value.",
	//         "Includes only basic metadata about the execution. Following fields are returned: name, start_time, end_time, state and workflow_revision_id.",
	//         "Includes all data."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta/{+parent}/executions",
	//   "response": {
	//     "$ref": "ListExecutionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Pages(ctx context.Context, f func(*ListExecutionsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
