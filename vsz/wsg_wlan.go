package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWLANService struct {
	apiClient *APIClient
}

func NewWSGWLANService(c *APIClient) *WSGWLANService {
	s := new(WSGWLANService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWLANService() *WSGWLANService {
	return NewWSGWLANService(ss.apiClient)
}

type WSGWLANCreateGuestAccessWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateGuestAccessWlan() *WSGWLANCreateGuestAccessWlan {
	m := new(WSGWLANCreateGuestAccessWlan)
	return m
}

type WSGWLANCreateHotspot20OpenWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateHotspot20OpenWlan() *WSGWLANCreateHotspot20OpenWlan {
	m := new(WSGWLANCreateHotspot20OpenWlan)
	return m
}

type WSGWLANCreateHotspot20Wlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	// Hotspot20Profile
	// Constraints:
	//    - required
	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile" validate:"required"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateHotspot20Wlan() *WSGWLANCreateHotspot20Wlan {
	m := new(WSGWLANCreateHotspot20Wlan)
	return m
}

type WSGWLANCreateHotspotWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateHotspotWlan() *WSGWLANCreateHotspotWlan {
	m := new(WSGWLANCreateHotspotWlan)
	return m
}

type WSGWLANCreateStandard80211Wlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable  for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable  for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateStandard80211Wlan() *WSGWLANCreateStandard80211Wlan {
	m := new(WSGWLANCreateStandard80211Wlan)
	return m
}

type WSGWLANCreateStandardOpenWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - default:'APLBO'
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateStandardOpenWlan() *WSGWLANCreateStandardOpenWlan {
	m := new(WSGWLANCreateStandardOpenWlan)
	return m
}

type WSGWLANCreateWebAuthWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateWebAuthWlan() *WSGWLANCreateWebAuthWlan {
	m := new(WSGWLANCreateWebAuthWlan)
	return m
}

type WSGWLANCreateWechatWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name" validate:"required,max=32,min=1"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid" validate:"required,max=32,min=1"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateWechatWlan() *WSGWLANCreateWechatWlan {
	m := new(WSGWLANCreateWechatWlan)
	return m
}

type WSGWLANModifyWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	Name *WSGWLANNameSSID `json:"name,omitempty"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WSGWLANNameSSID `json:"ssid,omitempty"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANModifyWlan() *WSGWLANModifyWlan {
	m := new(WSGWLANModifyWlan)
	return m
}

type WSGWLANAccounting struct {
	// AccountingDelayEnabled
	// Indicates whether accounting delay time is enabled
	AccountingDelayEnabled *bool `json:"accountingDelayEnabled,omitempty"`

	// BackupAccountingId
	// Backup accounting service or profile ID. At least one backupAccountingId or backupAccountingName is required in the request when setting backup accounting service.
	BackupAccountingId *string `json:"backupAccountingId,omitempty"`

	// BackupAccountingName
	// Backup accounting service or profile name. At least one backupAccountingId or backupAccountingName is required in the request when setting backup accounting service.
	BackupAccountingName *string `json:"backupAccountingName,omitempty"`

	// Id
	// Accounting service or profile ID. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// InterimUpdateMin
	// Interval (in minutes) for sending interim updates
	// Constraints:
	//    - min:0
	//    - max:1440
	InterimUpdateMin *int `json:"interimUpdateMin,omitempty" validate:"gte=0,lte=1440"`

	// Name
	// Accounting service or profile name. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	RealmBasedAcct *bool `json:"realmBasedAcct,omitempty"`

	// ThroughController
	// Indicates whether accounting messages were sent through the controller
	ThroughController *bool `json:"throughController,omitempty"`
}

func NewWSGWLANAccounting() *WSGWLANAccounting {
	m := new(WSGWLANAccounting)
	return m
}

type WSGWLANAdvanced struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	// AssocRssiThr
	// Assoc RSSI threshold.
	// Constraints:
	//    - min:-90
	//    - max:-60
	AssocRssiThr *int `json:"assocRssiThr,omitempty" validate:"gte=-90,lte=-60"`

	// AuthRssiThr
	// Auth RSSI threshold.
	// Constraints:
	//    - min:-90
	//    - max:-60
	AuthRssiThr *int `json:"authRssiThr,omitempty" validate:"gte=-90,lte=-60"`

	// AvcEnabled
	// Indicator of whether AVC support is enabled or disabled
	AvcEnabled *bool `json:"avcEnabled,omitempty"`

	// BandBalancing
	// Indicates whether band balancing is enabled or disabled
	// Constraints:
	//    - default:'UseZoneSetting'
	//    - oneof:[Disabled,UseZoneSetting]
	BandBalancing *string `json:"bandBalancing,omitempty" validate:"oneof=Disabled UseZoneSetting"`

	BssMinRateMbps *WSGWLANBssMinRateMbps `json:"bssMinRateMbps,omitempty"`

	// ClientFingerprintingEnabled
	// Indicates whether client fingerprinting is enabled or disabled
	ClientFingerprintingEnabled *bool `json:"clientFingerprintingEnabled,omitempty"`

	// ClientIdleTimeoutSec
	// Client idle timeout in seconds
	// Constraints:
	//    - default:120
	//    - min:60
	//    - max:86400
	ClientIdleTimeoutSec *int `json:"clientIdleTimeoutSec,omitempty" validate:"gte=60,lte=86400"`

	// ClientIsolationAutoVrrpEnabled
	// Indicates whether Automatic support for VRRP of wireless client isolation is enabled or disabled
	ClientIsolationAutoVrrpEnabled *bool `json:"clientIsolationAutoVrrpEnabled,omitempty"`

	// ClientIsolationEnabled
	// Indicates whether wireless client isolation is enabled or disabled
	ClientIsolationEnabled *bool `json:"clientIsolationEnabled,omitempty"`

	// ClientIsolationMulticastEnabled
	// Indicates whether isolate multicast of wireless client isolation is enabled or disabled
	ClientIsolationMulticastEnabled *bool `json:"clientIsolationMulticastEnabled,omitempty"`

	// ClientIsolationUnicastEnabled
	// Indicates whether isolate unicast of wireless client isolation is enabled or disabled
	ClientIsolationUnicastEnabled *bool `json:"clientIsolationUnicastEnabled,omitempty"`

	ClientIsolationWhitelist *WSGCommonGenericRef `json:"clientIsolationWhitelist,omitempty"`

	// ClientLoadBalancingEnabled
	// Indicates whether Client Load Balancing is enabled or disabled
	ClientLoadBalancingEnabled *bool `json:"clientLoadBalancingEnabled,omitempty"`

	// DgafEnabled
	// Indicates whether dgaf is enabled or disabled
	DgafEnabled *bool `json:"dgafEnabled,omitempty"`

	// Dhcp82Format
	// DHCP Option 82 format. This variable no longer supports from v8_1 and only kept for backward compatibility.
	// Constraints:
	//    - oneof:[RUCKUS_DEFAULT,RADIUS_DHCP,RADIUS_NAT,RADIUS_DHCP_NAT,RADIUS_NAT_RUCKUS,RADIUS_NAT_SOFTGRE,SOFTGRE_CUSTOMIZED]
	Dhcp82Format *string `json:"dhcp82Format,omitempty" validate:"oneof=RUCKUS_DEFAULT RADIUS_DHCP RADIUS_NAT RADIUS_DHCP_NAT RADIUS_NAT_RUCKUS RADIUS_NAT_SOFTGRE SOFTGRE_CUSTOMIZED"`

	// Dhcp82MacFormat
	// AP and Client Mac format. If dhcpOption82Enabled is true, you have to set the dhcp82MacFormat ["COLON","HYPHEN","NODELIMITER"].
	// Constraints:
	//    - oneof:[COLON,HYPHEN,NODELIMITER]
	Dhcp82MacFormat *string `json:"dhcp82MacFormat,omitempty" validate:"oneof=COLON HYPHEN NODELIMITER"`

	// Dhcp82SubOpt1Format
	// Subopt-1 format
	// Constraints:
	//    - oneof:[NONE,SUBOPT1_AP_INFO_LOCATION,SUBOPT1_AP_INFO,SUBOPT1_AP_MAC_ESSID_PRIVACYTYPE,SUBOPT1_AP_MAC_hex,SUBOPT1_AP_MAC_hex_ESSID,SUBOPT1_ESSID,SUBOPT1_AP_MAC,SUBOPT1_AP_MAC_ESSID,SUBOPT1_AP_Name_ESSID]
	Dhcp82SubOpt1Format *string `json:"dhcp82SubOpt1Format,omitempty" validate:"oneof=NONE SUBOPT1_AP_INFO_LOCATION SUBOPT1_AP_INFO SUBOPT1_AP_MAC_ESSID_PRIVACYTYPE SUBOPT1_AP_MAC_hex SUBOPT1_AP_MAC_hex_ESSID SUBOPT1_ESSID SUBOPT1_AP_MAC SUBOPT1_AP_MAC_ESSID SUBOPT1_AP_Name_ESSID"`

	// Dhcp82SubOpt2Format
	// Subopt-2 format
	// Constraints:
	//    - oneof:[NONE,SUBOPT2_CLIENT_MAC,SUBOPT2_CLIENT_MAC_hex,SUBOPT2_CLIENT_MAC_hex_ESSID,SUBOPT2_AP_MAC,SUBOPT2_AP_MAC_hex,SUBOPT2_AP_MAC_hex_ESSID,SUBOPT2_AP_MAC_ESSID,SUBOPT2_AP_Name]
	Dhcp82SubOpt2Format *string `json:"dhcp82SubOpt2Format,omitempty" validate:"oneof=NONE SUBOPT2_CLIENT_MAC SUBOPT2_CLIENT_MAC_hex SUBOPT2_CLIENT_MAC_hex_ESSID SUBOPT2_AP_MAC SUBOPT2_AP_MAC_hex SUBOPT2_AP_MAC_hex_ESSID SUBOPT2_AP_MAC_ESSID SUBOPT2_AP_Name"`

	// Dhcp82SubOpt150Format
	// Subopt-150 with VLAN-Id
	// Constraints:
	//    - oneof:[NONE,SUBOPT150_VLAN_ID]
	Dhcp82SubOpt150Format *string `json:"dhcp82SubOpt150Format,omitempty" validate:"oneof=NONE SUBOPT150_VLAN_ID"`

	// Dhcp82SubOpt151AreaName
	// Subopt-151 Area Name value
	Dhcp82SubOpt151AreaName *string `json:"dhcp82SubOpt151AreaName,omitempty"`

	// Dhcp82SubOpt151Format
	// Subopt-151 format
	// Constraints:
	//    - oneof:[NONE,SUBOPT151_AREA_NAME,SUBOPT151_ESSID]
	Dhcp82SubOpt151Format *string `json:"dhcp82SubOpt151Format,omitempty" validate:"oneof=NONE SUBOPT151_AREA_NAME SUBOPT151_ESSID"`

	// Dhcp82SubOptRadiusFormat
	// Constraints:
	//    - oneof:[NONE,RUCKUS_DEFAULT,RADIUS_DHCP,RADIUS_NAT,RADIUS_DHCP_NAT,RADIUS_NAT_RUCKUS,RADIUS_NAT_SOFTGRE,SOFTGRE_CUSTOMIZED]
	Dhcp82SubOptRadiusFormat *string `json:"dhcp82SubOptRadiusFormat,omitempty" validate:"oneof=NONE RUCKUS_DEFAULT RADIUS_DHCP RADIUS_NAT RADIUS_DHCP_NAT RADIUS_NAT_RUCKUS RADIUS_NAT_SOFTGRE SOFTGRE_CUSTOMIZED"`

	// DhcpOption82Enabled
	// Indicates whether DCHP Option 82 is enabled or disabled. This variable no longer supports from v8_1 and only kept for backward compatibility.
	DhcpOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	// DirectedThreshold
	// Directed Threshold Setting, Defines the client count at which an AP will stop converting group addressed data traffic to unicast.
	// Constraints:
	//    - default:5
	//    - min:0
	//    - max:128
	DirectedThreshold *int `json:"directedThreshold,omitempty" validate:"gte=0,lte=128"`

	// DnsSpoofingProfileId
	// DNS Spoofing Profile ID
	DnsSpoofingProfileId *string `json:"dnsSpoofingProfileId,omitempty"`

	// DownlinkEnabled
	// SSID Rate Limiting downlink enabled.
	DownlinkEnabled *bool `json:"downlinkEnabled,omitempty"`

	// DownlinkRate
	// SSID Rate Limiting downlink.
	DownlinkRate *float64 `json:"downlinkRate,omitempty"`

	// DropRandomProbesEnabled
	// Drop Random Probes enabled.
	DropRandomProbesEnabled *bool `json:"dropRandomProbesEnabled,omitempty"`

	// DtimInterval
	// DTIM Interval
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:255
	DtimInterval *int `json:"dtimInterval,omitempty" validate:"gte=1,lte=255"`

	// FlowLogEnabled
	// Flow log enabled.
	FlowLogEnabled *bool `json:"flowLogEnabled,omitempty"`

	// ForceClientDHCPTimeoutSec
	// Force DHCP disconnects the client if the client does not obtain a valid IP address within the timeout peroid. To disable force DHCP, set this value to zero (0).
	// Constraints:
	//    - default:0
	//    - oneof:[0,5,6,7,8,9,10,11,12,13,14,15]
	ForceClientDHCPTimeoutSec *int `json:"forceClientDHCPTimeoutSec,omitempty" validate:"oneof=0 5 6 7 8 9 10 11 12 13 14 15"`

	// HdOverheadOptimizeEnable
	// Airtime decongestion enabled.
	HdOverheadOptimizeEnable *bool `json:"hdOverheadOptimizeEnable,omitempty"`

	// HideSsidEnabled
	// Indicates whether the SSID is hidden or broadcast
	HideSsidEnabled *bool `json:"hideSsidEnabled,omitempty"`

	// Hs20Onboarding
	// Allow WISPr WLAN for Hotspot 2.0 Onboarding
	Hs20Onboarding *bool `json:"hs20Onboarding,omitempty"`

	// JoinAcceptTimeout
	// Join expire time.
	// Constraints:
	//    - default:300
	//    - min:1
	//    - max:300
	JoinAcceptTimeout *int `json:"joinAcceptTimeout,omitempty" validate:"gte=1,lte=300"`

	// JoinIgnoreThr
	// Join wait threshold.
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:50
	JoinIgnoreThr *int `json:"joinIgnoreThr,omitempty" validate:"gte=1,lte=50"`

	// JoinIgnoreTimeout
	// Join wait time.
	// Constraints:
	//    - default:30
	//    - min:1
	//    - max:60
	JoinIgnoreTimeout *int `json:"joinIgnoreTimeout,omitempty" validate:"gte=1,lte=60"`

	// MaxAllowedRA
	// Max Allowed RAs
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:1440
	MaxAllowedRA *int `json:"maxAllowedRA,omitempty" validate:"gte=1,lte=1440"`

	// MaxClientsPerRadio
	// Maximum number of clients per radio
	// Constraints:
	//    - default:100
	//    - min:1
	//    - max:512
	MaxClientsPerRadio *int `json:"maxClientsPerRadio,omitempty" validate:"gte=1,lte=512"`

	MgmtTxRateMbps *WSGWLANMgmtTxRateMbps `json:"mgmtTxRateMbps,omitempty"`

	// MulticastDownlinkRateLimit
	// Multicast Rate Limiting downlink (mbps).
	// Constraints:
	//    - min:1
	//    - max:100
	MulticastDownlinkRateLimit *int `json:"multicastDownlinkRateLimit,omitempty" validate:"gte=1,lte=100"`

	// MulticastDownlinkRateLimitEnabled
	// Multicast Rate Limiting downlink enabled.
	MulticastDownlinkRateLimitEnabled *bool `json:"multicastDownlinkRateLimitEnabled,omitempty"`

	// MulticastFilterDrop
	// Drop the broadcast/multicast packets from associated clients.
	MulticastFilterDrop *bool `json:"multicastFilterDrop,omitempty"`

	// MulticastUplinkRateLimit
	// Multicast Rate Limiting uplink (mbps).
	// Constraints:
	//    - min:1
	//    - max:100
	MulticastUplinkRateLimit *int `json:"multicastUplinkRateLimit,omitempty" validate:"gte=1,lte=100"`

	// MulticastUplinkRateLimitEnabled
	// Multicast Rate Limiting uplink enabled.
	MulticastUplinkRateLimitEnabled *bool `json:"multicastUplinkRateLimitEnabled,omitempty"`

	// NdProxyEnabled
	// Indicates whether ND Proxy is enabled or disabled
	NdProxyEnabled *bool `json:"ndProxyEnabled,omitempty"`

	// OceBroadcastProbeResponseDelay
	// Broadcast probe response delay.
	// Constraints:
	//    - default:15
	//    - min:8
	//    - max:120
	OceBroadcastProbeResponseDelay *int `json:"oceBroadcastProbeResponseDelay,omitempty" validate:"gte=8,lte=120"`

	// OceEnabled
	// Optimized Connectivity Experience(OCE) enabled.
	OceEnabled *bool `json:"oceEnabled,omitempty"`

	// OceRssiBasedAssociationRejectionThreshold
	// RSSI-based association rejection threshold.
	// Constraints:
	//    - default:-75
	//    - min:-90
	//    - max:-60
	OceRssiBasedAssociationRejectionThreshold *int `json:"oceRssiBasedAssociationRejectionThreshold,omitempty" validate:"gte=-90,lte=-60"`

	// OfdmOnlyEnabled
	// Indicates whether OFDM only is enabled or disabled
	OfdmOnlyEnabled *bool `json:"ofdmOnlyEnabled,omitempty"`

	// OkcEnabled
	// Indicator of whether OKC support is enabled or disabled. The default value is true when the WLAN is WPA+AES non open WLAN.
	OkcEnabled *bool `json:"okcEnabled,omitempty"`

	// PmkCachingEnabled
	// Indicator of whether PKM caching support is enabled or disabled. The default value is true when the WLAN is WPA+AES non open WLAN.
	PmkCachingEnabled *bool `json:"pmkCachingEnabled,omitempty"`

	// Priority
	// Priority of the WLAN
	// Constraints:
	//    - default:'High'
	//    - oneof:[High,Low]
	Priority *string `json:"priority,omitempty" validate:"oneof=High Low"`

	// ProbeRssiThr
	// Join RSSI threshold.
	// Constraints:
	//    - nullable
	//    - default:-85
	//    - min:-90
	//    - max:-60
	ProbeRssiThr *int `json:"probeRssiThr,omitempty" validate:"omitempty,gte=-90,lte=-60"`

	// ProxyARPEnabled
	// Indicates whether proxy ARP is enabled or disabled
	ProxyARPEnabled *bool `json:"proxyARPEnabled,omitempty"`

	// RaInterval
	// A timer that RA proxy runs and once receives unsolicited RA checks against the configured time and allow/drop RA based on next timeout
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:256
	RaInterval *int `json:"raInterval,omitempty" validate:"gte=1,lte=256"`

	// RaProxyEnabled
	// Indicates whether RA proxy is enabled or disabled
	RaProxyEnabled *bool `json:"raProxyEnabled,omitempty"`

	// RatePerSTADownlink
	// UE Rate Limiting downlink.
	RatePerSTADownlink *string `json:"ratePerSTADownlink,omitempty"`

	// RatePerSTAUplink
	// UE Rate Limiting uplink.
	RatePerSTAUplink *string `json:"ratePerSTAUplink,omitempty"`

	// RaThrottlingEnabled
	// Indicates whether RA Throttling is enabled or disabled
	RaThrottlingEnabled *bool `json:"raThrottlingEnabled,omitempty"`

	// RsraGuardEnabled
	// Indicates whether RS/RA Guard is enabled or disabled
	RsraGuardEnabled *bool `json:"rsraGuardEnabled,omitempty"`

	// Support80211dEnabled
	// Indicates whether support for 802.11d is enabled or disabled
	Support80211dEnabled *bool `json:"support80211dEnabled,omitempty"`

	// Support80211kEnabled
	// Indicates whether support for 802.11k is enabled or disabled
	Support80211kEnabled *bool `json:"support80211kEnabled,omitempty"`

	// SuppressNsEnabled
	// Indicates whether supperssNS is enabled or disabled
	SuppressNsEnabled *bool `json:"suppressNsEnabled,omitempty"`

	// TransientClientMgmtEnable
	// Transient Client Management enabled.
	TransientClientMgmtEnable *bool `json:"transientClientMgmtEnable,omitempty"`

	// UnauthClientStatsEnabled
	// Indicates whether to send statistics of unauthorized clients or not
	UnauthClientStatsEnabled *bool `json:"unauthClientStatsEnabled,omitempty"`

	// UplinkEnabled
	// SSID Rate Limiting uplink enabled.
	UplinkEnabled *bool `json:"uplinkEnabled,omitempty"`

	// UplinkRate
	// SSID Rate Limiting uplink.
	UplinkRate *float64 `json:"uplinkRate,omitempty"`

	// UrlFilteringPolicyEnabled
	// Indicator of whether URL Filtering is enabled or disabled
	UrlFilteringPolicyEnabled *bool `json:"urlFilteringPolicyEnabled,omitempty"`

	// UrlFilteringPolicyId
	// The URL Filtering policy ID.
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`

	// WifiCallingPolicyEnabled
	// Indicator of whether Wi-Fi Calling is enabled or disabled
	WifiCallingPolicyEnabled *bool `json:"wifiCallingPolicyEnabled,omitempty"`

	// WifiCallingPolicyIds
	// The Wi-Fi Calling policy IDs. (Maximum allowed number is 5)
	WifiCallingPolicyIds []string `json:"wifiCallingPolicyIds"`
}

func NewWSGWLANAdvanced() *WSGWLANAdvanced {
	m := new(WSGWLANAdvanced)
	return m
}

type WSGWLANAuthentication struct {
	// AuthenticationOption
	// Option of the authentication service or profile, At least one ID or name or authenticationOption is required in the request. This only applies to hotspot and guest WLANs.
	// Constraints:
	//    - nullable
	//    - oneof:[Local DB,Guest,Always Accept]
	AuthenticationOption *string `json:"authenticationOption,omitempty" validate:"omitempty,oneof=Local DB Guest Always Accept"`

	// BackupAuthenticationId
	// Identifier of the backup authentication service or profile. At least one backupAuthenticationId or backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup authentication service.
	BackupAuthenticationId *string `json:"backupAuthenticationId,omitempty"`

	// BackupAuthenticationName
	// Name of the backup authentication service or profile. At least one backupAuthenticationId or backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup authentication service. Or could input the 'Always Accept'.
	BackupAuthenticationName *string `json:"backupAuthenticationName,omitempty"`

	// BackupAuthenticationOption
	// Option of the backup authentication service or profile, At least one backupAuthenticationId or backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup authentication service. This only applies to hotspot WLANs.
	// Constraints:
	//    - nullable
	//    - oneof:[Always Accept]
	BackupAuthenticationOption *string `json:"backupAuthenticationOption,omitempty" validate:"omitempty,oneof=Always Accept"`

	// Id
	// Identifier of the authentication service or profile. At least one ID or name or authenticationOption is required in the request.
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 location delivery support
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Name
	// Name of the authentication service or profile. At least one ID or name or authenticationOption is required in the request. Or could input the 'Always Accept' or 'Local DB'.
	Name *string `json:"name,omitempty"`

	RealmBasedAuth *bool `json:"realmBasedAuth,omitempty"`

	// ThroughController
	// Indicates whether authentication messages were sent through the controller or not
	ThroughController *bool `json:"throughController,omitempty"`
}

func NewWSGWLANAuthentication() *WSGWLANAuthentication {
	m := new(WSGWLANAuthentication)
	return m
}

type WSGWLANBssMinRateMbps string

func NewWSGWLANBssMinRateMbps() *WSGWLANBssMinRateMbps {
	m := new(WSGWLANBssMinRateMbps)
	return m
}

type WSGWLANConfiguration struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	// Description
	// Description of the WLAN
	Description *string `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	// Id
	// Identifier of the WLAN
	Id *string `json:"id,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the WLAN
	// Constraints:
	//    - oneof:[Standard_Open,Standard_8021X,Standard_Mac,Hotspot,Hotspot_MacByPass,Guest,WebAuth,Hotspot20,Hotspot20_Open,Hotspot20_OSEN]
	Type *string `json:"type,omitempty" validate:"oneof=Standard_Open Standard_8021X Standard_Mac Hotspot Hotspot_MacByPass Guest WebAuth Hotspot20 Hotspot20_Open Hotspot20_OSEN"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`

	// ZoneId
	// Identifier of the zone to which the WLAN belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGWLANConfiguration() *WSGWLANConfiguration {
	m := new(WSGWLANConfiguration)
	return m
}

type WSGWLANCoreTunnel struct {
	// Id
	// Identifier of the forwarding profile. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the forwarding profile. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	// Type
	// Tunnel type
	// Constraints:
	//    - required
	//    - oneof:[L2oGRE,Bridge,TTG_PDG]
	Type *string `json:"type" validate:"required,oneof=L2oGRE Bridge TTG_PDG"`
}

func NewWSGWLANCoreTunnel() *WSGWLANCoreTunnel {
	m := new(WSGWLANCoreTunnel)
	return m
}

type WSGWLANDSCPSetting struct {
	// Enable
	// Enabled or disabled
	// Constraints:
	//    - required
	Enable *bool `json:"enable" validate:"required"`

	Excepts []int `json:"excepts"`

	// High
	// DSCP range - high
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:255
	High *int `json:"high,omitempty" validate:"omitempty,gte=0,lte=255"`

	// Low
	// DSCP range - low
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:255
	Low *int `json:"low,omitempty" validate:"omitempty,gte=0,lte=255"`

	// Priority
	// Priority
	// Constraints:
	//    - required
	Priority *int `json:"priority" validate:"required"`
}

func NewWSGWLANDSCPSetting() *WSGWLANDSCPSetting {
	m := new(WSGWLANDSCPSetting)
	return m
}

type WSGWLANEncryption struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	// Constraints:
	//    - oneof:[AES,TKIP_AES,AES_GCMP_256]
	Algorithm *string `json:"algorithm,omitempty" validate:"oneof=AES TKIP_AES AES_GCMP_256"`

	// KeyIndex
	// Key index. This only applies to WEP64 and WEP128.
	KeyIndex *int `json:"keyIndex,omitempty"`

	// KeyInHex
	// Key in hex format. This only applies to WEP64 and WEP128.
	KeyInHex *string `json:"keyInHex,omitempty"`

	// Method
	// Encryption method
	// Constraints:
	//    - required
	//    - default:'None'
	//    - oneof:[WPA2,WPA_Mixed,WEP_64,WEP_128,None,WPA3,WPA23_Mixed,OWE]
	Method *string `json:"method" validate:"required,oneof=WPA2 WPA_Mixed WEP_64 WEP_128 None WPA3 WPA23_Mixed OWE"`

	// Mfp
	// Management frame protection. This only applies to WPA2 + AES or OWE method.
	// Constraints:
	//    - oneof:[disabled,capable,required]
	Mfp *string `json:"mfp,omitempty" validate:"oneof=disabled capable required"`

	// MobilityDomainId
	// mobility Domain Id.
	// Constraints:
	//    - min:1
	//    - max:65535
	MobilityDomainId *int `json:"mobilityDomainId,omitempty" validate:"gte=1,lte=65535"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	Passphrase *string `json:"passphrase,omitempty"`

	// SaePassphrase
	// saePassphrase. This only applies to WPA3 and WPA23 mixed mode.
	SaePassphrase *string `json:"saePassphrase,omitempty"`

	// Support80211rEnabled
	// Enable 802.11r Fast BSS Transition, fast Romaing.
	Support80211rEnabled *bool `json:"support80211rEnabled,omitempty"`
}

