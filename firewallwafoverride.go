// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallWAFOverrideService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallWAFOverrideService]
// method instead.
type FirewallWAFOverrideService struct {
	Options []option.RequestOption
}

// NewFirewallWAFOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFOverrideService(opts ...option.RequestOption) (r *FirewallWAFOverrideService) {
	r = &FirewallWAFOverrideService{}
	r.Options = opts
	return
}

// Updates an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallWAFOverrideUpdateParams, opts ...option.RequestOption) (res *FirewallWAFOverrideUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallWAFOverrideDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallWAFOverrideGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) WAFOverridesNewAWAFOverride(ctx context.Context, zoneIdentifier string, body FirewallWAFOverrideWAFOverridesNewAWAFOverrideParams, opts ...option.RequestOption) (res *FirewallWAFOverrideWAFOverridesNewAwafOverrideResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) WAFOverridesListWAFOverrides(ctx context.Context, zoneIdentifier string, query FirewallWAFOverrideWAFOverridesListWAFOverridesParams, opts ...option.RequestOption) (res *[]FirewallWAFOverrideWAFOverridesListWAFOverridesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallWAFOverrideUpdateResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction FirewallWAFOverrideUpdateResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]FirewallWAFOverrideUpdateResponseRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                              `json:"urls"`
	JSON firewallWAFOverrideUpdateResponseJSON `json:"-"`
}

// firewallWAFOverrideUpdateResponseJSON contains the JSON metadata for the struct
// [FirewallWAFOverrideUpdateResponse]
type firewallWAFOverrideUpdateResponseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type FirewallWAFOverrideUpdateResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     FirewallWAFOverrideUpdateResponseRewriteActionBlock `json:"block"`
	Challenge interface{}                                         `json:"challenge"`
	Default   interface{}                                         `json:"default"`
	// The WAF rule action to apply.
	Disable  FirewallWAFOverrideUpdateResponseRewriteActionDisable `json:"disable"`
	Simulate interface{}                                           `json:"simulate"`
	JSON     firewallWAFOverrideUpdateResponseRewriteActionJSON    `json:"-"`
}

// firewallWAFOverrideUpdateResponseRewriteActionJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideUpdateResponseRewriteAction]
type firewallWAFOverrideUpdateResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type FirewallWAFOverrideUpdateResponseRewriteActionBlock string

const (
	FirewallWAFOverrideUpdateResponseRewriteActionBlockChallenge FirewallWAFOverrideUpdateResponseRewriteActionBlock = "challenge"
	FirewallWAFOverrideUpdateResponseRewriteActionBlockBlock     FirewallWAFOverrideUpdateResponseRewriteActionBlock = "block"
	FirewallWAFOverrideUpdateResponseRewriteActionBlockSimulate  FirewallWAFOverrideUpdateResponseRewriteActionBlock = "simulate"
	FirewallWAFOverrideUpdateResponseRewriteActionBlockDisable   FirewallWAFOverrideUpdateResponseRewriteActionBlock = "disable"
	FirewallWAFOverrideUpdateResponseRewriteActionBlockDefault   FirewallWAFOverrideUpdateResponseRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideUpdateResponseRewriteActionDisable string

const (
	FirewallWAFOverrideUpdateResponseRewriteActionDisableChallenge FirewallWAFOverrideUpdateResponseRewriteActionDisable = "challenge"
	FirewallWAFOverrideUpdateResponseRewriteActionDisableBlock     FirewallWAFOverrideUpdateResponseRewriteActionDisable = "block"
	FirewallWAFOverrideUpdateResponseRewriteActionDisableSimulate  FirewallWAFOverrideUpdateResponseRewriteActionDisable = "simulate"
	FirewallWAFOverrideUpdateResponseRewriteActionDisableDisable   FirewallWAFOverrideUpdateResponseRewriteActionDisable = "disable"
	FirewallWAFOverrideUpdateResponseRewriteActionDisableDefault   FirewallWAFOverrideUpdateResponseRewriteActionDisable = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideUpdateResponseRule string

const (
	FirewallWAFOverrideUpdateResponseRuleChallenge FirewallWAFOverrideUpdateResponseRule = "challenge"
	FirewallWAFOverrideUpdateResponseRuleBlock     FirewallWAFOverrideUpdateResponseRule = "block"
	FirewallWAFOverrideUpdateResponseRuleSimulate  FirewallWAFOverrideUpdateResponseRule = "simulate"
	FirewallWAFOverrideUpdateResponseRuleDisable   FirewallWAFOverrideUpdateResponseRule = "disable"
	FirewallWAFOverrideUpdateResponseRuleDefault   FirewallWAFOverrideUpdateResponseRule = "default"
)

type FirewallWAFOverrideDeleteResponse struct {
	// The unique identifier of the WAF override.
	ID   string                                `json:"id"`
	JSON firewallWAFOverrideDeleteResponseJSON `json:"-"`
}

// firewallWAFOverrideDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallWAFOverrideDeleteResponse]
type firewallWAFOverrideDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction FirewallWAFOverrideGetResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]FirewallWAFOverrideGetResponseRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                           `json:"urls"`
	JSON firewallWAFOverrideGetResponseJSON `json:"-"`
}

