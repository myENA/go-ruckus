package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/urlfiltering"
)

type WSGURLFilteringPolicyService struct {
	client *Client
}

func NewWSGURLFilteringPolicyService(client *Client) *WSGURLFilteringPolicyService {
	s := new(WSGURLFilteringPolicyService)
	s.client = client
	return s
}

func (ss *WSGService) WSGURLFilteringPolicyService() *WSGURLFilteringPolicyService {
	serv := new(WSGURLFilteringPolicyService)
	serv.client = ss.client
	return serv
}

// FindUrlFilteringBlockCategories
//
// Use this API command to retrieve the block categories of URL Filtering.
func (s *WSGURLFilteringPolicyService) FindUrlFilteringBlockCategories(ctx context.Context) (*urlfiltering.UrlFilteringBlockCategoriesList, error) {
}

// FindUrlFilteringByQueryCriteria
//
// Use this API command to retrieve a list of URL Filtering policies by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGURLFilteringPolicyService) FindUrlFilteringByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*urlfiltering.UrlFilteringPolicyList, error) {
}

// FindUrlFilteringUrlFilteringPolicy
//
// Use this API command to retrieve list of URL Filtering policies.
//
// Query Parameters:
// - qDomainId string
// - qIndex string
// - qListSize string
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicy(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*urlfiltering.UrlFilteringPolicyList, error) {
}

// FindUrlFilteringUrlFilteringPolicyById
//
// Use this API command to retrieve an URL Filtering policy.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicyById(ctx context.Context, pId string) (*urlfiltering.UrlFilteringPolicy, error) {
}
