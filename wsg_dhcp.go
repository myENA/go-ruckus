package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDHCPService struct {
	apiClient *VSZClient
}

func NewWSGDHCPService(c *VSZClient) *WSGDHCPService {
	s := new(WSGDHCPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDHCPService() *WSGDHCPService {
	return NewWSGDHCPService(ss.apiClient)
}

// AddRkszonesDhcpSiteDhcpProfileByZoneId
//
// Operation ID: addRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to create DHCP Pool.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGProfileCreateDhcpProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Operation ID: addRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *WSGCommonDoAssignIp
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRefAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// DeleteRkszonesDhcpSiteDhcpProfileById
//
// Operation ID: deleteRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to delete DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}

// DeleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Operation ID: deleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to delete multiple DHCP Pools.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Operation ID: findDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPMessageStatsDhcpMsgStatsAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPMessageStatsDhcpMsgStats
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDhcpDataDhcpMsgStatsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDHCPMessageStatsDhcpMsgStats()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindDhcpDataDhcpPoolsByApMac
//
// Operation ID: findDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPPoolsAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPPools
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDHCPPools()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindDhcpDataDhcpPoolsByPoolIndex
//
// Operation ID: findDhcpDataDhcpPoolsByPoolIndex
//
// Use this API command to get AP DHCP Pool Usage by pool's index.
//
// Required Parameters:
// - apMac string
//		- required
// - poolIndex string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, apMac string, poolIndex string, mutators ...RequestMutator) (*WSGDHCPPoolsDhcpPoolInfoAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPPoolsDhcpPoolInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByPoolIndex, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("apMac", apMac)
	req.PathParams.Set("poolIndex", poolIndex)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDHCPPoolsDhcpPoolInfo()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpProfileById
//
// Operation ID: findRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to get DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpProfileRefAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpProfileRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonDhcpProfileRef()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Operation ID: findRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileDhcpProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileDhcpProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Operation ID: findRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRefAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpSiteConfigByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Operation ID: findRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGDHCPService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria, mutators ...RequestMutator) (*WSGZoneDhcpSiteConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGZoneDhcpSiteConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRkszonesServicesDhcpSiteConfigByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGZoneDhcpSiteConfigList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindServicesDhcpProfileByQueryCriteria
//
// Operation ID: findServicesDhcpProfileByQueryCriteria
//
// Query DHCP Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDHCPService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileDhcpProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesDhcpProfileByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileDhcpProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// PartialUpdateRkszonesDhcpSiteDhcpProfileById
//
// Operation ID: partialUpdateRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to modify DHCP Pool by pool's ID.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) PartialUpdateRkszonesDhcpSiteDhcpProfileById(ctx context.Context, body *WSGProfileCreateDhcpProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}
