// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/apijson"
	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
	"github.com/stainless-sdks/mockhttp-go/packages/respjson"
)

// GetService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetService] method instead.
type GetService struct {
	Options []option.RequestOption
}

// NewGetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGetService(opts ...option.RequestOption) (r GetService) {
	r = GetService{}
	r.Options = opts
	return
}

// Returns request information for GET requests
func (r *GetService) RequestInfo(ctx context.Context, opts ...option.RequestOption) (res *GetRequestInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GetRequestInfoResponse struct {
	Headers     map[string]string `json:"headers,required"`
	Method      string            `json:"method,required"`
	QueryParams map[string]string `json:"queryParams,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Method      respjson.Field
		QueryParams respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRequestInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *GetRequestInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
