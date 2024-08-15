// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnalyticsEventBytimeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsEventBytimeService] method instead.
type AnalyticsEventBytimeService struct {
	Options []option.RequestOption
}

// NewAnalyticsEventBytimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsEventBytimeService(opts ...option.RequestOption) (r *AnalyticsEventBytimeService) {
	r = &AnalyticsEventBytimeService{}
	r.Options = opts
	return
}
