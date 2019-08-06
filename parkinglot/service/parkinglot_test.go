package service

import (
	"reflect"
	"testing"

	"github.com/parking_lot/parkinglot"
)

func TestService_CreateParkingLot(t *testing.T) {
	type fields struct {
		store parkinglot.Store
	}
	type args struct {
		slot int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			if err := s.CreateParkingLot(tt.args.slot); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateParkingLot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Park(t *testing.T) {
	type fields struct {
		store parkinglot.Store
	}
	type args struct {
		car parkinglot.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			if err := s.Park(tt.args.car); (err != nil) != tt.wantErr {
				t.Errorf("Service.Park() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Leave(t *testing.T) {
	type fields struct {
		store parkinglot.Store
	}
	type args struct {
		slotID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			if err := s.Leave(tt.args.slotID); (err != nil) != tt.wantErr {
				t.Errorf("Service.Leave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetRegistrationNumbersByColor(t *testing.T) {
	type fields struct {
		store parkinglot.Store
	}
	type args struct {
		color string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			got, err := s.GetRegistrationNumbersByColor(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetRegistrationNumbersByColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetRegistrationNumbersByColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetSlotNumbersByColor(t *testing.T) {
	type fields struct {
		store parkinglot.Store
	}
	type args struct {
		color string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			got, err := s.GetSlotNumbersByColor(tt.args.color)
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
	type fields struct {
		store parkinglot.Store
	}
	type args struct {
		registrationNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				store: tt.fields.store,
			}
			got, err := s.GetSlotNumberByRegistrationNumber(tt.args.registrationNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSlotNumberByRegistrationNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetSlotNumberByRegistrationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
