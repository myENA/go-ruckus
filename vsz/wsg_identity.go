package vsz

// API Version: v8_1

type WSGIdentityAaaServer struct {
	// Id
	// the identifier of the AAA server
	Id *string `json:"id,omitempty"`

	// Name
	// the identifier of the AAA server
	Name *string `json:"name,omitempty"`
}

func NewWSGIdentityAaaServer() *WSGIdentityAaaServer {
	m := new(WSGIdentityAaaServer)
	return m
}

type WSGIdentityAaaServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentityAaaServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentityAaaServerList() *WSGIdentityAaaServerList {
	m := new(WSGIdentityAaaServerList)
	return m
}

type WSGIdentityAuthenticationServerConfig struct {
	// AuthenticationServerID
	// Authentication server id
	AuthenticationServerID *string `json:"AUTHENTICATION_SERVER_ID,omitempty"`

	// AuthenticationServerName
	// Authentication server name
	AuthenticationServerName *string `json:"AUTHENTICATION_SERVER_NAME,omitempty"`

	// AuthenticationServerType
	// Authentication server type
	AuthenticationServerType *string `json:"AUTHENTICATION_SERVER_TYPE,omitempty"`

	// Id
	// server id
	Id *string `json:"id,omitempty"`

	// Local
	// Is local server
	Local *bool `json:"local,omitempty"`

	// Name
	// server name
	Name *string `json:"name,omitempty"`

	// Type
	// server type
	Type *string `json:"type,omitempty"`
}

func NewWSGIdentityAuthenticationServerConfig() *WSGIdentityAuthenticationServerConfig {
	m := new(WSGIdentityAuthenticationServerConfig)
	return m
}

type WSGIdentityCountryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentityCountrySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentityCountryList() *WSGIdentityCountryList {
	m := new(WSGIdentityCountryList)
	return m
}

type WSGIdentityCountrySummary struct {
	// CountryName
	// Full name of country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Short name of country
	CountryShortName *string `json:"countryShortName,omitempty"`
}

func NewWSGIdentityCountrySummary() *WSGIdentityCountrySummary {
	m := new(WSGIdentityCountrySummary)
	return m
}

type WSGIdentityCreateIdentityGuestPass struct {
	// AutoGeneratedPassword
	// Pass generation
	AutoGeneratedPassword *bool `json:"autoGeneratedPassword,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// GuestName
	// Constraints:
	//    - required
	GuestName *WSGCommonNormalName `json:"guestName" validate:"required,max=32,min=2"`

	// MaxDevices
	// Constraints:
	//    - required
	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices" validate:"required"`

	// NumberOfPasses
	// Number of passes
	// Constraints:
	//    - required
	NumberOfPasses *int `json:"numberOfPasses" validate:"required"`

	// PassEffectSince
	// Pass effective since
	// Constraints:
	//    - oneof:[CREATION_TIME,FIRST_USE]
	PassEffectSince *string `json:"passEffectSince,omitempty" validate:"oneof=CREATION_TIME FIRST_USE"`

	// PassUseDays
	// Expire new guest pass if not used within
	PassUseDays *int `json:"passUseDays,omitempty"`

	// PassValidFor
	// Constraints:
	//    - required
	PassValidFor *WSGIdentityPassValidFor `json:"passValidFor" validate:"required"`

	// PassValue
	// Pass value
	PassValue *string `json:"passValue,omitempty"`

	// Remarks
	// Remarks
	Remarks *string `json:"remarks,omitempty"`

	SessionDuration *WSGIdentitySessionDuration `json:"sessionDuration,omitempty"`

	// Wlan
	// Constraints:
	//    - required
	Wlan *WSGCommonGenericRef `json:"wlan" validate:"required"`

	// Zone
	// Constraints:
	//    - required
	Zone *WSGCommonGenericRef `json:"zone" validate:"required"`
}

func NewWSGIdentityCreateIdentityGuestPass() *WSGIdentityCreateIdentityGuestPass {
	m := new(WSGIdentityCreateIdentityGuestPass)
	return m
}

