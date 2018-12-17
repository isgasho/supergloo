// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kube/namespace.go

// Package mock_kube is a generated GoMock package.
package mock_kube

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNamespaceClient is a mock of NamespaceClient interface
type MockNamespaceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceClientMockRecorder
}

// MockNamespaceClientMockRecorder is the mock recorder for MockNamespaceClient
type MockNamespaceClientMockRecorder struct {
	mock *MockNamespaceClient
}

// NewMockNamespaceClient creates a new mock instance
func NewMockNamespaceClient(ctrl *gomock.Controller) *MockNamespaceClient {
	mock := &MockNamespaceClient{ctrl: ctrl}
	mock.recorder = &MockNamespaceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceClient) EXPECT() *MockNamespaceClientMockRecorder {
	return m.recorder
}

// CreateNamespaceIfNotExist mocks base method
func (m *MockNamespaceClient) CreateNamespaceIfNotExist(namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespaceIfNotExist", namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespaceIfNotExist indicates an expected call of CreateNamespaceIfNotExist
func (mr *MockNamespaceClientMockRecorder) CreateNamespaceIfNotExist(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespaceIfNotExist", reflect.TypeOf((*MockNamespaceClient)(nil).CreateNamespaceIfNotExist), namespace)
}

// TryDeleteInstallNamespace mocks base method
func (m *MockNamespaceClient) TryDeleteInstallNamespace(namespaceName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TryDeleteInstallNamespace", namespaceName)
}

// TryDeleteInstallNamespace indicates an expected call of TryDeleteInstallNamespace
func (mr *MockNamespaceClientMockRecorder) TryDeleteInstallNamespace(namespaceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryDeleteInstallNamespace", reflect.TypeOf((*MockNamespaceClient)(nil).TryDeleteInstallNamespace), namespaceName)
}
