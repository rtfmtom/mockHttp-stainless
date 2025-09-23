// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/apiquery"
	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
	"github.com/stainless-sdks/mockhttp-go/packages/param"
)

// ResponseHeaderService contains methods and other services that help with
// interacting with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseHeaderService] method instead.
type ResponseHeaderService struct {
	Options []option.RequestOption
}

// NewResponseHeaderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResponseHeaderService(opts ...option.RequestOption) (r ResponseHeaderService) {
	r = ResponseHeaderService{}
	r.Options = opts
	return
}

// Returns a set of response headers based on the query string.
func (r *ResponseHeaderService) New(ctx context.Context, body ResponseHeaderNewParams, opts ...option.RequestOption) (res *ResponseHeaderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "response-headers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a set of response headers based on the query string.
func (r *ResponseHeaderService) Get(ctx context.Context, query ResponseHeaderGetParams, opts ...option.RequestOption) (res *ResponseHeaderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "response-headers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ResponseHeaderNewResponse map[string]any

type ResponseHeaderGetResponse map[string]any

type ResponseHeaderNewParams struct {
	// Enable additional properties
	AdditionalProperties param.Opt[bool] `query:"additionalProperties,omitzero" json:"-"`
	// Freeform query string parameters to be added as response headers
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Type parameter for response headers
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResponseHeaderNewParams]'s query parameters as
// `url.Values`.
func (r ResponseHeaderNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ResponseHeaderGetParams struct {
	// Enable additional properties
	AdditionalProperties param.Opt[bool] `query:"additionalProperties,omitzero" json:"-"`
	// Freeform query string parameters to be added as response headers
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Type parameter for response headers
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResponseHeaderGetParams]'s query parameters as
// `url.Values`.
func (r ResponseHeaderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
