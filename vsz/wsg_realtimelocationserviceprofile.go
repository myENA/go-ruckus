package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGRealTimeLocationServiceProfileService struct {
	apiClient *APIClient
}

func NewWSGRealTimeLocationServiceProfileService(c *APIClient) *WSGRealTimeLocationServiceProfileService {
	s := new(WSGRealTimeLocationServiceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService() *WSGRealTimeLocationServiceProfileService {
	return NewWSGRealTimeLocationServiceProfileService(ss.apiClient)
}

// AddRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to create RTLS Profile.
//
// Request Body:
//	 - body *WSGProfileCreateRtlsProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) AddRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, body *WSGProfileCreateRtlsProfile, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesRealTimeLocationServiceByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesRealTimeLocationServiceById
//
// Use this API command to Delete RTLS Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) DeleteRkszonesRealTimeLocationServiceById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRealTimeLocationServiceById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindRkszonesRealTimeLocationServiceById
//
// Use this API command to Get RTLS Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceById(ctx context.Context, id string, zoneId string) (*WSGProfileCreateRtlsProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileCreateRtlsProfile
		httpResp *http.Response
		err      error
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRealTimeLocationServiceById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileCreateRtlsProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to Get RTLS Profile by zone ID.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, zoneId string) (*WSGProfileRtlsProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileRtlsProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRealTimeLocationServiceByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRtlsProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateRkszonesRealTimeLocationServiceById
//
// Use this API command to Modify RTLS Profile by profile's ID.
//
// Request Body:
//	 - body *WSGProfileUpdateRtlsProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) UpdateRkszonesRealTimeLocationServiceById(ctx context.Context, body *WSGProfileUpdateRtlsProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesRealTimeLocationServiceById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
