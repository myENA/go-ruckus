package vsz

// API Version: v9_0

type WSGZoneApFirmware struct {
	// FirmwareVersion
	// version of the AP firmare
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Supported
	// version of the AP firmare is supported for Upgrade or Downgrade.
	Supported *bool `json:"supported,omitempty"`

	// UnsupportedApModelSummary
	// summary of the AP Model is unsupported for AP firmware version.
	UnsupportedApModelSummary []*WSGZoneUnsupportedApModel `json:"unsupportedApModelSummary,omitempty"`
}

func NewWSGZoneApFirmware() *WSGZoneApFirmware {
	m := new(WSGZoneApFirmware)
	return m
}

type WSGZoneApFirmwareList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGZoneApFirmware `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneApFirmwareList() *WSGZoneApFirmwareList {
	m := new(WSGZoneApFirmwareList)
	return m
}

type WSGZoneApLogin struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName *WSGCommonApLoginName `json:"apLoginName"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword *WSGCommonApLoginPassword `json:"apLoginPassword"`
}

func NewWSGZoneApLogin() *WSGZoneApLogin {
	m := new(WSGZoneApLogin)
	return m
}

type WSGZoneApSnmpOptions struct {
	// ApSnmpEnabled
	// Enable AP SNMP
	ApSnmpEnabled *bool `json:"apSnmpEnabled,omitempty"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	SnmpV2Agent []*WSGCommonSnmpCommunity `json:"snmpV2Agent,omitempty"`

	// SnmpV3Agent
	// User List of the SNMP V3 Agent.
	SnmpV3Agent []*WSGZoneSnmpUser `json:"snmpV3Agent,omitempty"`
}

func NewWSGZoneApSnmpOptions() *WSGZoneApSnmpOptions {
	m := new(WSGZoneApSnmpOptions)
	return m
}

type WSGZoneAvailableTunnelProfile struct {
	// AaaAffinityEnabled
	// Enable AAA affinity (Soft GRE only)
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// Id
	// Tunnel Profile ID
	Id *string `json:"id,omitempty"`

	IpMode *WSGCommonIpMode `json:"ipMode,omitempty"`

	// Name
	// Tunnel Profile Name
	Name *string `json:"name,omitempty"`

	// TunnelType
	// Tunnel Profile Type ("RuckusGRE", "SoftGRE",or "Ipsec")
	// Constraints:
	//    - oneof:[RuckusGRE,SoftGRE,Ipsec]
	TunnelType *string `json:"tunnelType,omitempty"`
}

func NewWSGZoneAvailableTunnelProfile() *WSGZoneAvailableTunnelProfile {
	m := new(WSGZoneAvailableTunnelProfile)
	return m
}

type WSGZoneAvailableTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGZoneAvailableTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneAvailableTunnelProfileList() *WSGZoneAvailableTunnelProfileList {
	m := new(WSGZoneAvailableTunnelProfileList)
	return m
}

type WSGZoneBackgroundScanning struct {
	// FrequencyInSec
	// Frequency in second
	// Constraints:
	//    - default:20
	//    - min:1
	//    - max:65535
	FrequencyInSec *int `json:"frequencyInSec,omitempty"`
}

func NewWSGZoneBackgroundScanning() *WSGZoneBackgroundScanning {
	m := new(WSGZoneBackgroundScanning)
	return m
}

type WSGZoneBandBalancing struct {
	// Wifi24Percentage
	// Percentage of client load on 2.4GHz radio band
	// Constraints:
	//    - default:25
	//    - min:0
	//    - max:100
	Wifi24Percentage *int `json:"wifi24Percentage,omitempty"`
}

func NewWSGZoneBandBalancing() *WSGZoneBandBalancing {
	m := new(WSGZoneBandBalancing)
	return m
}

type WSGZoneBonjourGatewayPolicyConfiguration struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*WSGZoneBonjourPolicyRuleConfiguration `json:"bonjourPolicyRuleList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGZoneBonjourGatewayPolicyConfiguration() *WSGZoneBonjourGatewayPolicyConfiguration {
	m := new(WSGZoneBonjourGatewayPolicyConfiguration)
	return m
}

type WSGZoneBonjourGatewayPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGZoneBonjourGatewayPolicySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneBonjourGatewayPolicyList() *WSGZoneBonjourGatewayPolicyList {
	m := new(WSGZoneBonjourGatewayPolicyList)
	return m
}

type WSGZoneBonjourGatewayPolicySummary struct {
	// Description
	// Description of the bonjour gateway policy
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the bonjour gateway policy
	Id *string `json:"id,omitempty"`

	// LastModifiedBy
	// Last modified user of the bonjour gateway policy
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedOn
	// Last modified time of the bonjour gateway policy
	LastModifiedOn *string `json:"lastModifiedOn,omitempty"`

	// Name
	// Name of the bonjour gateway policy
	Name *string `json:"name,omitempty"`
}

func NewWSGZoneBonjourGatewayPolicySummary() *WSGZoneBonjourGatewayPolicySummary {
	m := new(WSGZoneBonjourGatewayPolicySummary)
	return m
}

// WSGZoneBonjourPolicyRule
//
// Bonjour policy rule
type WSGZoneBonjourPolicyRule struct {
	// BridgeService
	// Constraints:
	//    - required
	BridgeService *WSGProfileBridgeService `json:"bridgeService"`

	// FromVlan
	// From VLAN
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	FromVlan *int `json:"fromVlan"`

	// Notes
	// Notes
	Notes *string `json:"notes,omitempty"`

	// Protocol
	// protocol. This is only available when bridgeService is OTHER
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	ToVlan *int `json:"toVlan"`
}

func NewWSGZoneBonjourPolicyRule() *WSGZoneBonjourPolicyRule {
	m := new(WSGZoneBonjourPolicyRule)
	return m
}

// WSGZoneBonjourPolicyRuleConfiguration
//
// Bonjour policy rule
type WSGZoneBonjourPolicyRuleConfiguration struct {
	// BridgeService
	// Bridge service
	BridgeService *string `json:"bridgeService,omitempty"`

	// FromVlan
	// From VLAN
	FromVlan *int `json:"fromVlan,omitempty"`

	// Notes
	// Notes
	Notes *string `json:"notes,omitempty"`

	// Priority
	// Priority
	Priority *string `json:"priority,omitempty"`

	// Protocol
	// protocol
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	ToVlan *int `json:"toVlan,omitempty"`
}

func NewWSGZoneBonjourPolicyRuleConfiguration() *WSGZoneBonjourPolicyRuleConfiguration {
	m := new(WSGZoneBonjourPolicyRuleConfiguration)
	return m
}

type WSGZoneClientLoadBalancing struct {
	// AdjacentRadioThreshold
	// Adjacent radio threshold
	// Constraints:
	//    - min:1
	//    - max:100
	AdjacentRadioThreshold *int `json:"adjacentRadioThreshold,omitempty"`
}

func NewWSGZoneClientLoadBalancing() *WSGZoneClientLoadBalancing {
	m := new(WSGZoneClientLoadBalancing)
	return m
}

type WSGZoneCreateBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*WSGZoneBonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGZoneCreateBonjourGatewayPolicy() *WSGZoneCreateBonjourGatewayPolicy {
	m := new(WSGZoneCreateBonjourGatewayPolicy)
	return m
}

type WSGZoneCreateDiffServProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkDiffServ *WSGZoneDownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *WSGZoneUplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewWSGZoneCreateDiffServProfile() *WSGZoneCreateDiffServProfile {
	m := new(WSGZoneCreateDiffServProfile)
	return m
}

