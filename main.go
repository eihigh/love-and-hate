package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/eihigh/love-and-hate/internal/audio"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

var g *game

func main() {
	rand.Seed(time.Now().UnixNano())

	// if env.DebugMode {
	// 	logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// 	if err != nil {
	// 		panic("cannot open test.log: " + err.Error())
	// 	}
	// 	defer logfile.Close()
	// 	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	// 	log.SetFlags(log.Ldate | log.Ltime)
	// 	log.Println("logging start")
	// }

	audio.Load()
	audio.SetBgmVolume(0.5)
	audio.SetSeVolume(0.2)
	defer audio.Finalize()

	g = newGame()
	err := ebiten.Run(update, int(env.View.W), int(env.View.H), 2, env.GameTitle)
	if err != nil && err != sio.ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	draw.Screen = screen
	return g.update()
}
