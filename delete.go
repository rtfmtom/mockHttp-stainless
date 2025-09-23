// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/apijson"
	shimjson "github.com/stainless-sdks/mockhttp-go/internal/encoding/json"
	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
	"github.com/stainless-sdks/mockhttp-go/packages/respjson"
)

// DeleteService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteService] method instead.
type DeleteService struct {
	Options []option.RequestOption
}

// NewDeleteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDeleteService(opts ...option.RequestOption) (r DeleteService) {
	r = DeleteService{}
	r.Options = opts
	return
}

// Handles a DELETE request and returns request information
func (r *DeleteService) Perform(ctx context.Context, body DeletePerformParams, opts ...option.RequestOption) (res *DeletePerformResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type DeletePerformResponse struct {
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
func (r DeletePerformResponse) RawJSON() string { return r.JSON.raw }
func (r *DeletePerformResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeletePerformParams struct {
	// The body of the DELETE request
	Body map[string]any
	paramObj
}

func (r DeletePerformParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *DeletePerformParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
