package parkinglot

// Car denotes car object
type Car struct {
	RegistrationNumber string
	Color              string
}

// Parking denotes parking data
type Parking struct {
	Slot int64
	Car
}
