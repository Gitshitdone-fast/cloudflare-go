// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// GatewayLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayLocationService] method
// instead.
type GatewayLocationService struct {
	Options []option.RequestOption
}

// NewGatewayLocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayLocationService(opts ...option.RequestOption) (r *GatewayLocationService) {
	r = &GatewayLocationService{}
	r.Options = opts
	return
}

// Updates a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Update(ctx context.Context, accountID interface{}, locationID interface{}, body GatewayLocationUpdateParams, opts ...option.RequestOption) (res *GatewayLocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", accountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Delete(ctx context.Context, accountID interface{}, locationID interface{}, opts ...option.RequestOption) (res *GatewayLocationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", accountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway location.
func (r *GatewayLocationService) Get(ctx context.Context, accountID interface{}, locationID interface{}, opts ...option.RequestOption) (res *GatewayLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", accountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Zero Trust Gateway location.
func (r *GatewayLocationService) ZeroTrustGatewayLocationsNewZeroTrustGatewayLocation(ctx context.Context, accountID interface{}, body GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams, opts ...option.RequestOption) (res *GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zero Trust Gateway locations for an account.
func (r *GatewayLocationService) ZeroTrustGatewayLocationsListZeroTrustGatewayLocations(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayLocationUpdateResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []GatewayLocationUpdateResponseNetwork `json:"networks"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      gatewayLocationUpdateResponseJSON      `json:"-"`
}

// gatewayLocationUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayLocationUpdateResponse]
type gatewayLocationUpdateResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationUpdateResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                   `json:"network,required"`
	JSON    gatewayLocationUpdateResponseNetworkJSON `json:"-"`
}

// gatewayLocationUpdateResponseNetworkJSON contains the JSON metadata for the
// struct [GatewayLocationUpdateResponseNetwork]
type gatewayLocationUpdateResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [GatewayLocationDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayLocationDeleteResponse interface {
	ImplementsGatewayLocationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayLocationDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayLocationGetResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []GatewayLocationGetResponseNetwork `json:"networks"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      gatewayLocationGetResponseJSON      `json:"-"`
}

// gatewayLocationGetResponseJSON contains the JSON metadata for the struct
// [GatewayLocationGetResponse]
type gatewayLocationGetResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationGetResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                `json:"network,required"`
	JSON    gatewayLocationGetResponseNetworkJSON `json:"-"`
}

// gatewayLocationGetResponseNetworkJSON contains the JSON metadata for the struct
// [GatewayLocationGetResponseNetwork]
type gatewayLocationGetResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetwork `json:"networks"`
	UpdatedAt time.Time                                                                            `json:"updated_at" format:"date-time"`
	JSON      gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseJSON      `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse]
type gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                                                                 `json:"network,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetworkJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetworkJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetwork]
type gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetwork `json:"networks"`
	UpdatedAt time.Time                                                                              `json:"updated_at" format:"date-time"`
	JSON      gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseJSON      `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse]
type gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                                                                   `json:"network,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetworkJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetworkJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetwork]
type gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationUpdateParams struct {
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]GatewayLocationUpdateParamsNetwork] `json:"networks"`
}

func (r GatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateResponseEnvelope struct {
	Errors   []GatewayLocationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLocationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationUpdateResponseEnvelope]
type gatewayLocationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayLocationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayLocationUpdateResponseEnvelopeErrors]
type gatewayLocationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayLocationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationUpdateResponseEnvelopeMessages]
type gatewayLocationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLocationUpdateResponseEnvelopeSuccess bool

const (
	GatewayLocationUpdateResponseEnvelopeSuccessTrue GatewayLocationUpdateResponseEnvelopeSuccess = true
)

type GatewayLocationDeleteResponseEnvelope struct {
	Errors   []GatewayLocationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLocationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationDeleteResponseEnvelope]
type gatewayLocationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayLocationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayLocationDeleteResponseEnvelopeErrors]
type gatewayLocationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayLocationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationDeleteResponseEnvelopeMessages]
type gatewayLocationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLocationDeleteResponseEnvelopeSuccess bool

const (
	GatewayLocationDeleteResponseEnvelopeSuccessTrue GatewayLocationDeleteResponseEnvelopeSuccess = true
)

type GatewayLocationGetResponseEnvelope struct {
	Errors   []GatewayLocationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLocationGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLocationGetResponseEnvelope]
type gatewayLocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    gatewayLocationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayLocationGetResponseEnvelopeErrors]
type gatewayLocationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayLocationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationGetResponseEnvelopeMessages]
type gatewayLocationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLocationGetResponseEnvelopeSuccess bool

const (
	GatewayLocationGetResponseEnvelopeSuccessTrue GatewayLocationGetResponseEnvelopeSuccess = true
)

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams struct {
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParamsNetwork] `json:"networks"`
}

func (r GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelope struct {
	Errors   []GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelope]
type gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrors]
type gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessages]
type gatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeSuccess bool

const (
	GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeSuccessTrue GatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseEnvelopeSuccess = true
)

type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelope struct {
	Errors   []GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeJSON       `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelope]
type gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrors struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrors]
type gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessages struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessages]
type gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeSuccess bool

const (
	GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeSuccessTrue GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeSuccess = true
)

type GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                             `json:"total_count"`
	JSON       gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfo]
type gatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
