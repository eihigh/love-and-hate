package draw

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/bitmapfont"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

var (
	Screen *ebiten.Image

	op            = &ebiten.DrawImageOptions{}
	top           = &ebiten.DrawTrianglesOptions{}
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	fface         = bitmapfont.Gothic12r
)

func init() {
	emptyImage.Fill(color.White)
}

type Group struct {
	Dst    *ebiten.Image
	Mode   ebiten.CompositeMode
	ColorM ebiten.ColorM

	geoM ebiten.GeoM
}

type OptionFn func(*Group)

func Shift(x, y float64) OptionFn {
	return func(g *Group) {
		g.geoM.Translate(x, y)
	}
}

func Scale(x, y float64) OptionFn {
	return func(g *Group) {
		g.geoM.Scale(x, y)
	}
}

func Paint(r, g, b, a float64) OptionFn {
	return func(gr *Group) {
		gr.ColorM.Scale(r, g, b, a)
	}
}

func (g *Group) DrawImage(src *ebiten.Image, fns ...OptionFn) {
	g.geoM.Reset()
	g.ColorM.Reset()

	for _, fn := range fns {
		fn(g)
	}

	op.GeoM = g.geoM
	op.ColorM = g.ColorM
	op.CompositeMode = g.Mode
	g.dst().DrawImage(src, op)
}

func (g *Group) DrawSprite(src *ebiten.Image, fns ...OptionFn) {
	g.geoM.Reset()
	g.ColorM.Reset()

	w, h := src.Size()
	g.geoM.Translate(-float64(w)/2, -float64(h)/2)
	for _, fn := range fns {
		fn(g)
	}

	op.GeoM = g.geoM
	op.ColorM = g.ColorM
	op.CompositeMode = g.Mode
	g.dst().DrawImage(src, op)
}

func (g *Group) DrawText(str string, re *sio.Rect, clr color.Color) {
	ofsX := int(sio.DefaultEmWidth / 2) // dot position
	ofsY := int(sio.DefaultEmHeight)    // ditto
	rows := sio.TextRows(str, re)
	for _, row := range rows {
		text.Draw(g.dst(), row.Text, fface, row.X+ofsX, row.Y+ofsY, clr)
	}
}

func (g *Group) DrawRect(re *sio.Rect, clr color.Color) {
	vs := corners(re, clr)
	indices := []uint16{
		0, 1, 3, 1, 3, 2,
	}

	top.CompositeMode = g.Mode
	g.dst().DrawTriangles(vs, indices, emptyImage, top)
}

func (g *Group) DrawBorder(re *sio.Rect, clr color.Color) {
	vs := corners(re, clr)
	indices := []uint16{
		0, 1, 0,
		1, 2, 1,
		2, 3, 2,
		3, 0, 3,
	}

	top.CompositeMode = g.Mode
	g.dst().DrawTriangles(vs, indices, emptyImage, top)
}

func corners(re *sio.Rect, clr color.Color) []ebiten.Vertex {
	vs := []ebiten.Vertex{
		{
			DstX: float32(re.X),
			DstY: float32(re.Y),
		},
		{
			DstX: float32(re.X + re.W),
			DstY: float32(re.Y),
		},
		{
			DstX: float32(re.X + re.W),
			DstY: float32(re.Y + re.H),
		},
		{
			DstX: float32(re.X),
			DstY: float32(re.Y + re.H),
		},
	}

	r, g, b, _ := clr.RGBA()
	for i := range vs {
		vs[i].ColorR = float32(r / 0xffff)
		vs[i].ColorG = float32(g / 0xffff)
		vs[i].ColorB = float32(b / 0xffff)
		vs[i].ColorA = 1
	}

	return vs
}

func (g *Group) dst() *ebiten.Image {
	if g.Dst == nil {
		return Screen
	}
	return g.Dst
}
