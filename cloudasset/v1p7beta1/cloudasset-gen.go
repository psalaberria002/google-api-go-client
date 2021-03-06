// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package cloudasset provides access to the Cloud Asset API.
//
// For product documentation, see: https://cloud.google.com/asset-inventory/docs/quickstart
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/cloudasset/v1p7beta1"
//   ...
//   ctx := context.Background()
//   cloudassetService, err := cloudasset.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   cloudassetService, err := cloudasset.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   cloudassetService, err := cloudasset.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package cloudasset // import "google.golang.org/api/cloudasset/v1p7beta1"

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

const apiId = "cloudasset:v1p7beta1"
const apiName = "cloudasset"
const apiVersion = "v1p7beta1"
const basePath = "https://cloudasset.googleapis.com/"
const mtlsBasePath = "https://cloudasset.mtls.googleapis.com/"

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
	s.V1p7beta1 = NewV1p7beta1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	V1p7beta1 *V1p7beta1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewV1p7beta1Service(s *Service) *V1p7beta1Service {
	rs := &V1p7beta1Service{s: s}
	return rs
}

type V1p7beta1Service struct {
	s *Service
}

// BigQueryDestination: A BigQuery destination for exporting assets to.
type BigQueryDestination struct {
	// Dataset: Required. The BigQuery dataset in format
	// "projects/projectId/datasets/datasetId", to which the snapshot result
	// should be exported. If this dataset does not exist, the export call
	// returns an INVALID_ARGUMENT error.
	Dataset string `json:"dataset,omitempty"`

	// Force: If the destination table already exists and this flag is
	// `TRUE`, the table will be overwritten by the contents of assets
	// snapshot. If the flag is `FALSE` or unset and the destination table
	// already exists, the export call returns an INVALID_ARGUMEMT error.
	Force bool `json:"force,omitempty"`

	// PartitionSpec: [partition_spec] determines whether to export to
	// partitioned table(s) and how to partition the data. If
	// [partition_spec] is unset or [partition_spec.partition_key] is unset
	// or `PARTITION_KEY_UNSPECIFIED`, the snapshot results will be exported
	// to non-partitioned table(s). [force] will decide whether to overwrite
	// existing table(s). If [partition_spec] is specified. First, the
	// snapshot results will be written to partitioned table(s) with two
	// additional timestamp columns, readTime and requestTime, one of which
	// will be the partition key. Secondly, in the case when any destination
	// table already exists, it will first try to update existing table's
	// schema as necessary by appending additional columns. Then, if [force]
	// is `TRUE`, the corresponding partition will be overwritten by the
	// snapshot results (data in different partitions will remain intact);
	// if [force] is unset or `FALSE`, it will append the data. An error
	// will be returned if the schema update or data appension fails.
	PartitionSpec *PartitionSpec `json:"partitionSpec,omitempty"`

	// SeparateTablesPerAssetType: If this flag is `TRUE`, the snapshot
	// results will be written to one or multiple tables, each of which
	// contains results of one asset type. The [force] and [partition_spec]
	// fields will apply to each of them. Field [table] will be concatenated
	// with "_" and the asset type names (see
	// https://cloud.google.com/asset-inventory/docs/supported-asset-types
	// for supported asset types) to construct per-asset-type table names,
	// in which all non-alphanumeric characters like "." and "/" will be
	// substituted by "_". Example: if field [table] is "mytable" and
	// snapshot results contain "storage.googleapis.com/Bucket" assets, the
	// corresponding table name will be
	// "mytable_storage_googleapis_com_Bucket". If any of these tables does
	// not exist, a new table with the concatenated name will be created.
	// When [content_type] in the ExportAssetsRequest is `RESOURCE`, the
	// schema of each table will include RECORD-type columns mapped to the
	// nested fields in the Asset.resource.data field of that asset type (up
	// to the 15 nested level BigQuery supports
	// (https://cloud.google.com/bigquery/docs/nested-repeated#limitations)).
	// The fields in >15 nested levels will be stored in JSON format string
	// as a child column of its parent RECORD column. If error occurs when
	// exporting to any table, the whole export call will return an error
	// but the export results that already succeed will persist. Example: if
	// exporting to table_type_A succeeds when exporting to table_type_B
	// fails during one export call, the results in table_type_A will
	// persist and there will not be partial results persisting in a table.
	SeparateTablesPerAssetType bool `json:"separateTablesPerAssetType,omitempty"`

	// Table: Required. The BigQuery table to which the snapshot result
	// should be written. If this table does not exist, a new table with the
	// given name will be created.
	Table string `json:"table,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Dataset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Dataset") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BigQueryDestination) MarshalJSON() ([]byte, error) {
	type NoMethod BigQueryDestination
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExportAssetsRequest: Export asset request.
type ExportAssetsRequest struct {
	// AssetTypes: A list of asset types to take a snapshot for. For
	// example: "compute.googleapis.com/Disk". Regular expressions are also
	// supported. For example: * "compute.googleapis.com.*" snapshots
	// resources whose asset type starts with "compute.googleapis.com". *
	// ".*Instance" snapshots resources whose asset type ends with
	// "Instance". * ".*Instance.*" snapshots resources whose asset type
	// contains "Instance". See RE2
	// (https://github.com/google/re2/wiki/Syntax) for all supported regular
	// expression syntax. If the regular expression does not match any
	// supported asset type, an INVALID_ARGUMENT error will be returned. If
	// specified, only matching assets will be returned, otherwise, it will
	// snapshot all asset types. See Introduction to Cloud Asset Inventory
	// (https://cloud.google.com/asset-inventory/docs/overview) for all
	// supported asset types.
	AssetTypes []string `json:"assetTypes,omitempty"`

	// ContentType: Asset content type. If not specified, no content but the
	// asset name will be returned.
	//
	// Possible values:
	//   "CONTENT_TYPE_UNSPECIFIED" - Unspecified content type.
	//   "RESOURCE" - Resource metadata.
	//   "IAM_POLICY" - The actual IAM policy set on a resource.
	//   "ORG_POLICY" - The Cloud Organization Policy set on an asset.
	//   "ACCESS_POLICY" - The Cloud Access context manager Policy set on an
	// asset.
	//   "RELATIONSHIP" - The related resources.
	ContentType string `json:"contentType,omitempty"`

	// OutputConfig: Required. Output configuration indicating where the
	// results will be output to.
	OutputConfig *OutputConfig `json:"outputConfig,omitempty"`

	// ReadTime: Timestamp to take an asset snapshot. This can only be set
	// to a timestamp between the current time and the current time minus 35
	// days (inclusive). If not specified, the current time will be used.
	// Due to delays in resource data collection and indexing, there is a
	// volatile window during which running the same query may get different
	// results.
	ReadTime string `json:"readTime,omitempty"`

	// RelationshipTypes: A list of relationship types to export, for
	// example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be
	// specified if content_type=RELATIONSHIP. If specified, it will
	// snapshot [asset_types]' specified relationships, or give errors if
	// any relationship_types' supported types are not in [asset_types]. If
	// not specified, it will snapshot all [asset_types]' supported
	// relationships. An unspecified [asset_types] field means all supported
	// asset_types. See Introduction to Cloud Asset Inventory
	// (https://cloud.google.com/asset-inventory/docs/overview) for all
	// supported asset types and relationship types.
	RelationshipTypes []string `json:"relationshipTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AssetTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AssetTypes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExportAssetsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ExportAssetsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GcsDestination: A Cloud Storage location.
type GcsDestination struct {
	// Uri: The uri of the Cloud Storage object. It's the same uri that is
	// used by gsutil. Example: "gs://bucket_name/object_name". See Viewing
	// and Editing Object Metadata
	// (https://cloud.google.com/storage/docs/viewing-editing-metadata) for
	// more information.
	Uri string `json:"uri,omitempty"`

	// UriPrefix: The uri prefix of all generated Cloud Storage objects.
	// Example: "gs://bucket_name/object_name_prefix". Each object uri is in
	// format:
	// "gs://bucket_name/object_name_prefix/{ASSET_TYPE}/{SHARD_NUMBER} and
	// only contains assets for that type. starts from 0. Example:
	// "gs://bucket_name/object_name_prefix/compute.googleapis.com/Disk/0"
	// is the first shard of output objects containing all
	// compute.googleapis.com/Disk assets. An INVALID_ARGUMENT error will be
	// returned if file with the same name
	// "gs://bucket_name/object_name_prefix" already exists.
	UriPrefix string `json:"uriPrefix,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Uri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Uri") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GcsDestination) MarshalJSON() ([]byte, error) {
	type NoMethod GcsDestination
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress. If `true`, the operation is completed, and either `error`
	// or `response` is available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation. It
	// typically contains progress information and common metadata such as
	// create time. Some services might not provide such metadata. Any
	// method that returns a long-running operation should document the
	// metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that originally returns it. If you use the default HTTP
	// mapping, the `name` should be a resource name ending with
	// `operations/{unique_id}`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success. If
	// the original method returns no data on success, such as `Delete`, the
	// response is `google.protobuf.Empty`. If the original method is
	// standard `Get`/`Create`/`Update`, the response should be the
	// resource. For other methods, the response should have the type
	// `XxxResponse`, where `Xxx` is the original method name. For example,
	// if the original method name is `TakeSnapshot()`, the inferred
	// response type is `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OutputConfig: Output configuration for export assets destination.
type OutputConfig struct {
	// BigqueryDestination: Destination on BigQuery. The output table stores
	// the fields in asset proto as columns in BigQuery.
	BigqueryDestination *BigQueryDestination `json:"bigqueryDestination,omitempty"`

	// GcsDestination: Destination on Cloud Storage.
	GcsDestination *GcsDestination `json:"gcsDestination,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BigqueryDestination")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BigqueryDestination") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *OutputConfig) MarshalJSON() ([]byte, error) {
	type NoMethod OutputConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PartitionSpec: Specifications of BigQuery partitioned table as export
// destination.
type PartitionSpec struct {
	// PartitionKey: The partition key for BigQuery partitioned table.
	//
	// Possible values:
	//   "PARTITION_KEY_UNSPECIFIED" - Unspecified partition key. If used,
	// it means using non-partitioned table.
	//   "READ_TIME" - The time when the snapshot is taken. If specified as
	// partition key, the result table(s) is partitoned by the additional
	// timestamp column, readTime. If [read_time] in ExportAssetsRequest is
	// specified, the readTime column's value will be the same as it.
	// Otherwise, its value will be the current time that is used to take
	// the snapshot.
	//   "REQUEST_TIME" - The time when the request is received and started
	// to be processed. If specified as partition key, the result table(s)
	// is partitoned by the requestTime column, an additional timestamp
	// column representing when the request was received.
	PartitionKey string `json:"partitionKey,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PartitionKey") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PartitionKey") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PartitionSpec) MarshalJSON() ([]byte, error) {
	type NoMethod PartitionSpec
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different programming environments, including REST APIs
// and RPC APIs. It is used by gRPC (https://github.com/grpc). Each
// `Status` message contains three pieces of data: error code, error
// message, and error details. You can find out more about this error
// model and how to work with it in the API Design Guide
// (https://cloud.google.com/apis/design/errors).
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There is a
	// common set of message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudasset.exportAssets":

type V1p7beta1ExportAssetsCall struct {
	s                   *Service
	parent              string
	exportassetsrequest *ExportAssetsRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// ExportAssets: Exports assets with time and resource types to a given
// Cloud Storage location/BigQuery table. For Cloud Storage location
// destinations, the output format is newline-delimited JSON. Each line
// represents a google.cloud.asset.v1p7beta1.Asset in the JSON format;
// for BigQuery table destinations, the output table stores the fields
// in asset proto as columns. This API implements the
// google.longrunning.Operation API , which allows you to keep track of
// the export. We recommend intervals of at least 2 seconds with
// exponential retry to poll the export operation result. For
// regular-size resource parent, the export operation usually finishes
// within 5 minutes.
func (r *V1p7beta1Service) ExportAssets(parent string, exportassetsrequest *ExportAssetsRequest) *V1p7beta1ExportAssetsCall {
	c := &V1p7beta1ExportAssetsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.exportassetsrequest = exportassetsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1p7beta1ExportAssetsCall) Fields(s ...googleapi.Field) *V1p7beta1ExportAssetsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1p7beta1ExportAssetsCall) Context(ctx context.Context) *V1p7beta1ExportAssetsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1p7beta1ExportAssetsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1p7beta1ExportAssetsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210217")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportassetsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1p7beta1/{+parent}:exportAssets")
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

// Do executes the "cloudasset.exportAssets" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *V1p7beta1ExportAssetsCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Exports assets with time and resource types to a given Cloud Storage location/BigQuery table. For Cloud Storage location destinations, the output format is newline-delimited JSON. Each line represents a google.cloud.asset.v1p7beta1.Asset in the JSON format; for BigQuery table destinations, the output table stores the fields in asset proto as columns. This API implements the google.longrunning.Operation API , which allows you to keep track of the export. We recommend intervals of at least 2 seconds with exponential retry to poll the export operation result. For regular-size resource parent, the export operation usually finishes within 5 minutes.",
	//   "flatPath": "v1p7beta1/{v1p7beta1Id}/{v1p7beta1Id1}:exportAssets",
	//   "httpMethod": "POST",
	//   "id": "cloudasset.exportAssets",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The relative name of the root asset. This can only be an organization number (such as \"organizations/123\"), a project ID (such as \"projects/my-project-id\"), or a project number (such as \"projects/12345\"), or a folder number (such as \"folders/123\").",
	//       "location": "path",
	//       "pattern": "^[^/]+/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1p7beta1/{+parent}:exportAssets",
	//   "request": {
	//     "$ref": "ExportAssetsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
