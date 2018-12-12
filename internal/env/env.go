package env

import "github.com/eihigh/sio"

const (
	GameTitle      = "love and hate"
	DebugMode      = true
	AssetsEmbedded = true
)

var (
	View     = sio.NewRect(7, 0, 0, 320, 240)
	PlayArea = sio.NewRect(7, 0, 0, 400, 480)
)
