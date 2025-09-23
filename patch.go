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

// PatchService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPatchService] method instead.
type PatchService struct {
	Options []option.RequestOption
}

// NewPatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPatchService(opts ...option.RequestOption) (r PatchService) {
	r = PatchService{}
	r.Options = opts
	return
}

// Handles a PATCH request and returns request information
func (r *PatchService) Handle(ctx context.Context, body PatchHandleParams, opts ...option.RequestOption) (res *PatchHandleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "patch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type PatchHandleResponse struct {
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
func (r PatchHandleResponse) RawJSON() string { return r.JSON.raw }
func (r *PatchHandleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PatchHandleParams struct {
	// The body of the PATCH request
	Body map[string]any
	paramObj
}

func (r PatchHandleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PatchHandleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
