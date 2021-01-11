package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGQueryWithFilterService struct {
	apiClient *VSZClient
}

func NewWSGQueryWithFilterService(c *VSZClient) *WSGQueryWithFilterService {
	s := new(WSGQueryWithFilterService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGQueryWithFilterService() *WSGQueryWithFilterService {
	return NewWSGQueryWithFilterService(ss.apiClient)
}

// FindGgsnGtpcConStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN Connection.
//
// Operation ID: findGgsnGtpcConStatsByQueryCriteria
// Operation path: /query/ggsnGtpcCon/stats
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindGgsnGtpcConStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsGgsnGtpcConListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGRACStatsGgsnGtpcConListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGRACStatsGgsnGtpcConListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindGgsnGtpcConStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGRACStatsGgsnGtpcConListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRACStatsGgsnGtpcConListAPIResponse), err
}

// FindGgsnGtpStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN/PGW GTP-C Sessions.
//
// Operation ID: findGgsnGtpStatsByQueryCriteria
// Operation path: /query/ggsnGtp/stats
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindGgsnGtpStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsGgsnGtpListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGRACStatsGgsnGtpListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGRACStatsGgsnGtpListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindGgsnGtpStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGRACStatsGgsnGtpListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRACStatsGgsnGtpListAPIResponse), err
}

// FindRadiusProxyStatsByQueryCriteria
//
// Use this API command to retrieve a list of Radius Proxy.
//
// Operation ID: findRadiusProxyStatsByQueryCriteria
// Operation path: /query/radiusProxy/stats
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindRadiusProxyStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsRadiusProxyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGRACStatsRadiusProxyListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGRACStatsRadiusProxyListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindRadiusProxyStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGRACStatsRadiusProxyListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRACStatsRadiusProxyListAPIResponse), err
}
