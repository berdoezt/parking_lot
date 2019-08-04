package parkinglot

var defaultService Service

// Init initialize parkinglot service
func Init(svc Service) {
	defaultService = svc
}

// GetService get parkinglot service
func GetService() Service {
	return defaultService
}
