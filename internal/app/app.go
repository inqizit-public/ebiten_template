package app

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

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
	audio         Audio
	width, height int
}

type Audio struct {
	Context *audio.Context
	players map[string]*audio.Player
}

func NewGame() ebiten.Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	if g.audio.Context == nil {
		g.audio.Context = audio.NewContext(48000)
	}
	if g.audio.players == nil {
		g.audio.players = map[string]*audio.Player{}
	}
	jumpD, err := vorbis.DecodeWithoutResampling(bytes.NewReader(raudio.Jump_ogg))
	if err != nil {
		log.Fatal(err)
	}
	jumpPlayer, err := g.audio.Context.NewPlayer(jumpD)
	if err != nil {
		log.Fatal(err)
	}
	g.audio.players["jump"] = jumpPlayer

	jabD, err := wav.DecodeWithoutResampling(bytes.NewReader(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}
	hitPlayer, err := g.audio.Context.NewPlayer(jabD)
	if err != nil {
		log.Fatal(err)
	}
	g.audio.players["hit"] = hitPlayer
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.width = outsideWidth
	g.height = outsideHeight
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	screen.Fill(color.White)
	// S := 100
	// dc := gg.NewContext(S, S)

	// dc.DrawCircle(50, 50, 50.0)
	// dc.SetRGBA255(255, 0, 0, 255)
	// dc.Fill()
	// img := ebiten.NewImageFromImage(dc.Image())
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(100, 100)
	// screen.DrawImage(img, op)
	// scoreStr := fmt.Sprintf("Something")
	// text.Draw(screen, scoreStr, arcadeFont, g.width-len(scoreStr)*fontSize, fontSize, color.Black)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
