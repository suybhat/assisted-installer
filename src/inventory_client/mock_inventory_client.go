// Code generated by MockGen. DO NOT EDIT.
// Source: inventory_client.go

// Package inventory_client is a generated GoMock package.
package inventory_client

import (
	reflect "reflect"

	models "github.com/filanov/bm-inventory/models"
	gomock "github.com/golang/mock/gomock"
)

// MockInventoryClient is a mock of InventoryClient interface
type MockInventoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryClientMockRecorder
}

// MockInventoryClientMockRecorder is the mock recorder for MockInventoryClient
type MockInventoryClientMockRecorder struct {
	mock *MockInventoryClient
}

// NewMockInventoryClient creates a new mock instance
func NewMockInventoryClient(ctrl *gomock.Controller) *MockInventoryClient {
	mock := &MockInventoryClient{ctrl: ctrl}
	mock.recorder = &MockInventoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInventoryClient) EXPECT() *MockInventoryClientMockRecorder {
	return m.recorder
}

// DownloadFile mocks base method
func (m *MockInventoryClient) DownloadFile(filename, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", filename, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadFile indicates an expected call of DownloadFile
func (mr *MockInventoryClientMockRecorder) DownloadFile(filename, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockInventoryClient)(nil).DownloadFile), filename, dest)
}

// UpdateHostStatus mocks base method
func (m *MockInventoryClient) UpdateHostStatus(newStatus, hostId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostStatus", newStatus, hostId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHostStatus indicates an expected call of UpdateHostStatus
func (mr *MockInventoryClientMockRecorder) UpdateHostStatus(newStatus, hostId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostStatus", reflect.TypeOf((*MockInventoryClient)(nil).UpdateHostStatus), newStatus, hostId)
}

// GetEnabledHostsNamesIds mocks base method
func (m *MockInventoryClient) GetEnabledHostsNamesIds() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnabledHostsNamesIds")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnabledHostsNamesIds indicates an expected call of GetEnabledHostsNamesIds
func (mr *MockInventoryClientMockRecorder) GetEnabledHostsNamesIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabledHostsNamesIds", reflect.TypeOf((*MockInventoryClient)(nil).GetEnabledHostsNamesIds))
}

// UploadIngressCa mocks base method
func (m *MockInventoryClient) UploadIngressCa(ingressCA, clusterId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadIngressCa", ingressCA, clusterId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadIngressCa indicates an expected call of UploadIngressCa
func (mr *MockInventoryClientMockRecorder) UploadIngressCa(ingressCA, clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadIngressCa", reflect.TypeOf((*MockInventoryClient)(nil).UploadIngressCa), ingressCA, clusterId)
}

// GetCluster mocks base method
func (m *MockInventoryClient) GetCluster() (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster")
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockInventoryClientMockRecorder) GetCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockInventoryClient)(nil).GetCluster))
}
