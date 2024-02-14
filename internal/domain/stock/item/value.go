package item

import (
	"errors"
)

type (
	Name struct {
		string
	}
)

var (
	ErrInvalidName = errors.New("invalid name")
)

// NewName creates a new Name object.
//
// It takes a string parameter and returns a Name object and an error.
func NewName(v string) (Name, error) {
	if v == "" {
		return Name{}, ErrInvalidName
	}
	return Name{v}, nil
}

// String returns the string representation of the Name type.
//
// No parameters.
// Returns a string.
func (v Name) String() string {
	return v.string
}
