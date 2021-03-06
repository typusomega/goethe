// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typusomega/goethe/pkg/storage (interfaces: Metrics)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	spec "github.com/typusomega/goethe/pkg/spec"
	storage "github.com/typusomega/goethe/pkg/storage"
)

// MockStorageMetrics is a mock of Metrics interface
type MockStorageMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMetricsMockRecorder
}

// MockStorageMetricsMockRecorder is the mock recorder for MockStorageMetrics
type MockStorageMetricsMockRecorder struct {
	mock *MockStorageMetrics
}

// NewMockStorageMetrics creates a new mock instance
func NewMockStorageMetrics(ctrl *gomock.Controller) *MockStorageMetrics {
	mock := &MockStorageMetrics{ctrl: ctrl}
	mock.recorder = &MockStorageMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageMetrics) EXPECT() *MockStorageMetricsMockRecorder {
	return m.recorder
}

// MeasureGetIterator mocks base method
func (m *MockStorageMetrics) MeasureGetIterator(arg0 func() (storage.EventsIterator, error)) (storage.EventsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeasureGetIterator", arg0)
	ret0, _ := ret[0].(storage.EventsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MeasureGetIterator indicates an expected call of MeasureGetIterator
func (mr *MockStorageMetricsMockRecorder) MeasureGetIterator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeasureGetIterator", reflect.TypeOf((*MockStorageMetrics)(nil).MeasureGetIterator), arg0)
}

// MeasurePersistEvent mocks base method
func (m *MockStorageMetrics) MeasurePersistEvent(arg0 func() (*spec.Event, error)) (*spec.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeasurePersistEvent", arg0)
	ret0, _ := ret[0].(*spec.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MeasurePersistEvent indicates an expected call of MeasurePersistEvent
func (mr *MockStorageMetricsMockRecorder) MeasurePersistEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeasurePersistEvent", reflect.TypeOf((*MockStorageMetrics)(nil).MeasurePersistEvent), arg0)
}
