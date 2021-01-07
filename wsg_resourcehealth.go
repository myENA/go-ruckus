package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGResourceHealthService struct {
	apiClient *VSZClient
}

func NewWSGResourceHealthService(c *VSZClient) *WSGResourceHealthService {
	s := new(WSGResourceHealthService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGResourceHealthService() *WSGResourceHealthService {
	return NewWSGResourceHealthService(ss.apiClient)
}

// FindResourceHealthSummaryByQueryCriteria
//
// Operation ID: findResourceHealthSummaryByQueryCriteria
//
// Get current connected count of requested resource
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[client,ap]
func (s *WSGResourceHealthService) FindResourceHealthSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, mutators ...RequestMutator) (*WSGCommonMonitoringSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonMonitoringSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindResourceHealthSummaryByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonMonitoringSummary()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
