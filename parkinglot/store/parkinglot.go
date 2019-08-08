package store

import (
	"github.com/parking_lot/parkinglot"
)

// CreateSlots create parking slot
func (s *Store) CreateSlots(sum int64) error {
	return nil
}

// FillSlot fill available slot with car
func (s *Store) FillSlot(slotID int64, car parkinglot.Car) error {
	return nil
}

// FreeSlot empty the slot
func (s *Store) FreeSlot(slotID int64) error {
	return nil
}

// GetAvailableSlot get available (empty) slot
func (s *Store) GetAvailableSlot() (int64, error) {
	return 0, nil
}

// GetStatus get all the parking status
func (s *Store) GetStatus() ([]parkinglot.Parking, error) {
	return nil, nil
}

// GetCars get cars based on filter
func (s *Store) GetCars(filter parkinglot.FilterType, value interface{}) ([]parkinglot.Car, error) {
	return []parkinglot.Car{}, nil
}

// GetSlotNumbers get slot number with filter
func (s *Store) GetSlotNumbers(filter parkinglot.FilterType, value interface{}) ([]int64, error) {
	return nil, nil
}
