package store

import (
	"reflect"
	"testing"

	"github.com/parking_lot/parkinglot"
)

func generateData(isCar bool) {
	d := parkinglot.Parking{
		Slot: 5,
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

func generateAvailableSlot() {
	AvailableSlot = []int{1, 2, 3, 4, 5}
}

func flushAvailableSlot() {
	AvailableSlot = nil
}

func TestStore_CreateSlots(t *testing.T) {
	type args struct {
		sum int
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

			flushData()
			flushAvailableSlot()
		})
	}
}

func TestStore_FillSlot(t *testing.T) {
	type args struct {
		slot int
		car  parkinglot.Car
	}
	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		isGenerateData          bool
		isCar                   bool
		isGenerateAvailableSlot bool
		remainingAvailableSlot  []int
		isAvailableSlot         bool
	}{
		{
			name: "#1 slot not found",
			args: args{
				slot: 5,
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
				slot: 5,
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
				slot: 5,
				car: parkinglot.Car{
					RegistrationNumber: "KH-1234",
					Color:              "White",
				},
			},
			wantErr:        false,
			isGenerateData: true,
			isCar:          false,
			isGenerateAvailableSlot: true,
			remainingAvailableSlot:  []int{2, 3, 4, 5},
			isAvailableSlot:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &Store{}
			if tt.isGenerateData {
				generateData(tt.isCar)
			}
			if tt.isGenerateAvailableSlot {
				generateAvailableSlot()
			}

			err := s.FillSlot(tt.args.slot, tt.args.car)

			if tt.isAvailableSlot {
				if !reflect.DeepEqual(AvailableSlot, tt.remainingAvailableSlot) {
					t.Errorf("Store.GetAvailableSlot() = %v, want %v", AvailableSlot, tt.remainingAvailableSlot)
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FillSlot() error = %v, wantErr %v", err, tt.wantErr)
			}

			flushData()
			flushAvailableSlot()
		})
	}
}

func TestStore_FreeSlot(t *testing.T) {
	type args struct {
		slot int
	}
	tests := []struct {
		name                   string
		args                   args
		wantErr                bool
		isGenerateData         bool
		isCar                  bool
		isAvailableSlot        bool
		remainingAvailableSlot []int
	}{
		{
			name: "#1 slot not found",
			args: args{
				slot: 5,
			},
			wantErr: true,
		},
		{
			name: "#2 slot found, car not empty",
			args: args{
				slot: 5,
			},
			wantErr:                false,
			isGenerateData:         true,
			isCar:                  true,
			isAvailableSlot:        true,
			remainingAvailableSlot: []int{5},
		},
		{
			name: "#3 slot found, car empty",
			args: args{
				slot: 5,
			},
			wantErr:        true,
			isGenerateData: true,
			isCar:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if tt.isGenerateData {
				generateData(tt.isCar)
			}

			err := s.FreeSlot(tt.args.slot)

			if tt.isAvailableSlot {
				if !reflect.DeepEqual(AvailableSlot, tt.remainingAvailableSlot) {
					t.Errorf("Store.GetAvailableSlot() = %v, want %v", AvailableSlot, tt.remainingAvailableSlot)
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FreeSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
			flushData()
			flushAvailableSlot()
		})
	}
}

func TestStore_GetStatus(t *testing.T) {
	tests := []struct {
		name           string
		want           []parkinglot.Parking
		wantErr        bool
		isGenerateData bool
	}{
		{
			name:    "#1 error empty data",
			want:    []parkinglot.Parking{},
			wantErr: true,
		},
		{
			name:           "#2 data not empty",
			want:           Data,
			wantErr:        false,
			isGenerateData: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if tt.isGenerateData {
				generateData(true)
				tt.want = Data
			}
			got, err := s.GetStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetStatus() = %v, want %v", got, tt.want)
			}

			flushData()
		})
	}
}

func TestStore_GetAvailableSlot(t *testing.T) {
	tests := []struct {
		name                    string
		want                    int
		wantErr                 bool
		isGenerateAvailableSlot bool
	}{
		{
			name:    "#1 empty available slot",
			want:    0,
			wantErr: true,
		},
		{
			name:                    "#2 available slot exist",
			want:                    1,
			wantErr:                 false,
			isGenerateAvailableSlot: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}

			if tt.isGenerateAvailableSlot {
				generateAvailableSlot()
			}
			got, err := s.GetAvailableSlot()
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetAvailableSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store.GetAvailableSlot() = %v, want %v", got, tt.want)
			}

			flushAvailableSlot()
		})
	}
}

func TestStore_GetCars(t *testing.T) {
	type args struct {
		filter parkinglot.FilterType
		value  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []parkinglot.Car
		wantErr bool
	}{
		{
			name: "#1 error car not found with color",
			args: args{
				filter: parkinglot.FilterTypeColor,
				value:  "White",
			},
			want:    []parkinglot.Car{},
			wantErr: true,
		},
		{
			name: "#2 car found with color",
			args: args{
				filter: parkinglot.FilterTypeColor,
				value:  "Black",
			},
			want: []parkinglot.Car{
				parkinglot.Car{
					RegistrationNumber: "KH-2345",
					Color:              "Black",
				},
			},
			wantErr: false,
		},
		{
			name: "#3 car found with registration number",
			args: args{
				filter: parkinglot.FilterTypeRegistrationNumber,
				value:  "KH-2345",
			},
			want: []parkinglot.Car{
				parkinglot.Car{
					RegistrationNumber: "KH-2345",
					Color:              "Black",
				},
			},
			wantErr: false,
		},
		{
			name: "#3 car not found with registration number",
			args: args{
				filter: parkinglot.FilterTypeRegistrationNumber,
				value:  "KH-23456",
			},
			want:    []parkinglot.Car{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			generateData(true)
			got, err := s.GetCars(tt.args.filter, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetCars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetCars() = %v, want %v", got, tt.want)
			}

			flushData()
		})
	}
}

func TestStore_GetSlotNumbers(t *testing.T) {
	type args struct {
		filter parkinglot.FilterType
		value  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "#1 slot not found with color",
			args: args{
				filter: parkinglot.FilterTypeColor,
				value:  "White",
			},
			want:    []int{},
			wantErr: true,
		},
		{
			name: "#2 slot not found with registration number",
			args: args{
				filter: parkinglot.FilterTypeRegistrationNumber,
				value:  "KH-5678",
			},
			want:    []int{},
			wantErr: true,
		},
		{
			name: "#3 slot found with registration number",
			args: args{
				filter: parkinglot.FilterTypeRegistrationNumber,
				value:  "KH-2345",
			},
			want:    []int{5},
			wantErr: false,
		},
		{
			name: "#4 slot found with color",
			args: args{
				filter: parkinglot.FilterTypeColor,
				value:  "Black",
			},
			want:    []int{5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			generateData(true)
			got, err := s.GetSlotNumbers(tt.args.filter, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetSlotNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetSlotNumbers() = %v, want %v", got, tt.want)
			}
			flushData()
		})
	}
}
