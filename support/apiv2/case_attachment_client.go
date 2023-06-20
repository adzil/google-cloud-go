// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package support

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	supportpb "cloud.google.com/go/support/apiv2/supportpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newCaseAttachmentClientHook clientHook

// CaseAttachmentCallOptions contains the retry settings for each method of CaseAttachmentClient.
type CaseAttachmentCallOptions struct {
	ListAttachments []gax.CallOption
}

func defaultCaseAttachmentGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudsupport.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudsupport.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudsupport.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCaseAttachmentCallOptions() *CaseAttachmentCallOptions {
	return &CaseAttachmentCallOptions{
		ListAttachments: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultCaseAttachmentRESTCallOptions() *CaseAttachmentCallOptions {
	return &CaseAttachmentCallOptions{
		ListAttachments: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalCaseAttachmentClient is an interface that defines the methods available from Google Cloud Support API.
type internalCaseAttachmentClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAttachments(context.Context, *supportpb.ListAttachmentsRequest, ...gax.CallOption) *AttachmentIterator
}

// CaseAttachmentClient is a client for interacting with Google Cloud Support API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service to manage file attachment for Google Cloud support cases.
type CaseAttachmentClient struct {
	// The internal transport-dependent client.
	internalClient internalCaseAttachmentClient

	// The call options for this service.
	CallOptions *CaseAttachmentCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CaseAttachmentClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CaseAttachmentClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CaseAttachmentClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAttachments retrieve all attachments associated with a support case.
func (c *CaseAttachmentClient) ListAttachments(ctx context.Context, req *supportpb.ListAttachmentsRequest, opts ...gax.CallOption) *AttachmentIterator {
	return c.internalClient.ListAttachments(ctx, req, opts...)
}

// caseAttachmentGRPCClient is a client for interacting with Google Cloud Support API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type caseAttachmentGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CaseAttachmentClient
	CallOptions **CaseAttachmentCallOptions

	// The gRPC API client.
	caseAttachmentClient supportpb.CaseAttachmentServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCaseAttachmentClient creates a new case attachment service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service to manage file attachment for Google Cloud support cases.
func NewCaseAttachmentClient(ctx context.Context, opts ...option.ClientOption) (*CaseAttachmentClient, error) {
	clientOpts := defaultCaseAttachmentGRPCClientOptions()
	if newCaseAttachmentClientHook != nil {
		hookOpts, err := newCaseAttachmentClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CaseAttachmentClient{CallOptions: defaultCaseAttachmentCallOptions()}

	c := &caseAttachmentGRPCClient{
		connPool:             connPool,
		caseAttachmentClient: supportpb.NewCaseAttachmentServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *caseAttachmentGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *caseAttachmentGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *caseAttachmentGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type caseAttachmentRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing CaseAttachmentClient
	CallOptions **CaseAttachmentCallOptions
}

// NewCaseAttachmentRESTClient creates a new case attachment service rest client.
//
// A service to manage file attachment for Google Cloud support cases.
func NewCaseAttachmentRESTClient(ctx context.Context, opts ...option.ClientOption) (*CaseAttachmentClient, error) {
	clientOpts := append(defaultCaseAttachmentRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCaseAttachmentRESTCallOptions()
	c := &caseAttachmentRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &CaseAttachmentClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCaseAttachmentRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudsupport.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudsupport.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudsupport.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *caseAttachmentRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *caseAttachmentRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *caseAttachmentRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *caseAttachmentGRPCClient) ListAttachments(ctx context.Context, req *supportpb.ListAttachmentsRequest, opts ...gax.CallOption) *AttachmentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListAttachments[0:len((*c.CallOptions).ListAttachments):len((*c.CallOptions).ListAttachments)], opts...)
	it := &AttachmentIterator{}
	req = proto.Clone(req).(*supportpb.ListAttachmentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*supportpb.Attachment, string, error) {
		resp := &supportpb.ListAttachmentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.caseAttachmentClient.ListAttachments(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAttachments(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListAttachments retrieve all attachments associated with a support case.
func (c *caseAttachmentRESTClient) ListAttachments(ctx context.Context, req *supportpb.ListAttachmentsRequest, opts ...gax.CallOption) *AttachmentIterator {
	it := &AttachmentIterator{}
	req = proto.Clone(req).(*supportpb.ListAttachmentsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*supportpb.Attachment, string, error) {
		resp := &supportpb.ListAttachmentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v2/%v/attachments", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetAttachments(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// AttachmentIterator manages a stream of *supportpb.Attachment.
type AttachmentIterator struct {
	items    []*supportpb.Attachment
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*supportpb.Attachment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AttachmentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AttachmentIterator) Next() (*supportpb.Attachment, error) {
	var item *supportpb.Attachment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AttachmentIterator) bufLen() int {
	return len(it.items)
}

func (it *AttachmentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
