package main

import eb "github.com/hajimehoshi/ebiten"

var (
	scr *eb.Image
)

func update(s *eb.Image) error {
	scr = s

	op := &eb.DrawImageOptions{}
	scr.DrawImage(love, op)
	scr.DrawImage(hate, op)

	op.GeoM.Translate(x, y)
	scr.DrawImage(player, op)

	if onLeft() {
		x--
	}

	if onRight() {
		x++
	}

	if onUp() {
		y--
	}

	if onDown() {
		y++
	}

	if onQuit() {
		return errSuccess
	}

	generate()
	return nil
}
