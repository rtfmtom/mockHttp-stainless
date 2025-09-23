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

func TestDigestAuthGet(t *testing.T) {
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
	err := client.DigestAuth.Get(
		context.TODO(),
		"passwd",
		mockhttp.DigestAuthGetParams{
			Qop:  "qop",
			User: "user",
		},
	)
	if err != nil {
		var apierr *mockhttp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDigestAuthGetWithAlgorithm(t *testing.T) {
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
	err := client.DigestAuth.GetWithAlgorithm(
		context.TODO(),
		"algorithm",
		mockhttp.DigestAuthGetWithAlgorithmParams{
			Qop:    "qop",
			User:   "user",
			Passwd: "passwd",
		},
	)
	if err != nil {
		var apierr *mockhttp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
