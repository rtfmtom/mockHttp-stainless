// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/apijson"
	"github.com/stainless-sdks/mockhttp-go/internal/apiquery"
	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
	"github.com/stainless-sdks/mockhttp-go/packages/param"
	"github.com/stainless-sdks/mockhttp-go/packages/respjson"
)

// BearerService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBearerService] method instead.
type BearerService struct {
	Options []option.RequestOption
}

// NewBearerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBearerService(opts ...option.RequestOption) (r BearerService) {
	r = BearerService{}
	r.Options = opts
	return
}

// HTTP Bearer authentication. Succeeds if any Bearer token is present unless
// ?required=true is provided.
func (r *BearerService) Get(ctx context.Context, query BearerGetParams, opts ...option.RequestOption) (res *BearerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "bearer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BearerGetResponse struct {
	Token         string `json:"token,required"`
	Authenticated bool   `json:"authenticated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token         respjson.Field
		Authenticated respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BearerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BearerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BearerGetParams struct {
	Required param.Opt[bool] `query:"required,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BearerGetParams]'s query parameters as `url.Values`.
func (r BearerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