func NewWSGWLANEncryption() *WSGWLANEncryption {
	m := new(WSGWLANEncryption)
	return m
}

type WSGWLANHESSID string

func NewWSGWLANHESSID() *WSGWLANHESSID {
	m := new(WSGWLANHESSID)
	return m
}

type WSGWLANList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANSummary `json:"list"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANList() *WSGWLANList {
	m := new(WSGWLANList)
	return m
}

type WSGWLANMACAuth struct {
	// CustomizedPassword
	// User defined password. When this field is set to an empty string, the MAC address is used as password.
	// Constraints:
	//    - max:64
	CustomizedPassword *string `json:"customizedPassword,omitempty" validate:"max=64"`

	// MacAuthMacFormat
	// MAC address format. The default format is 0010a42319c0 and the 802.1X format is 00-10-A4-23-19-C0.
	// Constraints:
	//    - oneof:[Default,802.1X,UpperColon,Upper,LowerDash,LowerColon]
	MacAuthMacFormat *string `json:"macAuthMacFormat,omitempty" validate:"oneof=Default 802.1X UpperColon Upper LowerDash LowerColon"`
}

func NewWSGWLANMACAuth() *WSGWLANMACAuth {
	m := new(WSGWLANMACAuth)
	return m
}

type WSGWLANMgmtTxRateMbps string

func NewWSGWLANMgmtTxRateMbps() *WSGWLANMgmtTxRateMbps {
	m := new(WSGWLANMgmtTxRateMbps)
	return m
}

type WSGWLANNameSSID string

func NewWSGWLANNameSSID() *WSGWLANNameSSID {
	m := new(WSGWLANNameSSID)
	return m
}

type WSGWLANRadius struct {
	// CalledStaIdType
	// Called station ID type
	// Constraints:
	//    - default:'WLAN_BSSID'
	//    - oneof:[WLAN_BSSID,AP_MAC,NONE,AP_GROUP]
	CalledStaIdType *string `json:"calledStaIdType,omitempty" validate:"oneof=WLAN_BSSID AP_MAC NONE AP_GROUP"`

	// CustomizedNasId
	// User defined NAS ID
	// Constraints:
	//    - max:64
	CustomizedNasId *string `json:"customizedNasId,omitempty" validate:"max=64"`

	// NasIdType
	// NAS ID type
	// Constraints:
	//    - default:'WLAN_BSSID'
	//    - oneof:[WLAN_BSSID,AP_MAC,Customized]
	NasIdType *string `json:"nasIdType,omitempty" validate:"oneof=WLAN_BSSID AP_MAC Customized"`

	// NasIpType
	// NAS IP type
	// Constraints:
	//    - default:'disabled'
	//    - oneof:[disabled,control,management,userDefined]
	NasIpType *string `json:"nasIpType,omitempty" validate:"oneof=disabled control management userDefined"`

	// NasIpUserDefined
	// User-defined NAS IP
	// Constraints:
	//    - max:45
	NasIpUserDefined *string `json:"nasIpUserDefined,omitempty" validate:"max=45"`

	// NasMaxRetry
	// NAS request maximum retry
	// Constraints:
	//    - default:2
	//    - min:2
	//    - max:10
	NasMaxRetry *int `json:"nasMaxRetry,omitempty" validate:"gte=2,lte=10"`

	// NasReconnectPrimaryMin
	// NAS reconnect primary time in minutes
	// Constraints:
	//    - default:5
	//    - min:1
	//    - max:60
	NasReconnectPrimaryMin *int `json:"nasReconnectPrimaryMin,omitempty" validate:"gte=1,lte=60"`

	// NasRequestTimeoutSec
	// NAS request timeout in seconds
	// Constraints:
	//    - default:3
	//    - min:2
	//    - max:20
	NasRequestTimeoutSec *int `json:"nasRequestTimeoutSec,omitempty" validate:"gte=2,lte=20"`

	// SingleSessionIdAcctEnabled
	// When Single Accounting Session ID is enabled, APs will maintain one accounting session for client roaming
	SingleSessionIdAcctEnabled *bool `json:"singleSessionIdAcctEnabled,omitempty"`

	// VendorSpecificAttributeProfileId
	// Vendor Specific Attribute Profile ID
	VendorSpecificAttributeProfileId *string `json:"vendorSpecificAttributeProfileId,omitempty"`
}

func NewWSGWLANRadius() *WSGWLANRadius {
	m := new(WSGWLANRadius)
	return m
}

type WSGWLANSchedule struct {
	// Id
	// Identifier of the schedule profile. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the schedule profile. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	// Type
	// Type of WLAN schedule
	// Constraints:
	//    - required
	//    - default:'AlwaysOn'
	//    - oneof:[AlwaysOn,AlwaysOff,Customized]
	Type *string `json:"type" validate:"required,oneof=AlwaysOn AlwaysOff Customized"`
}

func NewWSGWLANSchedule() *WSGWLANSchedule {
	m := new(WSGWLANSchedule)
	return m
}

type WSGWLANSummary struct {
	// Id
	// Identifier of the WLAN
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGWLANSummary() *WSGWLANSummary {
	m := new(WSGWLANSummary)
	return m
}

type WSGWLANVlan struct {
	// AaaVlanOverride
	// Indicates whether the AAA VLAN settings can be overriden or not
	AaaVlanOverride *bool `json:"aaaVlanOverride,omitempty"`

	// AccessVlan
	// Access VLAN ID
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:4094
	AccessVlan *int `json:"accessVlan,omitempty" validate:"gte=1,lte=4094"`

	// CoreQinQEnabled
	// Indicates whether Q-in-Q is allowed at the core network or not
	CoreQinQEnabled *bool `json:"coreQinQEnabled,omitempty"`

	// CoreSVlan
	// Core SVLAN ID. This only applies when core Q-in-Q is enabled
	// Constraints:
	//    - min:1
	//    - max:4094
	CoreSVlan *int `json:"coreSVlan,omitempty" validate:"gte=1,lte=4094"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGWLANVlan() *WSGWLANVlan {
	m := new(WSGWLANVlan)
	return m
}

