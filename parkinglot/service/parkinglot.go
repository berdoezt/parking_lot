package service

import (
	"errors"

	"github.com/parking_lot/parkinglot"
)

// Park parking a car to available slot
func (s *Service) Park(car parkinglot.Car) (int, error) {
	slotID, err := s.store.GetAvailableSlot()
	if err != nil {
		return 0, err
	}

	if err := s.store.FillSlot(slotID, car); err != nil {
		return 0, err
	}

	return slotID, nil
}

// Leave car leaving slot
func (s *Service) Leave(slotID int) error {
	err := s.store.FreeSlot(slotID)
	if err != nil {
		return err
	}

	return nil
}

// CreateParkingLot crete parking lot with specific numbers
func (s *Service) CreateParkingLot(count int) error {
	err := s.store.CreateSlots(count)
	if err != nil {
		return err
	}

	return nil
}

// GetStatus get all the parking status
func (s *Service) GetStatus() ([]parkinglot.Parking, error) {
	parkingData, err := s.store.GetStatus()
	if err != nil {
		return []parkinglot.Parking{}, err
	}

	return parkingData, nil
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
func (s *Service) GetSlotNumbersByColor(color string) ([]int, error) {
	slots, err := s.store.GetSlotNumbers(parkinglot.FilterTypeColor, color)
	if err != nil {
		return []int{}, err
	}

	return slots, nil
}

// GetSlotNumberByRegistrationNumber get slot number by registration number
func (s *Service) GetSlotNumberByRegistrationNumber(registrationNumber string) (int, error) {
	slots, err := s.store.GetSlotNumbers(parkinglot.FilterTypeRegistrationNumber, registrationNumber)
	if err != nil {
		return 0, err
	}

	if len(slots) > 1 {
		return 0, errors.New("duplicate registration number")
	}

	return slots[0], nil
}