type WSGIdentityCreateIdentityUserRole struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// MaxDevices
	// Constraints:
	//    - required
	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName2to64 `json:"name" validate:"required,max=64,min=2"`

	// UserTrafficProfile
	// Constraints:
	//    - required
	UserTrafficProfile *WSGCommonGenericRef `json:"userTrafficProfile" validate:"required"`

	// VlanId
	// vlan id
	// Constraints:
	//    - min:1
	//    - max:4096
	VlanId *int `json:"vlanId,omitempty" validate:"gte=1,lte=4096"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGIdentityCreateIdentityUserRole() *WSGIdentityCreateIdentityUserRole {
	m := new(WSGIdentityCreateIdentityUserRole)
	return m
}

type WSGIdentityCreateSubscriptionPackage struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - required
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval" validate:"required,oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	// Constraints:
	//    - required
	ExpirationValue *int `json:"expirationValue" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGIdentityCreateSubscriptionPackage() *WSGIdentityCreateSubscriptionPackage {
	m := new(WSGIdentityCreateSubscriptionPackage)
	return m
}

type WSGIdentityCreateUser struct {
	// Address
	// Address
	Address *string `json:"address,omitempty"`

	// City
	// City
	City *string `json:"city,omitempty"`

	// CountryName
	// Country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Country
	CountryShortName *string `json:"countryShortName,omitempty"`

	// DomainId
	// Domain ID
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// Email
	Email *string `json:"email,omitempty"`

	// FirstName
	// First name
	// Constraints:
	//    - required
	FirstName *string `json:"firstName" validate:"required"`

	// IsDisabled
	// Is Disabled
	// Constraints:
	//    - required
	//    - oneof:[NO,YES]
	IsDisabled *string `json:"isDisabled" validate:"required,oneof=NO YES"`

	// LastName
	// Last Name
	// Constraints:
	//    - required
	LastName *string `json:"lastName" validate:"required"`

	// Password
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// Phone
	// Phone
	Phone *string `json:"phone,omitempty"`

	// Remark
	// Remark
	Remark *string `json:"remark,omitempty"`

	// State
	// State
	State *string `json:"state,omitempty"`

	SubscriberPackage *WSGCommonGenericRef `json:"subscriberPackage,omitempty"`

	// UserName
	// User Name
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`

	// ZipCode
	// Zip Code
	ZipCode *string `json:"zipCode,omitempty"`
}

func NewWSGIdentityCreateUser() *WSGIdentityCreateUser {
	m := new(WSGIdentityCreateUser)
	return m
}

type WSGIdentityDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGIdentityDeleteBulk() *WSGIdentityDeleteBulk {
	m := new(WSGIdentityDeleteBulk)
	return m
}

type WSGIdentityGuestPassConfiguration struct {
	// AutoGeneratedPassword
	// Pass generation
	AutoGeneratedPassword *bool `json:"autoGeneratedPassword,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationDate
	// Expiration date and time
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// GeneratedOn
	// Generated date and time
	GeneratedOn *string `json:"generatedOn,omitempty"`

	GuestName *WSGCommonNormalName `json:"guestName,omitempty"`

	// Id
	// ID of the identity guest pass
	Id *string `json:"id,omitempty"`

	// Key
	// Identifier of the identity guest pass
	Key *string `json:"key,omitempty"`

	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices,omitempty"`

	// PassEffectSince
	// Pass effective since
	// Constraints:
	//    - oneof:[CREATION_TIME,FIRST_USE]
	PassEffectSince *string `json:"passEffectSince,omitempty" validate:"oneof=CREATION_TIME FIRST_USE"`

	// PassUseDays
	// Expire new guest pass if not used within
	PassUseDays *int `json:"passUseDays,omitempty"`

	PassValidFor *WSGIdentityPassValidFor `json:"passValidFor,omitempty"`

	// Remarks
	// Remarks
	Remarks *string `json:"remarks,omitempty"`

	SessionDuration *WSGIdentitySessionDuration `json:"sessionDuration,omitempty"`

	// Ssid
	// SSID
	Ssid *string `json:"ssid,omitempty"`

	// UserId
	// user ID of the identity guest pass
	UserId *string `json:"userId,omitempty"`

	Wlan *WSGCommonGenericRef `json:"wlan,omitempty"`

	// WlanRestrition
	// Wlan description
	WlanRestrition *string `json:"wlanRestrition,omitempty"`

	Zone *WSGCommonGenericRef `json:"zone,omitempty"`
}

