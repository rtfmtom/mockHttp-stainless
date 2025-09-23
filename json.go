// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
)

// JsonService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJsonService] method instead.
type JsonService struct {
	Options []option.RequestOption
}

// NewJsonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJsonService(opts ...option.RequestOption) (r JsonService) {
	r = JsonService{}
	r.Options = opts
	return
}

// Returns random JSON content
func (r *JsonService) New(ctx context.Context, opts ...option.RequestOption) (res *JsonNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Returns random JSON content
func (r *JsonService) Get(ctx context.Context, opts ...option.RequestOption) (res *JsonGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns random JSON content
func (r *JsonService) Update(ctx context.Context, opts ...option.RequestOption) (res *JsonUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Returns random JSON content
func (r *JsonService) Delete(ctx context.Context, opts ...option.RequestOption) (res *JsonDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns random JSON content
func (r *JsonService) Modify(ctx context.Context, opts ...option.RequestOption) (res *JsonModifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type JsonNewResponse map[string]any

type JsonGetResponse map[string]any

type JsonUpdateResponse map[string]any

type JsonDeleteResponse map[string]any

type JsonModifyResponse map[string]any
