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

// UserAgentService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAgentService] method instead.
type UserAgentService struct {
	Options []option.RequestOption
}

// NewUserAgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAgentService(opts ...option.RequestOption) (r UserAgentService) {
	r = UserAgentService{}
	r.Options = opts
	return
}

// Returns the User-Agent header of the client request
func (r *UserAgentService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserAgentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user-agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserAgentGetResponse struct {
	// The User-Agent header of the client request
	UserAgent string `json:"userAgent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserAgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
