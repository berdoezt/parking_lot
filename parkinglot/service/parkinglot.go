package service

import (
	"errors"

	"github.com/parking_lot/parkinglot"
)

// Park parking a car to available slot
func (s *Service) Park(car parkinglot.Car) error {
	slotID, err := s.store.GetAvailableSlot()
	if err != nil {
		return err
	}

	if err := s.store.FillSlot(slotID, car); err != nil {
		return err
	}

	return nil
}

// Leave car leaving slot
func (s *Service) Leave(slotID int64) error {
	err := s.store.FreeSlot(slotID)
	if err != nil {
		return err
	}

	return nil
}

// CreateParkingLot crete parking lot with specific slot
func (s *Service) CreateParkingLot(slot int64) error {
	err := s.store.CreateSlots(slot)
	if err != nil {
		return err
	}

	return nil
}

// GetStatus get all the parking status
func (s *Service) GetStatus() ([]parkinglot.Parking, error) {
	return []parkinglot.Parking{}, nil
}

// GetRegistrationNumbersByColor get registration numbers by color
func (s *Service) GetRegistrationNumbersByColor(color string) ([]string, error) {
	result := make([]string, 0)

	cars, err := s.store.GetCars(parkinglot.FilterTypeColor, color)
	if err != nil {
		return result, err
	}

	for _, car := range cars {
		result = append(result, car.RegistrationNumber)
	}

	return result, nil
}

// GetSlotNumbersByColor get slot numbers by color
func (s *Service) GetSlotNumbersByColor(color string) ([]int64, error) {
	result := make([]int64, 0)

	slots, err := s.store.GetSlotNumbers(parkinglot.FilterTypeColor, color)
	if err != nil {
		return result, err
	}

	result = append(result, slots...)

	return result, nil
}

// GetSlotNumberByRegistrationNumber get slot number by registration number
func (s *Service) GetSlotNumberByRegistrationNumber(registrationNumber string) (int64, error) {
	var result int64

	slots, err := s.store.GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber, registrationNumber)
	if err != nil {
		return result, err
	}

	if len(slots) > 1 {
		return result, errors.New("duplicate registration number")
	}

	result = slots[0]
	return result, nil
}
