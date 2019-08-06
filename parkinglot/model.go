package parkinglot

type (
	// Car car object
	Car struct {
		RegistrationNumber string
		Color              string
	}

	// Parking parking data object
	Parking struct {
		Slot int64
		Car
	}
)

type (
	// FilterType denotes the parking filter type
	FilterType int
	// AttributeType denotes the car attribute
	AttributeType int
)

const (
	_ FilterType = iota

	// FilterTypeColor color filter
	FilterTypeColor

	// FilterTypeRegistrationNumber registration number filter
	FilterTypeRegistrationNumber
)

const (
	_ AttributeType = iota

	// AttributeTypeColor color attribute
	AttributeTypeColor

	// AttributeTypeRegistrationNumber registration number attribute
	AttributeTypeRegistrationNumber
)
