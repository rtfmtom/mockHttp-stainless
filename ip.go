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

// IPService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r IPService) {
	r = IPService{}
	r.Options = opts
	return
}

// Returns the IP address of the client
func (r *IPService) Get(ctx context.Context, opts ...option.RequestOption) (res *IPGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IPGetResponse struct {
	// The IP address of the client
	IP string `json:"ip,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IP          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IPGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
