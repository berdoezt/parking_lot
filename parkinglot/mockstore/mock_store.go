// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/parking_lot/parkinglot (interfaces: Store)

package mockstore

import (
	gomock "github.com/golang/mock/gomock"
	parkinglot "github.com/parking_lot/parkinglot"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStore) EXPECT() *MockStoreMockRecorder {
	return _m.recorder
}

// CreateSlots mocks base method
func (_m *MockStore) CreateSlots(_param0 int64) error {
	ret := _m.ctrl.Call(_m, "CreateSlots", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSlots indicates an expected call of CreateSlots
func (_mr *MockStoreMockRecorder) CreateSlots(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateSlots", reflect.TypeOf((*MockStore)(nil).CreateSlots), arg0)
}

// FillSlot mocks base method
func (_m *MockStore) FillSlot(_param0 int64, _param1 parkinglot.Car) error {
	ret := _m.ctrl.Call(_m, "FillSlot", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FillSlot indicates an expected call of FillSlot
func (_mr *MockStoreMockRecorder) FillSlot(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FillSlot", reflect.TypeOf((*MockStore)(nil).FillSlot), arg0, arg1)
}

// FreeSlot mocks base method
func (_m *MockStore) FreeSlot(_param0 int64) error {
	ret := _m.ctrl.Call(_m, "FreeSlot", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FreeSlot indicates an expected call of FreeSlot
func (_mr *MockStoreMockRecorder) FreeSlot(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FreeSlot", reflect.TypeOf((*MockStore)(nil).FreeSlot), arg0)
}

// GetAvailableSlot mocks base method
func (_m *MockStore) GetAvailableSlot() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetAvailableSlot")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableSlot indicates an expected call of GetAvailableSlot
func (_mr *MockStoreMockRecorder) GetAvailableSlot() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAvailableSlot", reflect.TypeOf((*MockStore)(nil).GetAvailableSlot))
}

// GetCarAttribute mocks base method
func (_m *MockStore) GetCarAttribute(_param0 parkinglot.FilterType) (parkinglot.Car, error) {
	ret := _m.ctrl.Call(_m, "GetCarAttribute", _param0)
	ret0, _ := ret[0].(parkinglot.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCarAttribute indicates an expected call of GetCarAttribute
func (_mr *MockStoreMockRecorder) GetCarAttribute(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetCarAttribute", reflect.TypeOf((*MockStore)(nil).GetCarAttribute), arg0)
}

// GetSlotNumbers mocks base method
func (_m *MockStore) GetSlotNumbers(_param0 parkinglot.FilterType) ([]int64, error) {
	ret := _m.ctrl.Call(_m, "GetSlotNumbers", _param0)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotNumbers indicates an expected call of GetSlotNumbers
func (_mr *MockStoreMockRecorder) GetSlotNumbers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSlotNumbers", reflect.TypeOf((*MockStore)(nil).GetSlotNumbers), arg0)
}

// GetStatus mocks base method
func (_m *MockStore) GetStatus() ([]parkinglot.Parking, error) {
	ret := _m.ctrl.Call(_m, "GetStatus")
	ret0, _ := ret[0].([]parkinglot.Parking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (_mr *MockStoreMockRecorder) GetStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetStatus", reflect.TypeOf((*MockStore)(nil).GetStatus))
}