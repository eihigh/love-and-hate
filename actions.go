package main

// 大きく状態の変更が発生するときのやり取りを疎結合にするための
// アクション定数
type action int

const (
	noAction action = iota

	// main game actions
	gameNewGame
	gameShowTitle
	gameShowGameOver

	// in stage actions
	stagePause
)
