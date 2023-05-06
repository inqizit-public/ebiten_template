package game

import (
	"github.com/inqizit-public/ebiten_template/internal/core/draw"
	"github.com/inqizit-public/ebiten_template/internal/core/input"
)

type App struct {
	input Input
}

type Input interface {
	input.MouseInput
	input.KeyInput
}

func NewApp(input Input) *App {
	return &App{
		input: input,
	}
}

func (a *App) Init() error {
	return nil
}

func (a *App) Update() error {
	return nil
}

func (a *App) Draw() []*draw.DrawOpts {

}
