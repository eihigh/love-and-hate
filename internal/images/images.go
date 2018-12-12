package images

import (
	"fmt"
	"image"
	"log"

	_ "image/png"

	"github.com/eihigh/love-and-hate/internal/assets"
	"github.com/hajimehoshi/ebiten"
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
		f, err := assets.Assets.Open(fmt.Sprintf("img/%s.png", name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		i, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		Images[name], _ = ebiten.NewImageFromImage(i, ebiten.FilterDefault)
	}
}
