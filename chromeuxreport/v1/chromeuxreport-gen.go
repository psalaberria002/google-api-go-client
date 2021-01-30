// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package chromeuxreport provides access to the Chrome UX Report API.
//
// For product documentation, see: https://developers.google.com/web/tools/chrome-user-experience-report/api/reference
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/chromeuxreport/v1"
//   ...
//   ctx := context.Background()
//   chromeuxreportService, err := chromeuxreport.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   chromeuxreportService, err := chromeuxreport.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   chromeuxreportService, err := chromeuxreport.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package chromeuxreport // import "google.golang.org/api/chromeuxreport/v1"

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

const apiId = "chromeuxreport:v1"
const apiName = "chromeuxreport"
const apiVersion = "v1"
const basePath = "https://chromeuxreport.googleapis.com/"
const mtlsBasePath = "https://chromeuxreport.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
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
	s.Records = NewRecordsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Records *RecordsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewRecordsService(s *Service) *RecordsService {
	rs := &RecordsService{s: s}
	return rs
}

type RecordsService struct {
	s *Service
}

// Bin: A bin is a discrete portion of data spanning from start to end,
// or if no end is given, then from start to +inf. A bin's start and end
// values are given in the value type of the metric it represents. For
// example, "first contentful paint" is measured in milliseconds and
// exposed as ints, therefore its metric bins will use int32s for its
// start and end types. However, "cumulative layout shift" is measured
// in unitless decimals and is exposed as a decimal encoded as a string,
// therefore its metric bins will use strings for its value type.
type Bin struct {
	// Density: The proportion of users that experienced this bin's value
	// for the given metric.
	Density float64 `json:"density,omitempty"`

	// End: End is the end of the data bin. If end is not populated, then
	// the bin has no end and is valid from start to +inf.
	End interface{} `json:"end,omitempty"`

	// Start: Start is the beginning of the data bin.
	Start interface{} `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Density") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Density") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Bin) MarshalJSON() ([]byte, error) {
	type NoMethod Bin
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Bin) UnmarshalJSON(data []byte) error {
	type NoMethod Bin
	var s1 struct {
		Density gensupport.JSONFloat64 `json:"density"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Density = float64(s1.Density)
	return nil
}

