// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// GatewayRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayRuleService] method
// instead.
type GatewayRuleService struct {
	Options []option.RequestOption
}

// NewGatewayRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayRuleService(opts ...option.RequestOption) (r *GatewayRuleService) {
	r = &GatewayRuleService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway rule.
func (r *GatewayRuleService) New(ctx context.Context, params GatewayRuleNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *GatewayRuleService) Update(ctx context.Context, ruleID string, params GatewayRuleUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) List(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustGatewayRules], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/rules", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) ListAutoPaging(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustGatewayRules] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Zero Trust Gateway rule.
func (r *GatewayRuleService) Delete(ctx context.Context, ruleID string, params GatewayRuleDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *GatewayRuleService) Get(ctx context.Context, ruleID string, query GatewayRuleGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSResolverSettingsV4 struct {
	// IPv4 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                    `json:"vnet_id"`
	JSON   dnsResolverSettingsV4JSON `json:"-"`
}

// dnsResolverSettingsV4JSON contains the JSON metadata for the struct
// [DNSResolverSettingsV4]
type dnsResolverSettingsV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DNSResolverSettingsV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResolverSettingsV4JSON) RawJSON() string {
	return r.raw
}

type DNSResolverSettingsV4Param struct {
	// IPv4 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r DNSResolverSettingsV4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSResolverSettingsV6 struct {
	// IPv6 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                    `json:"vnet_id"`
	JSON   dnsResolverSettingsV6JSON `json:"-"`
}

// dnsResolverSettingsV6JSON contains the JSON metadata for the struct
// [DNSResolverSettingsV6]
type dnsResolverSettingsV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DNSResolverSettingsV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResolverSettingsV6JSON) RawJSON() string {
	return r.raw
}

type DNSResolverSettingsV6Param struct {
	// IPv6 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r DNSResolverSettingsV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol or layer to use.
type FitlerItem string

const (
	FitlerItemHTTP   FitlerItem = "http"
	FitlerItemDNS    FitlerItem = "dns"
	FitlerItemL4     FitlerItem = "l4"
	FitlerItemEgress FitlerItem = "egress"
)

func (r FitlerItem) IsKnown() bool {
	switch r {
	case FitlerItemHTTP, FitlerItemDNS, FitlerItemL4, FitlerItemEgress:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type RuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH RuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls RuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession RuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin. Only valid when a rule's action is
	// set to 'resolve'.
	DNSResolvers RuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress RuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override RuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings RuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog RuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified. Only valid when a
	// rule's action is set to 'resolve'.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT RuleSettingsUntrustedCERT `json:"untrusted_cert"`
	JSON          ruleSettingsJSON          `json:"-"`
}

// ruleSettingsJSON contains the JSON metadata for the struct [RuleSettings]
type ruleSettingsJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BisoAdminControls               apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	DNSResolvers                    apijson.Field
	Egress                          apijson.Field
	InsecureDisableDNSSECValidation apijson.Field
	IPCategories                    apijson.Field
	IPIndicatorFeeds                apijson.Field
	L4override                      apijson.Field
	NotificationSettings            apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	ResolveDNSThroughCloudflare     apijson.Field
	UntrustedCERT                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *RuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the Audit SSH action.
type RuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                     `json:"command_logging"`
	JSON           ruleSettingsAuditSSHJSON `json:"-"`
}

// ruleSettingsAuditSSHJSON contains the JSON metadata for the struct
// [RuleSettingsAuditSSH]
type ruleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure how browser isolation behaves.
type RuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                              `json:"du"`
	JSON ruleSettingsBisoAdminControlsJSON `json:"-"`
}

// ruleSettingsBisoAdminControlsJSON contains the JSON metadata for the struct
// [RuleSettingsBisoAdminControls]
type ruleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure how session check behaves.
type RuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                         `json:"enforce"`
	JSON    ruleSettingsCheckSessionJSON `json:"-"`
}