// firewallWAFOverrideGetResponseJSON contains the JSON metadata for the struct
// [FirewallWAFOverrideGetResponse]
type firewallWAFOverrideGetResponseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type FirewallWAFOverrideGetResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     FirewallWAFOverrideGetResponseRewriteActionBlock `json:"block"`
	Challenge interface{}                                      `json:"challenge"`
	Default   interface{}                                      `json:"default"`
	// The WAF rule action to apply.
	Disable  FirewallWAFOverrideGetResponseRewriteActionDisable `json:"disable"`
	Simulate interface{}                                        `json:"simulate"`
	JSON     firewallWAFOverrideGetResponseRewriteActionJSON    `json:"-"`
}

// firewallWAFOverrideGetResponseRewriteActionJSON contains the JSON metadata for
// the struct [FirewallWAFOverrideGetResponseRewriteAction]
type firewallWAFOverrideGetResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type FirewallWAFOverrideGetResponseRewriteActionBlock string

const (
	FirewallWAFOverrideGetResponseRewriteActionBlockChallenge FirewallWAFOverrideGetResponseRewriteActionBlock = "challenge"
	FirewallWAFOverrideGetResponseRewriteActionBlockBlock     FirewallWAFOverrideGetResponseRewriteActionBlock = "block"
	FirewallWAFOverrideGetResponseRewriteActionBlockSimulate  FirewallWAFOverrideGetResponseRewriteActionBlock = "simulate"
	FirewallWAFOverrideGetResponseRewriteActionBlockDisable   FirewallWAFOverrideGetResponseRewriteActionBlock = "disable"
	FirewallWAFOverrideGetResponseRewriteActionBlockDefault   FirewallWAFOverrideGetResponseRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideGetResponseRewriteActionDisable string

const (
	FirewallWAFOverrideGetResponseRewriteActionDisableChallenge FirewallWAFOverrideGetResponseRewriteActionDisable = "challenge"
	FirewallWAFOverrideGetResponseRewriteActionDisableBlock     FirewallWAFOverrideGetResponseRewriteActionDisable = "block"
	FirewallWAFOverrideGetResponseRewriteActionDisableSimulate  FirewallWAFOverrideGetResponseRewriteActionDisable = "simulate"
	FirewallWAFOverrideGetResponseRewriteActionDisableDisable   FirewallWAFOverrideGetResponseRewriteActionDisable = "disable"
	FirewallWAFOverrideGetResponseRewriteActionDisableDefault   FirewallWAFOverrideGetResponseRewriteActionDisable = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideGetResponseRule string

const (
	FirewallWAFOverrideGetResponseRuleChallenge FirewallWAFOverrideGetResponseRule = "challenge"
	FirewallWAFOverrideGetResponseRuleBlock     FirewallWAFOverrideGetResponseRule = "block"
	FirewallWAFOverrideGetResponseRuleSimulate  FirewallWAFOverrideGetResponseRule = "simulate"
	FirewallWAFOverrideGetResponseRuleDisable   FirewallWAFOverrideGetResponseRule = "disable"
	FirewallWAFOverrideGetResponseRuleDefault   FirewallWAFOverrideGetResponseRule = "default"
)

type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                                                   `json:"urls"`
	JSON firewallWAFOverrideWAFOverridesNewAwafOverrideResponseJSON `json:"-"`
}

// firewallWAFOverrideWAFOverridesNewAwafOverrideResponseJSON contains the JSON
// metadata for the struct [FirewallWAFOverrideWAFOverridesNewAwafOverrideResponse]
type firewallWAFOverrideWAFOverridesNewAwafOverrideResponseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesNewAwafOverrideResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock `json:"block"`
	Challenge interface{}                                                              `json:"challenge"`
	Default   interface{}                                                              `json:"default"`
	// The WAF rule action to apply.
	Disable  FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable `json:"disable"`
	Simulate interface{}                                                                `json:"simulate"`
	JSON     firewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionJSON    `json:"-"`
}

// firewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionJSON contains
// the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteAction]
type firewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock string

