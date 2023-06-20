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

package logging

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newMetricsClientHook clientHook

// MetricsCallOptions contains the retry settings for each method of MetricsClient.
type MetricsCallOptions struct {
	ListLogMetrics  []gax.CallOption
	GetLogMetric    []gax.CallOption
	CreateLogMetric []gax.CallOption
	UpdateLogMetric []gax.CallOption
	DeleteLogMetric []gax.CallOption
	CancelOperation []gax.CallOption
	GetOperation    []gax.CallOption
	ListOperations  []gax.CallOption
}

func defaultMetricsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("logging.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("logging.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://logging.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMetricsCallOptions() *MetricsCallOptions {
	return &MetricsCallOptions{
		ListLogMetrics: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetLogMetric: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateLogMetric: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateLogMetric: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteLogMetric: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CancelOperation: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

// internalMetricsClient is an interface that defines the methods available from Cloud Logging API.
type internalMetricsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListLogMetrics(context.Context, *loggingpb.ListLogMetricsRequest, ...gax.CallOption) *LogMetricIterator
	GetLogMetric(context.Context, *loggingpb.GetLogMetricRequest, ...gax.CallOption) (*loggingpb.LogMetric, error)
	CreateLogMetric(context.Context, *loggingpb.CreateLogMetricRequest, ...gax.CallOption) (*loggingpb.LogMetric, error)
	UpdateLogMetric(context.Context, *loggingpb.UpdateLogMetricRequest, ...gax.CallOption) (*loggingpb.LogMetric, error)
	DeleteLogMetric(context.Context, *loggingpb.DeleteLogMetricRequest, ...gax.CallOption) error
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// MetricsClient is a client for interacting with Cloud Logging API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for configuring logs-based metrics.
type MetricsClient struct {
	// The internal transport-dependent client.
	internalClient internalMetricsClient

	// The call options for this service.
	CallOptions *MetricsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetricsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetricsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MetricsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListLogMetrics lists logs-based metrics.
func (c *MetricsClient) ListLogMetrics(ctx context.Context, req *loggingpb.ListLogMetricsRequest, opts ...gax.CallOption) *LogMetricIterator {
	return c.internalClient.ListLogMetrics(ctx, req, opts...)
}

// GetLogMetric gets a logs-based metric.
func (c *MetricsClient) GetLogMetric(ctx context.Context, req *loggingpb.GetLogMetricRequest, opts ...gax.CallOption) (*loggingpb.LogMetric, error) {
	return c.internalClient.GetLogMetric(ctx, req, opts...)
}

// CreateLogMetric creates a logs-based metric.
func (c *MetricsClient) CreateLogMetric(ctx context.Context, req *loggingpb.CreateLogMetricRequest, opts ...gax.CallOption) (*loggingpb.LogMetric, error) {
	return c.internalClient.CreateLogMetric(ctx, req, opts...)
}

// UpdateLogMetric creates or updates a logs-based metric.
func (c *MetricsClient) UpdateLogMetric(ctx context.Context, req *loggingpb.UpdateLogMetricRequest, opts ...gax.CallOption) (*loggingpb.LogMetric, error) {
	return c.internalClient.UpdateLogMetric(ctx, req, opts...)
}

// DeleteLogMetric deletes a logs-based metric.
func (c *MetricsClient) DeleteLogMetric(ctx context.Context, req *loggingpb.DeleteLogMetricRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteLogMetric(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *MetricsClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *MetricsClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *MetricsClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// metricsGRPCClient is a client for interacting with Cloud Logging API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metricsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing MetricsClient
	CallOptions **MetricsCallOptions

	// The gRPC API client.
	metricsClient loggingpb.MetricsServiceV2Client

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMetricsClient creates a new metrics service v2 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for configuring logs-based metrics.
func NewMetricsClient(ctx context.Context, opts ...option.ClientOption) (*MetricsClient, error) {
	clientOpts := defaultMetricsGRPCClientOptions()
	if newMetricsClientHook != nil {
		hookOpts, err := newMetricsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MetricsClient{CallOptions: defaultMetricsCallOptions()}

	c := &metricsGRPCClient{
		connPool:         connPool,
		metricsClient:    loggingpb.NewMetricsServiceV2Client(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *metricsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metricsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metricsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *metricsGRPCClient) ListLogMetrics(ctx context.Context, req *loggingpb.ListLogMetricsRequest, opts ...gax.CallOption) *LogMetricIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListLogMetrics[0:len((*c.CallOptions).ListLogMetrics):len((*c.CallOptions).ListLogMetrics)], opts...)
	it := &LogMetricIterator{}
	req = proto.Clone(req).(*loggingpb.ListLogMetricsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*loggingpb.LogMetric, string, error) {
		resp := &loggingpb.ListLogMetricsResponse{}
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
			resp, err = c.metricsClient.ListLogMetrics(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMetrics(), resp.GetNextPageToken(), nil
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

func (c *metricsGRPCClient) GetLogMetric(ctx context.Context, req *loggingpb.GetLogMetricRequest, opts ...gax.CallOption) (*loggingpb.LogMetric, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "metric_name", url.QueryEscape(req.GetMetricName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetLogMetric[0:len((*c.CallOptions).GetLogMetric):len((*c.CallOptions).GetLogMetric)], opts...)
	var resp *loggingpb.LogMetric
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsClient.GetLogMetric(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsGRPCClient) CreateLogMetric(ctx context.Context, req *loggingpb.CreateLogMetricRequest, opts ...gax.CallOption) (*loggingpb.LogMetric, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateLogMetric[0:len((*c.CallOptions).CreateLogMetric):len((*c.CallOptions).CreateLogMetric)], opts...)
	var resp *loggingpb.LogMetric
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsClient.CreateLogMetric(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsGRPCClient) UpdateLogMetric(ctx context.Context, req *loggingpb.UpdateLogMetricRequest, opts ...gax.CallOption) (*loggingpb.LogMetric, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "metric_name", url.QueryEscape(req.GetMetricName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateLogMetric[0:len((*c.CallOptions).UpdateLogMetric):len((*c.CallOptions).UpdateLogMetric)], opts...)
	var resp *loggingpb.LogMetric
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsClient.UpdateLogMetric(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsGRPCClient) DeleteLogMetric(ctx context.Context, req *loggingpb.DeleteLogMetricRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "metric_name", url.QueryEscape(req.GetMetricName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteLogMetric[0:len((*c.CallOptions).DeleteLogMetric):len((*c.CallOptions).DeleteLogMetric)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.metricsClient.DeleteLogMetric(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metricsGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metricsGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// LogMetricIterator manages a stream of *loggingpb.LogMetric.
type LogMetricIterator struct {
	items    []*loggingpb.LogMetric
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
	InternalFetch func(pageSize int, pageToken string) (results []*loggingpb.LogMetric, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *LogMetricIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *LogMetricIterator) Next() (*loggingpb.LogMetric, error) {
	var item *loggingpb.LogMetric
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *LogMetricIterator) bufLen() int {
	return len(it.items)
}

func (it *LogMetricIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
