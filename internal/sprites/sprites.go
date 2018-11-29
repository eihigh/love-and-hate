package sprites

import (
	"fmt"
	"log"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	Sprites                = map[string]*Sprite{}
	LoveSprite, HateSprite *Sprite

	images = []string{
		"love",
		"hate",
		"player",
		"ripple",
		"cross",
	}
)

func Load() {

	for _, image := range images {
		i, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("i/img/%s.png", image), ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		Sprites[image] = NewSprite(i)
	}

	LoveSprite = Sprites["love"]
	HateSprite = Sprites["hate"]
}

type Sprite struct {
	Image *ebiten.Image
	Rect  *sio.Rect
}

func NewSprite(i *ebiten.Image) *Sprite {
	w, h := i.Size()
	return &Sprite{
		Image: i,
		Rect:  sio.NewRect(5, 0, 0, float64(w), float64(h)),
	}
}

func (s *Sprite) Bring(op *ebiten.DrawImageOptions, pos complex128) {
	op.GeoM.Translate(s.Rect.Pos(7))
	op.GeoM.Translate(real(pos), imag(pos))
}
