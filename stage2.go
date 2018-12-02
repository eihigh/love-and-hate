package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	green = color.RGBA{0, 255, 0, 255}
)

type stage2 struct {
	objs *objects.Objects

	r struct {
		love, loveBar *sio.Rect
		hate, hateBar *sio.Rect
	}

	state sio.Stm
	spawn sio.Stm

	phase phase
}

type phase struct {
	message                       string
	minLoves, maxLoves, showLoves int
	minHates, maxHates, showHates int
}

const (
	phaseStart int = iota
	phaseMain
	phaseResult
)

func newStage2() *stage2 {
	s := &stage2{
		objs: objects.NewObjects(),
	}

	s.phase = phase{
		minLoves:  20,
		maxLoves:  100,
		minHates:  0,
		maxHates:  10,
		showLoves: 100,
		showHates: 100,
	}

	// TODO: add mirror method into rect
	sym := view.Clone(8, 8)
	sym.SetSize(16*2, 16)
	s.r.love = sym.Clone(4, 4).Scale(0.5, 1)
	s.r.hate = sym.Clone(6, 6).Scale(0.5, 1)

	w, _ := s.r.love.Pos(7)
	h := 16.0
	x, y := view.Pos(7)
	s.r.loveBar = sio.NewRect(7, x, y, w, h)
	s.r.hateBar = s.r.loveBar.Clone(9, 9).Move(view.Pos(9)).Drive(6).Scale(1, 0.3)
	s.r.loveBar.Drive(4).Scale(1, 0.3)
	return s
}

func (s *stage2) update() action {

	o := s.objs
	s.spawn.Update()

	if debugMode {
		// 		pl := o.Player
		// 		dmsg := fmt.Sprintf("L %d, H %d, len %d, FPS %d", pl.Loves, pl.Hates, len(o.Symbols), int(ebiten.CurrentFPS()))
		// 		ebitenutil.DebugPrint(scr, dmsg)
	}

	// ここからレベル特有の処理
	if s.spawn.HasCounted(7) {
		o.Symbols = append(o.Symbols, newUp())
		s.spawn.Reset()
	}
	// ここまでレベル特有の処理

	o.UpdatePlayer()

	for _, sym := range o.Symbols {
		sym.Update()
	}
	o.Collision(view)

	s.draw()

	return noAction
}

func (s *stage2) draw() {
	o := s.objs
	dg := &draw.Group{Dst: scr}

	pl := sprites.Sprites["player"]
	yellow := 1 - 0.8*sio.UWave(s.state.RatioTo(20))
	pl.Draw(dg, draw.Shift(c2f(o.Player.Pos)), draw.Paint(1, 1, yellow, 1))

	for _, sym := range o.Symbols {
		b := sym.Base()
		spr := sprites.HateSprite
		if b.IsLove {
			spr = sprites.LoveSprite
		}
		spr.Draw(dg, draw.Shift(c2f(b.Pos)), draw.Paint(1, 1, 1, sym.Alpha()))
	}

	// UIs
	bar := s.r.loveBar
	show := float64(s.phase.showLoves)
	ratio := float64(o.Player.Loves) / show
	ax, ay := bar.Pos(9)
	bx, by := ax-bar.W*ratio, ay
	ebitenutil.DrawLine(scr, ax, ay, bx, by, color.White)

	min := float64(s.phase.minLoves) / show
	max := float64(s.phase.maxLoves) / show
	ax, ay = bar.Pos(3)
	bx, by = ax-bar.W*max, ay
	ax -= bar.W * min
	ebitenutil.DrawLine(scr, ax, ay, bx, by, green)

	bar = s.r.hateBar
	show = float64(s.phase.showHates)
	ratio = float64(o.Player.Hates) / show
	ax, ay = bar.Pos(7)
	bx, by = ax+bar.W*ratio, ay
	ebitenutil.DrawLine(scr, ax, ay, bx, by, color.White)

	min = float64(s.phase.minHates) / show
	max = float64(s.phase.maxHates) / show
	ax, ay = bar.Pos(1)
	bx, by = ax+bar.W*max, ay
	ax += bar.W * min
	ebitenutil.DrawLine(scr, ax, ay, bx, by, green)

	sprites.LoveSprite.Draw(dg, draw.Shift(s.r.love.Pos(5)))
	sprites.HateSprite.Draw(dg, draw.Shift(s.r.hate.Pos(5)))
}

// ------------------------------------------------------------
//  Symbols
// ------------------------------------------------------------

type up struct {
	objects.SymbolBase
	vec   complex128
	state sio.Stm
}

func newUp() *up {
	u := &up{}
	u.Pos = complex(50, 200)
	u.IsLove = true
	return u
}

func (u *up) Alpha() float64 {
	return u.state.RatioTo(10)
}

func (u *up) Update() {
	u.state.Update()
	u.Pos += complex(0, -1)
}
