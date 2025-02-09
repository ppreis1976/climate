package model

import "errors"

type ZipCodeID string

const (
	ErrZipCodeIDInvalid = "invalid zipcode"
	ErrZipCodeNotFound  = "can not find zipcode"
	ErrMethodNotAllowed = "Method not allowed"
)

func (z ZipCodeID) String() string {
	return string(z)
}

func (z ZipCodeID) Validate() error {
	if len(z) != 8 {
		return errors.New(ErrZipCodeIDInvalid)
	}
	return nil
}