func NewWSGIdentityGuestPassConfiguration() *WSGIdentityGuestPassConfiguration {
	m := new(WSGIdentityGuestPassConfiguration)
	return m
}

type WSGIdentityGuestPassList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentityGuestPassConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentityGuestPassList() *WSGIdentityGuestPassList {
	m := new(WSGIdentityGuestPassList)
	return m
}

type WSGIdentityList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentityListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentityList() *WSGIdentityList {
	m := new(WSGIdentityList)
	return m
}

type WSGIdentityListType struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	UserTrafficProfile *WSGCommonGenericRef `json:"userTrafficProfile,omitempty"`

	// VlanId
	// vlan id
	VlanId *int `json:"vlanId,omitempty"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGIdentityListType() *WSGIdentityListType {
	m := new(WSGIdentityListType)
	return m
}

type WSGIdentityUserRole struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// the identifier of the object
	Id *string `json:"id,omitempty"`

	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`

	UserTrafficProfile *WSGCommonGenericRef `json:"userTrafficProfile,omitempty"`

	// VlanId
	// vlan id
	VlanId *int `json:"vlanId,omitempty"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGIdentityUserRole() *WSGIdentityUserRole {
	m := new(WSGIdentityUserRole)
	return m
}

type WSGIdentityUserSummary struct {
	// CreatedOn
	// Created on
	CreatedOn *string `json:"createdOn,omitempty"`

	DisplayName *WSGCommonNormalName `json:"displayName,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the identity user
	Id *string `json:"id,omitempty"`

	// IsDisabled
	// Is disalbed
	IsDisabled *string `json:"isDisabled,omitempty"`

	// UserName
	// User Name
	UserName *string `json:"userName,omitempty"`

	// UserSource
	// User Source
	UserSource *string `json:"userSource,omitempty"`

	// UserType
	// User Type
	UserType *string `json:"userType,omitempty"`
}

func NewWSGIdentityUserSummary() *WSGIdentityUserSummary {
	m := new(WSGIdentityUserSummary)
	return m
}

type WSGIdentityImportIdentityGuestPass struct {
	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// MaxDevices
	// Constraints:
	//    - required
	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices" validate:"required"`

	// PassEffectSince
	// Pass effective since
	// Constraints:
	//    - oneof:[CREATION_TIME,FIRST_USE]
	PassEffectSince *string `json:"passEffectSince,omitempty" validate:"oneof=CREATION_TIME FIRST_USE"`

	// PassUseDays
	// Expire new guest pass if not used within
	PassUseDays *int `json:"passUseDays,omitempty"`

	// PassValidFor
	// Constraints:
	//    - required
	PassValidFor *WSGIdentityPassValidFor `json:"passValidFor" validate:"required"`

	SessionDuration *WSGIdentitySessionDuration `json:"sessionDuration,omitempty"`

	// Wlan
	// Constraints:
	//    - required
	Wlan *WSGCommonGenericRef `json:"wlan" validate:"required"`

	// Zone
	// Constraints:
	//    - required
	Zone *WSGCommonGenericRef `json:"zone" validate:"required"`
}

func NewWSGIdentityImportIdentityGuestPass() *WSGIdentityImportIdentityGuestPass {
	m := new(WSGIdentityImportIdentityGuestPass)
	return m
}

type WSGIdentityMaxDevices struct {
	// MaxDevicesAllowed
	// Max devices allowed
	// Constraints:
	//    - default:'LIMITED'
	//    - oneof:[UNLIMITED,LIMITED]
	MaxDevicesAllowed *string `json:"maxDevicesAllowed,omitempty" validate:"oneof=UNLIMITED LIMITED"`

	// MaxDevicesNumber
	// max devices number
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:10
	MaxDevicesNumber *int `json:"maxDevicesNumber,omitempty" validate:"gte=1,lte=10"`
}

func NewWSGIdentityMaxDevices() *WSGIdentityMaxDevices {
	m := new(WSGIdentityMaxDevices)
	return m
}

