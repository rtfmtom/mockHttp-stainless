// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
)

// AbsoluteRedirectService contains methods and other services that help with
// interacting with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAbsoluteRedirectService] method instead.
type AbsoluteRedirectService struct {
	Options []option.RequestOption
}

// NewAbsoluteRedirectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAbsoluteRedirectService(opts ...option.RequestOption) (r AbsoluteRedirectService) {
	r = AbsoluteRedirectService{}
	r.Options = opts
	return
}

// Redirects the request to the target URL using an absolute URL
func (r *AbsoluteRedirectService) Get(ctx context.Context, value int64, query AbsoluteRedirectGetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(query.Host) {
		opts = append(opts, option.WithHeader("host", fmt.Sprintf("%s", query.Host)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("absolute-redirect/%v", value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type AbsoluteRedirectGetParams struct {
	Host string `header:"host,required" json:"-"`
	paramObj
}
