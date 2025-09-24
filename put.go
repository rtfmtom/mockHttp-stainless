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

// PutService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPutService] method instead.
type PutService struct {
	Options []option.RequestOption
}

// NewPutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPutService(opts ...option.RequestOption) (r PutService) {
	r = PutService{}
	r.Options = opts
	return
}

// Handles a PUT request and returns request information
func (r *PutService) Handle(ctx context.Context, body PutHandleParams, opts ...option.RequestOption) (res *PutHandleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "put"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PutHandleResponse struct {
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
func (r PutHandleResponse) RawJSON() string { return r.JSON.raw }
func (r *PutHandleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PutHandleParams struct {
	// The body of the PUT request
	Body map[string]any
	paramObj
}

func (r PutHandleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PutHandleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
