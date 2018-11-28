package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

const (
	titleName = "prototype"
)

var (
	view = sio.NewRect(7, 0, 0, 320, 240)

	debugMode = true
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if debugMode {
		logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("cannot open test.log: " + err.Error())
		}
		defer logfile.Close()
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("logging start")
	}

	g := &game{}
	g.state.To(sceneOpening)

	w, h := view.Width(), view.Height()
	err := ebiten.Run(g.update, int(w), int(h), 2, titleName)
	if err != nil && err != sio.ErrSuccess {
		log.Fatal(err)
	}
}
