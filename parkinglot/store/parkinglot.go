package store

import (
	"errors"

	"github.com/parking_lot/parkinglot"
)

var (
	// Data to store data on the fly
	Data []parkinglot.Parking

	// AvailableSlot contains free slot that can be park with car, sorted from lowest to highest ID
	AvailableSlot []int64
)

// CreateSlots create parking slot
// we can always create slots to extend parking slot while the program running
func (s *Store) CreateSlots(sum int64) error {
	id := int64(len(Data) + 1)

	for i := int64(0); i < sum; i++ {
		p := parkinglot.Parking{
			Slot: parkinglot.Slot(id),
		}

		Data = append(Data, p)
		AvailableSlot = append(AvailableSlot, id)
		id++
	}

	return nil
}

// FillSlot fill available slot with car
func (s *Store) FillSlot(slot parkinglot.Slot, car parkinglot.Car) error {

	for _, d := range Data {
		if d.Slot == slot {
			if d.RegistrationNumber == "" {
				d.Car = car
				return nil
			}
		}
	}

	return errors.New("slot not found")
}

// FreeSlot empty the slot
func (s *Store) FreeSlot(slot parkinglot.Slot) error {

	for _, d := range Data {
		if d.Slot == slot {
			if d.RegistrationNumber != "" {
				d.Car = parkinglot.Car{}
				return nil
			}
		}
	}

	return errors.New("slot not found")
}

// GetAvailableSlot get available (empty) slot
func (s *Store) GetAvailableSlot() (parkinglot.Slot, error) {
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
func (s *Store) GetSlotNumbers(filter parkinglot.FilterType, value interface{}) ([]parkinglot.Slot, error) {
	return nil, nil
}
