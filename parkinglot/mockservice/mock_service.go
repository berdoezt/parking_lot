// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/parking_lot/parkinglot (interfaces: Service)

// Package mockservice is a generated GoMock package.
package mockservice

import (
	gomock "github.com/golang/mock/gomock"
	parkinglot "github.com/parking_lot/parkinglot"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateParkingLot mocks base method
func (m *MockService) CreateParkingLot(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParkingLot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateParkingLot indicates an expected call of CreateParkingLot
func (mr *MockServiceMockRecorder) CreateParkingLot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParkingLot", reflect.TypeOf((*MockService)(nil).CreateParkingLot), arg0)
}

// GetRegistrationNumbersByColor mocks base method
func (m *MockService) GetRegistrationNumbersByColor(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistrationNumbersByColor", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistrationNumbersByColor indicates an expected call of GetRegistrationNumbersByColor
func (mr *MockServiceMockRecorder) GetRegistrationNumbersByColor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistrationNumbersByColor", reflect.TypeOf((*MockService)(nil).GetRegistrationNumbersByColor), arg0)
}

// GetSlotNumberByRegistrationNumber mocks base method
func (m *MockService) GetSlotNumberByRegistrationNumber(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotNumberByRegistrationNumber", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotNumberByRegistrationNumber indicates an expected call of GetSlotNumberByRegistrationNumber
func (mr *MockServiceMockRecorder) GetSlotNumberByRegistrationNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotNumberByRegistrationNumber", reflect.TypeOf((*MockService)(nil).GetSlotNumberByRegistrationNumber), arg0)
}

// GetSlotNumbersByColor mocks base method
func (m *MockService) GetSlotNumbersByColor(arg0 string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotNumbersByColor", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotNumbersByColor indicates an expected call of GetSlotNumbersByColor
func (mr *MockServiceMockRecorder) GetSlotNumbersByColor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotNumbersByColor", reflect.TypeOf((*MockService)(nil).GetSlotNumbersByColor), arg0)
}

// GetStatus mocks base method
func (m *MockService) GetStatus() ([]parkinglot.Parking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].([]parkinglot.Parking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockServiceMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockService)(nil).GetStatus))
}

// Leave mocks base method
func (m *MockService) Leave(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Leave indicates an expected call of Leave
func (mr *MockServiceMockRecorder) Leave(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockService)(nil).Leave), arg0)
}

// Park mocks base method
func (m *MockService) Park(arg0 parkinglot.Car) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Park", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Park indicates an expected call of Park
func (mr *MockServiceMockRecorder) Park(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Park", reflect.TypeOf((*MockService)(nil).Park), arg0)
}
