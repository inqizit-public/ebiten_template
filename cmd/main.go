// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"

	"github.com/fogleman/gg"
	"github.com/inqizit/focus/internal/ui/components/flower"

	_ "embed"
	"flag"
	"fmt"
	"image/color"

	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var flagCRT = flag.Bool("crt", false, "enable the CRT effect")

//go:embed crt.go
var crtGo []byte

func init() {
	rand.Seed(time.Now().UnixNano())
}

func floorDiv(x, y int) int {
	d := x / y
	if d*y == x || x >= 0 {
		return d
	}
	return d - 1
}

func floorMod(x, y int) int {
	return x - floorDiv(x, y)*y
}

const (
	// screenWidth  = 1366
	// screenHeight = 768
	screenWidth      = 640
	screenHeight     = 480
	tileSize         = 32
	titleFontSize    = fontSize * 1.5
	fontSize         = 24
	smallFontSize    = fontSize / 2
	pipeWidth        = tileSize * 2
	pipeStartOffsetX = 8
	pipeIntervalX    = 8
	pipeGapY         = 5
)

var (
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	touchIDs     []ebiten.TouchID
	audioContext *audio.Context
	jumpPlayer   *audio.Player
	hitPlayer    *audio.Player
	w, h         int
}

func NewGame() ebiten.Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	if g.audioContext == nil {
		g.audioContext = audio.NewContext(48000)
	}

	jumpD, err := vorbis.DecodeWithoutResampling(bytes.NewReader(raudio.Jump_ogg))
	if err != nil {
		log.Fatal(err)
	}
	g.jumpPlayer, err = g.audioContext.NewPlayer(jumpD)
	if err != nil {
		log.Fatal(err)
	}

	jabD, err := wav.DecodeWithoutResampling(bytes.NewReader(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}
	g.hitPlayer, err = g.audioContext.NewPlayer(jabD)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// s := ebiten.DeviceScaleFactor()
	// fmt.Println(outsideWidth, outsideHeight)
	g.w = outsideWidth
	g.h = outsideHeight
	return outsideWidth, outsideHeight
	// return screenWidth, screenHeight
	// return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func (g *Game) GetPosClicked() []*flower.Vector2D {
	res := make([]*flower.Vector2D, 0)
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		res = append(res, &flower.Vector2D{0, 0})
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		res = append(res, &flower.Vector2D{x, y})
	}
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	for _, t := range g.touchIDs {
		x, y := ebiten.TouchPosition(t)
		res = append(res, &flower.Vector2D{x, y})
	}

	return res

}
func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	screen.Fill(color.White)
	S := 100
	dc := gg.NewContext(S, S)

	dc.DrawCircle(50, 50, 50.0)
	dc.SetRGBA255(255, 0, 0, 255)
	dc.Fill()
	img := ebiten.NewImageFromImage(dc.Image())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100)
	screen.DrawImage(img, op)
	scoreStr := fmt.Sprintf("Something")
	text.Draw(screen, scoreStr, arcadeFont, g.w-len(scoreStr)*fontSize, fontSize, color.Black)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Focus")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
