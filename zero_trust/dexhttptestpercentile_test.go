// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
  "context"
  "errors"
  "os"
  "testing"

  "github.com/cloudflare/cloudflare-go/v2"
  "github.com/cloudflare/cloudflare-go/v2/internal/testutil"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDEXHTTPTestPercentileGetWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := cloudflare.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
    option.WithAPIEmail("user@example.com"),
  )
  _, err := client.ZeroTrust.DEX.HTTPTests.Percentiles.Get(
    context.TODO(),
    "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
    zero_trust.DEXHTTPTestPercentileGetParams{
      AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
      From: cloudflare.F("2023-09-20T17:00:00Z"),
      To: cloudflare.F("2023-09-20T17:00:00Z"),
      Colo: cloudflare.F("colo"),
      DeviceID: cloudflare.F([]string{"string", "string", "string"}),
    },
  )
  if err != nil {
    var apierr *cloudflare.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
