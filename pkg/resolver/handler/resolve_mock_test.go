// Code generated by MockGen. DO NOT EDIT.
// Source: resolve.go

// Package handler is a generated GoMock package.
package handler

import (
	reflect "reflect"

	model "github.com/authgear/authgear-server/pkg/api/model"
	identity "github.com/authgear/authgear-server/pkg/lib/authn/identity"
	accesscontrol "github.com/authgear/authgear-server/pkg/util/accesscontrol"
	gomock "github.com/golang/mock/gomock"
)

// MockIdentityService is a mock of IdentityService interface.
type MockIdentityService struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityServiceMockRecorder
}

// MockIdentityServiceMockRecorder is the mock recorder for MockIdentityService.
type MockIdentityServiceMockRecorder struct {
	mock *MockIdentityService
}

// NewMockIdentityService creates a new mock instance.
func NewMockIdentityService(ctrl *gomock.Controller) *MockIdentityService {
	mock := &MockIdentityService{ctrl: ctrl}
	mock.recorder = &MockIdentityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityService) EXPECT() *MockIdentityServiceMockRecorder {
	return m.recorder
}

// ListByUser mocks base method.
func (m *MockIdentityService) ListByUser(userID string) ([]*identity.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUser", userID)
	ret0, _ := ret[0].([]*identity.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUser indicates an expected call of ListByUser.
func (mr *MockIdentityServiceMockRecorder) ListByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUser", reflect.TypeOf((*MockIdentityService)(nil).ListByUser), userID)
}

// MockVerificationService is a mock of VerificationService interface.
type MockVerificationService struct {
	ctrl     *gomock.Controller
	recorder *MockVerificationServiceMockRecorder
}

// MockVerificationServiceMockRecorder is the mock recorder for MockVerificationService.
type MockVerificationServiceMockRecorder struct {
	mock *MockVerificationService
}

// NewMockVerificationService creates a new mock instance.
func NewMockVerificationService(ctrl *gomock.Controller) *MockVerificationService {
	mock := &MockVerificationService{ctrl: ctrl}
	mock.recorder = &MockVerificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerificationService) EXPECT() *MockVerificationServiceMockRecorder {
	return m.recorder
}

// IsUserVerified mocks base method.
func (m *MockVerificationService) IsUserVerified(identities []*identity.Info) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserVerified", identities)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserVerified indicates an expected call of IsUserVerified.
func (mr *MockVerificationServiceMockRecorder) IsUserVerified(identities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserVerified", reflect.TypeOf((*MockVerificationService)(nil).IsUserVerified), identities)
}

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// ReadOnly mocks base method.
func (m *MockDatabase) ReadOnly(arg0 func() error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnly", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadOnly indicates an expected call of ReadOnly.
func (mr *MockDatabaseMockRecorder) ReadOnly(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnly", reflect.TypeOf((*MockDatabase)(nil).ReadOnly), arg0)
}

// MockUserProvider is a mock of UserProvider interface.
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider.
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance.
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserProvider) Get(id string, role accesscontrol.Role) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id, role)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserProviderMockRecorder) Get(id, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserProvider)(nil).Get), id, role)
}

// MockRolesAndGroupsProvider is a mock of RolesAndGroupsProvider interface.
type MockRolesAndGroupsProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRolesAndGroupsProviderMockRecorder
}

// MockRolesAndGroupsProviderMockRecorder is the mock recorder for MockRolesAndGroupsProvider.
type MockRolesAndGroupsProviderMockRecorder struct {
	mock *MockRolesAndGroupsProvider
}

// NewMockRolesAndGroupsProvider creates a new mock instance.
func NewMockRolesAndGroupsProvider(ctrl *gomock.Controller) *MockRolesAndGroupsProvider {
	mock := &MockRolesAndGroupsProvider{ctrl: ctrl}
	mock.recorder = &MockRolesAndGroupsProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRolesAndGroupsProvider) EXPECT() *MockRolesAndGroupsProviderMockRecorder {
	return m.recorder
}

// ListComputedRolesByUserID mocks base method.
func (m *MockRolesAndGroupsProvider) ListComputedRolesByUserID(userID string) ([]*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComputedRolesByUserID", userID)
	ret0, _ := ret[0].([]*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComputedRolesByUserID indicates an expected call of ListComputedRolesByUserID.
func (mr *MockRolesAndGroupsProviderMockRecorder) ListComputedRolesByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComputedRolesByUserID", reflect.TypeOf((*MockRolesAndGroupsProvider)(nil).ListComputedRolesByUserID), userID)
}
