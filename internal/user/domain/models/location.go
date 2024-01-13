package model

// Location is a value object referencing user's default location
type Location struct {
	Longitude float64
	Latitude  float64
}

func NewLocation() Location {
	return Location{}
}

func (l *Location) SetCoordinates(long, lat float64) {
	l.Latitude = lat
	l.Longitude = long
}

func (l Location) IsValid() bool {
	return true
}
