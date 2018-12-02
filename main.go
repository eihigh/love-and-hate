package main

import (
	"image/color"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	titleName = "love and hate"
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

	g := newGame()

	err := ebiten.Run(g.update, int(view.W), int(view.H), 2, titleName)
	if err != nil && err != sio.ErrSuccess {
		log.Fatal(err)
	}
}

func bd(r *sio.Rect) {
	x, y := r.Pos(7)
	w, h := r.W, r.H
	ebitenutil.DrawLine(scr, x, y, x+w, y, color.White)
	ebitenutil.DrawLine(scr, x+w, y, x+w, y+h, color.White)
	ebitenutil.DrawLine(scr, x+w, y+h, x, y+h, color.White)
	ebitenutil.DrawLine(scr, x, y+h, x, y, color.White)
}
