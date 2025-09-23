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

// RedirectToService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRedirectToService] method instead.
type RedirectToService struct {
	Options []option.RequestOption
}

// NewRedirectToService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRedirectToService(opts ...option.RequestOption) (r RedirectToService) {
	r = RedirectToService{}
	r.Options = opts
	return
}

// Redirects the request to the target URL
func (r *RedirectToService) New(ctx context.Context, body RedirectToNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "redirect-to"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Redirects the request to the target URL
func (r *RedirectToService) Get(ctx context.Context, query RedirectToGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "redirect-to"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Redirects the request to the target URL
func (r *RedirectToService) Update(ctx context.Context, body RedirectToUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "redirect-to"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Redirects the request to the target URL
func (r *RedirectToService) Delete(ctx context.Context, body RedirectToDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "redirect-to"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Redirects the request to the target URL
func (r *RedirectToService) Modify(ctx context.Context, body RedirectToModifyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "redirect-to"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

type RedirectToNewParams struct {
	// The URL to redirect to
	URL string `query:"url,required" format:"uri" json:"-"`
	// The HTTP status code to use for redirection
	StatusCode param.Opt[string] `query:"status_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RedirectToNewParams]'s query parameters as `url.Values`.
func (r RedirectToNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RedirectToGetParams struct {
	// The URL to redirect to
	URL string `query:"url,required" format:"uri" json:"-"`
	// The HTTP status code to use for redirection
	StatusCode param.Opt[string] `query:"status_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RedirectToGetParams]'s query parameters as `url.Values`.
func (r RedirectToGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RedirectToUpdateParams struct {
	// The URL to redirect to
	URL string `query:"url,required" format:"uri" json:"-"`
	// The HTTP status code to use for redirection
	StatusCode param.Opt[string] `query:"status_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RedirectToUpdateParams]'s query parameters as `url.Values`.
func (r RedirectToUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RedirectToDeleteParams struct {
	// The URL to redirect to
	URL string `query:"url,required" format:"uri" json:"-"`
	// The HTTP status code to use for redirection
	StatusCode param.Opt[string] `query:"status_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RedirectToDeleteParams]'s query parameters as `url.Values`.
func (r RedirectToDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RedirectToModifyParams struct {
	// The URL to redirect to
	URL string `query:"url,required" format:"uri" json:"-"`
	// The HTTP status code to use for redirection
	StatusCode param.Opt[string] `query:"status_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RedirectToModifyParams]'s query parameters as `url.Values`.
func (r RedirectToModifyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