// ruleSettingsCheckSessionJSON contains the JSON metadata for the struct
// [RuleSettingsCheckSession]
type ruleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin. Only valid when a rule's action is
// set to 'resolve'.
type RuleSettingsDNSResolvers struct {
	IPV4 []DNSResolverSettingsV4      `json:"ipv4"`
	IPV6 []DNSResolverSettingsV6      `json:"ipv6"`
	JSON ruleSettingsDNSResolversJSON `json:"-"`
}

// ruleSettingsDNSResolversJSON contains the JSON metadata for the struct
// [RuleSettingsDNSResolvers]
type ruleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type RuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                 `json:"ipv6"`
	JSON ruleSettingsEgressJSON `json:"-"`
}

// ruleSettingsEgressJSON contains the JSON metadata for the struct
// [RuleSettingsEgress]
type ruleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port.
type RuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                      `json:"port"`
	JSON ruleSettingsL4overrideJSON `json:"-"`
}

// ruleSettingsL4overrideJSON contains the JSON metadata for the struct
// [RuleSettingsL4override]
type ruleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule is
// matched.
type RuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                               `json:"support_url"`
	JSON       ruleSettingsNotificationSettingsJSON `json:"-"`
}

// ruleSettingsNotificationSettingsJSON contains the JSON metadata for the struct
// [RuleSettingsNotificationSettings]
type ruleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging.
type RuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                       `json:"enabled"`
	JSON    ruleSettingsPayloadLogJSON `json:"-"`
}

// ruleSettingsPayloadLogJSON contains the JSON metadata for the struct
// [RuleSettingsPayloadLog]
type ruleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type RuleSettingsUntrustedCERT struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action RuleSettingsUntrustedCERTAction `json:"action"`
	JSON   ruleSettingsUntrustedCERTJSON   `json:"-"`
}

