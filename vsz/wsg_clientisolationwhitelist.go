package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGClientIsolationWhitelistService struct {
	apiClient *APIClient
}

func NewWSGClientIsolationWhitelistService(c *APIClient) *WSGClientIsolationWhitelistService {
	s := new(WSGClientIsolationWhitelistService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService() *WSGClientIsolationWhitelistService {
	return NewWSGClientIsolationWhitelistService(ss.apiClient)
}

// AddRkszonesClientIsolationWhitelistByZoneId
//
// Create a new ClientIsolationWhitelist.
//
// Request Body:
//	 - body *WSGProfileCreateClientIsolationWhitelist
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) AddRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, body *WSGProfileCreateClientIsolationWhitelist, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesClientIsolationWhitelistByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteRkszonesClientIsolationWhitelist
//
// Use this API command to delete Bulk Client Isolation Whitelist.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelist, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteRkszonesClientIsolationWhitelistById
//
// Delete a Client Isolation Whitelist.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById(ctx context.Context, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelistById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesClientIsolationWhitelistById
//
// Retrieve an Client Isolation Whitelist.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById(ctx context.Context, id string, zoneId string) (*WSGProfileClientIsolationWhitelist, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileClientIsolationWhitelist
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClientIsolationWhitelist()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesClientIsolationWhitelistByZoneId
//
// Retrieve a list of Client Isolation Whitelist.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, zoneId string) (*WSGProfileClientIsolationWhitelistArray, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileClientIsolationWhitelistArray
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClientIsolationWhitelistArray()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindServicesClientIsolationWhitelistByQueryCriteria
//
// Retrieve a list of Client Isolation Whitelist.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileClientIsolationWhitelistArray, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileClientIsolationWhitelistArray
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesClientIsolationWhitelistByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClientIsolationWhitelistArray()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesClientIsolationWhitelistById
//
// Modify a specific Client Isolation Whitelist basic.
//
// Request Body:
//	 - body *WSGProfileModifyClientIsolationWhitelist
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) PartialUpdateRkszonesClientIsolationWhitelistById(ctx context.Context, body *WSGProfileModifyClientIsolationWhitelist, id string, zoneId string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesClientIsolationWhitelistById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
