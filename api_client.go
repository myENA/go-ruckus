package ruckus

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
)

const (
	VSZResourceNameCluster       = "CLUSTER_CATEGORY"
	VSZResourceNameAP            = "AP_CATEGORY"
	VSZResourceNameWLAN          = "WLAN_CATEGORY"
	VSZResourceNameDevice        = "DEVICE_CATEGORY"
	VSZResourceNameAdministrator = "ADMINISTRATOR_CATEGORY"
	VSZResourceNameTenant        = "TENANT_CATEGORY"
	VSZResourceNameICX           = "ICX_CATEGORY"

	VSZResourceDisplayCluster       = "SZ"
	VSZResourceDisplayAP            = "AP"
	VSZResourceDisplayWLAN          = "WLAN"
	VSZResourceDisplayDevice        = "User/Device/App"
	VSZResourceDisplayAdministrator = "Admin"
	VSZResourceDisplayTenant        = "MVNO"
	VSZResourceDisplayICX           = "ICX Switch"

	VSZResourceAccessFullAccess = "FULL_ACCESS"
	VSZResourceAccessModify     = "MODIFY"
	VSZResourceAccessReadOnly   = "READ_ONLY"
	VSZResourceAccessRead       = "READ"
	VSZResourceAccessNoAccess   = "NO_ACCESS"

	VSZRoleNameSuperAdmin           = "SUPER_ADMIN"
	VSZRoleNameSystemAdmin          = "SYSTEM_ADMIN"
	VSZRoleNameReadOnlySystemAdmin  = "RO_SYSTEM_ADMIN"
	VSZRoleNameNetworkAdmin         = "NETWORK_ADMIN"
	VSZRoleNameReadOnlyNetworkAdmin = "RO_NETWORK_ADMIN"
	VSZRoleNameAPAdmin              = "AP_ADMIN"
	VSZRoleNameGuestPassAdmin       = "GUEST_PASS_ADMIN"
	VSZRoleNameCustom               = "CUSTOM"

	VSZRoleLabelSuperAdmin           = "Super Admin"
	VSZRoleLabelSystemAdmin          = "System Admin"
	VSZRoleLabelReadOnlySystemAdmin  = "Read-Only System Admin"
	VSZRoleLabelNetworkAdmin         = "Network Admin"
	VSZRoleLabelReadOnlyNetworkAdmin = "Read-Only Network Admin"
	VSZRoleLabelAPAdmin              = "AP Admin"
	VSZRoleLabelGuestPassAdmin       = "Guest Pass Admin"
	VSZRoleLabelCustom               = "SCGAdmin"

	VSZSecurityProfileNameDefault    = "Default"
	VSZSecurityProfileNameMoreSecure = "More Secure"

	VSZDomainTypeRegular = "REGULAR"
	VSZDomainTypePartner = "PARTNER"

	VSZServiceTicketQueryParameter = "serviceTicket"

	SCIAccessTokenQueryParameter = "access_token"

	logDebugAPIRequestPrepFormat     = "Preparing api request #%d \"%s %s\""
	logDebugAPIRequestNoBodyFormat   = "%s without body"
	logDebugAPIRequestWithBodyFormat = "%s with body: %s"
)

type AuthParam struct {
	Name  string
	Value string
}

type baseClient struct {
	log   *log.Logger
	debug bool

	addr       string
	pathPrefix string

	client *http.Client
}

func newBaseClient(l *log.Logger, debug bool, addr, pathPrefix string, client *http.Client) *baseClient {
	c := new(baseClient)

	if l == nil {
		c.log = log.New(ioutil.Discard, "", 0)
	} else {
		c.log = l
	}

	c.debug = debug
	c.addr = addr
	c.pathPrefix = pathPrefix

	if client != nil {
		c.client = client
	} else {
		// pooled transport config shamelessly borrowed from
		// https://github.com/hashicorp/go-cleanhttp/blob/v0.5.1/cleanhttp.go
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			},
		}
	}

	return c
}

// VSZAddress returns the address of the VSZ node currently being connected to
func (c *baseClient) Address() string {
	return c.addr
}

// VSZPathPrefix returns the prefix that will be applied to all api calls.  It is valid for this to be blank, and is by
// default
func (c *baseClient) PathPrefix() string {
	return c.pathPrefix
}

