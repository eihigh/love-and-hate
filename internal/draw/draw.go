package draw

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/bitmapfont"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

var (
	op  = &ebiten.DrawImageOptions{}
	top = &ebiten.DrawTrianglesOptions{}

	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	fface         = bitmapfont.Gothic12r
)

func init() {
	emptyImage.Fill(color.White)
}

type Group struct {
	Dst *ebiten.Image
	// ColorM ebiten.ColorM
	Mode   ebiten.CompositeMode
	geoM   ebiten.GeoM
	colorM ebiten.ColorM
}

func (g *Group) Draw(src *ebiten.Image, fns ...OptionFn) {
	g.geoM.Reset()
	g.colorM.Reset()
	for _, fn := range fns {
		fn(g)
	}

	op.GeoM = g.geoM
	op.ColorM = g.colorM
	op.CompositeMode = g.Mode
	g.Dst.DrawImage(src, op)
}

func (g *Group) DrawRect(re *sio.Rect, clr color.Color) {
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

	r, gr, b, _ := clr.RGBA()
	for i := range vs {
		vs[i].ColorR = float32(r / 0xffff)
		vs[i].ColorG = float32(gr / 0xffff)
		vs[i].ColorB = float32(b / 0xffff)
		vs[i].ColorA = 1
	}

	indices := []uint16{
		0, 1, 3, 1, 3, 2,
	}

	g.Dst.DrawTriangles(vs, indices, emptyImage, top)
}

func (g *Group) Border(re *sio.Rect, clr color.Color) {
	x, y := re.Pos(7)
	w, h := re.W, re.H
	ebitenutil.DrawLine(g.Dst, x, y, x+w, y, clr)
	ebitenutil.DrawLine(g.Dst, x+w, y, x+w, y+h, clr)
	ebitenutil.DrawLine(g.Dst, x+w, y+h, x, y+h, clr)
	ebitenutil.DrawLine(g.Dst, x, y+h, x, y, clr)
}

func (g *Group) DrawText(str string, re *sio.Rect, clr color.Color) {
	ofsX := int(sio.DefaultEmWidth / 2) // dot position
	ofsY := int(sio.DefaultEmHeight)    // ditto
	rows := sio.TextRows(str, re)
	for _, row := range rows {
		text.Draw(g.Dst, row.Text, fface, row.X+ofsX, row.Y+ofsY, clr)
	}
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
		gr.colorM.Scale(r, g, b, a)
	}
}
