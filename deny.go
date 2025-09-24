// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"slices"

	"github.com/rtfmtom/mockHttp-stainless/internal/requestconfig"
	"github.com/rtfmtom/mockHttp-stainless/option"
)

// DenyService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDenyService] method instead.
type DenyService struct {
	Options []option.RequestOption
}

// NewDenyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDenyService(opts ...option.RequestOption) (r DenyService) {
	r = DenyService{}
	r.Options = opts
	return
}

// Returns page denied by robots.txt rules
func (r *DenyService) New(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "deny"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Returns page denied by robots.txt rules
func (r *DenyService) Get(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "deny"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Returns page denied by robots.txt rules
func (r *DenyService) Update(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "deny"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Returns page denied by robots.txt rules
func (r *DenyService) Delete(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "deny"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns page denied by robots.txt rules
func (r *DenyService) Modify(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "deny"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}
