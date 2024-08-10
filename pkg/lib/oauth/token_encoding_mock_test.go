// Code generated by MockGen. DO NOT EDIT.
// Source: token_encoding.go

// Package oauth is a generated GoMock package.
package oauth

import (
	url "net/url"
	reflect "reflect"

	event "github.com/authgear/authgear-server/pkg/api/event"
	gomock "github.com/golang/mock/gomock"
	jwt "github.com/lestrrat-go/jwx/v2/jwt"
)

// MockIDTokenIssuer is a mock of IDTokenIssuer interface.
type MockIDTokenIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenIssuerMockRecorder
}

// MockIDTokenIssuerMockRecorder is the mock recorder for MockIDTokenIssuer.
type MockIDTokenIssuerMockRecorder struct {
	mock *MockIDTokenIssuer
}

// NewMockIDTokenIssuer creates a new mock instance.
func NewMockIDTokenIssuer(ctrl *gomock.Controller) *MockIDTokenIssuer {
	mock := &MockIDTokenIssuer{ctrl: ctrl}
	mock.recorder = &MockIDTokenIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDTokenIssuer) EXPECT() *MockIDTokenIssuerMockRecorder {
	return m.recorder
}

// Iss mocks base method.
func (m *MockIDTokenIssuer) Iss() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iss")
	ret0, _ := ret[0].(string)
	return ret0
}

// Iss indicates an expected call of Iss.
func (mr *MockIDTokenIssuerMockRecorder) Iss() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iss", reflect.TypeOf((*MockIDTokenIssuer)(nil).Iss))
}

// PopulateUserClaimsInIDToken mocks base method.
func (m *MockIDTokenIssuer) PopulateUserClaimsInIDToken(token jwt.Token, userID string, clientLike *ClientLike) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopulateUserClaimsInIDToken", token, userID, clientLike)
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateUserClaimsInIDToken indicates an expected call of PopulateUserClaimsInIDToken.
func (mr *MockIDTokenIssuerMockRecorder) PopulateUserClaimsInIDToken(token, userID, clientLike interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateUserClaimsInIDToken", reflect.TypeOf((*MockIDTokenIssuer)(nil).PopulateUserClaimsInIDToken), token, userID, clientLike)
}

// MockBaseURLProvider is a mock of BaseURLProvider interface.
type MockBaseURLProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBaseURLProviderMockRecorder
}

// MockBaseURLProviderMockRecorder is the mock recorder for MockBaseURLProvider.
type MockBaseURLProviderMockRecorder struct {
	mock *MockBaseURLProvider
}

// NewMockBaseURLProvider creates a new mock instance.
func NewMockBaseURLProvider(ctrl *gomock.Controller) *MockBaseURLProvider {
	mock := &MockBaseURLProvider{ctrl: ctrl}
	mock.recorder = &MockBaseURLProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseURLProvider) EXPECT() *MockBaseURLProviderMockRecorder {
	return m.recorder
}

// Origin mocks base method.
func (m *MockBaseURLProvider) Origin() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Origin")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// Origin indicates an expected call of Origin.
func (mr *MockBaseURLProviderMockRecorder) Origin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Origin", reflect.TypeOf((*MockBaseURLProvider)(nil).Origin))
}

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// DispatchEventOnCommit mocks base method.
func (m *MockEventService) DispatchEventOnCommit(payload event.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchEventOnCommit", payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchEventOnCommit indicates an expected call of DispatchEventOnCommit.
func (mr *MockEventServiceMockRecorder) DispatchEventOnCommit(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchEventOnCommit", reflect.TypeOf((*MockEventService)(nil).DispatchEventOnCommit), payload)
}
