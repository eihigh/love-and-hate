package images

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	Images = map[string]*ebiten.Image{}
)

func Load() {
	for _, name := range []string{
		"love",
		"hate",
		"player",
		"ripple",
		"cross",
	} {
		i, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("i/img/%s.png", name), ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		Images[name] = i
	}
}
