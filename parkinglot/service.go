package parkinglot

// Service provides mechanism to access parkinglot features
//
//go:generate mockgen -destination ./mockservice/mock_service.go -package mockservice github.com/parking_lot/parkinglot Service
type Service interface {
	CreateParkingLot(sum int) error
	Park(car Car) (int, error)
	Leave(slotID int) error
	GetStatus() ([]Parking, error)
	GetRegistrationNumbersByColor(color string) ([]string, error)
	GetSlotNumbersByColor(color string) ([]int, error)
	GetSlotNumberByRegistrationNumber(registrationNumber string) (int, error)
}
