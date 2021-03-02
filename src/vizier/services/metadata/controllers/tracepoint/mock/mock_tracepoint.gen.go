// Code generated by MockGen. DO NOT EDIT.
// Source: tracepoint.go

// Package mock_tracepoint is a generated GoMock package.
package mock_tracepoint

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	storepb "pixielabs.ai/pixielabs/src/vizier/services/metadata/storepb"
)

// MockagentMessenger is a mock of agentMessenger interface.
type MockagentMessenger struct {
	ctrl     *gomock.Controller
	recorder *MockagentMessengerMockRecorder
}

// MockagentMessengerMockRecorder is the mock recorder for MockagentMessenger.
type MockagentMessengerMockRecorder struct {
	mock *MockagentMessenger
}

// NewMockagentMessenger creates a new mock instance.
func NewMockagentMessenger(ctrl *gomock.Controller) *MockagentMessenger {
	mock := &MockagentMessenger{ctrl: ctrl}
	mock.recorder = &MockagentMessengerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockagentMessenger) EXPECT() *MockagentMessengerMockRecorder {
	return m.recorder
}

// MessageActiveAgents mocks base method.
func (m *MockagentMessenger) MessageActiveAgents(msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageActiveAgents", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// MessageActiveAgents indicates an expected call of MessageActiveAgents.
func (mr *MockagentMessengerMockRecorder) MessageActiveAgents(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageActiveAgents", reflect.TypeOf((*MockagentMessenger)(nil).MessageActiveAgents), msg)
}

// MessageAgents mocks base method.
func (m *MockagentMessenger) MessageAgents(agentIDs []uuid.UUID, msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageAgents", agentIDs, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// MessageAgents indicates an expected call of MessageAgents.
func (mr *MockagentMessengerMockRecorder) MessageAgents(agentIDs, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageAgents", reflect.TypeOf((*MockagentMessenger)(nil).MessageAgents), agentIDs, msg)
}

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

// DeleteTracepoint mocks base method.
func (m *MockStore) DeleteTracepoint(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTracepoint", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTracepoint indicates an expected call of DeleteTracepoint.
func (mr *MockStoreMockRecorder) DeleteTracepoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTracepoint", reflect.TypeOf((*MockStore)(nil).DeleteTracepoint), arg0)
}

// DeleteTracepointTTLs mocks base method.
func (m *MockStore) DeleteTracepointTTLs(arg0 []uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTracepointTTLs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTracepointTTLs indicates an expected call of DeleteTracepointTTLs.
func (mr *MockStoreMockRecorder) DeleteTracepointTTLs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTracepointTTLs", reflect.TypeOf((*MockStore)(nil).DeleteTracepointTTLs), arg0)
}

// DeleteTracepointsForAgent mocks base method.
func (m *MockStore) DeleteTracepointsForAgent(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTracepointsForAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTracepointsForAgent indicates an expected call of DeleteTracepointsForAgent.
func (mr *MockStoreMockRecorder) DeleteTracepointsForAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTracepointsForAgent", reflect.TypeOf((*MockStore)(nil).DeleteTracepointsForAgent), arg0)
}

// GetTracepoint mocks base method.
func (m *MockStore) GetTracepoint(arg0 uuid.UUID) (*storepb.TracepointInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracepoint", arg0)
	ret0, _ := ret[0].(*storepb.TracepointInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracepoint indicates an expected call of GetTracepoint.
func (mr *MockStoreMockRecorder) GetTracepoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracepoint", reflect.TypeOf((*MockStore)(nil).GetTracepoint), arg0)
}

// GetTracepointStates mocks base method.
func (m *MockStore) GetTracepointStates(arg0 uuid.UUID) ([]*storepb.AgentTracepointStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracepointStates", arg0)
	ret0, _ := ret[0].([]*storepb.AgentTracepointStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracepointStates indicates an expected call of GetTracepointStates.
func (mr *MockStoreMockRecorder) GetTracepointStates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracepointStates", reflect.TypeOf((*MockStore)(nil).GetTracepointStates), arg0)
}

// GetTracepointTTLs mocks base method.
func (m *MockStore) GetTracepointTTLs() ([]uuid.UUID, []time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracepointTTLs")
	ret0, _ := ret[0].([]uuid.UUID)
	ret1, _ := ret[1].([]time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTracepointTTLs indicates an expected call of GetTracepointTTLs.
func (mr *MockStoreMockRecorder) GetTracepointTTLs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracepointTTLs", reflect.TypeOf((*MockStore)(nil).GetTracepointTTLs))
}

// GetTracepoints mocks base method.
func (m *MockStore) GetTracepoints() ([]*storepb.TracepointInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracepoints")
	ret0, _ := ret[0].([]*storepb.TracepointInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracepoints indicates an expected call of GetTracepoints.
func (mr *MockStoreMockRecorder) GetTracepoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracepoints", reflect.TypeOf((*MockStore)(nil).GetTracepoints))
}

// GetTracepointsForIDs mocks base method.
func (m *MockStore) GetTracepointsForIDs(arg0 []uuid.UUID) ([]*storepb.TracepointInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracepointsForIDs", arg0)
	ret0, _ := ret[0].([]*storepb.TracepointInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracepointsForIDs indicates an expected call of GetTracepointsForIDs.
func (mr *MockStoreMockRecorder) GetTracepointsForIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracepointsForIDs", reflect.TypeOf((*MockStore)(nil).GetTracepointsForIDs), arg0)
}

// GetTracepointsWithNames mocks base method.
func (m *MockStore) GetTracepointsWithNames(arg0 []string) ([]*uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracepointsWithNames", arg0)
	ret0, _ := ret[0].([]*uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracepointsWithNames indicates an expected call of GetTracepointsWithNames.
func (mr *MockStoreMockRecorder) GetTracepointsWithNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracepointsWithNames", reflect.TypeOf((*MockStore)(nil).GetTracepointsWithNames), arg0)
}

// SetTracepointTTL mocks base method.
func (m *MockStore) SetTracepointTTL(arg0 uuid.UUID, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTracepointTTL", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTracepointTTL indicates an expected call of SetTracepointTTL.
func (mr *MockStoreMockRecorder) SetTracepointTTL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTracepointTTL", reflect.TypeOf((*MockStore)(nil).SetTracepointTTL), arg0, arg1)
}

// SetTracepointWithName mocks base method.
func (m *MockStore) SetTracepointWithName(arg0 string, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTracepointWithName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTracepointWithName indicates an expected call of SetTracepointWithName.
func (mr *MockStoreMockRecorder) SetTracepointWithName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTracepointWithName", reflect.TypeOf((*MockStore)(nil).SetTracepointWithName), arg0, arg1)
}

// UpdateTracepointState mocks base method.
func (m *MockStore) UpdateTracepointState(arg0 *storepb.AgentTracepointStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTracepointState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTracepointState indicates an expected call of UpdateTracepointState.
func (mr *MockStoreMockRecorder) UpdateTracepointState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTracepointState", reflect.TypeOf((*MockStore)(nil).UpdateTracepointState), arg0)
}

// UpsertTracepoint mocks base method.
func (m *MockStore) UpsertTracepoint(arg0 uuid.UUID, arg1 *storepb.TracepointInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTracepoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertTracepoint indicates an expected call of UpsertTracepoint.
func (mr *MockStoreMockRecorder) UpsertTracepoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTracepoint", reflect.TypeOf((*MockStore)(nil).UpsertTracepoint), arg0, arg1)
}
