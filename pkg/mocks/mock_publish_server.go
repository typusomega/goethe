// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typusomega/goethe/pkg/spec (interfaces: Goethe_PublishServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	spec "github.com/typusomega/goethe/pkg/spec"
	metadata "google.golang.org/grpc/metadata"
)

// MockGoethe_PublishServer is a mock of Goethe_PublishServer interface
type MockGoethe_PublishServer struct {
	ctrl     *gomock.Controller
	recorder *MockGoethe_PublishServerMockRecorder
}

// MockGoethe_PublishServerMockRecorder is the mock recorder for MockGoethe_PublishServer
type MockGoethe_PublishServerMockRecorder struct {
	mock *MockGoethe_PublishServer
}

// NewMockGoethe_PublishServer creates a new mock instance
func NewMockGoethe_PublishServer(ctrl *gomock.Controller) *MockGoethe_PublishServer {
	mock := &MockGoethe_PublishServer{ctrl: ctrl}
	mock.recorder = &MockGoethe_PublishServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGoethe_PublishServer) EXPECT() *MockGoethe_PublishServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockGoethe_PublishServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockGoethe_PublishServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockGoethe_PublishServer)(nil).Context))
}

// Recv mocks base method
func (m *MockGoethe_PublishServer) Recv() (*spec.PublishRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*spec.PublishRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockGoethe_PublishServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockGoethe_PublishServer)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockGoethe_PublishServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockGoethe_PublishServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockGoethe_PublishServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockGoethe_PublishServer) Send(arg0 *spec.PublishResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockGoethe_PublishServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockGoethe_PublishServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockGoethe_PublishServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockGoethe_PublishServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockGoethe_PublishServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockGoethe_PublishServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockGoethe_PublishServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockGoethe_PublishServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockGoethe_PublishServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockGoethe_PublishServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockGoethe_PublishServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockGoethe_PublishServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockGoethe_PublishServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockGoethe_PublishServer)(nil).SetTrailer), arg0)
}