// ruleSettingsUntrustedCERTJSON contains the JSON metadata for the struct
// [RuleSettingsUntrustedCERT]
type ruleSettingsUntrustedCERTJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsUntrustedCERT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsUntrustedCERTJSON) RawJSON() string {
	return r.raw
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type RuleSettingsUntrustedCERTAction string

const (
	RuleSettingsUntrustedCERTActionPassThrough RuleSettingsUntrustedCERTAction = "pass_through"
	RuleSettingsUntrustedCERTActionBlock       RuleSettingsUntrustedCERTAction = "block"
	RuleSettingsUntrustedCERTActionError       RuleSettingsUntrustedCERTAction = "error"
)

func (r RuleSettingsUntrustedCERTAction) IsKnown() bool {
	switch r {
	case RuleSettingsUntrustedCERTActionPassThrough, RuleSettingsUntrustedCERTActionBlock, RuleSettingsUntrustedCERTActionError:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type RuleSettingsParam struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[RuleSettingsAuditSSHParam] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[RuleSettingsBisoAdminControlsParam] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[RuleSettingsCheckSessionParam] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin. Only valid when a rule's action is
	// set to 'resolve'.
	DNSResolvers param.Field[RuleSettingsDNSResolversParam] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[RuleSettingsEgressParam] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[RuleSettingsL4overrideParam] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[RuleSettingsNotificationSettingsParam] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[RuleSettingsPayloadLogParam] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified. Only valid when a
	// rule's action is set to 'resolve'.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT param.Field[RuleSettingsUntrustedCERTParam] `json:"untrusted_cert"`
}

func (r RuleSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type RuleSettingsAuditSSHParam struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r RuleSettingsAuditSSHParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type RuleSettingsBisoAdminControlsParam struct {
	// Set to true to enable copy-pasting.
	Dcp param.Field[bool] `json:"dcp"`
	// Set to true to enable downloading.
	Dd param.Field[bool] `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk param.Field[bool] `json:"dk"`
	// Set to true to enable printing.
	Dp param.Field[bool] `json:"dp"`
	// Set to true to enable uploading.
	Du param.Field[bool] `json:"du"`
}

func (r RuleSettingsBisoAdminControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type RuleSettingsCheckSessionParam struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r RuleSettingsCheckSessionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin. Only valid when a rule's action is
// set to 'resolve'.
type RuleSettingsDNSResolversParam struct {
	IPV4 param.Field[[]DNSResolverSettingsV4Param] `json:"ipv4"`
	IPV6 param.Field[[]DNSResolverSettingsV6Param] `json:"ipv6"`
}

func (r RuleSettingsDNSResolversParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type RuleSettingsEgressParam struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r RuleSettingsEgressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type RuleSettingsL4overrideParam struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r RuleSettingsL4overrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type RuleSettingsNotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r RuleSettingsNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type RuleSettingsPayloadLogParam struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleSettingsPayloadLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type RuleSettingsUntrustedCERTParam struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[RuleSettingsUntrustedCERTAction] `json:"action"`
}

func (r RuleSettingsUntrustedCERTParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type Schedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri string `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon string `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat string `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun string `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu string `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone string `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue string `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed  string       `json:"wed"`
	JSON scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	Fri         apijson.Field
	Mon         apijson.Field
	Sat         apijson.Field
	Sun         apijson.Field
	Thu         apijson.Field
	TimeZone    apijson.Field
	Tue         apijson.Field
	Wed         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type ScheduleParam struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri param.Field[string] `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon param.Field[string] `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat param.Field[string] `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun param.Field[string] `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu param.Field[string] `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone param.Field[string] `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue param.Field[string] `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed param.Field[string] `json:"wed"`
}

func (r ScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRules struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    ZeroTrustGatewayRulesAction `json:"action"`
	CreatedAt time.Time                   `json:"created_at" format:"date-time"`
	// Date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The description of the rule.
	Description string `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture string `json:"device_posture"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters []ZeroTrustGatewayRulesFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings RuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule Schedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                    `json:"traffic"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayRulesJSON `json:"-"`
}

// zeroTrustGatewayRulesJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayRules]
type zeroTrustGatewayRulesJSON struct {
	ID            apijson.Field
	Action        apijson.Field
	CreatedAt     apijson.Field
	DeletedAt     apijson.Field
	Description   apijson.Field
	DevicePosture apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Identity      apijson.Field
	Name          apijson.Field
	Precedence    apijson.Field
	RuleSettings  apijson.Field
	Schedule      apijson.Field
	Traffic       apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesJSON) RawJSON() string {
	return r.raw
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type ZeroTrustGatewayRulesAction string

const (
	ZeroTrustGatewayRulesActionOn           ZeroTrustGatewayRulesAction = "on"
	ZeroTrustGatewayRulesActionOff          ZeroTrustGatewayRulesAction = "off"
	ZeroTrustGatewayRulesActionAllow        ZeroTrustGatewayRulesAction = "allow"
	ZeroTrustGatewayRulesActionBlock        ZeroTrustGatewayRulesAction = "block"
	ZeroTrustGatewayRulesActionScan         ZeroTrustGatewayRulesAction = "scan"
	ZeroTrustGatewayRulesActionNoscan       ZeroTrustGatewayRulesAction = "noscan"
	ZeroTrustGatewayRulesActionSafesearch   ZeroTrustGatewayRulesAction = "safesearch"
	ZeroTrustGatewayRulesActionYtrestricted ZeroTrustGatewayRulesAction = "ytrestricted"
	ZeroTrustGatewayRulesActionIsolate      ZeroTrustGatewayRulesAction = "isolate"
	ZeroTrustGatewayRulesActionNoisolate    ZeroTrustGatewayRulesAction = "noisolate"
	ZeroTrustGatewayRulesActionOverride     ZeroTrustGatewayRulesAction = "override"
	ZeroTrustGatewayRulesActionL4Override   ZeroTrustGatewayRulesAction = "l4_override"
	ZeroTrustGatewayRulesActionEgress       ZeroTrustGatewayRulesAction = "egress"
	ZeroTrustGatewayRulesActionAuditSSH     ZeroTrustGatewayRulesAction = "audit_ssh"
	ZeroTrustGatewayRulesActionResolve      ZeroTrustGatewayRulesAction = "resolve"
)

func (r ZeroTrustGatewayRulesAction) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayRulesActionOn, ZeroTrustGatewayRulesActionOff, ZeroTrustGatewayRulesActionAllow, ZeroTrustGatewayRulesActionBlock, ZeroTrustGatewayRulesActionScan, ZeroTrustGatewayRulesActionNoscan, ZeroTrustGatewayRulesActionSafesearch, ZeroTrustGatewayRulesActionYtrestricted, ZeroTrustGatewayRulesActionIsolate, ZeroTrustGatewayRulesActionNoisolate, ZeroTrustGatewayRulesActionOverride, ZeroTrustGatewayRulesActionL4Override, ZeroTrustGatewayRulesActionEgress, ZeroTrustGatewayRulesActionAuditSSH, ZeroTrustGatewayRulesActionResolve:
		return true
	}
	return false
}

// The protocol or layer to use.
type ZeroTrustGatewayRulesFilter string

const (
	ZeroTrustGatewayRulesFilterHTTP   ZeroTrustGatewayRulesFilter = "http"
	ZeroTrustGatewayRulesFilterDNS    ZeroTrustGatewayRulesFilter = "dns"
	ZeroTrustGatewayRulesFilterL4     ZeroTrustGatewayRulesFilter = "l4"
	ZeroTrustGatewayRulesFilterEgress ZeroTrustGatewayRulesFilter = "egress"
)

func (r ZeroTrustGatewayRulesFilter) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayRulesFilterHTTP, ZeroTrustGatewayRulesFilterDNS, ZeroTrustGatewayRulesFilterL4, ZeroTrustGatewayRulesFilterEgress:
		return true
	}
	return false
}

type GatewayRuleNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleNewParamsAction] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]FitlerItem] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[RuleSettingsParam] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleNewParamsAction string

const (
	GatewayRuleNewParamsActionOn           GatewayRuleNewParamsAction = "on"
	GatewayRuleNewParamsActionOff          GatewayRuleNewParamsAction = "off"
	GatewayRuleNewParamsActionAllow        GatewayRuleNewParamsAction = "allow"
	GatewayRuleNewParamsActionBlock        GatewayRuleNewParamsAction = "block"
	GatewayRuleNewParamsActionScan         GatewayRuleNewParamsAction = "scan"
	GatewayRuleNewParamsActionNoscan       GatewayRuleNewParamsAction = "noscan"
	GatewayRuleNewParamsActionSafesearch   GatewayRuleNewParamsAction = "safesearch"
	GatewayRuleNewParamsActionYtrestricted GatewayRuleNewParamsAction = "ytrestricted"
	GatewayRuleNewParamsActionIsolate      GatewayRuleNewParamsAction = "isolate"
	GatewayRuleNewParamsActionNoisolate    GatewayRuleNewParamsAction = "noisolate"
	GatewayRuleNewParamsActionOverride     GatewayRuleNewParamsAction = "override"
	GatewayRuleNewParamsActionL4Override   GatewayRuleNewParamsAction = "l4_override"
	GatewayRuleNewParamsActionEgress       GatewayRuleNewParamsAction = "egress"
	GatewayRuleNewParamsActionAuditSSH     GatewayRuleNewParamsAction = "audit_ssh"
	GatewayRuleNewParamsActionResolve      GatewayRuleNewParamsAction = "resolve"
)

func (r GatewayRuleNewParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleNewParamsActionOn, GatewayRuleNewParamsActionOff, GatewayRuleNewParamsActionAllow, GatewayRuleNewParamsActionBlock, GatewayRuleNewParamsActionScan, GatewayRuleNewParamsActionNoscan, GatewayRuleNewParamsActionSafesearch, GatewayRuleNewParamsActionYtrestricted, GatewayRuleNewParamsActionIsolate, GatewayRuleNewParamsActionNoisolate, GatewayRuleNewParamsActionOverride, GatewayRuleNewParamsActionL4Override, GatewayRuleNewParamsActionEgress, GatewayRuleNewParamsActionAuditSSH, GatewayRuleNewParamsActionResolve:
		return true
	}
	return false
}

type GatewayRuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ZeroTrustGatewayRules                                     `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponseEnvelope]
type gatewayRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleNewResponseEnvelopeSuccess bool

const (
	GatewayRuleNewResponseEnvelopeSuccessTrue GatewayRuleNewResponseEnvelopeSuccess = true
)

func (r GatewayRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleUpdateParamsAction] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]FitlerItem] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[RuleSettingsParam] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleUpdateParamsAction string

const (
	GatewayRuleUpdateParamsActionOn           GatewayRuleUpdateParamsAction = "on"
	GatewayRuleUpdateParamsActionOff          GatewayRuleUpdateParamsAction = "off"
	GatewayRuleUpdateParamsActionAllow        GatewayRuleUpdateParamsAction = "allow"
	GatewayRuleUpdateParamsActionBlock        GatewayRuleUpdateParamsAction = "block"
	GatewayRuleUpdateParamsActionScan         GatewayRuleUpdateParamsAction = "scan"
	GatewayRuleUpdateParamsActionNoscan       GatewayRuleUpdateParamsAction = "noscan"
	GatewayRuleUpdateParamsActionSafesearch   GatewayRuleUpdateParamsAction = "safesearch"
	GatewayRuleUpdateParamsActionYtrestricted GatewayRuleUpdateParamsAction = "ytrestricted"
	GatewayRuleUpdateParamsActionIsolate      GatewayRuleUpdateParamsAction = "isolate"
	GatewayRuleUpdateParamsActionNoisolate    GatewayRuleUpdateParamsAction = "noisolate"
	GatewayRuleUpdateParamsActionOverride     GatewayRuleUpdateParamsAction = "override"
	GatewayRuleUpdateParamsActionL4Override   GatewayRuleUpdateParamsAction = "l4_override"
	GatewayRuleUpdateParamsActionEgress       GatewayRuleUpdateParamsAction = "egress"
	GatewayRuleUpdateParamsActionAuditSSH     GatewayRuleUpdateParamsAction = "audit_ssh"
	GatewayRuleUpdateParamsActionResolve      GatewayRuleUpdateParamsAction = "resolve"
)

func (r GatewayRuleUpdateParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateParamsActionOn, GatewayRuleUpdateParamsActionOff, GatewayRuleUpdateParamsActionAllow, GatewayRuleUpdateParamsActionBlock, GatewayRuleUpdateParamsActionScan, GatewayRuleUpdateParamsActionNoscan, GatewayRuleUpdateParamsActionSafesearch, GatewayRuleUpdateParamsActionYtrestricted, GatewayRuleUpdateParamsActionIsolate, GatewayRuleUpdateParamsActionNoisolate, GatewayRuleUpdateParamsActionOverride, GatewayRuleUpdateParamsActionL4Override, GatewayRuleUpdateParamsActionEgress, GatewayRuleUpdateParamsActionAuditSSH, GatewayRuleUpdateParamsActionResolve:
		return true
	}
	return false
}

type GatewayRuleUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ZeroTrustGatewayRules                                     `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleUpdateResponseEnvelope]
type gatewayRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleUpdateResponseEnvelopeSuccess bool

const (
	GatewayRuleUpdateResponseEnvelopeSuccessTrue GatewayRuleUpdateResponseEnvelopeSuccess = true
)

func (r GatewayRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r GatewayRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GatewayRuleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleDeleteResponseEnvelope]
type gatewayRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleDeleteResponseEnvelopeSuccess bool

const (
	GatewayRuleDeleteResponseEnvelopeSuccessTrue GatewayRuleDeleteResponseEnvelopeSuccess = true
)

func (r GatewayRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ZeroTrustGatewayRules                                     `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponseEnvelope]
type gatewayRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleGetResponseEnvelopeSuccess bool

const (
	GatewayRuleGetResponseEnvelopeSuccessTrue GatewayRuleGetResponseEnvelopeSuccess = true
)

func (r GatewayRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
