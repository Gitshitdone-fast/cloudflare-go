// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// TeamnetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTeamnetService] method instead.
type TeamnetService struct {
	Options         []option.RequestOption
	Routes          *TeamnetRouteService
	VirtualNetworks *TeamnetVirtualNetworkService
}

// NewTeamnetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamnetService(opts ...option.RequestOption) (r *TeamnetService) {
	r = &TeamnetService{}
	r.Options = opts
	r.Routes = NewTeamnetRouteService(opts...)
	r.VirtualNetworks = NewTeamnetVirtualNetworkService(opts...)
	return
}
