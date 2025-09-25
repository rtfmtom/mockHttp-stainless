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

// AnythingService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnythingService] method instead.
type AnythingService struct {
	Options []option.RequestOption
}

// NewAnythingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnythingService(opts ...option.RequestOption) (r AnythingService) {
	r = AnythingService{}
	r.Options = opts
	return
}

// Will return anything that you send it
func (r *AnythingService) New(ctx context.Context, opts ...option.RequestOption) (res *AnythingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "anything"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Will return anything that you send it
func (r *AnythingService) Update(ctx context.Context, opts ...option.RequestOption) (res *AnythingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "anything"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Will return anything that you send it
func (r *AnythingService) List(ctx context.Context, opts ...option.RequestOption) (res *AnythingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "anything"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Will return anything that you send it
func (r *AnythingService) Delete(ctx context.Context, opts ...option.RequestOption) (res *AnythingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "anything"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Will return anything that you send it
func (r *AnythingService) Patch(ctx context.Context, opts ...option.RequestOption) (res *AnythingPatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "anything"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type AnythingNewResponse struct {
	// Route parameters
	Args map[string]any `json:"args"`
	// Request body data
	Data map[string]any `json:"data"`
	// Uploaded files
	Files map[string]any `json:"files"`
	// Form data
	Form map[string]any `json:"form"`
	// Request headers
	Headers map[string]any `json:"headers"`
	// Parsed JSON body
	Json map[string]any `json:"json"`
	// HTTP method used
	Method string `json:"method"`
	// Origin of the request
	Origin string `json:"origin"`
	// Request URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Data        respjson.Field
		Files       respjson.Field
		Form        respjson.Field
		Headers     respjson.Field
		Json        respjson.Field
		Method      respjson.Field
		Origin      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnythingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AnythingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnythingUpdateResponse struct {
	// Route parameters
	Args map[string]any `json:"args"`
	// Request body data
	Data map[string]any `json:"data"`
	// Uploaded files
	Files map[string]any `json:"files"`
	// Form data
	Form map[string]any `json:"form"`
	// Request headers
	Headers map[string]any `json:"headers"`
	// Parsed JSON body
	Json map[string]any `json:"json"`
	// HTTP method used
	Method string `json:"method"`
	// Origin of the request
	Origin string `json:"origin"`
	// Request URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Data        respjson.Field
		Files       respjson.Field
		Form        respjson.Field
		Headers     respjson.Field
		Json        respjson.Field
		Method      respjson.Field
		Origin      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnythingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AnythingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnythingListResponse struct {
	// Route parameters
	Args map[string]any `json:"args"`
	// Request body data
	Data map[string]any `json:"data"`
	// Uploaded files
	Files map[string]any `json:"files"`
	// Form data
	Form map[string]any `json:"form"`
	// Request headers
	Headers map[string]any `json:"headers"`
	// Parsed JSON body
	Json map[string]any `json:"json"`
	// HTTP method used
	Method string `json:"method"`
	// Origin of the request
	Origin string `json:"origin"`
	// Request URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Data        respjson.Field
		Files       respjson.Field
		Form        respjson.Field
		Headers     respjson.Field
		Json        respjson.Field
		Method      respjson.Field
		Origin      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnythingListResponse) RawJSON() string { return r.JSON.raw }
func (r *AnythingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnythingDeleteResponse struct {
	// Route parameters
	Args map[string]any `json:"args"`
	// Request body data
	Data map[string]any `json:"data"`
	// Uploaded files
	Files map[string]any `json:"files"`
	// Form data
	Form map[string]any `json:"form"`
	// Request headers
	Headers map[string]any `json:"headers"`
	// Parsed JSON body
	Json map[string]any `json:"json"`
	// HTTP method used
	Method string `json:"method"`
	// Origin of the request
	Origin string `json:"origin"`
	// Request URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Data        respjson.Field
		Files       respjson.Field
		Form        respjson.Field
		Headers     respjson.Field
		Json        respjson.Field
		Method      respjson.Field
		Origin      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnythingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AnythingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnythingPatchResponse struct {
	// Route parameters
	Args map[string]any `json:"args"`
	// Request body data
	Data map[string]any `json:"data"`
	// Uploaded files
	Files map[string]any `json:"files"`
	// Form data
	Form map[string]any `json:"form"`
	// Request headers
	Headers map[string]any `json:"headers"`
	// Parsed JSON body
	Json map[string]any `json:"json"`
	// HTTP method used
	Method string `json:"method"`
	// Origin of the request
	Origin string `json:"origin"`
	// Request URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Data        respjson.Field
		Files       respjson.Field
		Form        respjson.Field
		Headers     respjson.Field
		Json        respjson.Field
		Method      respjson.Field
		Origin      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnythingPatchResponse) RawJSON() string { return r.JSON.raw }
func (r *AnythingPatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
