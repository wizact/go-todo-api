package model

// Location is a value object referencing user's default location
type Location struct {
	longitude float64
	latitude  float64
}

func NewLocation(longitude, latitude float64) Location {
	return Location{longitude: longitude, latitude: latitude}
}

func (l Location) Coordinates() (float64, float64) {
	return l.longitude, l.latitude
}

func (l Location) IsValid() bool {
	return true
}
