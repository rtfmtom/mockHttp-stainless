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

// EtagService contains methods and other services that help with interacting with
// the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEtagService] method instead.
type EtagService struct {
	Options []option.RequestOption
}

// NewEtagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEtagService(opts ...option.RequestOption) (r EtagService) {
	r = EtagService{}
	r.Options = opts
	return
}

// Handles ETag-based conditional requests.
func (r *EtagService) Get(ctx context.Context, etag string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if etag == "" {
		err = errors.New("missing required etag parameter")
		return
	}
	path := fmt.Sprintf("etag/%s", etag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
