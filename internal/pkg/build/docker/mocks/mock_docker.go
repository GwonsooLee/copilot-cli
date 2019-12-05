// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/build/docker/docker.go

// Package mocks is a generated GoMock package.
package mocks

import (
	command "github.com/aws/amazon-ecs-cli-v2/internal/pkg/term/command"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockcommandService is a mock of commandService interface
type MockcommandService struct {
	ctrl     *gomock.Controller
	recorder *MockcommandServiceMockRecorder
}

// MockcommandServiceMockRecorder is the mock recorder for MockcommandService
type MockcommandServiceMockRecorder struct {
	mock *MockcommandService
}

// NewMockcommandService creates a new mock instance
func NewMockcommandService(ctrl *gomock.Controller) *MockcommandService {
	mock := &MockcommandService{ctrl: ctrl}
	mock.recorder = &MockcommandServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcommandService) EXPECT() *MockcommandServiceMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockcommandService) Run(name string, args []string, options ...command.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, args}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockcommandServiceMockRecorder) Run(name, args interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, args}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockcommandService)(nil).Run), varargs...)
}
