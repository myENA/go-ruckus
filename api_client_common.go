package bigdog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
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
)

const (
	SCIAccessTokenQueryParameter = "access_token"

	// 2016-04-06T16:04:46+00:00
	SCIFilterTimestampFormat = "2006-01-02T15:04:05-07:00"
)

const (
	logDebugAPIRequestPrepFormat     = "Preparing api request #%d \"%s %s\""
	logDebugAPIRequestNoBodyFormat   = "%s without body"
	logDebugAPIRequestWithBodyFormat = "%s with body"
)

const (
	uriPathParameterSearchFormat  = "{%s}"
	uriQueryParameterPrefixFormat = "%s?%s"

	apiRequestURLFormat = "%s%s%s"

	headerKeyContentType         = "Content-Type"
	headerKeyContentEncoding     = "Content-Encoding"
	headerKeyContentDisposition  = "Content-Disposition"
	headerKeyContentLength       = "Content-Length"
	headerKeyAccept              = "Accept"
	headerValueApplicationJSON   = "application/json"
	headerValueMultipartFormData = "multipart/form-data"
)

// todo: export this?
type APIAuthProviderError struct {
	meta APIResponseMeta
	err  error
}

func NewAPIAuthProviderError(meta APIResponseMeta, err error) *APIAuthProviderError {
	ape := new(APIAuthProviderError)
	ape.meta = meta
	ape.err = err
	return ape
}

func (ape *APIAuthProviderError) ResponseMeta() APIResponseMeta {
	if ape == nil {
		return APIResponseMeta{}
	}
	return ape.meta
}

func (ape *APIAuthProviderError) Is(err error) bool {
	return ape != nil && ape.err != nil && errors.Is(err, ape.err)
}

func (ape *APIAuthProviderError) Unwrap() error {
	if ape == nil || ape.err == nil {
		return nil
	}
	return ape.err
}

func (ape *APIAuthProviderError) Error() string {
	if ape == nil || ape.err == nil {
		return ""
	}
	return ape.err.Error()
}

type baseClientConfig struct {
	// Address [required]
	//
	// Full address of service, including scheme and port
	Address string

	// PathPrefix [optional]
	//
	// Custom path prefix to prepend to all api calls.  Default is to leave this blank.
	PathPrefix string

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool

	// DisableAutoHydrate [bool]
	//
	// If true, response bodies will no longer automatically hydrated into the returned response models.  This enables
	// you to instead use the response models as Readers to receive the raw bytes of the response in favor of having
	// then unmarshalled if you don't need it.
	DisableAutoHydrate bool

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger

	// HTTPClient [optional]
	HTTPClient *http.Client
}

type baseClient struct {
	log   *log.Logger
	debug bool

	addr       string
	pathPrefix string

	client *http.Client

	autoHydrate bool
}

func newBaseClient(cfg baseClientConfig) *baseClient {
	c := new(baseClient)

	if cfg.Logger == nil {
		c.log = log.New(ioutil.Discard, "", 0)
	} else {
		c.log = cfg.Logger
	}

	c.debug = cfg.Debug
	c.addr = cfg.Address
	c.pathPrefix = cfg.PathPrefix
	c.autoHydrate = !cfg.DisableAutoHydrate

	if cfg.HTTPClient != nil {
		c.client = cfg.HTTPClient
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

func (c *baseClient) do(ctx context.Context, request *APIRequest, authParamName, authParamValue string, mutators ...RequestMutator) (*http.Response, error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		err          error
	)

	if c.debug {
		logMsg := fmt.Sprintf(logDebugAPIRequestPrepFormat, request.ID(), request.Method, request.CompiledURI())

		body := request.Body()

		if body != nil {
			logMsg = fmt.Sprintf(logDebugAPIRequestWithBodyFormat, logMsg)
		} else {
			logMsg = fmt.Sprintf(logDebugAPIRequestNoBodyFormat, logMsg)
		}

		c.log.Print(logMsg)
	}
	if httpRequest, err = request.ToHTTP(ctx, c.addr, c.pathPrefix, authParamName, authParamValue); err != nil {
		return nil, err
	}

	if ml := len(mutators); ml > 0 {
		if c.debug {
			c.log.Printf("Executing %d mutators...", ml)
		}
		for _, fn := range mutators {
			fn(httpRequest)
		}
	}

	httpResponse, err = c.client.Do(httpRequest)
	return httpResponse, err
}

func CleanupReadCloser(rc io.ReadCloser) {
	if rc == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, rc)
	_ = rc.Close()
}

func handleAPIResponse(req *APIRequest, successCode int, httpResp *http.Response, respFunc ResponseFactoryFunc, autoHydrate bool, sourceErr error) (APIResponse, error) {
	var (
		finalErr    error
		cleanupBody bool

		apiResp = respFunc(req, successCode, httpResp)
	)

	defer func() {
		if cleanupBody {
			CleanupReadCloser(httpResp.Body)
		}
	}()

	if sourceErr != nil {
		// if the incoming error is from a service ticket provider, return as-is
		if aerr, ok := sourceErr.(*APIAuthProviderError); ok && aerr != nil {
			cleanupBody = true
			return newErrRawAPIResponse(aerr.ResponseMeta()), aerr
		}

		if httpResp != nil {
			// otherwise, build response meta and return source error
			return apiResp, sourceErr
		}
		return apiResp, sourceErr
	}

	if httpResp == nil {
		panic(fmt.Sprintf("severe problem: nil *http.Response seen with nil error. meta: %v", newAPIResponseMeta(req, successCode, httpResp)))
	}

	// if the response code matches the expected "success" code...
	if httpResp.StatusCode == successCode {
		if modelPtr != nil {
			if w, ok := modelPtr.(io.Writer); ok {
				if _, err := io.Copy(w, httpResp.Body); err != nil {
					finalErr = fmt.Errorf("error copying bytes from response to provided writer: %w", err)
				}
			} else if b, ok := modelPtr.(*[]byte); ok {
				if tmp, err := ioutil.ReadAll(httpResp.Body); err != nil && err != io.EOF {
					finalErr = fmt.Errorf("error reading bytes from response: %w", err)
				} else {
					*b = tmp
				}
			} else if rr, ok := modelPtr.(*RawAPIResponse); ok {
				cleanupBody = false
				rr.Body = httpResp.Body
				rr.Header = httpResp.Header
			} else if err := json.NewDecoder(httpResp.Body).Decode(modelPtr); err != nil && err != io.EOF {
				// ... and this query has a modeled response, attempt to unmarshal into that type
				finalErr = fmt.Errorf("error unmarshalling response body into %T: %w", modelPtr, err)
			}
		}
	} else {
		var apiErr error
		// todo: do something better here.
		if strings.HasPrefix(req.URI, "/wsg/") || strings.HasPrefix(req.URI, "/switchm/") {
			apiErr = new(VSZAPIError)
		} else {
			apiErr = new(SCIAPIError)
		}
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

func SCIFilterFormatTimestamp(t time.Time) string {
	return t.Format(SCIFilterTimestampFormat)
}
