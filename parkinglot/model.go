package parkinglot

type (
	// Slot denotes parking slot ID
	Slot int64

	// Car car object
	Car struct {
		RegistrationNumber string
		Color              string
	}

	// Parking parking data object
	Parking struct {
		Slot
		Car
	}
)

type (
	// FilterType denotes the parking filter type
	FilterType int
)

const (
	_ FilterType = iota

	// FilterTypeColor color filter
	FilterTypeColor

	// FilterTypeRegistrationNumber registration number filter
	FilterTypeRegistrationNumber
)
