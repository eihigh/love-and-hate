package main

import (
	"image/color"

	ko "github.com/eihigh/koromo"
	eb "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/mattn/go-runewidth"
)

var (
	scr *eb.Image

	scene  ko.Stm
	action string
)

type gameScene = int

const (
	title gameScene = iota
	play
	gameOver
)

type gameAction = int

const (
	actionToTitle = iota
	actionNewGame
)

func update(s *eb.Image) error {

	scr = s

	if onQuit() {
		return ko.ErrSuccess
	}

	action = ""
	scene.Update()
	switch scene.Current() {

	case title:
		updateTitle()

		if action == "new game" {
			scene.To(play)
		}

	case play:
		updatePlay()

		if action == "back to title" {
			scene.To(title)
		}

	case gameOver:

	}

	return nil
}

func updateTitle() {
	if onDecide() {
		action = "new game"
	}

	data := []struct {
		t string
		y int
	}{
		{
			t: "裏切られるのは一瞬だった。",
			y: 80,
		},
		{
			t: "だから、僕はその恨みを晴らす権利がある。",
			y: 96,
		},
		{
			t: "STAGE 0",
			y: 140,
		},
		{
			t: "MISSION -- LOVE < 35",
			y: 156,
		},
	}

	for _, d := range data {
		l := runewidth.StringWidth(d.t)
		x := 166 - l*6/2
		text.Draw(scr, d.t, fface, x, d.y, color.White)
	}
}

func updatePlay() {
	op := &eb.DrawImageOptions{}
	a := 0.8 * ko.UWave(scene.ElapsedRatio(60), 0.25)
	op.ColorM.Scale(1.0, 1.0, a, 1.0)
	scr.DrawImage(images["player"], op)

	if onLeft() {
		x -= 2
	}
	if onRight() {
		x += 2
	}
	if onUp() {
		y -= 2
	}
	if onDown() {
		y += 2
	}
	generate()

	if onDecide() {
		action = "back to title"
	}
}
