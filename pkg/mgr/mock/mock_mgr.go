// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/mgr/interface.go

// Package mock_mgr is a generated GoMock package.
package mock_mgr

import (
	gomock "github.com/golang/mock/gomock"
	pipeline "github.com/packagrio/go-common/pipeline"
	config "github.com/packagrio/releasr/pkg/config"
	http "net/http"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockInterface) Init(pipelineData *pipeline.Data, myconfig config.Interface, client *http.Client) error {
	ret := m.ctrl.Call(m, "Init", pipelineData, myconfig, client)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockInterfaceMockRecorder) Init(pipelineData, myconfig, client interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockInterface)(nil).Init), pipelineData, myconfig, client)
}

// MgrValidateTools mocks base method
func (m *MockInterface) MgrValidateTools() error {
	ret := m.ctrl.Call(m, "MgrValidateTools")
	ret0, _ := ret[0].(error)
	return ret0
}

// MgrValidateTools indicates an expected call of MgrValidateTools
func (mr *MockInterfaceMockRecorder) MgrValidateTools() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MgrValidateTools", reflect.TypeOf((*MockInterface)(nil).MgrValidateTools))
}

// MgrAssembleStep mocks base method
func (m *MockInterface) MgrAssembleStep() error {
	ret := m.ctrl.Call(m, "MgrAssembleStep")
	ret0, _ := ret[0].(error)
	return ret0
}

// MgrAssembleStep indicates an expected call of MgrAssembleStep
func (mr *MockInterfaceMockRecorder) MgrAssembleStep() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MgrAssembleStep", reflect.TypeOf((*MockInterface)(nil).MgrAssembleStep))
}

// MgrDependenciesStep mocks base method
func (m *MockInterface) MgrDependenciesStep() error {
	ret := m.ctrl.Call(m, "MgrDependenciesStep")
	ret0, _ := ret[0].(error)
	return ret0
}

// MgrDependenciesStep indicates an expected call of MgrDependenciesStep
func (mr *MockInterfaceMockRecorder) MgrDependenciesStep() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MgrDependenciesStep", reflect.TypeOf((*MockInterface)(nil).MgrDependenciesStep))
}

// MgrPackageStep mocks base method
func (m *MockInterface) MgrPackageStep(currentMetadata, nextMetadata interface{}) error {
	ret := m.ctrl.Call(m, "MgrPackageStep", currentMetadata, nextMetadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// MgrPackageStep indicates an expected call of MgrPackageStep
func (mr *MockInterfaceMockRecorder) MgrPackageStep(currentMetadata, nextMetadata interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MgrPackageStep", reflect.TypeOf((*MockInterface)(nil).MgrPackageStep), currentMetadata, nextMetadata)
}

// MgrDistStep mocks base method
func (m *MockInterface) MgrDistStep(currentMetadata, nextMetadata interface{}) error {
	ret := m.ctrl.Call(m, "MgrDistStep", currentMetadata, nextMetadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// MgrDistStep indicates an expected call of MgrDistStep
func (mr *MockInterfaceMockRecorder) MgrDistStep(currentMetadata, nextMetadata interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MgrDistStep", reflect.TypeOf((*MockInterface)(nil).MgrDistStep), currentMetadata, nextMetadata)
}