type WSGZoneCreateZone struct {
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

	ApLatencyInterval *WSGCommonApLatencyInterval `json:"apLatencyInterval,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	ApRebootTimeout *WSGCommonApRebootTimeout `json:"apRebootTimeout,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	BackgroundScanning24 *WSGZoneBackgroundScanning `json:"backgroundScanning24,omitempty"`

	BackgroundScanning50 *WSGZoneBackgroundScanning `json:"backgroundScanning50,omitempty"`

	BonjourFencingPolicy *WSGCommonGenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
	Channel144Enabled *bool `json:"channel144Enabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	ClientAdmissionControl24 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone.
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	CountryCode *string `json:"countryCode,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code.
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DhcpSiteConfig *WSGCommonDhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	DomainId *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// EnforcePriorityZoneAffinityEnable
	// Enforced the priority of Affinity Profile.
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*WSGCommonGenericRef `json:"ipsecProfiles,omitempty"`

	// IpsecTunnelMode
	// Constraints:
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	LoadBalancing *WSGZoneLoadBalancing `json:"loadBalancing,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonGenericRef `json:"locationBasedService,omitempty"`

	// Login
	// Constraints:
	//    - required
	Login *WSGZoneApLogin `json:"login"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mesh *WSGZoneMeshConfiguration `json:"mesh,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	NodeAffinityProfile *WSGCommonGenericRef `json:"nodeAffinityProfile,omitempty"`

	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	Rogue *WSGZoneRogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	RuckusGreTunnelProfile *WSGCommonGenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	SmartMonitor *WSGCommonSmartMonitor `json:"smartMonitor,omitempty"`

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*WSGZoneSoftGreRef `json:"softGreTunnelProflies,omitempty"`

	// SshTunnelEncryption
	// Constraints:
	//    - default:'AES128'
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty"`

	Syslog *WSGZoneSyslog `json:"syslog,omitempty"`

	Timezone *WSGZoneTimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *WSGCommonZoneTunnelType `json:"tunnelType,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Version *WSGCommonFirmwareVersion `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// Zone affinity profile of the zone
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewWSGZoneCreateZone() *WSGZoneCreateZone {
	m := new(WSGZoneCreateZone)
	return m
}

type WSGZoneCustomizedTimeZone struct {
	// Abbreviation
	// Time zone abbreviation
	// Constraints:
	//    - required
	Abbreviation *string `json:"abbreviation"`

	End *WSGZoneDaylightSavingTime `json:"end,omitempty"`

	// GmtOffset
	// GMT offset
	// Constraints:
	//    - required
	//    - min:-11
	//    - max:14
	GmtOffset *int `json:"gmtOffset"`

	// GmtOffsetMinute
	// GMT offset minute
	// Constraints:
	//    - required
	//    - min:0
	//    - max:59
	GmtOffsetMinute *int `json:"gmtOffsetMinute"`

	Start *WSGZoneDaylightSavingTime `json:"start,omitempty"`
}

func NewWSGZoneCustomizedTimeZone() *WSGZoneCustomizedTimeZone {
	m := new(WSGZoneCustomizedTimeZone)
	return m
}

type WSGZoneDaylightSavingTime struct {
	// Day
	// Day of the week (0 for Sunday, 1 for Monday, 2 for Tuesday, and so on)
	// Constraints:
	//    - required
	//    - oneof:[0,1,2,3,4,5,6]
	Day *int `json:"day"`

	// Hour
	// Hour of the day
	// Constraints:
	//    - required
	//    - min:0
	//    - max:23
	Hour *int `json:"hour"`

	// Month
	// Month when daylight saving time begins
	// Constraints:
	//    - required
	//    - oneof:[1,2,3,4,5,6,7,8,9,10,11,12]
	Month *int `json:"month"`

	// Week
	// Week of the month (1 for the first week, 2 for the second week, and so on)
	// Constraints:
	//    - required
	//    - oneof:[1,2,3,4,5]
	Week *int `json:"week"`
}

func NewWSGZoneDaylightSavingTime() *WSGZoneDaylightSavingTime {
	m := new(WSGZoneDaylightSavingTime)
	return m
}

type WSGZoneDhcpSiteConfigList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCommonDhcpSiteConfigListRef `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneDhcpSiteConfigList() *WSGZoneDhcpSiteConfigList {
	m := new(WSGZoneDhcpSiteConfigList)
	return m
}

type WSGZoneDiffServConfiguration struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkDiffServ *WSGZoneDownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Id
	// Identifier of the zone
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *WSGZoneUplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewWSGZoneDiffServConfiguration() *WSGZoneDiffServConfiguration {
	m := new(WSGZoneDiffServConfiguration)
	return m
}

type WSGZoneDiffServList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGZoneDiffServSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneDiffServList() *WSGZoneDiffServList {
	m := new(WSGZoneDiffServList)
	return m
}

type WSGZoneDiffServSummary struct {
	// Id
	// Identifier of the diff serv
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the diff serv
	Name *string `json:"name,omitempty"`
}

func NewWSGZoneDiffServSummary() *WSGZoneDiffServSummary {
	m := new(WSGZoneDiffServSummary)
	return m
}

type WSGZoneDownlinkDiffServ struct {
	// Downlink
	// Downlink
	Downlink *string `json:"downlink,omitempty"`

	// DownlinkEnable
	// Downlink enable
	DownlinkEnable *bool `json:"downlinkEnable,omitempty"`
}

func NewWSGZoneDownlinkDiffServ() *WSGZoneDownlinkDiffServ {
	m := new(WSGZoneDownlinkDiffServ)
	return m
}

type WSGZoneLoadBalancing struct {
	BandBalancing *WSGZoneBandBalancing `json:"bandBalancing,omitempty"`

	ClientLoadBalancing24 *WSGZoneClientLoadBalancing `json:"clientLoadBalancing24,omitempty"`

	ClientLoadBalancing50 *WSGZoneClientLoadBalancing `json:"clientLoadBalancing50,omitempty"`

	// LoadBalancingMethod
	// Constraints:
	//    - default:'BASED_ON_CLIENT_COUNT'
	//    - oneof:[BASED_ON_CLIENT_COUNT,BASED_ON_CAPACITY,OFF]
	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty"`

	// SteeringMode
	// Steering Mode: BASIC-Withholds probe and authentication responses at connetcion time in heavily loaded band to balance clients to the other band, PROACTIVE-Uses BASIC functionality and actively rebalances clients via 802.11v BTM, STRICT-Uses PROACTIVE functionality and forcefully rebalances clients via 802.11v BTM
	// Constraints:
	//    - default:'BASIC'
	//    - oneof:[BASIC,PROACTIVE,STRICT]
	SteeringMode *string `json:"steeringMode,omitempty"`
}

