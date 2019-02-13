// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/secret/secret_syncer.go

// Package mock_secret is a generated GoMock package.
package mock_secret

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/supergloo/pkg2/api/external/istio/encryption/v1"
	v10 "github.com/solo-io/supergloo/pkg2/api/v1"
)

// MockSecretSyncer is a mock of SecretSyncer interface
type MockSecretSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockSecretSyncerMockRecorder
}

// MockSecretSyncerMockRecorder is the mock recorder for MockSecretSyncer
type MockSecretSyncerMockRecorder struct {
	mock *MockSecretSyncer
}

// NewMockSecretSyncer creates a new mock instance
func NewMockSecretSyncer(ctrl *gomock.Controller) *MockSecretSyncer {
	mock := &MockSecretSyncer{ctrl: ctrl}
	mock.recorder = &MockSecretSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretSyncer) EXPECT() *MockSecretSyncerMockRecorder {
	return m.recorder
}

// SyncSecret mocks base method
func (m *MockSecretSyncer) SyncSecret(ctx context.Context, installNamespace string, encryption *v10.Encryption, secretList v1.IstioCacertsSecretList, preinstall bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncSecret", ctx, installNamespace, encryption, secretList, preinstall)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncSecret indicates an expected call of SyncSecret
func (mr *MockSecretSyncerMockRecorder) SyncSecret(ctx, installNamespace, encryption, secretList, preinstall interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncSecret", reflect.TypeOf((*MockSecretSyncer)(nil).SyncSecret), ctx, installNamespace, encryption, secretList, preinstall)
}
