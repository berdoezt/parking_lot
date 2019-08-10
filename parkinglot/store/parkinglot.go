package store

import (
	"errors"
	"fmt"
	"sort"

	"github.com/parking_lot/parkinglot"
)

var (
	// Data to store data on the fly
	Data []parkinglot.Parking

	// AvailableSlot contains free slot that can be park with car, sorted from lowest to highest ID
	AvailableSlot []int
)

// CreateSlots create parking slot
// we can always create slots to extend parking slot while the program running
func (s *Store) CreateSlots(sum int) error {
	id := len(Data) + 1

	for i := 0; i < sum; i++ {
		p := parkinglot.Parking{
			Slot: id,
		}

		Data = append(Data, p)
		AvailableSlot = append(AvailableSlot, id)
		id++
	}

	fmt.Println(AvailableSlot)
	fmt.Println(Data)

	return nil
}

// FillSlot fill available slot with car
func (s *Store) FillSlot(slot int, car parkinglot.Car) error {
	for i, d := range Data {
		if d.Slot == slot {
			if d.Car.RegistrationNumber == "" && d.Car.Color == "" {
				temp := d
				temp.Car = car
				Data[i] = temp
				AvailableSlot = AvailableSlot[1:]

				fmt.Println(AvailableSlot)
				fmt.Println(Data)
				return nil
			}
		}
	}

	return errors.New("slot not found")
}

// FreeSlot empty the slot
func (s *Store) FreeSlot(slot int) error {
	for i, d := range Data {
		if d.Slot == slot {
			if d.Car.RegistrationNumber != "" || d.Car.Color != "" {
				temp := d
				temp.Car = parkinglot.Car{}
				Data[i] = temp
				AvailableSlot = append(AvailableSlot, slot)

				sort.Ints(AvailableSlot)

				fmt.Println(AvailableSlot)
				fmt.Println(Data)
				return nil
			}
		}
	}

	return errors.New("slot not found")
}

// GetAvailableSlot get available (empty) slot
func (s *Store) GetAvailableSlot() (int, error) {
	if len(AvailableSlot) == 0 {
		return 0, errors.New("Sorry, parking lot is full")
	}

	return AvailableSlot[0], nil
}

// GetStatus get all the parking status
func (s *Store) GetStatus() ([]parkinglot.Parking, error) {
	if len(Data) == 0 {
		return []parkinglot.Parking{}, errors.New("data not found")
	}

	return Data, nil
}

// GetCars get cars based on filter
func (s *Store) GetCars(filter parkinglot.FilterType, value interface{}) ([]parkinglot.Car, error) {
	result := make([]parkinglot.Car, 0)

	for _, d := range Data {
		if filter == parkinglot.FilterTypeColor {
			if d.Car.Color == value {
				result = append(result, d.Car)
			}
		} else if filter == parkinglot.FilterTypeRegistrationNumber {
			if d.Car.RegistrationNumber == value {
				result = append(result, d.Car)
			}
		}
	}

	if len(result) == 0 {
		return result, errors.New("data not found")
	}

	return result, nil
}

// GetSlotNumbers get slot number with filter
func (s *Store) GetSlotNumbers(filter parkinglot.FilterType, value interface{}) ([]int, error) {
	result := make([]int, 0)

	for _, d := range Data {
		if filter == parkinglot.FilterTypeColor {
			if d.Car.Color == value {
				result = append(result, d.Slot)
			}
		} else if filter == parkinglot.FilterTypeRegistrationNumber {
			if d.Car.RegistrationNumber == value {
				result = append(result, d.Slot)
			}
		}
	}

	if len(result) == 0 {
		return result, errors.New("data not found")
	}

	return result, nil
}
