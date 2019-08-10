package parkinglot

// Service provides mechanism to access parkinglot features
//
//go:generate mockgen -destination ./mockservice/mock_service.go -package mockservice github.com/parking_lot/parkinglot Service
type Service interface {
	CreateParkingLot(sum int64) error
	Park(car Car) (int64, error)
	Leave(slotID int64) error
	GetStatus() ([]Parking, error)
	GetRegistrationNumbersByColor(color string) ([]string, error)
	GetSlotNumbersByColor(color string) ([]Slot, error)
	GetSlotNumberByRegistrationNumber(registrationNumber string) (Slot, error)
}
