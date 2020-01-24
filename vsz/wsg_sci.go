package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGSCIService struct {
	apiClient *APIClient
}

func NewWSGSCIService(c *APIClient) *WSGSCIService {
	s := new(WSGSCIService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	return NewWSGSCIService(ss.apiClient)
}

type WSGSCICreateSciProfile struct {
	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost" validate:"required"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort" validate:"required"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId" validate:"required"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser" validate:"required"`
}

func NewWSGSCICreateSciProfile() *WSGSCICreateSciProfile {
	m := new(WSGSCICreateSciProfile)
	return m
}

type WSGSCIDeleteSciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id" validate:"required"`
}

func NewWSGSCIDeleteSciProfile() *WSGSCIDeleteSciProfile {
	m := new(WSGSCIDeleteSciProfile)
	return m
}

type WSGSCIDeleteSciProfileList struct {
	List []*WSGSCIDeleteSciProfile `json:"list,omitempty"`
}

func NewWSGSCIDeleteSciProfileList() *WSGSCIDeleteSciProfileList {
	m := new(WSGSCIDeleteSciProfileList)
	return m
}

type WSGSCIModifyEventCode struct {
	// SciAcceptedEventCodes
	// Constraints:
	//    - required
	SciAcceptedEventCodes []int `json:"sciAcceptedEventCodes" validate:"required,dive"`
}

func NewWSGSCIModifyEventCode() *WSGSCIModifyEventCode {
	m := new(WSGSCIModifyEventCode)
	return m
}

type WSGSCIModifySciEnabled struct {
	// SciEnabled
	// Is SZ/SCI interface enabled or disabled
	// Constraints:
	//    - required
	SciEnabled *bool `json:"sciEnabled" validate:"required"`
}

func NewWSGSCIModifySciEnabled() *WSGSCIModifySciEnabled {
	m := new(WSGSCIModifySciEnabled)
	return m
}

type WSGSCIModifySciPriorityList struct {
	List []*WSGSCIModifySciPriorityListType `json:"list,omitempty"`
}

func NewWSGSCIModifySciPriorityList() *WSGSCIModifySciPriorityList {
	m := new(WSGSCIModifySciPriorityList)
	return m
}

type WSGSCIModifySciPriorityListType struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id" validate:"required"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPriority *int `json:"sciPriority" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`
}

func NewWSGSCIModifySciPriorityListType() *WSGSCIModifySciPriorityListType {
	m := new(WSGSCIModifySciPriorityListType)
	return m
}

type WSGSCIModifySciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword" validate:"required"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile" validate:"required"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost" validate:"required"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort" validate:"required"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId" validate:"required"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser" validate:"required"`
}

func NewWSGSCIModifySciProfile() *WSGSCIModifySciProfile {
	m := new(WSGSCIModifySciProfile)
	return m
}

type WSGSCIEventCode struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SCI accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCIEventCodeListType `json:"list,omitempty"`

	// TotalCount
	// Total SCI accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCIEventCode() *WSGSCIEventCode {
	m := new(WSGSCIEventCode)
	return m
}

type WSGSCIEventCodeListType struct {
	// Code
	// SCI accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// SCI accepted event type
	Type *string `json:"type,omitempty"`
}

func NewWSGSCIEventCodeListType() *WSGSCIEventCodeListType {
	m := new(WSGSCIEventCodeListType)
	return m
}

type WSGSCIProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	SciPassword *string `json:"sciPassword,omitempty"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	SciPriority *int `json:"sciPriority,omitempty"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	SciProfile *string `json:"sciProfile,omitempty"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	SciServerHost *string `json:"sciServerHost,omitempty"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	SciServerPort *string `json:"sciServerPort,omitempty"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	SciSystemId *string `json:"sciSystemId,omitempty"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	SciUser *string `json:"sciUser,omitempty"`
}

func NewWSGSCIProfile() *WSGSCIProfile {
	m := new(WSGSCIProfile)
	return m
}

type WSGSCIProfileList struct {
	Extra *WSGSCIProfileListExtraType `json:"extra,omitempty"`

	List []*WSGSCIProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGSCIProfileList) UnmarshalJSON(b []byte) error {
	tmpt := new(WSGSCIProfileList)
	if err := json.Unmarshal(b, tmpt); err != nil {
		return err
	}
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	delete(tmp, "extra")
	delete(tmp, "list")
	tmpt.XAdditionalProperties = tmp
	*t = *tmpt
	return nil
}

func (t *WSGSCIProfileList) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	return json.Marshal(tmp)
}

func NewWSGSCIProfileList() *WSGSCIProfileList {
	m := new(WSGSCIProfileList)
	return m
}

type WSGSCIProfileListExtraType struct {
	// SciEnabled
	// SCI password of the SCI profile for SZ/SCI interface
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}

func NewWSGSCIProfileListExtraType() *WSGSCIProfileListExtraType {
	m := new(WSGSCIProfileListExtraType)
	return m
}

// AddSciSciEventCode
//
// Use this API command to modify SciAcceptedEventCodes.
//
// Request Body:
//	 - body *WSGSCIModifyEventCode
func (s *WSGSCIService) AddSciSciEventCode(ctx context.Context, body *WSGSCIModifyEventCode) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddSciSciEventCode, true)
}

// AddSciSciProfile
//
// Use this API command to create sciProfile.
//
// Request Body:
//	 - body *WSGSCICreateSciProfile
func (s *WSGSCIService) AddSciSciProfile(ctx context.Context, body *WSGSCICreateSciProfile) (*WSGCommonCreateResult, error) {
	var (
		resp *WSGCommonCreateResult
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddSciSciProfile, true)
}

// AddSciSciProfileSciPriority
//
// Use this API command to modify sciPriorities.
//
// Request Body:
//	 - body *WSGSCIModifySciPriorityList
func (s *WSGSCIService) AddSciSciProfileSciPriority(ctx context.Context, body *WSGSCIModifySciPriorityList) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddSciSciProfileSciPriority, true)
}

// DeleteSciSciProfile
//
// Use this API command to delete sciProfile list.
//
// Request Body:
//	 - body *WSGSCIDeleteSciProfileList
func (s *WSGSCIService) DeleteSciSciProfile(ctx context.Context, body *WSGSCIDeleteSciProfileList) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteSciSciProfile, true)
}

// DeleteSciSciProfileById
//
// Use this API command to delete sciProfile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) DeleteSciSciProfileById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteSciSciProfileById, true)
}

// FindSciSciEventCode
//
// Use this API command to retrieve SciAcceptedEventCodes.
func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context) (*WSGSCIEventCode, error) {
	var (
		resp *WSGSCIEventCode
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindSciSciEventCode, true)
}

// FindSciSciProfile
//
// Use this API command to retrieve sciProfile list.
func (s *WSGSCIService) FindSciSciProfile(ctx context.Context) (*WSGSCIProfileList, error) {
	var (
		resp *WSGSCIProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindSciSciProfile, true)
}

// FindSciSciProfileById
//
// Use this API command to retrieve sciProfile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, id string) (*WSGSCIProfile, error) {
	var (
		resp *WSGSCIProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindSciSciProfileById, true)
}

// PartialUpdateSciSciEnabled
//
// Use this API command to modify SCI settings is enabled or not.
//
// Request Body:
//	 - body *WSGSCIModifySciEnabled
func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context, body *WSGSCIModifySciEnabled) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSciSciEnabled, true)
}

// PartialUpdateSciSciProfileById
//
// Use this API command to modify sciProfile.
//
// Request Body:
//	 - body *WSGSCIModifySciProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) PartialUpdateSciSciProfileById(ctx context.Context, body *WSGSCIModifySciProfile, id string) (interface{}, error) {
	var (
		resp interface{}
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSciSciProfileById, true)
}