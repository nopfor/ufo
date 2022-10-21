// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "embed"
    "image"
    _ "image/png"
    "log"
    "math/rand"
    "time"
    "os"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
    width = 150
    height = 54
    titleBarSize = 32
)

var (
    ufo *ebiten.Image
    xSpeed = 256
    ySpeed = 256
    acceleration = 8
    initialized = false
)

//go:embed UFO.png
var f embed.FS

func init() {
    infile, err := f.Open("UFO.png")
    if err != nil {
        log.Fatal(err)
    }
    defer infile.Close()

    img, _, err := image.Decode(infile)
    if err != nil {
        log.Fatal(err)
    }
    ufo = ebiten.NewImageFromImage(img)
}

func init() {
    rand.Seed(time.Now().UnixNano())

    // Exit immediately 75% of the time
    if rand.Intn(4) > 0 {
        os.Exit(0)
    }
}

type mascot struct {
    x16  int
    y16  int
    vx16 int
    vy16 int
}

func (m *mascot) Layout(outsideWidth, outsideHeight int) (int, int) {
    return width, height
}

func (m *mascot) Update() error {

    // Get window size
    sw, sh := ebiten.ScreenSizeInFullscreen()

    // Initialize
    if !initialized {
        // Start between 25% and 75% vertically
        m.y16 = 16*((titleBarSize+height-sh)*(25+rand.Intn(50))/100)

        // Set x direction and starting side
        xDirection := rand.Intn(2)
        if xDirection == 1 {
            m.x16 = 0
        } else {
            xDirection = -1
            m.x16 = 16*(sw-width)
        }

        // Set y direction
        yDirection := rand.Intn(2)
        if yDirection == 0 {
            yDirection = -1
        }

        // Set static speed and get moving
        xSpeed = 128+rand.Intn(256)
        ySpeed = 16+rand.Intn(32)

        m.vx16 = xDirection * xSpeed
        m.vy16 = yDirection * ySpeed

        initialized = true
    }

    // Redraw
    ebiten.SetWindowPosition(m.x16/16, m.y16/16+sh-height)

    // Control it for awhile
    if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
        m.vx16 = -xSpeed
    } else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
        m.vx16 = xSpeed
    } else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
        m.vy16 = ySpeed
    } else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
        m.vy16 = -ySpeed
    }

    // Left/right motion
    m.x16 += m.vx16

    // Crash into the left/right walls
    if m.x16 <= 0 && m.vx16 < 0 {
        os.Exit(0)
    } else if m.x16/16 > sw-width && m.vx16 > 0 {
        os.Exit(0)
    }

    // Up/down motion
    m.y16 += m.vy16

    // Crash into the floor/ceiling
    if m.y16 >= 0 && m.vy16 > 0 {
        os.Exit(0)
    } else if m.y16/16 < titleBarSize+height-sh && m.vy16 < 0 {
        os.Exit(0)
    }

    return nil
}

func (m *mascot) Draw(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}
    if m.vx16 < 0 {
        op.GeoM.Scale(-1, 1)
        op.GeoM.Translate(width, 0)
    }
    screen.DrawImage(ufo, op)
}

func main() {
    ebiten.SetScreenTransparent(true)
    ebiten.SetWindowDecorated(false) // set to 'true' to show window borders
    ebiten.SetWindowFloating(true)
    ebiten.SetWindowSize(width, height)
    if err := ebiten.RunGame(&mascot{}); err != nil {
        log.Fatal(err)
    }
}
