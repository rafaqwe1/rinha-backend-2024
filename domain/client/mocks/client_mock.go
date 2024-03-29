// Code generated by MockGen. DO NOT EDIT.
// Source: domain/client/client.go
//
// Generated by this command:
//
//	mockgen -source=domain/client/client.go -destination=domain/client/mocks/client_mock.go
//

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	client "github.com/rafaqwe1/rinha-backend-2024/domain/client"
	gomock "go.uber.org/mock/gomock"
)

// MockClientRepositoryInterface is a mock of ClientRepositoryInterface interface.
type MockClientRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientRepositoryInterfaceMockRecorder
}

// MockClientRepositoryInterfaceMockRecorder is the mock recorder for MockClientRepositoryInterface.
type MockClientRepositoryInterfaceMockRecorder struct {
	mock *MockClientRepositoryInterface
}

// NewMockClientRepositoryInterface creates a new mock instance.
func NewMockClientRepositoryInterface(ctrl *gomock.Controller) *MockClientRepositoryInterface {
	mock := &MockClientRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockClientRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientRepositoryInterface) EXPECT() *MockClientRepositoryInterfaceMockRecorder {
	return m.recorder
}

// BalanceWithTransactions mocks base method.
func (m *MockClientRepositoryInterface) BalanceWithTransactions(id, limitTransactions int) (*client.BalanceWithTransactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceWithTransactions", id, limitTransactions)
	ret0, _ := ret[0].(*client.BalanceWithTransactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceWithTransactions indicates an expected call of BalanceWithTransactions.
func (mr *MockClientRepositoryInterfaceMockRecorder) BalanceWithTransactions(id, limitTransactions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceWithTransactions", reflect.TypeOf((*MockClientRepositoryInterface)(nil).BalanceWithTransactions), id, limitTransactions)
}

// DecreaseBalance mocks base method.
func (m *MockClientRepositoryInterface) DecreaseBalance(id, value int) (*client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseBalance", id, value)
	ret0, _ := ret[0].(*client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseBalance indicates an expected call of DecreaseBalance.
func (mr *MockClientRepositoryInterfaceMockRecorder) DecreaseBalance(id, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseBalance", reflect.TypeOf((*MockClientRepositoryInterface)(nil).DecreaseBalance), id, value)
}

// Exists mocks base method.
func (m *MockClientRepositoryInterface) Exists(id int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockClientRepositoryInterfaceMockRecorder) Exists(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockClientRepositoryInterface)(nil).Exists), id)
}

// IncreaseBalance mocks base method.
func (m *MockClientRepositoryInterface) IncreaseBalance(id, value int) (*client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseBalance", id, value)
	ret0, _ := ret[0].(*client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncreaseBalance indicates an expected call of IncreaseBalance.
func (mr *MockClientRepositoryInterfaceMockRecorder) IncreaseBalance(id, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseBalance", reflect.TypeOf((*MockClientRepositoryInterface)(nil).IncreaseBalance), id, value)
}
