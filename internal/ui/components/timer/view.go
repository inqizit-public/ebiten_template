package timer

import (
	"image"

	"github.com/inqizit/focus/internal/ui/components/component"
)

type View interface {
	draw(img image.Image, dim *component.Dimension, totalTime, elapsedTime int64) error
}

type ClockView struct {
}

func (c *ClockView) draw(img image.Image, dim *component.Dimension, totalTime, elapsedTime int64) error {

	return nil
}
