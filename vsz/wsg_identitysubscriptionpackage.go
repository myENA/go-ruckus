package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGIdentitySubscriptionPackageService struct {
	apiClient *APIClient
}

func NewWSGIdentitySubscriptionPackageService(c *APIClient) *WSGIdentitySubscriptionPackageService {
	s := new(WSGIdentitySubscriptionPackageService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentitySubscriptionPackageService() *WSGIdentitySubscriptionPackageService {
	return NewWSGIdentitySubscriptionPackageService(ss.apiClient)
}

// AddIdentityPackageList
//
// Use this API command to retrieve a list of subscription package.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackageList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentitySubscriptionPackageList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentitySubscriptionPackageList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityPackageList, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddIdentityPackages
//
// Use this API command to create subscription package.
//
// Request Body:
//	 - body *WSGIdentityCreateSubscriptionPackage
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackages(ctx context.Context, body *WSGIdentityCreateSubscriptionPackage) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityPackages, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteIdentityPackages
//
// Use this API command to delete multiple subscription packages.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackages(ctx context.Context, body *WSGIdentityDeleteBulk) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityPackages, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}

// DeleteIdentityPackagesById
//
// Use this API command to delete subscription package.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackagesById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityPackagesById, true)
	req.SetPathParameter("id", id)
}

// FindIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages(ctx context.Context) (*WSGIdentitySubscriptionPackageList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentitySubscriptionPackageList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityPackages, true)
}

// FindIdentityPackagesById
//
// Use this API command to retrieve subscription package.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById(ctx context.Context, id string) (*WSGIdentitySubscriptionPackage, error) {
	var (
		req  *APIRequest
		resp *WSGIdentitySubscriptionPackage
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityPackagesById, true)
	req.SetPathParameter("id", id)
}

// PartialUpdateIdentityPackagesById
//
// Use this API command to modify the basic information of subscription package.
//
// Request Body:
//	 - body *WSGIdentityModifySubscriptionPackage
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) PartialUpdateIdentityPackagesById(ctx context.Context, body *WSGIdentityModifySubscriptionPackage, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityPackagesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}
