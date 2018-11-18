package main

import (
	"errors"
	"log"

	ko "github.com/eihigh/koromo"
	eb "github.com/hajimehoshi/ebiten"
)

var (
	love, hate, chara *eb.Image
	x, y              float64

	errSuccess = errors.New("")
)

func init() {
	love, _ = ko.LoadImage("i/img/love.png")
	hate, _ = ko.LoadImage("i/img/hate.png")
	chara, _ = ko.LoadImage("i/img/player.png")
}

func main() {
	err := eb.Run(update, 320, 240, 2, "aaaaaa")
	if err != nil {
		log.Fatal(err)
	}
}
