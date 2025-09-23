// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/mockhttp-go"
	"github.com/stainless-sdks/mockhttp-go/internal/testutil"
	"github.com/stainless-sdks/mockhttp-go/option"
)

func TestResponseHeaderNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mockhttp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ResponseHeaders.New(context.TODO(), mockhttp.ResponseHeaderNewParams{
		AdditionalProperties: mockhttp.String("additionalProperties"),
		Description:          mockhttp.String("description"),
		Type:                 mockhttp.String("type"),
	})
	if err != nil {
		var apierr *mockhttp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponseHeaderGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mockhttp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ResponseHeaders.Get(context.TODO(), mockhttp.ResponseHeaderGetParams{
		AdditionalProperties: mockhttp.Bool(true),
		Description:          mockhttp.String("description"),
		Type:                 mockhttp.String("type"),
	})
	if err != nil {
		var apierr *mockhttp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