type WSGIdentityModifyIdentityUserRole struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	MaxDevices *WSGIdentityMaxDevices `json:"maxDevices,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`

	UserTrafficProfile *WSGCommonGenericRef `json:"userTrafficProfile,omitempty"`

	// VlanId
	// vlan id
	// Constraints:
	//    - min:1
	//    - max:4096
	VlanId *int `json:"vlanId,omitempty" validate:"gte=1,lte=4096"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGIdentityModifyIdentityUserRole() *WSGIdentityModifyIdentityUserRole {
	m := new(WSGIdentityModifyIdentityUserRole)
	return m
}

type WSGIdentityModifySubscriptionPackage struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval,omitempty" validate:"oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	ExpirationValue *int `json:"expirationValue,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGIdentityModifySubscriptionPackage() *WSGIdentityModifySubscriptionPackage {
	m := new(WSGIdentityModifySubscriptionPackage)
	return m
}

type WSGIdentityModifyUser struct {
	// Address
	// Address
	Address *string `json:"address,omitempty"`

	// City
	// City
	City *string `json:"city,omitempty"`

	// CountryName
	// Country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Country
	CountryShortName *string `json:"countryShortName,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// Email
	Email *string `json:"email,omitempty"`

	// FirstName
	// First name
	FirstName *string `json:"firstName,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// IsDisabled
	// Is Disabled
	// Constraints:
	//    - oneof:[NO,YES]
	IsDisabled *string `json:"isDisabled,omitempty" validate:"oneof=NO YES"`

	// LastName
	// Last Name
	LastName *string `json:"lastName,omitempty"`

	// Password
	// Password
	Password *string `json:"password,omitempty"`

	// Phone
	// Phone
	Phone *string `json:"phone,omitempty"`

	// Remark
	// Remark
	Remark *string `json:"remark,omitempty"`

	// State
	// State
	State *string `json:"state,omitempty"`

	SubscriberPackage *WSGCommonGenericRef `json:"subscriberPackage,omitempty"`

	// ZipCode
	// Zip Code
	ZipCode *string `json:"zipCode,omitempty"`
}

func NewWSGIdentityModifyUser() *WSGIdentityModifyUser {
	m := new(WSGIdentityModifyUser)
	return m
}

type WSGIdentityPackageConfiguration struct {
	// PackageExpiration
	// Package expiration interval and value
	PackageExpiration *string `json:"packageExpiration,omitempty"`

	SubscriberPackage *WSGCommonGenericRef `json:"subscriberPackage,omitempty"`
}

func NewWSGIdentityPackageConfiguration() *WSGIdentityPackageConfiguration {
	m := new(WSGIdentityPackageConfiguration)
	return m
}

type WSGIdentityPackageList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentityPackageConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentityPackageList() *WSGIdentityPackageList {
	m := new(WSGIdentityPackageList)
	return m
}

type WSGIdentityPassValidFor struct {
	// ExpirationUnit
	// Constraints:
	//    - oneof:[HOUR,DAY,WEEK]
	ExpirationUnit *string `json:"expirationUnit,omitempty" validate:"oneof=HOUR DAY WEEK"`

	ExpirationValue *int `json:"expirationValue,omitempty"`
}

func NewWSGIdentityPassValidFor() *WSGIdentityPassValidFor {
	m := new(WSGIdentityPassValidFor)
	return m
}

type WSGIdentityQueryCriteria struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []*WSGCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGCommonQueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"gte=1"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty" validate:"gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGIdentityQueryCriteria() *WSGIdentityQueryCriteria {
	m := new(WSGIdentityQueryCriteria)
	return m
}

type WSGIdentitySessionDuration struct {
	RequireLoginAgain *bool `json:"requireLoginAgain,omitempty"`

	// SessionUnit
	// Constraints:
	//    - oneof:[MIN,HOUR,DAY,WEEK]
	SessionUnit *string `json:"sessionUnit,omitempty" validate:"oneof=MIN HOUR DAY WEEK"`

	SessionValue *int `json:"sessionValue,omitempty"`
}

func NewWSGIdentitySessionDuration() *WSGIdentitySessionDuration {
	m := new(WSGIdentitySessionDuration)
	return m
}

