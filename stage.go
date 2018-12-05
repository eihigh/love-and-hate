package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/love-and-hate/internal/text"
	"github.com/eihigh/sio"
	"github.com/fogleman/ease"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func c2f(c complex128) (float64, float64) {
	return real(c), imag(c)
}

var (
	red   = color.RGBA{255, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
)

const (
	phaseTitleIn int = iota
	phaseTitleSus
	phaseTitleOut
	phaseBody
	phaseResult
)

type stage struct {
	objs *objects.Objects

	loveIcon, loveBar *sio.Rect
	hateIcon, hateBar *sio.Rect
	message           *sio.Rect

	states  map[string]*sio.Stm
	workers map[string]*sio.Worker

	phases     []phase
	phaseIndex int
}

type phase interface {
	update(*stage)
	base() *phaseBase
}

func newStage() *stage {
	s := &stage{}
	s.objs = objects.NewObjects()

	s.message = view.Clone(8, 8).Scale(1, 0.7)
	s.loveIcon = &sio.Rect{
		X: 160 - 16,
		Y: 4,
		W: 16,
		H: 16,
	}
	s.hateIcon = &sio.Rect{
		X: 160,
		Y: 4,
		W: 16,
		H: 16,
	}
	s.loveBar = s.loveIcon.Clone(4, 6).SetSize(160-16, 6)
	s.hateBar = s.hateIcon.Clone(6, 4).SetSize(160-16, 6)

	s.phases = append(s.phases, newPhase1())

	s.states = map[string]*sio.Stm{
		"stage": &sio.Stm{},
		"phase": &sio.Stm{},
	}

	s.workers = map[string]*sio.Worker{
		"phase": &sio.Worker{
			State: "begin",
		},
		"ui": &sio.Worker{
			State: "begin",
		},
	}
	return s
}

func (s *stage) update() action {
	o := s.objs
	for _, st := range s.states {
		st.Update()
	}

	for _, w := range s.workers {
		w.Count++
	}

	pw := s.workers["phase"]
	switch pw.State {
	case "begin":
		if pw.Count > 40 {
			pw.Switch("text")
		}

	case "text":
		pw.Do(0, 300, func(t float64) {
			mes := s.phases[s.phaseIndex].base().message
			l := 50.0
			x, y := 160.0, 100.0
			alpha := 1.0

			pw.Do(0, 100, func(t float64) {
				y -= l * (1 - ease.OutQuad(t))
				alpha = t
			})

			pw.Do(200, 300, func(t float64) {
				y += l * ease.InQuad(t)
				alpha = 1 - t
			})

			box := s.message.Clone(5, 5).Shift(x, y)
			tb := box.NewTextBox(mes, 5)
			text.Draw(scr, tb, color.Alpha{uint8(255 * alpha)})
		})
		if pw.Count >= 300 {
			pw.Switch("main")
		}

	case "main":
	case "result":
	}

	ph := s.phases[s.phaseIndex]
	ph.update(s)

	o.UpdatePlayer()

	for _, sym := range o.Symbols {
		sym.Update()
	}
	o.Collision(view)

	s.draw()

	return noAction
}

func (s *stage) draw() {
	o := s.objs
	dg := &draw.Group{Dst: scr}

	pl := sprites.Sprites["player"]
	yellow := 1 - 0.8*sio.UWave(s.states["stage"].RatioTo(20))
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
	pb := s.phases[s.phaseIndex].base()
	uw := s.workers["ui"]
	scale := 1.0
	if uw.State == "begin" {
		uw.Do(0, 60, func(t float64) {
			scale = ease.OutQuad(t)
		})
	}

	// draw love
	bk, fr := pb.love.ratios(o.Player.Loves)
	bc, fc := pb.love.colors()
	bar := s.loveBar.Clone(6, 6).Scale(bk*scale, 1)
	dg.DrawRect(bar, bc)
	bar = s.loveBar.Clone(6, 6).Scale(fr, 1)
	dg.DrawRect(bar, fc)

	// draw hate
	bk, fr = pb.hate.ratios(o.Player.Hates)
	bc, fc = pb.hate.colors()
	bar = s.hateBar.Clone(4, 4).Scale(bk*scale, 1)
	dg.DrawRect(bar, bc)
	bar = s.hateBar.Clone(4, 4).Scale(fr, 1)
	dg.DrawRect(bar, fc)

	alpha := 1.0
	if pb.love.isPoor(o.Player.Loves) {
		alpha = 1 - 0.5*sio.UWave(s.states["stage"].RatioTo(40))
	}
	sprites.LoveSprite.Draw(dg, draw.Shift(s.loveIcon.Pos(5)), draw.Paint(1, 1, 1, alpha))
	if pb.love.isOver(o.Player.Loves) {
		ax, ay := s.loveIcon.Pos(7)
		bx, by := s.loveIcon.Pos(3)
		ebitenutil.DrawLine(scr, ax, ay, bx, by, red)
		ax, ay = s.loveIcon.Pos(9)
		bx, by = s.loveIcon.Pos(1)
		ebitenutil.DrawLine(scr, ax, ay, bx, by, red)
	}

	alpha = 1.0
	if pb.hate.isPoor(o.Player.Hates) {
		alpha = 1 - 0.5*sio.UWave(s.states["stage"].RatioTo(40))
	}
	sprites.HateSprite.Draw(dg, draw.Shift(s.hateIcon.Pos(5)), draw.Paint(1, 1, 1, alpha))
	if pb.hate.isOver(o.Player.Hates) {
		ax, ay := s.hateIcon.Pos(7)
		bx, by := s.hateIcon.Pos(3)
		ebitenutil.DrawLine(scr, ax, ay, bx, by, red)
		ax, ay = s.hateIcon.Pos(9)
		bx, by = s.hateIcon.Pos(1)
		ebitenutil.DrawLine(scr, ax, ay, bx, by, red)
	}

	bd(s.loveIcon)
	bd(s.hateIcon)
}

type emo struct {
	target     int
	shown      int
	isPositive bool
}

func (e *emo) isOver(current int) bool {
	if e.isPositive {
		return false
	}
	return e.target <= current
}

func (e *emo) isPoor(current int) bool {
	if !e.isPositive {
		return false
	}
	return e.target > current
}

func (e *emo) colors() (back, front color.Color) {
	if e.isPositive {
		return red, white
	}
	return white, red
}

func (e *emo) ratios(current int) (back, front float64) {
	s := float64(e.shown)
	back = float64(e.target) / s
	front = float64(current) / s
	return back, front
}