func (c *baseClient) do(ctx context.Context, request *APIRequest, authParamName, authParamValue string) (*http.Response, error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		err          error
	)

	if c.debug {
		logMsg := fmt.Sprintf(logDebugAPIRequestPrepFormat, request.ID(), request.Method(), request.CompiledURI())

		body := request.Body()

		if len(body) == 0 {
			logMsg = fmt.Sprintf(logDebugAPIRequestNoBodyFormat, logMsg)
		} else {
			logMsg = fmt.Sprintf(logDebugAPIRequestWithBodyFormat, logMsg, string(body))
		}

		c.log.Print(logMsg)
	}
	if httpRequest, err = request.ToHTTP(ctx, c.addr, c.pathPrefix, authParamName, authParamValue); err != nil {
		return nil, err
	}
	httpResponse, err = c.client.Do(httpRequest)
	return httpResponse, err
}

type VSZClientConfig struct {
	// Address [required]
	//
	// Full address of VSZ, including scheme and port
	Address string

	// PathPrefix [optional]
	//
	// Custom path prefix to prepend to all api calls.  Default is to leave this blank.
	PathPrefix string

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool

	// ServiceTicketProvider [required]
	//
	// ServiceTicketProvider to use to handle request auth session
	ServiceTicketProvider VSZServiceTicketProvider

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger

	// HTTPClient [optional]
	HTTPClient *http.Client
}

type VSZClient struct {
	*baseClient
	stp VSZServiceTicketProvider
}

func NewVSZClient(config *VSZClientConfig) *VSZClient {
	c := new(VSZClient)
	c.baseClient = newBaseClient(config.Logger, config.Debug, config.Address, config.PathPrefix, config.HTTPClient)
	c.stp = config.ServiceTicketProvider

	return c
}

// VSZServiceTicketProvider returns the provider used by this client
func (c *VSZClient) ServiceTicketProvider() VSZServiceTicketProvider {
	return c.stp
}

func (c *VSZClient) WSG() *WSGService {
	return NewWSGService(c)
}

func (c *VSZClient) SwitchM() *SwitchMService {
	return NewSwitchMService(c)
}

func (c *VSZClient) SCGAdmin() *SCGAdminService {
	return NewSCGAdminService(c)
}

func (c *VSZClient) Do(ctx context.Context, request *APIRequest) (*http.Response, error) {
	var (
		cas           VSZServiceTicketCAS
		serviceTicket string
		err           error
	)
	if request.authenticated {
		if cas, serviceTicket, err = c.stp.Current(); err != nil {
			if c.debug {
				c.log.Printf("Error fetching current service ticket: %v", err)
			}
			if cas, err = c.stp.Refresh(ctx, c, cas); err != nil {
				if c.debug {
					c.log.Printf("Error refreshing service ticket: %v", err)
				}
				return nil, err
			} else if cas, serviceTicket, err = c.stp.Current(); err != nil {
				if c.debug {
					c.log.Printf("Error fetching current service ticket after refresh: %v", err)
				}
				return nil, err
			}
		}
	}
	return c.do(ctx, request, VSZServiceTicketQueryParameter, serviceTicket)
}

type SCIClientConfig struct {
	// Address [required]
	//
	// Full address of SCI, including scheme and port
	Address string

	// PathPrefix [optional]
	//
	// Custom path prefix to prepend to all api calls.  Default is to leave this blank.
	PathPrefix string

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool

	// ServiceTicketProvider [required]
	//
	// ServiceTicketProvider to use to handle request auth session
	AccessTokenProvider SCIAccessTokenProvider

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger

	// HTTPClient [optional]
	HTTPClient *http.Client
}

type SCIClient struct {
	*baseClient
	atp SCIAccessTokenProvider
}

func NewSCIClient(config *SCIClientConfig) *SCIClient {
	c := new(SCIClient)
	c.baseClient = newBaseClient(config.Logger, config.Debug, config.Address, config.PathPrefix, config.HTTPClient)
	c.atp = config.AccessTokenProvider
	return c
}

func (c *SCIClient) AccessTokenProvider() SCIAccessTokenProvider {
	return c.atp
}

func (c *SCIClient) SCI() *SCIService {
	return NewSCIService(c)
}

