package component

import (
	"image"
)

type Dimension struct {
	StartX, StartY, EndX, Endy int
}

type Component interface {
	Update(events []Event) error
	Draw(image *image.Image, dim *Dimension) error
}

type Event interface {
	Id() string
	Name() string
}
