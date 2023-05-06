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
	_ "embed"
	"flag"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inqizit-public/ebiten_template/internal/app"
)

var flagCRT = flag.Bool("crt", false, "enable the CRT effect")

//go:embed crt.go
var crtGo []byte

const (
	// screenWidth  = 1366
	// screenHeight = 768
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("ebiten_template")
	if err := ebiten.RunGame(app.NewGame()); err != nil {
		panic(err)
	}
}