func (c *SCIClient) Do(ctx context.Context, request *APIRequest) (*http.Response, error) {
	var (
		cas         SCIAccessTokenCAS
		accessToken string
		err         error
	)
	if request.authenticated {
		if cas, accessToken, err = c.atp.Current(); err != nil {
			if c.debug {
				c.log.Printf("Error fetching current access token: %v", err)
			}
			if cas, err = c.atp.Refresh(ctx, c, cas); err != nil {
				c.log.Printf("Error refreshing access token: %v", err)
			}
			return nil, err
		} else if cas, accessToken, err = c.atp.Current(); err != nil {
			if c.debug {
				c.log.Printf("Error fetching current access token after refresh: %v", err)
			}
			return nil, err
		}
	}
	return c.do(ctx, request, SCIAccessTokenQueryParameter, accessToken)
}

type APIResponseMeta struct {
	RequestMethod string `json:"requestMethod"`
	RequestURI    string `json:"requestURI"`

	SuccessCode int `json:"successCode"`

	ResponseCode   int    `json:"responseCode"`
	ResponseStatus string `json:"responseStatus"`
}

func newAPIResponseMeta(req *APIRequest, successCode int, httpResp *http.Response) *APIResponseMeta {
	rm := new(APIResponseMeta)
	rm.RequestMethod = req.Method()
	rm.RequestURI = req.CompiledURI()
	rm.SuccessCode = successCode
	if httpResp != nil {
		rm.ResponseCode = httpResp.StatusCode
		rm.ResponseStatus = http.StatusText(httpResp.StatusCode)
	}
	return rm
}

func newErrAPIResponseMeta() *APIResponseMeta {
	rm := new(APIResponseMeta)
	rm.ResponseCode = http.StatusInternalServerError
	rm.ResponseStatus = http.StatusText(http.StatusInternalServerError)
	return rm
}

func (rm *APIResponseMeta) String() string {
	var msg string
	if rm.SuccessCode == rm.ResponseCode {
		msg = "Successful"
	} else {
		msg = "Failed"
	}
	return fmt.Sprintf("%s response from request %s %s", msg, rm.RequestMethod, rm.RequestURI)
}

func cleanupHTTPResponseBody(hp *http.Response) {
	if hp == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, hp.Body)
	_ = hp.Body.Close()
}

func handleResponse(req *APIRequest, successCode int, httpResp *http.Response, modelPtr interface{}, sourceErr error) (*APIResponseMeta, error) {
	var finalErr error

	defer cleanupHTTPResponseBody(httpResp)

	if sourceErr != nil {
		// if the incoming error is from a service ticket provider, return as-is
		if sterr, ok := sourceErr.(*ServiceTicketProviderError); ok {
			return sterr.ResponseMeta(), sterr
		}

		if httpResp != nil {
			// otherwise, build response meta and return source error
			return newAPIResponseMeta(req, successCode, httpResp), sourceErr
		}

		return newErrAPIResponseMeta(), sourceErr
	}

	if httpResp == nil {
		panic(fmt.Sprintf("severe problem: nil *http.Response seen with nil error. meta: %v", newAPIResponseMeta(req, successCode, httpResp)))
	}

	if httpResp.StatusCode == successCode {
		// if the response code matches the expected "success" code...
		if modelPtr != nil {
			// ... and this query has a modeled response, attempt to unmarshal into that type
			if err := json.NewDecoder(httpResp.Body).Decode(modelPtr); err != nil && err != io.EOF {
				finalErr = fmt.Errorf("error unmarshalling response body into %T: %w", modelPtr, err)
			}
		}
	} else {
		apiErr := new(APIError)
		if err := json.NewDecoder(httpResp.Body).Decode(apiErr); err != nil {
			finalErr = fmt.Errorf("error unmarshalling error body into %T: %w", finalErr, err)
		} else {
			finalErr = apiErr
		}
	}

	// build response.
	return newAPIResponseMeta(req, successCode, httpResp), finalErr
}

type WSGService struct {
	apiClient *VSZClient
}

func NewWSGService(c *VSZClient) *WSGService {
	s := new(WSGService)
	s.apiClient = c
	return s
}

type SwitchMService struct {
	apiClient *VSZClient
}

func NewSwitchMService(c *VSZClient) *SwitchMService {
	s := new(SwitchMService)
	s.apiClient = c
	return s
}

type SCIService struct {
	apiClient *SCIClient
}

func NewSCIService(c *SCIClient) *SCIService {
	s := new(SCIService)
	s.apiClient = c
	return s
}