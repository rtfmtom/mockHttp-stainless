// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mockhttp-go/internal/requestconfig"
	"github.com/stainless-sdks/mockhttp-go/option"
)

// DigestAuthService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDigestAuthService] method instead.
type DigestAuthService struct {
	Options []option.RequestOption
}

// NewDigestAuthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDigestAuthService(opts ...option.RequestOption) (r DigestAuthService) {
	r = DigestAuthService{}
	r.Options = opts
	return
}

// HTTP Digest authentication. Mirrors httpbin behavior for testing clients.
func (r *DigestAuthService) Get(ctx context.Context, passwd string, query DigestAuthGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if query.Qop == "" {
		err = errors.New("missing required qop parameter")
		return
	}
	if query.User == "" {
		err = errors.New("missing required user parameter")
		return
	}
	if passwd == "" {
		err = errors.New("missing required passwd parameter")
		return
	}
	path := fmt.Sprintf("digest-auth/%s/%s/%s", query.Qop, query.User, passwd)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// HTTP Digest authentication. Mirrors httpbin behavior for testing clients.
func (r *DigestAuthService) GetWithAlgorithm(ctx context.Context, algorithm string, query DigestAuthGetWithAlgorithmParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if query.Qop == "" {
		err = errors.New("missing required qop parameter")
		return
	}
	if query.User == "" {
		err = errors.New("missing required user parameter")
		return
	}
	if query.Passwd == "" {
		err = errors.New("missing required passwd parameter")
		return
	}
	if algorithm == "" {
		err = errors.New("missing required algorithm parameter")
		return
	}
	path := fmt.Sprintf("digest-auth/%s/%s/%s/%s", query.Qop, query.User, query.Passwd, algorithm)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type DigestAuthGetParams struct {
	Qop  string `path:"qop,required" json:"-"`
	User string `path:"user,required" json:"-"`
	paramObj
}

type DigestAuthGetWithAlgorithmParams struct {
	Qop    string `path:"qop,required" json:"-"`
	User   string `path:"user,required" json:"-"`
	Passwd string `path:"passwd,required" json:"-"`
	paramObj
}
