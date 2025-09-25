// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"slices"

	"github.com/rtfmtom/mockHttp-stainless/internal/requestconfig"
	"github.com/rtfmtom/mockHttp-stainless/option"
)

// HTMLService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTMLService] method instead.
type HTMLService struct {
	Options []option.RequestOption
}

// NewHTMLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTMLService(opts ...option.RequestOption) (r HTMLService) {
	r = HTMLService{}
	r.Options = opts
	return
}

// Returns random HTML content
func (r *HTMLService) New(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Returns random HTML content
func (r *HTMLService) Get(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns random HTML content
func (r *HTMLService) Update(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Returns random HTML content
func (r *HTMLService) Delete(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns random HTML content
func (r *HTMLService) Modify(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}
