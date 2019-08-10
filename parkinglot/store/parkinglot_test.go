package store

import (
	"testing"

	"github.com/parking_lot/parkinglot"
)

func generateData(isCar bool) {
	d := parkinglot.Parking{
		Slot: parkinglot.Slot(5),
	}

	if isCar {
		d.Car = parkinglot.Car{
			RegistrationNumber: "KH-2345",
			Color:              "Black",
		}
	}

	Data = []parkinglot.Parking{
		d,
	}
}

func flushData() {
	Data = nil
}

func TestStore_CreateSlots(t *testing.T) {
	type args struct {
		sum int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#1 normal case",
			args: args{
				sum: 6,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.CreateSlots(tt.args.sum); (err != nil) != tt.wantErr {
				t.Errorf("Store.CreateSlots() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_FillSlot(t *testing.T) {
	type args struct {
		slot parkinglot.Slot
		car  parkinglot.Car
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		isGenerateData bool
		isCar          bool
	}{
		{
			name: "#1 error",
			args: args{
				slot: parkinglot.Slot(5),
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr: true,
		},
		{
			name: "#2 slot found, car not empty",
			args: args{
				slot: parkinglot.Slot(5),
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr:        true,
			isGenerateData: true,
			isCar:          true,
		},
		{
			name: "#3 slot found, car empty",
			args: args{
				slot: parkinglot.Slot(5),
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr:        false,
			isGenerateData: true,
			isCar:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flushData()

			s := &Store{}
			if tt.isGenerateData {
				generateData(tt.isCar)
			}
			if err := s.FillSlot(tt.args.slot, tt.args.car); (err != nil) != tt.wantErr {
				t.Errorf("Store.FillSlot() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestStore_FreeSlot(t *testing.T) {
	type args struct {
		slot parkinglot.Slot
		car  parkinglot.Car
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		isGenerateData bool
		isCar          bool
	}{
		{
			name: "#1 error",
			args: args{
				slot: parkinglot.Slot(5),
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr: true,
		},
		{
			name: "#2 slot found, car not empty",
			args: args{
				slot: parkinglot.Slot(5),
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr:        false,
			isGenerateData: true,
			isCar:          true,
		},
		{
			name: "#3 slot found, car empty",
			args: args{
				slot: parkinglot.Slot(5),
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr:        true,
			isGenerateData: true,
			isCar:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flushData()
			s := &Store{}
			if tt.isGenerateData {
				generateData(tt.isCar)
			}
			if err := s.FreeSlot(tt.args.slot); (err != nil) != tt.wantErr {
				t.Errorf("Store.FreeSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
