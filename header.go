// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"slices"

	"github.com/rtfmtom/mockHttp-stainless/internal/apijson"
	"github.com/rtfmtom/mockHttp-stainless/internal/requestconfig"
	"github.com/rtfmtom/mockHttp-stainless/option"
	"github.com/rtfmtom/mockHttp-stainless/packages/respjson"
)

// HeaderService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHeaderService] method instead.
type HeaderService struct {
	Options []option.RequestOption
}

// NewHeaderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHeaderService(opts ...option.RequestOption) (r HeaderService) {
	r = HeaderService{}
	r.Options = opts
	return
}

// Returns all headers of the client request
func (r *HeaderService) List(ctx context.Context, opts ...option.RequestOption) (res *HeaderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "headers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HeaderListResponse struct {
	// All headers of the client request
	Headers map[string]string `json:"headers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeaderListResponse) RawJSON() string { return r.JSON.raw }
func (r *HeaderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
