// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerPoolPreviewService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerPoolPreviewService] method instead.
type UserLoadBalancerPoolPreviewService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPoolPreviewService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolPreviewService(opts ...option.RequestOption) (r *UserLoadBalancerPoolPreviewService) {
	r = &UserLoadBalancerPoolPreviewService{}
	r.Options = opts
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *UserLoadBalancerPoolPreviewService) LoadBalancerPoolsPreviewPool(ctx context.Context, identifier string, body UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/preview", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                                                   `json:"pools"`
	PreviewID string                                                              `json:"preview_id"`
	JSON      userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseJSON `json:"-"`
}

// userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponse]
type userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType] `json:"type"`
}

func (r UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType string

const (
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeHTTP     UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "http"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeHTTPs    UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "https"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeTcp      UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "tcp"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeUdpIcmp  UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "udp_icmp"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeIcmpPing UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "icmp_ping"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeSmtp     UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "smtp"
)

type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelope]
type userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrors]
type userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessages]
type userLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeSuccessTrue UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolResponseEnvelopeSuccess = true
)
