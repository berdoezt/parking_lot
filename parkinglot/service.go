package parkinglot

// Service provides mechanism to access parkinglot features
type Service interface {
	CreateParkingLot(slot int64) error
	Park(car Car) error
	Leave(slotID int64) error
	GetStatus() ([]Parking, error)
	GetRegistrationNumbersByColor(color string) ([]string, error)
	GetSlotNumbersByColor(color string) ([]int64, error)
	GetSlotNumberByRegistrationNumber(registrationNumber string) (int64, error)
}
