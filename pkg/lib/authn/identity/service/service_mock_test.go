// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	authenticator "github.com/authgear/authgear-server/pkg/lib/authn/authenticator"
	anonymous "github.com/authgear/authgear-server/pkg/lib/authn/identity/anonymous"
	loginid "github.com/authgear/authgear-server/pkg/lib/authn/identity/loginid"
	oauth "github.com/authgear/authgear-server/pkg/lib/authn/identity/oauth"
	config "github.com/authgear/authgear-server/pkg/lib/config"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLoginIDIdentityProvider is a mock of LoginIDIdentityProvider interface
type MockLoginIDIdentityProvider struct {
	ctrl     *gomock.Controller
	recorder *MockLoginIDIdentityProviderMockRecorder
}

// MockLoginIDIdentityProviderMockRecorder is the mock recorder for MockLoginIDIdentityProvider
type MockLoginIDIdentityProviderMockRecorder struct {
	mock *MockLoginIDIdentityProvider
}

// NewMockLoginIDIdentityProvider creates a new mock instance
func NewMockLoginIDIdentityProvider(ctrl *gomock.Controller) *MockLoginIDIdentityProvider {
	mock := &MockLoginIDIdentityProvider{ctrl: ctrl}
	mock.recorder = &MockLoginIDIdentityProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoginIDIdentityProvider) EXPECT() *MockLoginIDIdentityProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockLoginIDIdentityProvider) Get(userID, id string) (*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID, id)
	ret0, _ := ret[0].(*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockLoginIDIdentityProviderMockRecorder) Get(userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).Get), userID, id)
}

// GetMany mocks base method
func (m *MockLoginIDIdentityProvider) GetMany(ids []string) ([]*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ids)
	ret0, _ := ret[0].([]*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany
func (mr *MockLoginIDIdentityProviderMockRecorder) GetMany(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).GetMany), ids)
}

// List mocks base method
func (m *MockLoginIDIdentityProvider) List(userID string) ([]*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userID)
	ret0, _ := ret[0].([]*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockLoginIDIdentityProviderMockRecorder) List(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).List), userID)
}

// GetByValue mocks base method
func (m *MockLoginIDIdentityProvider) GetByValue(loginIDValue string) ([]*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByValue", loginIDValue)
	ret0, _ := ret[0].([]*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByValue indicates an expected call of GetByValue
func (mr *MockLoginIDIdentityProviderMockRecorder) GetByValue(loginIDValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByValue", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).GetByValue), loginIDValue)
}

// ListByClaim mocks base method
func (m *MockLoginIDIdentityProvider) ListByClaim(name, value string) ([]*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByClaim", name, value)
	ret0, _ := ret[0].([]*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByClaim indicates an expected call of ListByClaim
func (mr *MockLoginIDIdentityProviderMockRecorder) ListByClaim(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByClaim", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).ListByClaim), name, value)
}

// New mocks base method
func (m *MockLoginIDIdentityProvider) New(userID string, loginID loginid.Spec) (*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", userID, loginID)
	ret0, _ := ret[0].(*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockLoginIDIdentityProviderMockRecorder) New(userID, loginID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).New), userID, loginID)
}

// WithValue mocks base method
func (m *MockLoginIDIdentityProvider) WithValue(iden *loginid.Identity, value string) (*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithValue", iden, value)
	ret0, _ := ret[0].(*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithValue indicates an expected call of WithValue
func (mr *MockLoginIDIdentityProviderMockRecorder) WithValue(iden, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValue", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).WithValue), iden, value)
}

// Create mocks base method
func (m *MockLoginIDIdentityProvider) Create(i *loginid.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockLoginIDIdentityProviderMockRecorder) Create(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).Create), i)
}

// Update mocks base method
func (m *MockLoginIDIdentityProvider) Update(i *loginid.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockLoginIDIdentityProviderMockRecorder) Update(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).Update), i)
}

// Delete mocks base method
func (m *MockLoginIDIdentityProvider) Delete(i *loginid.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockLoginIDIdentityProviderMockRecorder) Delete(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).Delete), i)
}

