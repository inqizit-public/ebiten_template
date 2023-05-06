package flower

import "image"

type Color struct {
	R, G, B, A int
}
type Vector2D struct {
	X, Y int
}
type Flower struct {
	Length int
	Color  *Color
	Pos    *Vector2D
	Speed  *Vector2D
	Mode   int
	Image  image.Image
}
