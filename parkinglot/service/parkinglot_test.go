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
		want     int64
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
			want:    5,
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
			want:    0,
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
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := service.Park(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Park() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Park() = %#v, want %#v", got, tt.want)
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
			name: "#2 error leave parking slot",
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
			Color:              "white",
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
				mockStore.EXPECT().GetCars(parkinglot.FilterTypeColor, gomock.Any()).Return(cars, nil)
			},
			args: args{
				color: "white",
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "#2 error",
			mockFunc: func() {
				mockStore.EXPECT().GetCars(parkinglot.FilterTypeColor, gomock.Any()).Return([]parkinglot.Car{}, err)
			},
			args: args{
				color: "black",
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
		want     []parkinglot.Slot
		wantErr  bool
	}{
		{
			name: "#1 slot numbers exists",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeColor, gomock.Any()).Return([]parkinglot.Slot{1, 2, 3}, nil)
			},
			args: args{
				color: "white",
			},
			want:    []parkinglot.Slot{1, 2, 3},
			wantErr: false,
		},
		{
			name: "#2 error",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeColor, gomock.Any()).Return([]parkinglot.Slot{}, err)
			},
			args: args{
				color: "black",
			},
			want:    []parkinglot.Slot{},
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
		registrationNumber string
	}
	tests := []struct {
		name     string
		mockFunc func()
		args     args
		want     parkinglot.Slot
		wantErr  bool
	}{
		{
			name: "#1 slot number exists",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber, gomock.Any()).Return([]parkinglot.Slot{5}, nil)
			},
			args: args{
				registrationNumber: "KH-1234",
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "#2 error",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber, gomock.Any()).Return([]parkinglot.Slot{}, err)
			},
			args: args{
				registrationNumber: "KH-1234",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "#3 duplicate registration number",
			mockFunc: func() {
				mockStore.EXPECT().GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber, gomock.Any()).Return([]parkinglot.Slot{1, 3}, nil)
			},
			args: args{
				registrationNumber: "KH-1234",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := service.GetSlotNumberByRegistrationNumber(tt.args.registrationNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSlotNumberByRegistrationNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetSlotNumberByRegistrationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_store.NewMockStore(ctrl)
	service := New(mockStore)

	tests := []struct {
		name     string
		mockFunc func()
		want     []parkinglot.Parking
		wantErr  bool
	}{
		{
			name: "#1 error",
			mockFunc: func() {
				mockStore.EXPECT().GetStatus().Return([]parkinglot.Parking{}, err)
			},
			want:    []parkinglot.Parking{},
			wantErr: true,
		},
		{
			name: "#2 parking data exists",
			mockFunc: func() {
				mockStore.EXPECT().GetStatus().Return([]parkinglot.Parking{
					parkinglot.Parking{
						Slot: 1,
						Car: parkinglot.Car{
							RegistrationNumber: "KH-1234",
							Color:              "Black",
						},
					},
					parkinglot.Parking{
						Slot: 2,
						Car: parkinglot.Car{
							RegistrationNumber: "KH-5678",
							Color:              "White",
						},
					},
				}, nil)
			},
			want: []parkinglot.Parking{
				parkinglot.Parking{
					Slot: 1,
					Car: parkinglot.Car{
						RegistrationNumber: "KH-1234",
						Color:              "Black",
					},
				},
				parkinglot.Parking{
					Slot: 2,
					Car: parkinglot.Car{
						RegistrationNumber: "KH-5678",
						Color:              "White",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := service.GetStatus()
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
