// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// MagicService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewMagicService] method instead.
type MagicService struct {
	Options         []option.RequestOption
	CfInterconnects *MagicCfInterconnectService
	GreTunnels      *MagicGreTunnelService
	IpsecTunnels    *MagicIpsecTunnelService
	Routes          *MagicRouteService
}

// NewMagicService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMagicService(opts ...option.RequestOption) (r *MagicService) {
	r = &MagicService{}
	r.Options = opts
	r.CfInterconnects = NewMagicCfInterconnectService(opts...)
	r.GreTunnels = NewMagicGreTunnelService(opts...)
	r.IpsecTunnels = NewMagicIpsecTunnelService(opts...)
	r.Routes = NewMagicRouteService(opts...)
	return
}
