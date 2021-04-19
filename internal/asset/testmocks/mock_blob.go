// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/pack/internal/asset (interfaces: Blob)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	dist "github.com/buildpacks/pack/internal/dist"
)

// MockBlob is a mock of Blob interface
type MockBlob struct {
	ctrl     *gomock.Controller
	recorder *MockBlobMockRecorder
}

// MockBlobMockRecorder is the mock recorder for MockBlob
type MockBlobMockRecorder struct {
	mock *MockBlob
}

// NewMockBlob creates a new mock instance
func NewMockBlob(ctrl *gomock.Controller) *MockBlob {
	mock := &MockBlob{ctrl: ctrl}
	mock.recorder = &MockBlobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlob) EXPECT() *MockBlobMockRecorder {
	return m.recorder
}

// AssetDescriptor mocks base method
func (m *MockBlob) AssetDescriptor() dist.AssetInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssetDescriptor")
	ret0, _ := ret[0].(dist.AssetInfo)
	return ret0
}

// AssetDescriptor indicates an expected call of AssetDescriptor
func (mr *MockBlobMockRecorder) AssetDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssetDescriptor", reflect.TypeOf((*MockBlob)(nil).AssetDescriptor))
}

// Open mocks base method
func (m *MockBlob) Open() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockBlobMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockBlob)(nil).Open))
}

// Size mocks base method
func (m *MockBlob) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockBlobMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockBlob)(nil).Size))
}