type WSGIdentitySubscriptionPackage struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval,omitempty" validate:"oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	ExpirationValue *int `json:"expirationValue,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGIdentitySubscriptionPackage() *WSGIdentitySubscriptionPackage {
	m := new(WSGIdentitySubscriptionPackage)
	return m
}

type WSGIdentitySubscriptionPackageList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentitySubscriptionPackageListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentitySubscriptionPackageList() *WSGIdentitySubscriptionPackageList {
	m := new(WSGIdentitySubscriptionPackageList)
	return m
}

type WSGIdentitySubscriptionPackageListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval,omitempty" validate:"oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	ExpirationValue *int `json:"expirationValue,omitempty"`

	// Id
	// the identifier of the subscription package
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGIdentitySubscriptionPackageListType() *WSGIdentitySubscriptionPackageListType {
	m := new(WSGIdentitySubscriptionPackageListType)
	return m
}

type WSGIdentityUserConfiguration struct {
	// Address
	// Address
	// Constraints:
	//    - max:256
	//    - min:2
	Address *string `json:"address,omitempty" validate:"max=256,min=2"`

	// City
	// City
	// Constraints:
	//    - max:50
	//    - min:2
	City *string `json:"city,omitempty" validate:"max=50,min=2"`

	// CountryName
	// Country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Country
	CountryShortName *string `json:"countryShortName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	CredentialsGuestPassDto *WSGIdentityUserConfigurationCredentialsGuestPassDtoType `json:"credentialsGuestPassDto,omitempty"`

	// Email
	// Email
	Email *string `json:"email,omitempty"`

	// FirstName
	// First name
	// Constraints:
	//    - max:32
	//    - min:2
	FirstName *string `json:"firstName,omitempty" validate:"max=32,min=2"`

	// IsDisabled
	// Is Disabled
	// Constraints:
	//    - oneof:[NO,YES]
	IsDisabled *string `json:"isDisabled,omitempty" validate:"oneof=NO YES"`

	// LastName
	// Last Name
	// Constraints:
	//    - max:32
	//    - min:2
	LastName *string `json:"lastName,omitempty" validate:"max=32,min=2"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// PackageExpirationDate
	// Package Expiration Date
	PackageExpirationDate *int `json:"packageExpirationDate,omitempty"`

	// PackageExpirationInterval
	// Package Expiration Interval
	// Constraints:
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	PackageExpirationInterval *string `json:"packageExpirationInterval,omitempty" validate:"oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// PackageExpirationValue
	// Package Expiration Value
	PackageExpirationValue *int `json:"packageExpirationValue,omitempty"`

	// PackageStatus
	// Package Status
	// Constraints:
	//    - oneof:[DEPLETED,AVAILABLE,EXPIRED,TERMINATED,REMOVED]
	PackageStatus *string `json:"packageStatus,omitempty" validate:"oneof=DEPLETED AVAILABLE EXPIRED TERMINATED REMOVED"`

	// Phone
	// Phone
	// Constraints:
	//    - max:32
	//    - min:2
	Phone *string `json:"phone,omitempty" validate:"max=32,min=2"`

	// Remark
	// Remark
	// Constraints:
	//    - max:32
	//    - min:2
	Remark *string `json:"remark,omitempty" validate:"max=32,min=2"`

	// State
	// State
	// Constraints:
	//    - max:32
	//    - min:2
	State *string `json:"state,omitempty" validate:"max=32,min=2"`

	SubscriberPackage *WSGCommonGenericRef `json:"subscriberPackage,omitempty"`

	// UserName
	// User Name
	// Constraints:
	//    - max:64
	//    - min:2
	UserName *string `json:"userName,omitempty" validate:"max=64,min=2"`

	UsernamePasswordCredentialsImplDto *WSGIdentityUsernamePasswordCredentialsImplDto `json:"usernamePasswordCredentialsImplDto,omitempty"`

	// ZipCode
	// Zip Code
	// Constraints:
	//    - max:32
	//    - min:2
	ZipCode *string `json:"zipCode,omitempty" validate:"max=32,min=2"`
}

func NewWSGIdentityUserConfiguration() *WSGIdentityUserConfiguration {
	m := new(WSGIdentityUserConfiguration)
	return m
}

