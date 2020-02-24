package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGFlexiVPNService struct {
	apiClient *APIClient
}

func NewWSGFlexiVPNService(c *APIClient) *WSGFlexiVPNService {
	s := new(WSGFlexiVPNService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFlexiVPNService() *WSGFlexiVPNService {
	return NewWSGFlexiVPNService(ss.apiClient)
}

type WSGFlexiVPNSetting struct {
	// ZoneAffinityId
	// Zone Affinity ID
	// Constraints:
	//    - required
	ZoneAffinityId *string `json:"zoneAffinityId" validate:"required"`
}

func NewWSGFlexiVPNSetting() *WSGFlexiVPNSetting {
	m := new(WSGFlexiVPNSetting)
	return m
}

// DeleteRkszonesWlansFlexiVpnProfileById
//
// Use this API command to delete Flexi-VPN on WLAN
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGFlexiVPNService) DeleteRkszonesWlansFlexiVpnProfileById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansFlexiVpnProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindServicesFlexiVpnProfileByQueryCriteria
//
// Use this API command to query Flexi-VPN profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFlexiVPNService) FindServicesFlexiVpnProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileFlexiVpnProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileFlexiVpnProfileList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesFlexiVpnProfileByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileFlexiVpnProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
