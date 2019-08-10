package parkinglot

// Store provide mechanism to manage parkinglot data
//
//go:generate mockgen -destination ./mockstore/mock_store.go -package mockstore github.com/parking_lot/parkinglot Store
type Store interface {
	CreateSlots(sum int64) error
	FillSlot(slot Slot, car Car) error
	FreeSlot(slot Slot) error
	GetAvailableSlot() (Slot, error)
	GetStatus() ([]Parking, error)
	GetCars(filter FilterType, value interface{}) ([]Car, error)
	GetSlotNumbers(filter FilterType, value interface{}) ([]Slot, error)
}
