// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mockhttp_test

import (
	"context"
	"os"
	"testing"

	"github.com/rtfmtom/mockHttp-stainless"
	"github.com/rtfmtom/mockHttp-stainless/internal/testutil"
	"github.com/rtfmtom/mockHttp-stainless/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Get.RequestInfo(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Headers)
}
