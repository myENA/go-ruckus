package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGSCGUserGroupService struct {
	apiClient *VSZClient
}

func NewWSGSCGUserGroupService(c *VSZClient) *WSGSCGUserGroupService {
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
// Operation ID: addUserGroups
// Operation path: /userGroups
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGSCGUserGroup
func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *WSGSCGUserGroup, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddUserGroups, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}

// DeleteUserGroups
//
// Delete multiple SCG user group.
//
// Operation ID: deleteUserGroups
// Operation path: /userGroups
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserGroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUserGroups, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Operation ID: deleteUserGroupsByUserGroupId
// Operation path: /userGroups/{userGroupId}
// Success code: 204 (No Content)
//
// Required parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, userGroupId string, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Operation ID: findUserGroupsByQueryCriteria
// Operation path: /userGroups/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria, mutators ...RequestMutator) (*WSGSCGUserGroupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindUserGroupsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSCGUserGroupListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupListAPIResponse), err
}

// FindUserGroupsByUserGroupId
//
// Get SCG user group.
//
// Operation ID: findUserGroupsByUserGroupId
// Operation path: /userGroups/{userGroupId}
// Success code: 200 (OK)
//
// Required parameters:
// - userGroupId string
//		- required
//
// Optional parameters:
// - includeUsers string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, userGroupId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSCGUserGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("userGroupId", userGroupId)
	if v, ok := optionalParams["includeUsers"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("includeUsers", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAPIResponse), err
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
//
// Operation ID: findUserGroupsCurrentUserPermissionCategories
// Operation path: /userGroups/currentUser/permissionCategories
// Success code: 200 (OK)
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context, mutators ...RequestMutator) (*WSGSCGUserGroupPermissionListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupPermissionListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsCurrentUserPermissionCategories, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
//
// Operation ID: findUserGroupsRoles
// Operation path: /userGroups/roles
// Success code: 200 (OK)
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context, mutators ...RequestMutator) (*WSGSCGUserGroupRoleLabelValueListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupRoleLabelValueListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupRoleLabelValueListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsRoles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupRoleLabelValueListAPIResponse), err
}

// FindUserGroupsRolesPermissionsByRole
//
// Get permission items of role.
//
// Operation ID: findUserGroupsRolesPermissionsByRole
// Operation path: /userGroups/roles/{role}/permissions
// Success code: 200 (OK)
//
// Required parameters:
// - role string
//		- required
//
// Optional parameters:
// - domainId string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, role string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSCGUserGroupPermissionListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupPermissionListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsRolesPermissionsByRole, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("role", role)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
}

// PartialUpdateUserGroupsByUserGroupId
//
// Update user groups.
//
// Operation ID: partialUpdateUserGroupsByUserGroupId
// Operation path: /userGroups/{userGroupId}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSCGUserPatchScgUserGroup
//
// Required parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, userGroupId string, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req.PathParams.Set("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}
