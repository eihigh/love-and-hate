package main

import "github.com/eihigh/sio"

// stageはステージ突入ごとに必ず作り直し
type stage struct {
	state sio.Stm
}

func newStage(level int) *stage {
	return &stage{}
}

func (s *stage) update() action {
	return noAction
}
