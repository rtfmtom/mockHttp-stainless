// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/rtfmtom/mockHttp-stainless/internal/apijson"
	shimjson "github.com/rtfmtom/mockHttp-stainless/internal/encoding/json"
	"github.com/rtfmtom/mockHttp-stainless/internal/requestconfig"
	"github.com/rtfmtom/mockHttp-stainless/option"
	"github.com/rtfmtom/mockHttp-stainless/packages/respjson"
)

// PostService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostService] method instead.
type PostService struct {
	Options []option.RequestOption
}

// NewPostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPostService(opts ...option.RequestOption) (r PostService) {
	r = PostService{}
	r.Options = opts
	return
}

// Handles a POST request and returns request information
func (r *PostService) New(ctx context.Context, body PostNewParams, opts ...option.RequestOption) (res *PostNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "post"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PostNewResponse struct {
	Body    map[string]any    `json:"body,required"`
	Headers map[string]string `json:"headers,required"`
	Method  string            `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PostNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostNewParams struct {
	// The body of the POST request
	Body map[string]any
	paramObj
}

func (r PostNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PostNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
