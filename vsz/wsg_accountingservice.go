package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGAccountingServiceService struct {
	apiClient *APIClient
}

func NewWSGAccountingServiceService(c *APIClient) *WSGAccountingServiceService {
	s := new(WSGAccountingServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountingServiceService() *WSGAccountingServiceService {
	return NewWSGAccountingServiceService(ss.apiClient)
}

// AddServicesAcctRadius
//
// Use this API command to create a new RADIUS accounting service.
//
// Request Body:
//	 - body *WSGServiceCreateRadiusAccounting
func (s *WSGAccountingServiceService) AddServicesAcctRadius(ctx context.Context, body *WSGServiceCreateRadiusAccounting) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAcctRadius, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddServicesAcctTestById
//
// Use this API command to test an accounting service.
//
// Request Body:
//	 - body *WSGServiceTestingConfig
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) AddServicesAcctTestById(ctx context.Context, body *WSGServiceTestingConfig, id string) (interface{}, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAcctTestById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}

// DeleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Request Body:
//	 - body *WSGServiceDeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *WSGServiceDeleteBulkAccountingService) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcct, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteServicesAcctById
//
// Use this API command to delete an accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctById, true)
	req.SetPathParameter("id", id)
}

// DeleteServicesAcctRadiusById
//
// Use this API command to delete a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusById, true)
	req.SetPathParameter("id", id)
}

// DeleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusSecondaryById, true)
	req.SetPathParameter("id", id)
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusStandbyPrimaryById, true)
	req.SetPathParameter("id", id)
}

// FindServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceCommonAccountingServiceList, error) {
	var (
		req  *APIRequest
		resp *WSGServiceCommonAccountingServiceList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAcctByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context) (*WSGServiceRadiusAccountingServiceList, error) {
	var (
		req  *APIRequest
		resp *WSGServiceRadiusAccountingServiceList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAcctRadius, true)
}

// FindServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) FindServicesAcctRadiusById(ctx context.Context, id string) (*WSGServiceRadiusAccountingService, error) {
	var (
		req  *APIRequest
		resp *WSGServiceRadiusAccountingService
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAcctRadiusById, true)
	req.SetPathParameter("id", id)
}

// FindServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceRadiusAccountingServiceList, error) {
	var (
		req  *APIRequest
		resp *WSGServiceRadiusAccountingServiceList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAcctRadiusByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateServicesAcctRadiusById
//
// Use this API command to modify the basic information of a RADIUS accounting service.
//
// Request Body:
//	 - body *WSGServiceModifyRadiusAccounting
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) PartialUpdateServicesAcctRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAccounting, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAcctRadiusById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}
