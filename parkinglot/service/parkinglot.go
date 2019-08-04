package service

import (
	"github.com/parking_lot/parkinglot"
)

// Park parking a car to available slot
func (s *Service) Park(car parkinglot.Car) error {
	return nil
}

// Leave car leaving slot
func (s *Service) Leave(slotID int64) error {
	return nil
}

// CreateParkingLot crete parking lot with specific slot
func (s *Service) CreateParkingLot(slot int64) error {
	return nil
}

// GetStatus get all the parking status
func (s *Service) GetStatus() ([]parkinglot.Parking, error) {
	return []parkinglot.Parking{}, nil
}

// GetRegistrationNumbersByColor get registration numbers by color
func (s *Service) GetRegistrationNumbersByColor(color string) ([]string, error) {
	return []string{""}, nil
}

// GetSlotNumbersByColor get slot numbers by color
func (s *Service) GetSlotNumbersByColor(color string) ([]int64, error) {
	return []int64{0}, nil
}

// GetSlotNumberByRegistrationNumber get slot number by registration number
func (s *Service) GetSlotNumberByRegistrationNumber(registrationNumber string) (int64, error) {
	return 0, nil
}
