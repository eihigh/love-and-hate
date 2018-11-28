package main

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/bitmapfont"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

var (
	fface = bitmapfont.Gothic12r
)

func DrawText(dst *ebiten.Image, tb *sio.TextBox, clr color.Color) {
	ofsX := tb.EmWidth / 2 // dot position
	ofsY := tb.EmHeight    // ditto
	lines := tb.Lines()

	for _, line := range lines {
		text.Draw(dst, line.Text, fface, line.X+int(ofsX), line.Y+int(ofsY), clr)
	}
}
