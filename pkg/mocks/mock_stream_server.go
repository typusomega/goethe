// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typusomega/goethe/pkg/spec (interfaces: Goethe_StreamServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	spec "github.com/typusomega/goethe/pkg/spec"
	metadata "google.golang.org/grpc/metadata"
)

// MockGoethe_StreamServer is a mock of Goethe_StreamServer interface
type MockGoethe_StreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockGoethe_StreamServerMockRecorder
}

// MockGoethe_StreamServerMockRecorder is the mock recorder for MockGoethe_StreamServer
type MockGoethe_StreamServerMockRecorder struct {
	mock *MockGoethe_StreamServer
}

// NewMockGoethe_StreamServer creates a new mock instance
func NewMockGoethe_StreamServer(ctrl *gomock.Controller) *MockGoethe_StreamServer {
	mock := &MockGoethe_StreamServer{ctrl: ctrl}
	mock.recorder = &MockGoethe_StreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGoethe_StreamServer) EXPECT() *MockGoethe_StreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockGoethe_StreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockGoethe_StreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockGoethe_StreamServer)(nil).Context))
}

// Recv mocks base method
func (m *MockGoethe_StreamServer) Recv() (*spec.Cursor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*spec.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockGoethe_StreamServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockGoethe_StreamServer)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockGoethe_StreamServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockGoethe_StreamServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockGoethe_StreamServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockGoethe_StreamServer) Send(arg0 *spec.Cursor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockGoethe_StreamServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockGoethe_StreamServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockGoethe_StreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockGoethe_StreamServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockGoethe_StreamServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockGoethe_StreamServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockGoethe_StreamServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockGoethe_StreamServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockGoethe_StreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockGoethe_StreamServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockGoethe_StreamServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockGoethe_StreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockGoethe_StreamServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockGoethe_StreamServer)(nil).SetTrailer), arg0)
}