func NewWSGZoneLoadBalancing() *WSGZoneLoadBalancing {
	m := new(WSGZoneLoadBalancing)
	return m
}

type WSGZoneMeshConfiguration struct {
	// MeshRadioIdx
	// Mesh radio index
	// Constraints:
	//    - default:'Radio5G'
	//    - oneof:[Radio24G,Radio5G]
	MeshRadioIdx *string `json:"meshRadioIdx,omitempty"`

	// Passphrase
	// Passphrase for the mesh network
	Passphrase *string `json:"passphrase,omitempty"`

	// Ssid
	// SSID of the mesh network
	Ssid *string `json:"ssid,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewWSGZoneMeshConfiguration() *WSGZoneMeshConfiguration {
	m := new(WSGZoneMeshConfiguration)
	return m
}

type WSGZoneModfiyApFirmware struct {
	// FirmwareVersion
	// new version of the AP firmare
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
}

func NewWSGZoneModfiyApFirmware() *WSGZoneModfiyApFirmware {
	m := new(WSGZoneModfiyApFirmware)
	return m
}

type WSGZoneModifyBonjourGatewayEnable struct {
	// EnabledBonjourGateway
	// Enable Bonjour gateway on th AP
	// Constraints:
	//    - required
	EnabledBonjourGateway *bool `json:"enabledBonjourGateway"`
}

func NewWSGZoneModifyBonjourGatewayEnable() *WSGZoneModifyBonjourGatewayEnable {
	m := new(WSGZoneModifyBonjourGatewayEnable)
	return m
}

type WSGZoneModifyBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*WSGZoneBonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGZoneModifyBonjourGatewayPolicy() *WSGZoneModifyBonjourGatewayPolicy {
	m := new(WSGZoneModifyBonjourGatewayPolicy)
	return m
}

type WSGZoneModifyDiffServProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkDiffServ *WSGZoneDownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *WSGZoneUplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewWSGZoneModifyDiffServProfile() *WSGZoneModifyDiffServProfile {
	m := new(WSGZoneModifyDiffServProfile)
	return m
}

type WSGZoneModifyZone struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

	ApLatencyInterval *WSGCommonApLatencyInterval `json:"apLatencyInterval,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	ApRebootTimeout *WSGCommonApRebootTimeout `json:"apRebootTimeout,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	BackgroundScanning24 *WSGZoneBackgroundScanning `json:"backgroundScanning24,omitempty"`

	BackgroundScanning50 *WSGZoneBackgroundScanning `json:"backgroundScanning50,omitempty"`

	BonjourFencingPolicy *WSGCommonGenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
	Channel144Enabled *bool `json:"channel144Enabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	ClientAdmissionControl24 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	CountryCode *string `json:"countryCode,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DhcpSiteConfig *WSGCommonDhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	DomainId *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// EnforcePriorityZoneAffinityEnable
	// Enforce the priority of zone affinity
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*WSGCommonGenericRef `json:"ipsecProfiles,omitempty"`

	// IpsecTunnelMode
	// Constraints:
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	LoadBalancing *WSGZoneLoadBalancing `json:"loadBalancing,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonGenericRef `json:"locationBasedService,omitempty"`

	Login *WSGZoneApLogin `json:"login,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mesh *WSGZoneMeshConfiguration `json:"mesh,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	NodeAffinityProfile *WSGCommonGenericRef `json:"nodeAffinityProfile,omitempty"`

	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RestrictedApAccessEnabled
	// Enable Restricted AP Access of the zone.
	RestrictedApAccessEnabled *bool `json:"restrictedApAccessEnabled,omitempty"`

	// RestrictedApAccessProfileId
	// Restricted AP Access Profile Id of the zone.
	RestrictedApAccessProfileId *string `json:"restrictedApAccessProfileId,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	Rogue *WSGZoneRogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	RuckusGreTunnelProfile *WSGCommonGenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	SmartMonitor *WSGCommonSmartMonitor `json:"smartMonitor,omitempty"`

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*WSGZoneSoftGreRef `json:"softGreTunnelProflies,omitempty"`

	// SshTunnelEncryption
	// Constraints:
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty"`

	Syslog *WSGZoneSyslog `json:"syslog,omitempty"`

	Timezone *WSGZoneTimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *WSGCommonZoneTunnelType `json:"tunnelType,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// Identifier of the ZoneAffinityProfile
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewWSGZoneModifyZone() *WSGZoneModifyZone {
	m := new(WSGZoneModifyZone)
	return m
}

type WSGZoneQueryCriteria struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []*WSGZoneQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGZoneQueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information.
	Options *WSGZoneQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGZoneQueryCriteria() *WSGZoneQueryCriteria {
	m := new(WSGZoneQueryCriteria)
	return m
}

type WSGZoneQueryCriteriaExtraFiltersType struct {
	Operator *WSGCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attribute
	// Constraints:
	//    - required
	//    - oneof:[VERSION]
	Type *string `json:"type"`

	// Value
	// Value for filtering
	// Constraints:
	//    - required
	Value *string `json:"value"`
}

func NewWSGZoneQueryCriteriaExtraFiltersType() *WSGZoneQueryCriteriaExtraFiltersType {
	m := new(WSGZoneQueryCriteriaExtraFiltersType)
	return m
}

type WSGZoneQueryCriteriaFiltersType struct {
	// Operator
	// Operator for filtering
	// Constraints:
	//    - oneof:[eq]
	Operator *string `json:"operator,omitempty"`

	// Type
	// Group type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,SWITCH_GROUP]
	Type *string `json:"type"`

	// Value
	// Group ID
	// Constraints:
	//    - required
	Value *string `json:"value"`
}

func NewWSGZoneQueryCriteriaFiltersType() *WSGZoneQueryCriteriaFiltersType {
	m := new(WSGZoneQueryCriteriaFiltersType)
	return m
}

// WSGZoneQueryCriteriaOptionsType
//
// Specified feature required information.
type WSGZoneQueryCriteriaOptionsType struct {
	// IncludeSharedResources
	// Include the resources of parent domain as well while querying.
	IncludeSharedResources *bool `json:"includeSharedResources,omitempty"`

	Zoneipmode *WSGCommonIpMode `json:"zone_ipmode,omitempty"`
}

func NewWSGZoneQueryCriteriaOptionsType() *WSGZoneQueryCriteriaOptionsType {
	m := new(WSGZoneQueryCriteriaOptionsType)
	return m
}

type WSGZoneRogue struct {
	// MaliciousTypes
	// Malicious type when reportType is Malicious
	// Constraints:
	//    - nullable
	MaliciousTypes []string `json:"maliciousTypes,omitempty"`

	// ProtectionEnabled
	// Protection enabled
	ProtectionEnabled *bool `json:"protectionEnabled,omitempty"`

	// ReportType
	// Report type
	// Constraints:
	//    - oneof:[All,Malicious]
	ReportType *string `json:"reportType,omitempty"`
}

func NewWSGZoneRogue() *WSGZoneRogue {
	m := new(WSGZoneRogue)
	return m
}

type WSGZoneSnmpUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	// Constraints:
	//    - min:8
	AuthPassword *string `json:"authPassword,omitempty"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	// Constraints:
	//    - oneof:[NONE,MD5,SHA]
	AuthProtocol *string `json:"authProtocol,omitempty"`

	// NotificationEnabled
	// notification privilege of the SNMP User
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	NotificationTarget []*WSGCommonTargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	// Constraints:
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty"`

	// PrivPassword
	// privPassword of the SNMP User.
	// Constraints:
	//    - min:8
	PrivPassword *string `json:"privPassword,omitempty"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	// Constraints:
	//    - oneof:[NONE,DES,AES]
	PrivProtocol *string `json:"privProtocol,omitempty"`

	// ReadEnabled
	// read privilege of the SNMP User
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	// Constraints:
	//    - required
	UserName *string `json:"userName"`

	// WriteEnabled
	// write privilege of the SNMP User
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

func NewWSGZoneSnmpUser() *WSGZoneSnmpUser {
	m := new(WSGZoneSnmpUser)
	return m
}

type WSGZoneSoftGreRef struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGZoneSoftGreRef() *WSGZoneSoftGreRef {
	m := new(WSGZoneSoftGreRef)
	return m
}

type WSGZoneSyslog struct {
	Address *WSGCommonIpAddress `json:"address,omitempty"`

	// Facility
	// Facility of the syslog server
	// Constraints:
	//    - default:'Keep_Original'
	//    - oneof:[Keep_Original,Local0,Local1,Local2,Local3,Local4,Local5,Local6,Local7]
	Facility *string `json:"facility,omitempty"`

	// FlowLevel
	// Flow Level of the syslog
	// Constraints:
	//    - default:'GENERAL_LOGS'
	//    - oneof:[GENERAL_LOGS,CLIENT_FLOW,ALL]
	FlowLevel *string `json:"flowLevel,omitempty"`

	// Port
	// Port number of the syslog server
	// Constraints:
	//    - default:514
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// Priority
	// Priority of the log messages
	// Constraints:
	//    - default:'Error'
	//    - oneof:[Emergency,Alert,Critical,Error,Warning,Notice,Info,All]
	Priority *string `json:"priority,omitempty"`

	// Protocol
	// Protocol of the syslog server
	// Constraints:
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	Protocol *string `json:"protocol,omitempty"`

	SecondaryAddress *WSGCommonIpAddress `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	// Constraints:
	//    - default:514
	//    - min:1
	//    - max:65535
	SecondaryPort *int `json:"secondaryPort,omitempty"`

	// SecondaryProtocol
	// Secondary Server Protocol of the syslog server
	// Constraints:
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty"`
}

