package vsz

// API Version: v8_1

// WSGDHCPPoolsDhcpClientInfo
//
// DHCP Pool Client Information List
type WSGDHCPPoolsDhcpClientInfo struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	LeaseExpiryTime *int `json:"leaseExpiryTime,omitempty"`

	LeaseTime *int `json:"leaseTime,omitempty"`

	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`
}

func NewWSGDHCPPoolsDhcpClientInfo() *WSGDHCPPoolsDhcpClientInfo {
	m := new(WSGDHCPPoolsDhcpClientInfo)
	return m
}

// WSGDHCPPoolsDhcpPoolInfo
//
// DHCP Pool Information List
type WSGDHCPPoolsDhcpPoolInfo struct {
	ApIp *string `json:"apIp,omitempty"`

	AvailableIpCount *int `json:"availableIpCount,omitempty"`

	ClientInfoList *WSGDHCPPoolsDhcpPoolInfoClientInfoListType `json:"clientInfoList,omitempty"`

	Name *string `json:"name,omitempty"`

	PoolEndIp *string `json:"poolEndIp,omitempty"`

	PoolIndex *int `json:"poolIndex,omitempty"`

	PoolStartIp *string `json:"poolStartIp,omitempty"`

	SubnetMask *string `json:"subnetMask,omitempty"`

	TotalIpCount *int `json:"totalIpCount,omitempty"`

	UsedIpCount *int `json:"usedIpCount,omitempty"`

	VlanId *int `json:"vlanId,omitempty"`
}

func NewWSGDHCPPoolsDhcpPoolInfo() *WSGDHCPPoolsDhcpPoolInfo {
	m := new(WSGDHCPPoolsDhcpPoolInfo)
	return m
}

type WSGDHCPPoolsDhcpPoolInfoClientInfoListType struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDHCPPoolsDhcpClientInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDHCPPoolsDhcpPoolInfoClientInfoListType() *WSGDHCPPoolsDhcpPoolInfoClientInfoListType {
	m := new(WSGDHCPPoolsDhcpPoolInfoClientInfoListType)
	return m
}

// WSGDHCPPools
//
// DHCP Pools Usage Per AP
type WSGDHCPPools struct {
	ApMac *string `json:"apMac,omitempty"`

	// DomainEntity
	// Constraints:
	//    - nullable
	DomainEntity *string `json:"domainEntity,omitempty" validate:"omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty" validate:"omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty" validate:"omitempty"`

	PoolInfoList []*WSGDHCPPoolsDhcpPoolInfo `json:"poolInfoList,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty" validate:"omitempty"`
}

func NewWSGDHCPPools() *WSGDHCPPools {
	m := new(WSGDHCPPools)
	return m
}