// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/rtfmtom/mockHttp-stainless"
	"github.com/rtfmtom/mockHttp-stainless/internal/testutil"
	"github.com/rtfmtom/mockHttp-stainless/option"
)

func TestCacheGet(t *testing.T) {
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
	_, err := client.Cache.Get(context.TODO())
	if err != nil {
		var apierr *mockhttp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCacheSetCacheControl(t *testing.T) {
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
	_, err := client.Cache.SetCacheControl(context.TODO(), 0)
	if err != nil {
		var apierr *mockhttp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
