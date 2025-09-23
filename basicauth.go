// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/apijson"
	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
	"github.com/stainless-sdks/mockhttp-go/packages/respjson"
)

// BasicAuthService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBasicAuthService] method instead.
type BasicAuthService struct {
	Options []option.RequestOption
}

// NewBasicAuthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBasicAuthService(opts ...option.RequestOption) (r BasicAuthService) {
	r = BasicAuthService{}
	r.Options = opts
	return
}

// HTTP Basic authentication. Succeeds only if the user/pass provided in the path
// matches the Basic Authorization header.
func (r *BasicAuthService) Authenticate(ctx context.Context, passwd string, query BasicAuthAuthenticateParams, opts ...option.RequestOption) (res *BasicAuthAuthenticateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.User == "" {
		err = errors.New("missing required user parameter")
		return
	}
	if passwd == "" {
		err = errors.New("missing required passwd parameter")
		return
	}
	path := fmt.Sprintf("basic-auth/%s/%s", query.User, passwd)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BasicAuthAuthenticateResponse struct {
	Authenticated bool   `json:"authenticated,required"`
	User          string `json:"user,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authenticated respjson.Field
		User          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BasicAuthAuthenticateResponse) RawJSON() string { return r.JSON.raw }
func (r *BasicAuthAuthenticateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BasicAuthAuthenticateParams struct {
	User string `path:"user,required" json:"-"`
	paramObj
}