const (
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlockChallenge FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock = "challenge"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlockBlock     FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock = "block"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlockSimulate  FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock = "simulate"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlockDisable   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock = "disable"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlockDefault   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable string

const (
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisableChallenge FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable = "challenge"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisableBlock     FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable = "block"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisableSimulate  FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable = "simulate"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisableDisable   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable = "disable"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisableDefault   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRewriteActionDisable = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule string

const (
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRuleChallenge FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule = "challenge"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRuleBlock     FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule = "block"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRuleSimulate  FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule = "simulate"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRuleDisable   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule = "disable"
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRuleDefault   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseRule = "default"
)

type FirewallWAFOverrideWAFOverridesListWAFOverridesResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id,required"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused,required"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string `json:"urls,required"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule `json:"rules"`
	JSON  firewallWAFOverrideWAFOverridesListWAFOverridesResponseJSON            `json:"-"`
}

// firewallWAFOverrideWAFOverridesListWAFOverridesResponseJSON contains the JSON
// metadata for the struct
// [FirewallWAFOverrideWAFOverridesListWAFOverridesResponse]
type firewallWAFOverrideWAFOverridesListWAFOverridesResponseJSON struct {
	ID            apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	URLs          apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesListWAFOverridesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock `json:"block"`
	Challenge interface{}                                                               `json:"challenge"`
	Default   interface{}                                                               `json:"default"`
	// The WAF rule action to apply.
	Disable  FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable `json:"disable"`
	Simulate interface{}                                                                 `json:"simulate"`
	JSON     firewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionJSON    `json:"-"`
}

// firewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionJSON
// contains the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteAction]
type firewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock string

const (
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlockChallenge FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock = "challenge"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlockBlock     FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock = "block"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlockSimulate  FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock = "simulate"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlockDisable   FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock = "disable"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlockDefault   FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable string

const (
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisableChallenge FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable = "challenge"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisableBlock     FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable = "block"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisableSimulate  FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable = "simulate"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisableDisable   FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable = "disable"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisableDefault   FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRewriteActionDisable = "default"
)

// The WAF rule action to apply.
type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule string

const (
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRuleChallenge FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule = "challenge"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRuleBlock     FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule = "block"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRuleSimulate  FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule = "simulate"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRuleDisable   FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule = "disable"
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRuleDefault   FirewallWAFOverrideWAFOverridesListWAFOverridesResponseRule = "default"
)

type FirewallWAFOverrideUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallWAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallWAFOverrideUpdateResponseEnvelope struct {
	Errors   []FirewallWAFOverrideUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFOverrideUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFOverrideUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFOverrideUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFOverrideUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideUpdateResponseEnvelope]
type firewallWAFOverrideUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    firewallWAFOverrideUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideUpdateResponseEnvelopeErrors]
type firewallWAFOverrideUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    firewallWAFOverrideUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideUpdateResponseEnvelopeMessages]
type firewallWAFOverrideUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideUpdateResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideUpdateResponseEnvelopeSuccessTrue FirewallWAFOverrideUpdateResponseEnvelopeSuccess = true
)

type FirewallWAFOverrideDeleteResponseEnvelope struct {
	Result FirewallWAFOverrideDeleteResponse             `json:"result"`
	JSON   firewallWAFOverrideDeleteResponseEnvelopeJSON `json:"-"`
}

// firewallWAFOverrideDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideDeleteResponseEnvelope]
type firewallWAFOverrideDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponseEnvelope struct {
	Errors   []FirewallWAFOverrideGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFOverrideGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFOverrideGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFOverrideGetResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFOverrideGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideGetResponseEnvelope]
type firewallWAFOverrideGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallWAFOverrideGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallWAFOverrideGetResponseEnvelopeErrors]
type firewallWAFOverrideGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallWAFOverrideGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideGetResponseEnvelopeMessages]
type firewallWAFOverrideGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideGetResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideGetResponseEnvelopeSuccessTrue FirewallWAFOverrideGetResponseEnvelopeSuccess = true
)

type FirewallWAFOverrideWAFOverridesNewAWAFOverrideParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallWAFOverrideWAFOverridesNewAWAFOverrideParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelope struct {
	Errors   []FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFOverrideWAFOverridesNewAwafOverrideResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelope]
type firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrors]
type firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessages]
type firewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeSuccessTrue FirewallWAFOverrideWAFOverridesNewAwafOverrideResponseEnvelopeSuccess = true
)

type FirewallWAFOverrideWAFOverridesListWAFOverridesParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallWAFOverrideWAFOverridesListWAFOverridesParams]'s
// query parameters as `url.Values`.
func (r FirewallWAFOverrideWAFOverridesListWAFOverridesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelope struct {
	Errors   []FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallWAFOverrideWAFOverridesListWAFOverridesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeJSON       `json:"-"`
}

// firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelope]
type firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrors]
type firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessages]
type firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeSuccessTrue FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeSuccess = true
)

type FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfo]
type firewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideWAFOverridesListWAFOverridesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