type WSGIdentityUserConfigurationCredentialsGuestPassDtoType struct {
	// AuthenticationMethod
	// Authentication method of credential
	// Constraints:
	//    - oneof:[GUEST_PASS]
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" validate:"oneof=GUEST_PASS"`

	// AutoGeneratePassword
	// Pass generation
	AutoGeneratePassword *bool `json:"autoGeneratePassword,omitempty"`

	Comment *string `json:"comment,omitempty"`

	// CreationDate
	// Creation Date
	CreationDate *int `json:"creationDate,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// DisplayName
	// filter identity user list by display name
	DisplayName *string `json:"displayName,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationDate
	// Expiration date and time
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// ExpirationUnit
	// Constraints:
	//    - oneof:[HOUR,DAY,WEEK]
	ExpirationUnit *string `json:"expirationUnit,omitempty" validate:"oneof=HOUR DAY WEEK"`

	ExpirationValue *int `json:"expirationValue,omitempty"`

	ExpireAfterIfNotUsed *int `json:"expireAfterIfNotUsed,omitempty"`

	// GuestExpiration
	// Expiration time of guest pass
	GuestExpiration *int `json:"guestExpiration,omitempty"`

	// Key
	// Key of this guest pass
	Key *string `json:"key,omitempty"`

	// LoginPassword
	// Login Password
	LoginPassword *string `json:"loginPassword,omitempty"`

	MacAddressList []string `json:"macAddressList,omitempty"`

	// MaxDevices
	// Maximum number of allowed device
	MaxDevices *int `json:"maxDevices,omitempty"`

	// ServiceProviderId
	// Service Provider Id
	ServiceProviderId *string `json:"serviceProviderId,omitempty"`

	// SessionUnit
	// Constraints:
	//    - oneof:[MIN,HOUR,DAY,WEEK]
	SessionUnit *string `json:"sessionUnit,omitempty" validate:"oneof=MIN HOUR DAY WEEK"`

	SessionValue *int `json:"sessionValue,omitempty"`

	// UserKey
	// user ID of the identity guest pass
	UserKey *string `json:"userKey,omitempty"`

	// UserName
	// Username of this guest pass
	UserName *string `json:"userName,omitempty"`

	// Wlan
	// WLAN Id
	Wlan *string `json:"wlan,omitempty"`

	// WlanName
	// WLAN Name
	WlanName *string `json:"wlanName,omitempty"`
}

func NewWSGIdentityUserConfigurationCredentialsGuestPassDtoType() *WSGIdentityUserConfigurationCredentialsGuestPassDtoType {
	m := new(WSGIdentityUserConfigurationCredentialsGuestPassDtoType)
	return m
}

type WSGIdentityUserList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIdentityUserSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIdentityUserList() *WSGIdentityUserList {
	m := new(WSGIdentityUserList)
	return m
}

type WSGIdentityUsernamePasswordCredentialsImplDto struct {
	// AuthenticationMethod
	// Authentication Method
	// Constraints:
	//    - oneof:[USERNAME_PASSWORD,GUEST_PASS,MAC_WLAN_DPSK,MO,REMOTE,OAUTH2]
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" validate:"oneof=USERNAME_PASSWORD GUEST_PASS MAC_WLAN_DPSK MO REMOTE OAUTH2"`

	AuthenticationServerConfig *WSGIdentityAuthenticationServerConfig `json:"authenticationServerConfig,omitempty"`

	// CreationDate
	// Creation Date
	CreationDate *int `json:"creationDate,omitempty"`

	// ExpirationDate
	// Expiration Date
	ExpirationDate *int `json:"expirationDate,omitempty"`

	// Key
	// identifier of the UsernamePasswordCredentialsImplDto
	Key *string `json:"key,omitempty"`

	// LoginName
	// Login Name
	LoginName *string `json:"loginName,omitempty"`

	// LoginPassword
	// Login Password
	LoginPassword *string `json:"loginPassword,omitempty"`

	// PasswordCreation
	// Creation Date of Password
	PasswordCreation *string `json:"passwordCreation,omitempty"`

	// PasswordExpiration
	// Expiration Date of Password
	PasswordExpiration *string `json:"passwordExpiration,omitempty"`

	// ServiceProviderId
	// Service Provider Id
	ServiceProviderId *string `json:"serviceProviderId,omitempty"`
}

func NewWSGIdentityUsernamePasswordCredentialsImplDto() *WSGIdentityUsernamePasswordCredentialsImplDto {
	m := new(WSGIdentityUsernamePasswordCredentialsImplDto)
	return m
}