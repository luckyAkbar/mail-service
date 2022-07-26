// Code generated by MockGen. DO NOT EDIT.
// Source: mail-service/internal/model (interfaces: MailRepository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "mail-service/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMailRepository is a mock of MailRepository interface.
type MockMailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMailRepositoryMockRecorder
}

// MockMailRepositoryMockRecorder is the mock recorder for MockMailRepository.
type MockMailRepositoryMockRecorder struct {
	mock *MockMailRepository
}

// NewMockMailRepository creates a new mock instance.
func NewMockMailRepository(ctrl *gomock.Controller) *MockMailRepository {
	mock := &MockMailRepository{ctrl: ctrl}
	mock.recorder = &MockMailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailRepository) EXPECT() *MockMailRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMailRepository) Create(arg0 context.Context, arg1 *model.MailingList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMailRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMailRepository)(nil).Create), arg0, arg1)
}

// GetPendingMailingList mocks base method.
func (m *MockMailRepository) GetPendingMailingList(arg0 context.Context, arg1 int) ([]model.MailingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingMailingList", arg0, arg1)
	ret0, _ := ret[0].([]model.MailingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingMailingList indicates an expected call of GetPendingMailingList.
func (mr *MockMailRepositoryMockRecorder) GetPendingMailingList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingMailingList", reflect.TypeOf((*MockMailRepository)(nil).GetPendingMailingList), arg0, arg1)
}

// MarkAsFailed mocks base method.
func (m *MockMailRepository) MarkAsFailed(arg0 context.Context, arg1 *model.MailingList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAsFailed indicates an expected call of MarkAsFailed.
func (mr *MockMailRepositoryMockRecorder) MarkAsFailed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsFailed", reflect.TypeOf((*MockMailRepository)(nil).MarkAsFailed), arg0, arg1)
}

// MarkAsSent mocks base method.
func (m *MockMailRepository) MarkAsSent(arg0 context.Context, arg1 *model.MailingList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsSent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAsSent indicates an expected call of MarkAsSent.
func (mr *MockMailRepositoryMockRecorder) MarkAsSent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsSent", reflect.TypeOf((*MockMailRepository)(nil).MarkAsSent), arg0, arg1)
}