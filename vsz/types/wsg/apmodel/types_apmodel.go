package apmodel

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApModel struct {
	CellularSettings *CellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *ExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *ExternalAntenna `json:"externalAntenna50,omitempty"`

	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	Lacp *LacpSetting `json:"lacp,omitempty"`

	LanPorts []*LanPortSetting `json:"lanPorts,omitempty"`

	LedMode *string `json:"ledMode,omitempty"`

	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *LldpSetting `json:"lldp,omitempty"`

	PoeModeSetting *string `json:"poeModeSetting,omitempty"`

	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	PoeTxChain *int `json:"poeTxChain,omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	RadioBand *string `json:"radioBand,omitempty"`

	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

type AuthenticatorAAAServer struct {
	EnableUseSCGasProxy *bool `json:"enableUseSCGasProxy,omitempty" validate:"required"`

	Server *common.GenericRef `json:"server,omitempty"`
}

type CellularSettings struct {
	DataRoaming *int `json:"dataRoaming,omitempty"`

	DataRoaming2 *int `json:"dataRoaming2,omitempty"`

	MobileAPName *string `json:"mobileAPName,omitempty"`

	MobileAPName2 *string `json:"mobileAPName2,omitempty"`

	SelectGG *int `json:"select3g4g,omitempty" validate:"required"`

	SelectGG2 *int `json:"select3g4g2,omitempty" validate:"required"`

	SimCardUsage *int `json:"simCardUsage,omitempty"`

	WanConnection *int `json:"wanConnection,omitempty" validate:"required"`

	WanRecoveryTimer *int `json:"wanRecoveryTimer,omitempty" validate:"required"`
}

type CommonAttribute struct {
	AllowDfsCountry *string `json:"allowDfsCountry,omitempty"`

	CapabilityScore *float64 `json:"capabilityScore,omitempty"`

	CpuFrequency *int `json:"cpuFrequency,omitempty"`

	HasCablemodem *bool `json:"hasCablemodem,omitempty"`

	HasGps *bool `json:"hasGps,omitempty"`

	HasScanRadio *bool `json:"hasScanRadio,omitempty"`

	IsAllowDisableExtAnt *bool `json:"isAllowDisableExtAnt,omitempty"`

	IsDualRadio *bool `json:"isDualRadio,omitempty"`

	IsOutdoor *bool `json:"isOutdoor,omitempty"`

	MaxChannelization5G *int `json:"maxChannelization5G,omitempty"`

	MaxChannelization24G *int `json:"maxChannelization24G,omitempty"`

	MaxClientsUpper *int `json:"maxClientsUpper,omitempty"`

	MaxWlanNum5G *int `json:"maxWlanNum5G,omitempty"`

	MaxWlanNum24G *int `json:"maxWlanNum24G,omitempty"`

	MeshRadioCaps *string `json:"meshRadioCaps,omitempty"`

	NoAvc *bool `json:"noAvc,omitempty"`

	NoMesh *bool `json:"noMesh,omitempty"`

	NonEditablePorts []int `json:"nonEditablePorts,omitempty"`

	NonVisiblePorts []int `json:"nonVisiblePorts,omitempty"`

	NumOfCores *int `json:"numOfCores,omitempty"`

	PoeModeCaps *string `json:"poeModeCaps,omitempty"`

	Ram *int `json:"ram,omitempty"`

	Reserved5GWlanForMesh *int `json:"reserved5GWlanForMesh,omitempty"`

	ScalingFactor *int `json:"scalingFactor,omitempty"`

	Support11AC *bool `json:"support11AC,omitempty"`

	SupportAPUsbSoftwarePackage *bool `json:"supportAPUsbSoftwarePackage,omitempty"`

	SupportBandSwitch *bool `json:"supportBandSwitch,omitempty"`

	SupportBonjour *bool `json:"supportBonjour,omitempty"`

	SupportChannelization160 *bool `json:"supportChannelization160,omitempty"`

	SupportIpsec *bool `json:"supportIpsec,omitempty"`

	SupportLBS *bool `json:"supportLBS,omitempty"`

	SupportResetCablemodem *bool `json:"supportResetCablemodem,omitempty"`
}

type ExternalAntenna struct {
	ChainMask *string `json:"chainMask,omitempty"`

	Dbi *int `json:"dbi,omitempty"`

	Enabled *bool `json:"enabled,omitempty" validate:"required"`
}

type LacpSetting struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type LanPort8021X struct {
	Authenticator *LanPortAuthenticator `json:"authenticator,omitempty"`

	Supplicant *LanPortSupplicant `json:"supplicant,omitempty"`

	Type *string `json:"type,omitempty" validate:"required"`
}

type LanPortAuthenticator struct {
	Accounting *AuthenticatorAAAServer `json:"accounting,omitempty"`

	Authentication *AuthenticatorAAAServer `json:"authentication,omitempty"`

	DisabledAccounting *bool `json:"disabledAccounting,omitempty"`

	MacAuthByPassEnabled *bool `json:"macAuthByPassEnabled,omitempty" validate:"required"`
}

type LanPortSetting struct {
	Enabled *bool `json:"enabled,omitempty" validate:"required"`

	EthPortProfile *common.GenericRef `json:"ethPortProfile,omitempty"`

	Members *string `json:"members,omitempty"`

	OverwriteVlanEnabled *bool `json:"overwriteVlanEnabled,omitempty"`

	PortName *string `json:"portName,omitempty" validate:"required"`

	VlanUntagId *int `json:"vlanUntagId,omitempty"`
}

type LanPortSupplicant struct {
	Password *string `json:"password,omitempty"`

	Type *string `json:"type,omitempty" validate:"required"`

	UserName *string `json:"userName,omitempty"`
}

type LldpSetting struct {
	AdvertiseIntervalInSec *int `json:"advertiseIntervalInSec,omitempty"`

	Enabled *bool `json:"enabled,omitempty" validate:"required"`

	HoldTimeInSec *int `json:"holdTimeInSec,omitempty"`

	ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
}
