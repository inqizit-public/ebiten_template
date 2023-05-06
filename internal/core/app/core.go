package app

import (
	"github.com/inqizit-public/ebiten_template/internal/core/draw"
)

type App interface {
	Init() error
	Update() error
	Draw() []*draw.DrawOpts
	Layout(width, height int) (appliedWidth int, appliedHeight int, err error)
	Quit() error
}
