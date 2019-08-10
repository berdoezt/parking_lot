package parkinglot

// Store provide mechanism to manage parkinglot data
//
//go:generate mockgen -destination ./mockstore/mock_store.go -package mockstore github.com/parking_lot/parkinglot Store
type Store interface {
	CreateSlots(sum int) error
	FillSlot(slot int, car Car) error
	FreeSlot(slot int) error
	GetAvailableSlot() (int, error)
	GetStatus() ([]Parking, error)
	GetCars(filter FilterType, value interface{}) ([]Car, error)
	GetSlotNumbers(filter FilterType, value interface{}) ([]int, error)
}
