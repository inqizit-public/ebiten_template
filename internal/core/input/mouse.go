package input

type MouseInput interface {
	IsButtonPressed(MouseButton int) bool
	IsButtonJustPressed(MouseButton int) bool
	IsButtonJustReleased(MouseButton int) bool
	ButtonPressedDuration(MouseButton int) int // return how long the button has been pressed in ticks
	CurrentCursorPosition() (x, y int)

	// Wheel returns x and y offsets of the mouse wheel. It returns 0 if the wheel isn't being rolled.
	Wheel() (xoff, yoff float64)
}

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
)
