package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/parking_lot/parkinglot"
	mock_store "github.com/parking_lot/parkinglot/mockstore"
)

var err = errors.New("")

func TestService_CreateParkingLot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	type args struct {
		slot int64
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "#1 normal case",
			args: args{
				slot: 5,
			},
			mockFunc: func() {
				mockStore.EXPECT().CreateSlots(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "#2 error case",
			args: args{
				slot: 5,
			},
			mockFunc: func() {
				mockStore.EXPECT().CreateSlots(gomock.Any()).Return(err)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if err := service.CreateParkingLot(tt.args.slot); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateParkingLot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Park(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	type args struct {
		car parkinglot.Car
	}

	tests := []struct {
		name     string
		mockFunc func()
		args     args
		wantErr  bool
	}{
		{
			name: "#1 normal case",
			mockFunc: func() {
				mockStore.EXPECT().GetAvailableSlot().Return(int64(5), nil)
				mockStore.EXPECT().FillSlot(int64(5), gomock.Any()).Return(nil)
			},
			args: args{
				car: parkinglot.Car{
					Color:              "white",
					RegistrationNumber: "KA-01-HH-7777",
				},
			},
			wantErr: false,
		},
		{
			name: "#2 slot is empty",
			mockFunc: func() {
				mockStore.EXPECT().GetAvailableSlot().Return(int64(0), err)
			},
			args: args{
				car: parkinglot.Car{
					Color:              "white",
					RegistrationNumber: "KA-01-HH-7777",
				},
			},
			wantErr: true,
		},
		{
			name: "#3 can't fill parking slot",
			mockFunc: func() {
				mockStore.EXPECT().GetAvailableSlot().Return(int64(5), nil)
				mockStore.EXPECT().FillSlot(int64(5), gomock.Any()).Return(err)
			},
			args: args{
				car: parkinglot.Car{
					Color:              "white",
					RegistrationNumber: "KA-01-HH-7777",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if err := service.Park(tt.args.car); (err != nil) != tt.wantErr {
				t.Errorf("Service.Park() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Leave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	type args struct {
		slotID int64
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		wantErr  bool
	}{
		{
			name: "#1 normal case",
			mockFunc: func() {
				mockStore.EXPECT().FreeSlot(gomock.Any()).Return(nil)
			},
			args: args{
				slotID: 5,
			},
			wantErr: false,
		},
		{
			name: "#1 error leave parking slot",
			mockFunc: func() {
				mockStore.EXPECT().FreeSlot(gomock.Any()).Return(err)
			},
			args: args{
				slotID: 5,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if err := service.Leave(tt.args.slotID); (err != nil) != tt.wantErr {
				t.Errorf("Service.Leave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetRegistrationNumbersByColor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	cars := []parkinglot.Car{
		parkinglot.Car{
			Color:              "white",
			RegistrationNumber: "KA-01-BB-0001",
		},
		parkinglot.Car{
			Color:              "black",
			RegistrationNumber: "KA-01-HH-7777",
		},
	}

	var expected []string
	for _, c := range cars {
		expected = append(expected, c.RegistrationNumber)
	}

	type args struct {
		color string
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		want     []string
		wantErr  bool
	}{
		{
			name: "#1 car exists",
			mockFunc: func() {
				mockStore.EXPECT().GetCars(parkinglot.FilterTypeColor).Return(cars, nil)
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "#2 car don't exist",
			mockFunc: func() {
				mockStore.EXPECT().GetCars(parkinglot.FilterTypeColor).Return([]parkinglot.Car{}, nil)
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "#3 error",
			mockFunc: func() {
				mockStore.EXPECT().GetCars(parkinglot.FilterTypeColor).Return([]parkinglot.Car{}, err)
			},
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := service.GetRegistrationNumbersByColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetRegistrationNumbersByColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetRegistrationNumbersByColor() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestService_GetSlotNumbersByColor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	type args struct {
		color string
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		want     []int64
		wantErr  bool
	}{
		{
			name: "#1 slot numbers exists",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeColor).Return([]int64{1, 2, 3}, nil)
			},
			want:    []int64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "#2 slot numbers don't exists",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeColor).Return([]int64{}, nil)
			},
			want:    []int64{},
			wantErr: false,
		},
		{
			name: "#3 error",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeColor).Return([]int64{}, err)
			},
			want:    []int64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := service.GetSlotNumbersByColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSlotNumbersByColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetSlotNumbersByColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetSlotNumberByRegistrationNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	type args struct {
		color string
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		want     int64
		wantErr  bool
	}{
		{
			name: "#1 slot numbers exists",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber).Return(5, nil)
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "#2 error",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber).Return(0, err)
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := service.GetSlotNumbersByColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSlotNumbersByColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetSlotNumbersByColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetStatus(t *testing.T) {
	type fields struct {
		store parkinglot.Store
	}
	tests := []struct {
		name    string
		fields  fields
		want    []parkinglot.Parking
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			got, err := s.GetStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
