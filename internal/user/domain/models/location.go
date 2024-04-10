package model

// Location is a value object referencing user's default location
type Location struct {
	Longitude float64
	Latitude  float64
}

func NewLocation() Location {
	return Location{}
}

// TODO: Validation of Value Object should happen during the creation only, and new changes should return a new one
func (l *Location) SetCoordinates(long, lat float64) {
	l.Latitude = lat
	l.Longitude = long
}

func (l Location) IsValid() bool {
	return true
}
