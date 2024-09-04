// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
  "context"
  "errors"
  "os"
  "testing"

  "github.com/cloudflare/cloudflare-go/v2"
  "github.com/cloudflare/cloudflare-go/v2/internal/testutil"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestRobotsTXTTopDirectiveGetWithOptionalParams(t *testing.T) {
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
  _, err := client.Radar.RobotsTXT.Top.Directive.Get(
    context.TODO(),
    radar.RobotsTXTTopDirectiveGetParamsDirectiveAllow,
    radar.RobotsTXTTopDirectiveGetParams{
      AgentCategory: cloudflare.F(radar.RobotsTXTTopDirectiveGetParamsAgentCategoryAI),
      Date: cloudflare.F("2024-09-19"),
      Format: cloudflare.F(radar.RobotsTXTTopDirectiveGetParamsFormatJson),
      Limit: cloudflare.F(int64(5)),
      Name: cloudflare.F([]string{"string", "string", "string"}),
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
