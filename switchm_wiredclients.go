package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMWiredClientsService struct {
	apiClient *VSZClient
}

func NewSwitchMWiredClientsService(c *VSZClient) *SwitchMWiredClientsService {
	s := new(SwitchMWiredClientsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMWiredClientsService() *SwitchMWiredClientsService {
	return NewSwitchMWiredClientsService(ss.apiClient)
}

// AddSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClients(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchConnectedDevicesQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchConnectedDevicesQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchClients, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchConnectedDevicesQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchConnectedAPsQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchConnectedAPsQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchClientsAp, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchConnectedAPsQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchClientsAPExport
//
// Download CSV of AP's discovered via LLDP
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAPExport(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*FileResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *FileResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchClientsAPExport, true)
	if b, err := json.Marshal(body); err != nil {
		return resp, rm, err
	} else if err = req.SetBody(map[string]string{"json": string(b)}); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, "application/x-www-form-urlencoded")
	req.SetHeader(headerKeyAccept, "application/octet-stream")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(FileResponse)
	rm, err = handleFileResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
