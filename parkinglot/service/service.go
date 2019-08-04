package service

import (
	"github.com/parking_lot/parkinglot"
)

// Service denote service object
type Service struct {
	store parkinglot.Store
}

// New create new instance of Service
func New(store parkinglot.Store) parkinglot.Service {
	service := &Service{
		store: store,
	}

	return service
}
