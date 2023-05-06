package timer

import (
	"image"

	"github.com/inqizit/focus/internal/ui/components/component"
)

type Timer struct {
	*baseTimer
	View View
}

func NewTimer(seconds int64) *Timer {
	timer := &Timer{
		baseTimer: &baseTimer{
			total:   int64(seconds),
			elapsed: 0,
		},
		View: &ClockView{},
	}
	return timer
}

func (timer *Timer) Update(events []component.Event) error {
	timer.update()
	return nil
}

func (timer *Timer) Draw(img image.Image, dim *component.Dimension) error {
	return timer.View.draw(img, dim, timer.total, timer.elapsed)
}

// func (timer *Timer) GetElapsedTime() int {
// 	return 0
// }
