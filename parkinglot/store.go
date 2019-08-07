package parkinglot

// Store provide mechanism to manage parkinglot data
//
//go:generate mockgen -destination ./mockstore/mock_store.go -package mockstore github.com/parking_lot/parkinglot Store
type Store interface {
	CreateSlots(sum int64) error
	FillSlot(slotID int64, car Car) error
	FreeSlot(slotID int64) error
	GetAvailableSlot() (int64, error)
	GetStatus() ([]Parking, error)
	GetCars(filter FilterType) ([]Car, error)
	GetSlotNumbers(filter FilterType) ([]int64, error)
}
