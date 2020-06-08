package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

type WSGMeshnodeinfoHelperZoneInfo struct {
	HelperAPZoneId *string `json:"helperAPZoneId,omitempty"`

	HelperAPZoneName *string `json:"helperAPZoneName,omitempty"`
}

func NewWSGMeshnodeinfoHelperZoneInfo() *WSGMeshnodeinfoHelperZoneInfo {
	m := new(WSGMeshnodeinfoHelperZoneInfo)
	return m
}

type WSGMeshnodeinfo struct {
	// ApMac
	// The MAC of the mesh AP
	ApMac *string `json:"apMac,omitempty"`

	// ApModel
	// The model of the mesh AP
	ApModel *string `json:"apModel,omitempty"`

	// ApName
	// The name of the mesh AP
	ApName *string `json:"apName,omitempty"`

	// Channel
	// The channel of the mesh AP
	Channel *string `json:"channel,omitempty"`

	// ClientCount
	// The count of clients of the mesh AP
	ClientCount *int `json:"clientCount,omitempty"`

	// DownlinkRssi
	// The downlinkRssi of the mesh AP
	DownlinkRssi *int `json:"downlinkRssi,omitempty"`

	// ExternalIPAddress
	// The external IP of the mesh AP
	ExternalIPAddress *string `json:"externalIPAddress,omitempty"`

	// HasDownLink
	// The hasDownLink of the mesh AP
	HasDownLink *bool `json:"hasDownLink,omitempty"`

	HelperZoneInfo []*WSGMeshnodeinfoHelperZoneInfo `json:"helperZoneInfo,omitempty"`

	// Hops
	// The hop count of this mesh AP
	Hops *int `json:"hops,omitempty"`

	// IpAddress
	// The IP of the mesh AP
	IpAddress *int `json:"ipAddress,omitempty"`

	// MeshRole
	// The Role of the mesh AP
	MeshRole *string `json:"meshRole,omitempty"`

	// UplinkRssi
	// The uplinkRssi of the mesh AP
	UplinkRssi *int `json:"uplinkRssi,omitempty"`
}

func NewWSGMeshnodeinfo() *WSGMeshnodeinfo {
	m := new(WSGMeshnodeinfo)
	return m
}

type WSGMeshnodeinfoArray []*WSGMeshnodeinfo

func MakeWSGMeshnodeinfoArray() WSGMeshnodeinfoArray {
	m := make(WSGMeshnodeinfoArray, 0)
	return m
}

type WSGMeshnodeinfoList struct {
	// Extra
	// Any additional response data.
	Extra *WSGMeshnodeinfoListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first MeshNodeInfo returned out of the complete Rogue AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more MeshNodeInfo after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGMeshnodeinfo `json:"list,omitempty"`

	// RawDataTotalCount
	// MeshNodeInfos count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// MeshNodeInfos count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGMeshnodeinfoList() *WSGMeshnodeinfoList {
	m := new(WSGMeshnodeinfoList)
	return m
}

// WSGMeshnodeinfoListExtraType
//
// Any additional response data.
type WSGMeshnodeinfoListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGMeshnodeinfoListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGMeshnodeinfoListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGMeshnodeinfoListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGMeshnodeinfoListExtraType() *WSGMeshnodeinfoListExtraType {
	m := new(WSGMeshnodeinfoListExtraType)
	return m
}

type WSGMeshnodeinfoUpdateAPZeroTouch struct {
	ApMac *string `json:"apMac,omitempty"`

	HelperZoneId *string `json:"helperZoneId,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty"`

	// Status
	// Constraints:
	//    - oneof:[Accept,Reject]
	Status *string `json:"status,omitempty"`
}

func NewWSGMeshnodeinfoUpdateAPZeroTouch() *WSGMeshnodeinfoUpdateAPZeroTouch {
	m := new(WSGMeshnodeinfoUpdateAPZeroTouch)
	return m
}