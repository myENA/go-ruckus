package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGSCGUserGroupService struct {
	apiClient *APIClient
}

func NewWSGSCGUserGroupService(c *APIClient) *WSGSCGUserGroupService {
	s := new(WSGSCGUserGroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCGUserGroupService() *WSGSCGUserGroupService {
	return NewWSGSCGUserGroupService(ss.apiClient)
}

// AddUserGroups
//
// Add SCG user group.
//
// Request Body:
//	 - body *WSGSCGUserGroup
func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *WSGSCGUserGroup) (*WSGSCGUserGroupAuditId, error) {
	var (
		resp *WSGSCGUserGroupAuditId
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddUserGroups, true)
}

// DeleteUserGroups
//
// Delete multiple SCG user group.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserGroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteUserGroups, true)
}

// DeleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, userGroupId string) (*WSGSCGUserGroupAuditId, error) {
	var (
		resp *WSGSCGUserGroupAuditId
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, userGroupId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteUserGroupsByUserGroupId, true)
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria) (*WSGSCGUserGroupList, error) {
	var (
		resp *WSGSCGUserGroupList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindUserGroupsByQueryCriteria, true)
}

// FindUserGroupsByUserGroupId
//
// Get SCG user group.
//
// Required Parameters:
// - userGroupId string
//		- required
//
// Optional Parameters:
// - includeUsers string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, userGroupId string, optionalParams map[string]interface{}) (*WSGSCGUserGroup, error) {
	var (
		resp *WSGSCGUserGroup
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, userGroupId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsByUserGroupId, true)
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context) (*WSGSCGUserGroupPermissionList, error) {
	var (
		resp *WSGSCGUserGroupPermissionList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsCurrentUserPermissionCategories, true)
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context) (*WSGSCGUserGroupRoleLabelValueList, error) {
	var (
		resp *WSGSCGUserGroupRoleLabelValueList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsRoles, true)
}

// FindUserGroupsRolesPermissionsByRole
//
// Get permission items of role.
//
// Required Parameters:
// - role string
//		- required
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, role string, optionalParams map[string]interface{}) (*WSGSCGUserGroupPermissionList, error) {
	var (
		resp *WSGSCGUserGroupPermissionList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, role, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsRolesPermissionsByRole, true)
}

// PartialUpdateUserGroupsByUserGroupId
//
// Update user groups.
//
// Request Body:
//	 - body *WSGSCGUserPatchScgUserGroup
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, userGroupId string) (*WSGSCGUserGroupAuditId, error) {
	var (
		resp *WSGSCGUserGroupAuditId
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
	if err = pkgValidator.VarCtx(ctx, userGroupId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUserGroupsByUserGroupId, true)
}