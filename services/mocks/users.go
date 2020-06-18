// Code generated by MockGen. DO NOT EDIT.
// Source: users.go

// Package mock_services is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	models "restapi/models"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ListUser mocks base method
func (m *MockUserService) ListUser(arg0 *gin.Context) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", arg0)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser
func (mr *MockUserServiceMockRecorder) ListUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockUserService)(nil).ListUser), arg0)
}

// GetUser mocks base method
func (m *MockUserService) GetUser(arg0 *gin.Context, arg1 int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserServiceMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), arg0, arg1)
}

// CreateUser mocks base method
func (m *MockUserService) CreateUser(c *gin.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserServiceMockRecorder) CreateUser(c, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), c, user)
}

// IsExisted mocks base method
func (m *MockUserService) IsExisted(c *gin.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExisted", c, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExisted indicates an expected call of IsExisted
func (mr *MockUserServiceMockRecorder) IsExisted(c, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExisted", reflect.TypeOf((*MockUserService)(nil).IsExisted), c, username)
}