// Key: Key defines all the dimensions that identify this record as
// unique.
type Key struct {
	// EffectiveConnectionType: The effective connection type is the general
	// connection class that all users experienced for this record. This
	// field uses the values ["offline", "slow-2G", "2G", "3G", "4G"] as
	// specified in:
	// https://wicg.github.io/netinfo/#effective-connection-types If the
	// effective connection type is unspecified, then aggregated data over
	// all effective connection types will be returned.
	EffectiveConnectionType string `json:"effectiveConnectionType,omitempty"`

	// FormFactor: The form factor is the device class that all users used
	// to access the site for this record. If the form factor is
	// unspecified, then aggregated data over all form factors will be
	// returned.
	//
	// Possible values:
	//   "ALL_FORM_FACTORS" - The default value, representing all device
	// classes.
	//   "PHONE" - The device class representing a "mobile"/"phone" sized
	// client.
	//   "DESKTOP" - The device class representing a "desktop"/"laptop" type
	// full size client.
	//   "TABLET" - The device class representing a "tablet" type client.
	FormFactor string `json:"formFactor,omitempty"`

	// Origin: Origin specifies the origin that this record is for. Note:
	// When specifying an origin, data for loads under this origin over all
	// pages are aggregated into origin level user experience data.
	Origin string `json:"origin,omitempty"`

	// Url: Url specifies a specific url that this record is for. Note: When
	// specifying a "url" only data for that specific url will be
	// aggregated.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "EffectiveConnectionType") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EffectiveConnectionType")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Key) MarshalJSON() ([]byte, error) {
	type NoMethod Key
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metric: A `metric` is a set of user experience data for a single web
// performance metric, like "first contentful paint". It contains a
// summary histogram of real world Chrome usage as a series of `bins`.
type Metric struct {
	// Histogram: The histogram of user experiences for a metric. The
	// histogram will have at least one bin and the densities of all bins
	// will add up to ~1.
	Histogram []*Bin `json:"histogram,omitempty"`

	// Percentiles: Common useful percentiles of the Metric. The value type
	// for the percentiles will be the same as the value types given for the
	// Histogram bins.
	Percentiles *Percentiles `json:"percentiles,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Histogram") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Histogram") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Metric) MarshalJSON() ([]byte, error) {
	type NoMethod Metric
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Percentiles: Percentiles contains synthetic values of a metric at a
// given statistical percentile. These are used for estimating a
// metric's value as experienced by a percentage of users out of the
// total number of users.
type Percentiles struct {
	// P75: 75% of users experienced the given metric at or below this
	// value.
	P75 interface{} `json:"p75,omitempty"`

	// ForceSendFields is a list of field names (e.g. "P75") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "P75") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Percentiles) MarshalJSON() ([]byte, error) {
	type NoMethod Percentiles
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryRequest: Request payload sent by a physical web client. This
// request includes all necessary context to load a particular user
// experience record.
type QueryRequest struct {
	// EffectiveConnectionType: The effective connection type is a query
	// dimension that specifies the effective network class that the
	// record's data should belong to. This field uses the values
	// ["offline", "slow-2G", "2G", "3G", "4G"] as specified in:
	// https://wicg.github.io/netinfo/#effective-connection-types Note: If
	// no effective connection type is specified, then a special record with
	// aggregated data over all effective connection types will be returned.
	EffectiveConnectionType string `json:"effectiveConnectionType,omitempty"`

	// FormFactor: The form factor is a query dimension that specifies the
	// device class that the record's data should belong to. Note: If no
	// form factor is specified, then a special record with aggregated data
	// over all form factors will be returned.
	//
	// Possible values:
	//   "ALL_FORM_FACTORS" - The default value, representing all device
	// classes.
	//   "PHONE" - The device class representing a "mobile"/"phone" sized
	// client.
	//   "DESKTOP" - The device class representing a "desktop"/"laptop" type
	// full size client.
	//   "TABLET" - The device class representing a "tablet" type client.
	FormFactor string `json:"formFactor,omitempty"`

	// Metrics: The metrics that should be included in the response. If none
	// are specified then any metrics found will be returned. Allowed
	// values: ["first_contentful_paint", "first_input_delay",
	// "largest_contentful_paint", "cumulative_layout_shift"]
	Metrics []string `json:"metrics,omitempty"`

	// Origin: The url pattern "origin" refers to a url pattern that is the
	// origin of a website. Examples: "https://example.com",
	// "https://cloud.google.com"
	Origin string `json:"origin,omitempty"`

	// Url: The url pattern "url" refers to a url pattern that is any
	// arbitrary url. Examples: "https://example.com/",
	// "https://cloud.google.com/why-google-cloud/"
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "EffectiveConnectionType") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EffectiveConnectionType")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *QueryRequest) MarshalJSON() ([]byte, error) {
	type NoMethod QueryRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryResponse: Response payload sent back to a physical web client.
// This response contains the record found based on the identiers
// present in a `QueryRequest`. The returned response will have a
// record, and sometimes details on normalization actions taken on the
// request that were necessary to make the request successful.
type QueryResponse struct {
	// Record: The record that was found.
	Record *Record `json:"record,omitempty"`

	// UrlNormalizationDetails: These are details about automated
	// normalization actions that were taken in order to make the requested
	// `url_pattern` valid.
	UrlNormalizationDetails *UrlNormalization `json:"urlNormalizationDetails,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Record") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Record") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryResponse) MarshalJSON() ([]byte, error) {
	type NoMethod QueryResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Record: Record is a single Chrome UX report data record. It contains
// use experience statistics for a single url pattern and set of
// dimensions.
type Record struct {
	// Key: Key defines all of the unique querying parameters needed to look
	// up a user experience record.
	Key *Key `json:"key,omitempty"`

	// Metrics: Metrics is the map of user experience data available for the
	// record defined in the key field. Metrics are keyed on the metric
	// name. Allowed key values: ["first_contentful_paint",
	// "first_input_delay", "largest_contentful_paint",
	// "cumulative_layout_shift"]
	Metrics map[string]Metric `json:"metrics,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Record) MarshalJSON() ([]byte, error) {
	type NoMethod Record
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UrlNormalization: Object representing the normalization actions taken
// to normalize a url to achieve a higher chance of successful lookup.
// These are simple automated changes that are taken when looking up the
// provided `url_patten` would be known to fail. Complex actions like
// following redirects are not handled.
type UrlNormalization struct {
	// NormalizedUrl: The URL after any normalization actions. This is a
	// valid user experience URL that could reasonably be looked up.
	NormalizedUrl string `json:"normalizedUrl,omitempty"`

	// OriginalUrl: The original requested URL prior to any normalization
	// actions.
	OriginalUrl string `json:"originalUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NormalizedUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NormalizedUrl") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UrlNormalization) MarshalJSON() ([]byte, error) {
	type NoMethod UrlNormalization
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "chromeuxreport.records.queryRecord":

type RecordsQueryRecordCall struct {
	s            *Service
	queryrequest *QueryRequest
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// QueryRecord: Queries the Chrome User Experience for a single `record`
// for a given site. Returns a `record` that contains one or more
// `metrics` corresponding to performance data about the requested site.
func (r *RecordsService) QueryRecord(queryrequest *QueryRequest) *RecordsQueryRecordCall {
	c := &RecordsQueryRecordCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.queryrequest = queryrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RecordsQueryRecordCall) Fields(s ...googleapi.Field) *RecordsQueryRecordCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RecordsQueryRecordCall) Context(ctx context.Context) *RecordsQueryRecordCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *RecordsQueryRecordCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *RecordsQueryRecordCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210129")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.queryrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/records:queryRecord")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chromeuxreport.records.queryRecord" call.
// Exactly one of *QueryResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *QueryResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RecordsQueryRecordCall) Do(opts ...googleapi.CallOption) (*QueryResponse, error) {
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
	ret := &QueryResponse{
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
	//   "description": "Queries the Chrome User Experience for a single `record` for a given site. Returns a `record` that contains one or more `metrics` corresponding to performance data about the requested site.",
	//   "flatPath": "v1/records:queryRecord",
	//   "httpMethod": "POST",
	//   "id": "chromeuxreport.records.queryRecord",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/records:queryRecord",
	//   "request": {
	//     "$ref": "QueryRequest"
	//   },
	//   "response": {
	//     "$ref": "QueryResponse"
	//   }
	// }

}
