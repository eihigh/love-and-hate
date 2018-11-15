package main

import (
	"errors"
	"log"

	ko "github.com/ei1chi/koromo"
	eb "github.com/hajimehoshi/ebiten"
)

var (
	love, hate, player *eb.Image
	x, y               float64

	errSuccess = errors.New("")
)

func init() {
	love, _ = ko.LoadImage("i/img/love.png")
	hate, _ = ko.LoadImage("i/img/hate.png")
	player, _ = ko.LoadImage("i/img/player.png")
}

func main() {
	err := eb.Run(update, 320, 240, 2, "aaaaaa")
	if err != nil {
		log.Fatal(err)
	}
}