func NewWSGZoneSyslog() *WSGZoneSyslog {
	m := new(WSGZoneSyslog)
	return m
}

type WSGZoneTimezoneSetting struct {
	CustomizedTimezone *WSGZoneCustomizedTimeZone `json:"customizedTimezone,omitempty"`

	// SystemTimezone
	// System defined time zone, please refer to the 'Overview > Time Zone' list
	SystemTimezone *string `json:"systemTimezone,omitempty"`
}

func NewWSGZoneTimezoneSetting() *WSGZoneTimezoneSetting {
	m := new(WSGZoneTimezoneSetting)
	return m
}

type WSGZoneUnsupportedApModel struct {
	// Amount
	// amount of the AP Model
	Amount *int `json:"amount,omitempty"`

	// Model
	// name of the AP Model
	Model *string `json:"model,omitempty"`
}

func NewWSGZoneUnsupportedApModel() *WSGZoneUnsupportedApModel {
	m := new(WSGZoneUnsupportedApModel)
	return m
}

type WSGZoneUplinkDiffServ struct {
	// Uplink
	// Uplink
	Uplink *string `json:"uplink,omitempty"`

	// UplinkEnable
	// Uplink enable
	UplinkEnable *bool `json:"uplinkEnable,omitempty"`
}

func NewWSGZoneUplinkDiffServ() *WSGZoneUplinkDiffServ {
	m := new(WSGZoneUplinkDiffServ)
	return m
}

