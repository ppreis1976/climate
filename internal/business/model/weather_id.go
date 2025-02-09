package model

import "errors"

const ErrWeatherIDInvalid = "invalid weather"

type WeatherID string

func (z WeatherID) String() string {
	return string(z)
}

func (z WeatherID) Validate() error {
	if len(z) == 0 {
		return errors.New(ErrWeatherIDInvalid)
	}
	return nil
}
