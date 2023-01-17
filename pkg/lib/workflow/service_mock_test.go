// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package workflow is a generated GoMock package.
package workflow

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(session *Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), session)
}

// CreateWorkflow mocks base method.
func (m *MockStore) CreateWorkflow(workflow *Workflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkflow", workflow)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWorkflow indicates an expected call of CreateWorkflow.
func (mr *MockStoreMockRecorder) CreateWorkflow(workflow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkflow", reflect.TypeOf((*MockStore)(nil).CreateWorkflow), workflow)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(session *Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), session)
}

// DeleteWorkflow mocks base method.
func (m *MockStore) DeleteWorkflow(workflow *Workflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflow", workflow)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflow indicates an expected call of DeleteWorkflow.
func (mr *MockStoreMockRecorder) DeleteWorkflow(workflow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflow", reflect.TypeOf((*MockStore)(nil).DeleteWorkflow), workflow)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(workflowID string) (*Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", workflowID)
	ret0, _ := ret[0].(*Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(workflowID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), workflowID)
}

// GetWorkflowByInstanceID mocks base method.
func (m *MockStore) GetWorkflowByInstanceID(instanceID string) (*Workflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowByInstanceID", instanceID)
	ret0, _ := ret[0].(*Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowByInstanceID indicates an expected call of GetWorkflowByInstanceID.
func (mr *MockStoreMockRecorder) GetWorkflowByInstanceID(instanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowByInstanceID", reflect.TypeOf((*MockStore)(nil).GetWorkflowByInstanceID), instanceID)
}

// MockSavepoint is a mock of Savepoint interface.
type MockSavepoint struct {
	ctrl     *gomock.Controller
	recorder *MockSavepointMockRecorder
}

// MockSavepointMockRecorder is the mock recorder for MockSavepoint.
type MockSavepointMockRecorder struct {
	mock *MockSavepoint
}

// NewMockSavepoint creates a new mock instance.
func NewMockSavepoint(ctrl *gomock.Controller) *MockSavepoint {
	mock := &MockSavepoint{ctrl: ctrl}
	mock.recorder = &MockSavepointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSavepoint) EXPECT() *MockSavepointMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockSavepoint) Begin() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockSavepointMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockSavepoint)(nil).Begin))
}

// Commit mocks base method.
func (m *MockSavepoint) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockSavepointMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockSavepoint)(nil).Commit))
}

// Rollback mocks base method.
func (m *MockSavepoint) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockSavepointMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockSavepoint)(nil).Rollback))
}