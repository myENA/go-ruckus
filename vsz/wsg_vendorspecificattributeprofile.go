package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGVendorSpecificAttributeProfileService struct {
	apiClient *APIClient
}

func NewWSGVendorSpecificAttributeProfileService(c *APIClient) *WSGVendorSpecificAttributeProfileService {
	s := new(WSGVendorSpecificAttributeProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVendorSpecificAttributeProfileService() *WSGVendorSpecificAttributeProfileService {
	return NewWSGVendorSpecificAttributeProfileService(ss.apiClient)
}

type WSGVendorSpecificAttributeProfileCreateResult interface{}

func NewWSGVendorSpecificAttributeProfileCreateResult() WSGVendorSpecificAttributeProfileCreateResult {
	m := new(WSGVendorSpecificAttributeProfileCreateResult)
	return m
}

type WSGVendorSpecificAttributeProfileDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileDeleteBulk() *WSGVendorSpecificAttributeProfileDeleteBulk {
	m := new(WSGVendorSpecificAttributeProfileDeleteBulk)
	return m
}

type WSGVendorSpecificAttributeProfileEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGVendorSpecificAttributeProfileEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGVendorSpecificAttributeProfileEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGVendorSpecificAttributeProfileEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGVendorSpecificAttributeProfileEmptyResult() *WSGVendorSpecificAttributeProfileEmptyResult {
	m := new(WSGVendorSpecificAttributeProfileEmptyResult)
	return m
}

type WSGVendorSpecificAttributeProfileGet struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileGet() *WSGVendorSpecificAttributeProfileGet {
	m := new(WSGVendorSpecificAttributeProfileGet)
	return m
}

type WSGVendorSpecificAttributeProfileList struct {
	// FirstIndex
	// Index of the first profile returned out of the profile list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more profiles after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of the vendor specific attribute profile
	List []*WSGVendorSpecificAttributeProfileListType `json:"list,omitempty"`

	// TotalCount
	// Total number of the profiles
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileList() *WSGVendorSpecificAttributeProfileList {
	m := new(WSGVendorSpecificAttributeProfileList)
	return m
}

type WSGVendorSpecificAttributeProfileListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileListType() *WSGVendorSpecificAttributeProfileListType {
	m := new(WSGVendorSpecificAttributeProfileListType)
	return m
}

type WSGVendorSpecificAttributeProfilePersist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	// Constraints:
	//    - required
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes" validate:"required,dive"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGVendorSpecificAttributeProfilePersist() *WSGVendorSpecificAttributeProfilePersist {
	m := new(WSGVendorSpecificAttributeProfilePersist)
	return m
}

type WSGVendorSpecificAttributeProfileQueryCriteriaResult struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVendorSpecificAttributeProfileGet `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileQueryCriteriaResult() *WSGVendorSpecificAttributeProfileQueryCriteriaResult {
	m := new(WSGVendorSpecificAttributeProfileQueryCriteriaResult)
	return m
}

type WSGVendorSpecificAttributeProfileVendorSpecificAttribute struct {
	// KeyId
	// Key ID of vendor specific attribute
	// Constraints:
	//    - required
	KeyId *int `json:"keyId" validate:"required"`

	// SupportedRadiusProtocol
	// The radius protocol to which this given vendor specific attribute will attach
	// Constraints:
	//    - required
	//    - oneof:[ACCOUNTING,AUTHENTICATION,BOTH_AUTHENTICATION_AND_ACCOUNTING]
	SupportedRadiusProtocol *string `json:"supportedRadiusProtocol" validate:"required,oneof=ACCOUNTING AUTHENTICATION BOTH_AUTHENTICATION_AND_ACCOUNTING"`

	// Type
	// Type of vendor specific attribute
	// Constraints:
	//    - required
	//    - oneof:[INTEGER,STRING]
	Type *string `json:"type" validate:"required,oneof=INTEGER STRING"`

	// Value
	// Value of vendor specific attribute
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`

	// VendorId
	// Vendor ID of vendor specific attribute
	// Constraints:
	//    - required
	VendorId *int `json:"vendorId" validate:"required"`
}

func NewWSGVendorSpecificAttributeProfileVendorSpecificAttribute() *WSGVendorSpecificAttributeProfileVendorSpecificAttribute {
	m := new(WSGVendorSpecificAttributeProfileVendorSpecificAttribute)
	return m
}

// AddRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Create a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) AddRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, zoneId string) (WSGVendorSpecificAttributeProfileCreateResult, error) {
	var (
		resp WSGVendorSpecificAttributeProfileCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesVendorSpecificAttributeProfilesByZoneId, true)
}

// DeleteRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to delete a vendor specific attribute profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string) (*WSGVendorSpecificAttributeProfileEmptyResult, error) {
	var (
		resp *WSGVendorSpecificAttributeProfileEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesById, true)
}

// DeleteRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Use this API command to delete a list of vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfileDeleteBulk
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfileDeleteBulk, zoneId string) (*WSGVendorSpecificAttributeProfileEmptyResult, error) {
	var (
		resp *WSGVendorSpecificAttributeProfileEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesByZoneId, true)
}

// FindRkszonesVendorSpecificAttributeProfilesById
//
// Get a vendor specific attribute profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string) (*WSGVendorSpecificAttributeProfileGet, error) {
	var (
		resp *WSGVendorSpecificAttributeProfileGet
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesById, true)
}

// FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria
//
// Use this API command to retrieve a list of vendor specific attribute profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGVendorSpecificAttributeProfileQueryCriteriaResult, error) {
	var (
		resp *WSGVendorSpecificAttributeProfileQueryCriteriaResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByQueryCriteria, true)
}

// FindRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Get a ID list of vendor specific attribute profile in this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, zoneId string) (*WSGVendorSpecificAttributeProfileList, error) {
	var (
		resp *WSGVendorSpecificAttributeProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByZoneId, true)
}

// UpdateRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to modify entire information of a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) UpdateRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, id string, zoneId string) (*WSGVendorSpecificAttributeProfileEmptyResult, error) {
	var (
		resp *WSGVendorSpecificAttributeProfileEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesVendorSpecificAttributeProfilesById, true)
}