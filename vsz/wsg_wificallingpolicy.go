package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGWiFiCallingPolicyService struct {
	apiClient *APIClient
}

func NewWSGWiFiCallingPolicyService(c *APIClient) *WSGWiFiCallingPolicyService {
	s := new(WSGWiFiCallingPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWiFiCallingPolicyService() *WSGWiFiCallingPolicyService {
	return NewWSGWiFiCallingPolicyService(ss.apiClient)
}

// AddWifiCallingWifiCallingPolicy
//
// Use this API command to Create Wi-Fi Calling Policy.
//
// Request Body:
//	 - body *WSGWIFICallingCreateWifiCallingPolicy
func (s *WSGWiFiCallingPolicyService) AddWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingCreateWifiCallingPolicy) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddWifiCallingWifiCallingPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteWifiCallingWifiCallingPolicy
//
// Use this API command to Delete bulk Wi-Fi Calling policies.
//
// Request Body:
//	 - body *WSGWIFICallingDeleteBulk
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingDeleteBulk) (interface{}, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteWifiCallingWifiCallingPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteWifiCallingWifiCallingPolicyById
//
// Use this API command to Delete a Wi-Fi Calling policy by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicyById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteWifiCallingWifiCallingPolicyById, true)
	req.SetPathParameter("id", id)
}

// FindWifiCallingByQueryCriteria
//
// Use this API command to Query Wi-Fi Calling Policy List.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWiFiCallingPolicyService) FindWifiCallingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWIFICallingPolicyList, error) {
	var (
		req  *APIRequest
		resp *WSGWIFICallingPolicyList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindWifiCallingByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindWifiCallingWifiCallingPolicy
//
// Use this API command to Retrieve List of Wi-Fi Calling Policy.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicy(ctx context.Context, optionalParams map[string]interface{}) (*WSGWIFICallingPolicyList, error) {
	var (
		req  *APIRequest
		resp *WSGWIFICallingPolicyList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindWifiCallingWifiCallingPolicy, true)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
}

// FindWifiCallingWifiCallingPolicyById
//
// Use this API command to Retrieve Wi-Fi Calling Policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicyById(ctx context.Context, id string) (*WSGWIFICallingPolicy, error) {
	var (
		req  *APIRequest
		resp *WSGWIFICallingPolicy
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindWifiCallingWifiCallingPolicyById, true)
	req.SetPathParameter("id", id)
}

// PartialUpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify a Wi-Fi Calling policy.
//
// Request Body:
//	 - body *WSGWIFICallingModifyWifiCallingPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) PartialUpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateWifiCallingWifiCallingPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}

// UpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify Entire Wi-Fi Calling policy.
//
// Request Body:
//	 - body *WSGWIFICallingModifyEntireWifiCallingPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) UpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyEntireWifiCallingPolicy, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateWifiCallingWifiCallingPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}
