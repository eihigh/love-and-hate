package main

import ko "github.com/eihigh/koromo"

var (
	player struct {
		pos   complex128
		loves int
		hates int
	}

	stage struct {
		target int
		state  ko.Stm
	}

	symbols []symbol
)

type symbol interface {
	isLove() bool
	pos() (float64, float64)
	update()
}

type stageState = int

const (
	beginStage stageState = iota
	playStage
	resultStage
)

func updateStage() {

	switch stage.state.Current() {
	case beginStage:
	case playStage:
	case resultStage:

	}

	for _, sym := range symbols {
		sym.update()
	}

	// ここで当たり判定
}