// CheckDuplicated mocks base method
func (m *MockLoginIDIdentityProvider) CheckDuplicated(uniqueKey string, standardClaims map[string]string, userID string) (*loginid.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDuplicated", uniqueKey, standardClaims, userID)
	ret0, _ := ret[0].(*loginid.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDuplicated indicates an expected call of CheckDuplicated
func (mr *MockLoginIDIdentityProviderMockRecorder) CheckDuplicated(uniqueKey, standardClaims, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDuplicated", reflect.TypeOf((*MockLoginIDIdentityProvider)(nil).CheckDuplicated), uniqueKey, standardClaims, userID)
}

// MockOAuthIdentityProvider is a mock of OAuthIdentityProvider interface
type MockOAuthIdentityProvider struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthIdentityProviderMockRecorder
}

// MockOAuthIdentityProviderMockRecorder is the mock recorder for MockOAuthIdentityProvider
type MockOAuthIdentityProviderMockRecorder struct {
	mock *MockOAuthIdentityProvider
}

// NewMockOAuthIdentityProvider creates a new mock instance
func NewMockOAuthIdentityProvider(ctrl *gomock.Controller) *MockOAuthIdentityProvider {
	mock := &MockOAuthIdentityProvider{ctrl: ctrl}
	mock.recorder = &MockOAuthIdentityProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOAuthIdentityProvider) EXPECT() *MockOAuthIdentityProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockOAuthIdentityProvider) Get(userID, id string) (*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID, id)
	ret0, _ := ret[0].(*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockOAuthIdentityProviderMockRecorder) Get(userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).Get), userID, id)
}

// GetMany mocks base method
func (m *MockOAuthIdentityProvider) GetMany(ids []string) ([]*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ids)
	ret0, _ := ret[0].([]*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany
func (mr *MockOAuthIdentityProviderMockRecorder) GetMany(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).GetMany), ids)
}

// List mocks base method
func (m *MockOAuthIdentityProvider) List(userID string) ([]*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userID)
	ret0, _ := ret[0].([]*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockOAuthIdentityProviderMockRecorder) List(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).List), userID)
}

// GetByProviderSubject mocks base method
func (m *MockOAuthIdentityProvider) GetByProviderSubject(provider config.ProviderID, subjectID string) (*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByProviderSubject", provider, subjectID)
	ret0, _ := ret[0].(*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByProviderSubject indicates an expected call of GetByProviderSubject
func (mr *MockOAuthIdentityProviderMockRecorder) GetByProviderSubject(provider, subjectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByProviderSubject", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).GetByProviderSubject), provider, subjectID)
}

// GetByUserProvider mocks base method
func (m *MockOAuthIdentityProvider) GetByUserProvider(userID string, provider config.ProviderID) (*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserProvider", userID, provider)
	ret0, _ := ret[0].(*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserProvider indicates an expected call of GetByUserProvider
func (mr *MockOAuthIdentityProviderMockRecorder) GetByUserProvider(userID, provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserProvider", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).GetByUserProvider), userID, provider)
}

// ListByClaim mocks base method
func (m *MockOAuthIdentityProvider) ListByClaim(name, value string) ([]*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByClaim", name, value)
	ret0, _ := ret[0].([]*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByClaim indicates an expected call of ListByClaim
func (mr *MockOAuthIdentityProviderMockRecorder) ListByClaim(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByClaim", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).ListByClaim), name, value)
}

// New mocks base method
func (m *MockOAuthIdentityProvider) New(userID string, provider config.ProviderID, subjectID string, profile, claims map[string]interface{}) *oauth.Identity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", userID, provider, subjectID, profile, claims)
	ret0, _ := ret[0].(*oauth.Identity)
	return ret0
}

// New indicates an expected call of New
func (mr *MockOAuthIdentityProviderMockRecorder) New(userID, provider, subjectID, profile, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).New), userID, provider, subjectID, profile, claims)
}

// Create mocks base method
func (m *MockOAuthIdentityProvider) Create(i *oauth.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockOAuthIdentityProviderMockRecorder) Create(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).Create), i)
}

// Update mocks base method
func (m *MockOAuthIdentityProvider) Update(i *oauth.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockOAuthIdentityProviderMockRecorder) Update(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).Update), i)
}

// Delete mocks base method
func (m *MockOAuthIdentityProvider) Delete(i *oauth.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockOAuthIdentityProviderMockRecorder) Delete(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).Delete), i)
}

