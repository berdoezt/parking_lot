package store

import (
	"github.com/parking_lot/parkinglot"
)

// Store denotes store object
type Store struct{}

// New create new instance of store
func New() parkinglot.Store {
	return &Store{}
}
