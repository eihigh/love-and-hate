package env

import "github.com/eihigh/sio"

const (
	GameTitle      = "love and hate"
	DebugMode      = true
	AssetsEmbedded = true
)

var (
	View     = sio.NewRect(7, 0, 0, 320, 240)
	PlayArea = View.Clone(5, 5).Scale(2, 3)
)