// CheckDuplicated mocks base method
func (m *MockOAuthIdentityProvider) CheckDuplicated(standardClaims map[string]string, userID string) (*oauth.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDuplicated", standardClaims, userID)
	ret0, _ := ret[0].(*oauth.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDuplicated indicates an expected call of CheckDuplicated
func (mr *MockOAuthIdentityProviderMockRecorder) CheckDuplicated(standardClaims, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDuplicated", reflect.TypeOf((*MockOAuthIdentityProvider)(nil).CheckDuplicated), standardClaims, userID)
}

// MockAnonymousIdentityProvider is a mock of AnonymousIdentityProvider interface
type MockAnonymousIdentityProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAnonymousIdentityProviderMockRecorder
}

// MockAnonymousIdentityProviderMockRecorder is the mock recorder for MockAnonymousIdentityProvider
type MockAnonymousIdentityProviderMockRecorder struct {
	mock *MockAnonymousIdentityProvider
}

// NewMockAnonymousIdentityProvider creates a new mock instance
func NewMockAnonymousIdentityProvider(ctrl *gomock.Controller) *MockAnonymousIdentityProvider {
	mock := &MockAnonymousIdentityProvider{ctrl: ctrl}
	mock.recorder = &MockAnonymousIdentityProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnonymousIdentityProvider) EXPECT() *MockAnonymousIdentityProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAnonymousIdentityProvider) Get(userID, id string) (*anonymous.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID, id)
	ret0, _ := ret[0].(*anonymous.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAnonymousIdentityProviderMockRecorder) Get(userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).Get), userID, id)
}

// GetMany mocks base method
func (m *MockAnonymousIdentityProvider) GetMany(ids []string) ([]*anonymous.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ids)
	ret0, _ := ret[0].([]*anonymous.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany
func (mr *MockAnonymousIdentityProviderMockRecorder) GetMany(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).GetMany), ids)
}

// GetByKeyID mocks base method
func (m *MockAnonymousIdentityProvider) GetByKeyID(keyID string) (*anonymous.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKeyID", keyID)
	ret0, _ := ret[0].(*anonymous.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKeyID indicates an expected call of GetByKeyID
func (mr *MockAnonymousIdentityProviderMockRecorder) GetByKeyID(keyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKeyID", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).GetByKeyID), keyID)
}

// List mocks base method
func (m *MockAnonymousIdentityProvider) List(userID string) ([]*anonymous.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userID)
	ret0, _ := ret[0].([]*anonymous.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockAnonymousIdentityProviderMockRecorder) List(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).List), userID)
}

// ListByClaim mocks base method
func (m *MockAnonymousIdentityProvider) ListByClaim(name, value string) ([]*anonymous.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByClaim", name, value)
	ret0, _ := ret[0].([]*anonymous.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByClaim indicates an expected call of ListByClaim
func (mr *MockAnonymousIdentityProviderMockRecorder) ListByClaim(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByClaim", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).ListByClaim), name, value)
}

// New mocks base method
func (m *MockAnonymousIdentityProvider) New(userID, keyID string, key []byte) *anonymous.Identity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", userID, keyID, key)
	ret0, _ := ret[0].(*anonymous.Identity)
	return ret0
}

// New indicates an expected call of New
func (mr *MockAnonymousIdentityProviderMockRecorder) New(userID, keyID, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).New), userID, keyID, key)
}

// Create mocks base method
func (m *MockAnonymousIdentityProvider) Create(i *anonymous.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockAnonymousIdentityProviderMockRecorder) Create(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).Create), i)
}

// Delete mocks base method
func (m *MockAnonymousIdentityProvider) Delete(i *anonymous.Identity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAnonymousIdentityProviderMockRecorder) Delete(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAnonymousIdentityProvider)(nil).Delete), i)
}

// MockAuthenticatorService is a mock of AuthenticatorService interface
type MockAuthenticatorService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorServiceMockRecorder
}

// MockAuthenticatorServiceMockRecorder is the mock recorder for MockAuthenticatorService
type MockAuthenticatorServiceMockRecorder struct {
	mock *MockAuthenticatorService
}

// NewMockAuthenticatorService creates a new mock instance
func NewMockAuthenticatorService(ctrl *gomock.Controller) *MockAuthenticatorService {
	mock := &MockAuthenticatorService{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticatorService) EXPECT() *MockAuthenticatorServiceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockAuthenticatorService) List(userID string, filters ...authenticator.Filter) ([]*authenticator.Info, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{userID}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*authenticator.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockAuthenticatorServiceMockRecorder) List(userID interface{}, filters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{userID}, filters...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthenticatorService)(nil).List), varargs...)
}

// Delete mocks base method
func (m *MockAuthenticatorService) Delete(authenticatorInfo *authenticator.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", authenticatorInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAuthenticatorServiceMockRecorder) Delete(authenticatorInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthenticatorService)(nil).Delete), authenticatorInfo)
}