// AddRkszonesWlansByZoneId
//
// Use this API command to create a new standard, open and non-tunneled basic WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateStandardOpenWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansByZoneId(ctx context.Context, body *WSGWLANCreateStandardOpenWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansGuestByZoneId
//
// Use this API command to create a new guest access WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateGuestAccessWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansGuestByZoneId(ctx context.Context, body *WSGWLANCreateGuestAccessWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansGuestByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansHotspot20ByZoneId
//
// Use this API command to create a new Hotspot 2.0 access WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateHotspot20Wlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20ByZoneId(ctx context.Context, body *WSGWLANCreateHotspot20Wlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansHotspot20ByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansHotspot20openByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as 'Open'.
//
// Request Body:
//	 - body *WSGWLANCreateHotspot20OpenWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20openByZoneId(ctx context.Context, body *WSGWLANCreateHotspot20OpenWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansHotspot20openByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansHotspot20osenByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as '802.1X'.
//
// Request Body:
//	 - body *WSGWLANCreateHotspot20OpenWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20osenByZoneId(ctx context.Context, body *WSGWLANCreateHotspot20OpenWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansHotspot20osenByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansQosMapsById
//
// Use this API command to enable Qos Map Set of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansQosMapsById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansQosMapsById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// AddRkszonesWlansStandard8021XByZoneId
//
// Use this API command to create a new standard, 802.1X and non-tunneled WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateStandard80211Wlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandard8021XByZoneId(ctx context.Context, body *WSGWLANCreateStandard80211Wlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansStandard8021XByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansStandard8021XmacByZoneId
//
// Use this API command to create a new standard, 802.1X with MAC address and non-tunneled WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateStandard80211Wlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandard8021XmacByZoneId(ctx context.Context, body *WSGWLANCreateStandard80211Wlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansStandard8021XmacByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansStandardmacByZoneId
//
// Use this API command to create a new standard, MAC auth and non-tunneled WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateStandard80211Wlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandardmacByZoneId(ctx context.Context, body *WSGWLANCreateStandard80211Wlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansStandardmacByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansWebauthByZoneId
//
// Use this API command to creates new web authentication WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateWebAuthWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWebauthByZoneId(ctx context.Context, body *WSGWLANCreateWebAuthWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansWebauthByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansWechatByZoneId
//
// Use this API command to create a new wechat WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateWechatWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWechatByZoneId(ctx context.Context, body *WSGWLANCreateWechatWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansWechatByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansWispr8021XByZoneId
//
// Use this API command to create a new hotspot (WISPr) with 802.1X WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateHotspotWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWispr8021XByZoneId(ctx context.Context, body *WSGWLANCreateHotspotWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansWispr8021XByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansWisprByZoneId
//
// Use this API command to create new hotspot (WISPr) WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateHotspotWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWisprByZoneId(ctx context.Context, body *WSGWLANCreateHotspotWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansWisprByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesWlansWisprmacByZoneId
//
// Use this API command to create a new hotspot (WISPr) with MAC bypass WLAN.
//
// Request Body:
//	 - body *WSGWLANCreateHotspotWlan
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWisprmacByZoneId(ctx context.Context, body *WSGWLANCreateHotspotWlan, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlansWisprmacByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesWlansAccountingServiceOrProfileById
//
// Use this API command to disable the accounting of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansAccountingServiceOrProfileById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansAccountingServiceOrProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlansById
//
// Use this API command to delete a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlansDevicePolicyById
//
// Use this API command to disable the device policy of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDevicePolicyById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansDevicePolicyById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlansDiffServProfileById
//
// Use this API command to disable the DiffServ profile of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDiffServProfileById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansDiffServProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlansDnsServerProfileById
//
// Use this API command to disable DNS server profile of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDnsServerProfileById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansDnsServerProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlansL2ACLById
//
// Use this API command to disable the layer 2 access control list (ACL) configuration of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansL2ACLById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansL2ACLById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesWlansQosMapsById
//
// Use this API command to disable Qos Map Set of a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansQosMapsById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansQosMapsById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesWlansById
//
// Use this API command to retrieve a WLAN.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) FindRkszonesWlansById(ctx context.Context, id string, zoneId string) (*WSGWLANConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlansById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlansByZoneId
//
// Use this API command to retrieve a list of WLANs within a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWLANService) FindRkszonesWlansByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGWLANList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlansByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindWlanByQueryCriteria
//
// Query WLANs with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWLANService) FindWlanByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWLANQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindWlanByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesWlansById
//
// Use this API command to modify the configuration of a WLAN.
//
// Request Body:
//	 - body *WSGWLANModifyWlan
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) PartialUpdateRkszonesWlansById(ctx context.Context, body *WSGWLANModifyWlan, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlansById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesWlansById
//
// Use this API command to modify entire information of a WLAN.
//
// Request Body:
//	 - body *WSGWLANModifyWlan
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) UpdateRkszonesWlansById(ctx context.Context, body *WSGWLANModifyWlan, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesWlansById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
