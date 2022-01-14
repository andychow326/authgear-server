// Code generated by MockGen. DO NOT EDIT.
// Source: deliverer.go

// Package hook is a generated GoMock package.
package hook

import (
	reflect "reflect"

	accesscontrol "github.com/authgear/authgear-server/pkg/util/accesscontrol"
	gomock "github.com/golang/mock/gomock"
)

// MockStandardAttributesServiceNoEvent is a mock of StandardAttributesServiceNoEvent interface.
type MockStandardAttributesServiceNoEvent struct {
	ctrl     *gomock.Controller
	recorder *MockStandardAttributesServiceNoEventMockRecorder
}

// MockStandardAttributesServiceNoEventMockRecorder is the mock recorder for MockStandardAttributesServiceNoEvent.
type MockStandardAttributesServiceNoEventMockRecorder struct {
	mock *MockStandardAttributesServiceNoEvent
}

// NewMockStandardAttributesServiceNoEvent creates a new mock instance.
func NewMockStandardAttributesServiceNoEvent(ctrl *gomock.Controller) *MockStandardAttributesServiceNoEvent {
	mock := &MockStandardAttributesServiceNoEvent{ctrl: ctrl}
	mock.recorder = &MockStandardAttributesServiceNoEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStandardAttributesServiceNoEvent) EXPECT() *MockStandardAttributesServiceNoEventMockRecorder {
	return m.recorder
}

// UpdateStandardAttributes mocks base method.
func (m *MockStandardAttributesServiceNoEvent) UpdateStandardAttributes(role accesscontrol.Role, userID string, stdAttrs map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStandardAttributes", role, userID, stdAttrs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStandardAttributes indicates an expected call of UpdateStandardAttributes.
func (mr *MockStandardAttributesServiceNoEventMockRecorder) UpdateStandardAttributes(role, userID, stdAttrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStandardAttributes", reflect.TypeOf((*MockStandardAttributesServiceNoEvent)(nil).UpdateStandardAttributes), role, userID, stdAttrs)
}

// MockCustomAttributesServiceNoEvent is a mock of CustomAttributesServiceNoEvent interface.
type MockCustomAttributesServiceNoEvent struct {
	ctrl     *gomock.Controller
	recorder *MockCustomAttributesServiceNoEventMockRecorder
}

// MockCustomAttributesServiceNoEventMockRecorder is the mock recorder for MockCustomAttributesServiceNoEvent.
type MockCustomAttributesServiceNoEventMockRecorder struct {
	mock *MockCustomAttributesServiceNoEvent
}

// NewMockCustomAttributesServiceNoEvent creates a new mock instance.
func NewMockCustomAttributesServiceNoEvent(ctrl *gomock.Controller) *MockCustomAttributesServiceNoEvent {
	mock := &MockCustomAttributesServiceNoEvent{ctrl: ctrl}
	mock.recorder = &MockCustomAttributesServiceNoEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomAttributesServiceNoEvent) EXPECT() *MockCustomAttributesServiceNoEventMockRecorder {
	return m.recorder
}

// UpdateAllCustomAttributes mocks base method.
func (m *MockCustomAttributesServiceNoEvent) UpdateAllCustomAttributes(role accesscontrol.Role, userID string, reprForm map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllCustomAttributes", role, userID, reprForm)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAllCustomAttributes indicates an expected call of UpdateAllCustomAttributes.
func (mr *MockCustomAttributesServiceNoEventMockRecorder) UpdateAllCustomAttributes(role, userID, reprForm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllCustomAttributes", reflect.TypeOf((*MockCustomAttributesServiceNoEvent)(nil).UpdateAllCustomAttributes), role, userID, reprForm)
}
