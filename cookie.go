// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/mockhttp-go/internal/apijson"
	"github.com/stainless-sdks/mockhttp-go/internal/apiquery"
	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
	"github.com/stainless-sdks/mockhttp-go/packages/param"
	"github.com/stainless-sdks/mockhttp-go/packages/respjson"
)

// CookieService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCookieService] method instead.
type CookieService struct {
	Options []option.RequestOption
}

// NewCookieService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCookieService(opts ...option.RequestOption) (r CookieService) {
	r = CookieService{}
	r.Options = opts
	return
}

// Set a cookie
func (r *CookieService) New(ctx context.Context, body CookieNewParams, opts ...option.RequestOption) (res *CookieNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cookies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return cookies from the request
func (r *CookieService) List(ctx context.Context, opts ...option.RequestOption) (res *CookieListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cookies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a cookie
func (r *CookieService) Delete(ctx context.Context, body CookieDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cookies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CookieNewResponse struct {
	Cookies map[string]string `json:"cookies,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CookieNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CookieNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CookieListResponse struct {
	Cookies map[string]string `json:"cookies,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CookieListResponse) RawJSON() string { return r.JSON.raw }
func (r *CookieListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CookieNewParams struct {
	Name    string               `json:"name,required"`
	Value   string               `json:"value,required"`
	Expires param.Opt[time.Time] `json:"expires,omitzero" format:"date-time"`
	paramObj
}

func (r CookieNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CookieNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CookieNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CookieDeleteParams struct {
	Name string `query:"name,required" json:"-"`
	paramObj
}

// URLQuery serializes [CookieDeleteParams]'s query parameters as `url.Values`.
func (r CookieDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
