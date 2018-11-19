package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/bitmapfont"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	x, y    float64
	images  map[string]*ebiten.Image
	display *sio.Rect
	fface   = bitmapfont.Gothic12r
)

func init() {
	rand.Seed(time.Now().UnixNano())

	display = sio.NewRect(7, 0, 0, 320, 240)

	// load resouces
	names := []string{
		"love",
		"hate",
		"player",
	}
	for _, name := range names {
		i, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("i/img/%s.png", name), ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		images[name] = i
	}
}

func main() {
	w, h := display.Width(), display.Height()
	err := ebiten.Run(update, int(w), int(h), 2, "aaaaaa")
	if err != nil && err != sio.ErrSuccess {
		log.Fatal(err)
	}
}
