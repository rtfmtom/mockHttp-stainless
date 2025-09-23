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

// StatusService contains methods and other services that help with interacting
// with the mockhttp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusService] method instead.
type StatusService struct {
	Options []option.RequestOption
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r StatusService) {
	r = StatusService{}
	r.Options = opts
	return
}

// Return status code or random status code if more than one are given
func (r *StatusService) New(ctx context.Context, code string, opts ...option.RequestOption) (res *StatusNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("status/%s?", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Return status code or random status code if more than one are given
func (r *StatusService) Get(ctx context.Context, code string, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("status/%s?", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return status code or random status code if more than one are given
func (r *StatusService) Update(ctx context.Context, code string, opts ...option.RequestOption) (res *StatusUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("status/%s?", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Return status code or random status code if more than one are given
func (r *StatusService) Delete(ctx context.Context, code string, opts ...option.RequestOption) (res *StatusDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("status/%s?", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Return status code or random status code if more than one are given
func (r *StatusService) Patch(ctx context.Context, code string, opts ...option.RequestOption) (res *StatusPatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("status/%s?", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type StatusNewResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusUpdateResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusDeleteResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusPatchResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusPatchResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusPatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