type WSGZoneConfiguration struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

	ApLatencyInterval *WSGCommonApLatencyInterval `json:"apLatencyInterval,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	ApRebootTimeout *WSGCommonApRebootTimeout `json:"apRebootTimeout,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	BackgroundScanning24 *WSGZoneBackgroundScanning `json:"backgroundScanning24,omitempty"`

	BackgroundScanning50 *WSGZoneBackgroundScanning `json:"backgroundScanning50,omitempty"`

	BonjourFencingPolicy *WSGCommonGenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
	Channel144Enabled *bool `json:"channel144Enabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	ClientAdmissionControl24 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	CountryCode *string `json:"countryCode,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DhcpSiteConfig *WSGCommonDhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	DomainId *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	// Id
	// Identifier of the zone
	Id *string `json:"id,omitempty"`

	IpMode *WSGCommonIpMode `json:"ipMode,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*WSGCommonGenericRef `json:"ipsecProfiles,omitempty"`

	// IpsecTunnelMode
	// Constraints:
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	LoadBalancing *WSGZoneLoadBalancing `json:"loadBalancing,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonGenericRef `json:"locationBasedService,omitempty"`

	Login *WSGZoneApLogin `json:"login,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mesh *WSGZoneMeshConfiguration `json:"mesh,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	NodeAffinityProfile *WSGCommonGenericRef `json:"nodeAffinityProfile,omitempty"`

	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RestrictedApAccessEnabled
	// Enable Restricted AP Access of the zone.
	RestrictedApAccessEnabled *bool `json:"restrictedApAccessEnabled,omitempty"`

	// RestrictedApAccessProfileId
	// Restricted AP Access Profile Id of the zone.
	RestrictedApAccessProfileId *string `json:"restrictedApAccessProfileId,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	Rogue *WSGZoneRogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	RuckusGreTunnelProfile *WSGCommonGenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	SmartMonitor *WSGCommonSmartMonitor `json:"smartMonitor,omitempty"`

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*WSGZoneSoftGreRef `json:"softGreTunnelProflies,omitempty"`

	// SshTunnelEncryption
	// Constraints:
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty"`

	Syslog *WSGZoneSyslog `json:"syslog,omitempty"`

	Timezone *WSGZoneTimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *WSGCommonZoneTunnelType `json:"tunnelType,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Version *WSGCommonFirmwareVersion `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *WSGCommonRadio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50SuperSet `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// zone affinity profile Id
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewWSGZoneConfiguration() *WSGZoneConfiguration {
	m := new(WSGZoneConfiguration)
	return m
}

type WSGZoneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGZoneSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneList() *WSGZoneList {
	m := new(WSGZoneList)
	return m
}

type WSGZoneSummary struct {
	// Id
	// Identifier of the zone
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the zone
	Name *string `json:"name,omitempty"`

	// ServiceName
	// Name of the zone
	ServiceName *string `json:"serviceName,omitempty"`
}

func NewWSGZoneSummary() *WSGZoneSummary {
	m := new(WSGZoneSummary)
	return m
}
