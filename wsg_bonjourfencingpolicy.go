package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGBonjourfencingpolicyService struct {
	apiClient *VSZClient
}

func NewWSGBonjourfencingpolicyService(c *VSZClient) *WSGBonjourfencingpolicyService {
	s := new(WSGBonjourfencingpolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourfencingpolicyService() *WSGBonjourfencingpolicyService {
	return NewWSGBonjourfencingpolicyService(ss.apiClient)
}

// AddRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to create Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGProfileCreateBonjourFencingPolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourfencingpolicyService) AddRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, body *WSGProfileCreateBonjourFencingPolicy, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesBonjourFencingPolicyByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesBonjourFencingPolicy
//
// Use this API command to delete Bulk Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBonjourfencingpolicyService) DeleteRkszonesBonjourFencingPolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBonjourFencingPolicy, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesBonjourFencingPolicyById
//
// Use this API command to delete Bonjour Fencing Policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBonjourfencingpolicyService) DeleteRkszonesBonjourFencingPolicyById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBonjourFencingPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApsBonjourFencingStatisticByApMac
//
// Use this API command to get Bonjour Fencing Statistic per AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGBonjourfencingpolicyService) FindApsBonjourFencingStatisticByApMac(ctx context.Context, apMac string) (*WSGProfileBonjourFencingStatistic, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBonjourFencingStatistic
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsBonjourFencingStatisticByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBonjourFencingStatistic()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesBonjourFencingPolicyById
//
// Use this API command to retrieve Bonjour Fencing Policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourfencingpolicyService) FindRkszonesBonjourFencingPolicyById(ctx context.Context, id string, zoneId string) (*WSGProfileBonjourFencingPolicy, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBonjourFencingPolicy
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourFencingPolicyById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBonjourFencingPolicy()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourfencingpolicyService) FindRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, zoneId string) (*WSGProfileBonjourFencingPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBonjourFencingPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourFencingPolicyByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBonjourFencingPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesBonjourFencingPolicyByQueryCriteria
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBonjourfencingpolicyService) FindServicesBonjourFencingPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBonjourFencingPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBonjourFencingPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesBonjourFencingPolicyByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBonjourFencingPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesBonjourFencingPolicyById
//
// Use this API command to modify the configuration of Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGProfileModifyBonjourFencingPolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourfencingpolicyService) PartialUpdateRkszonesBonjourFencingPolicyById(ctx context.Context, body *WSGProfileModifyBonjourFencingPolicy, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesBonjourFencingPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}