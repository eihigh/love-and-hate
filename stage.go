package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/images"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
	"github.com/fogleman/ease"
)

var (
	phaseLists = map[int]func() []obj.Phase{}
)

type stage struct {
	objs       *obj.Objects
	timers     sio.TimersMap
	phases     []obj.Phase
	phaseIndex int

	result struct {
		isLoveOk, isHateOk bool
	}

	loveIcon, loveBar *sio.Rect
	hateIcon, hateBar *sio.Rect
	phaseText         *sio.Rect
}

func newStage(level int) *stage {

	bars := env.View.Clone(8, 8).SetSize(-1, 16).Shift(0, 4)
	// love
	lb := bars.Clone(4, 4).Scale(0.5, 1)
	hb := bars.Clone(6, 6).Scale(0.5, 1)

	return &stage{
		objs: &obj.Objects{},
		timers: sio.TimersMap{
			"ui":    &sio.Timer{},
			"stage": &sio.Timer{},
		},
		phases:    phaseLists[level](),
		loveIcon:  lb.Clone(6, 6).SetSize(16, 16),
		hateIcon:  hb.Clone(4, 4).SetSize(16, 16),
		loveBar:   lb.Clone(4, 4).Resize(-16, -10),
		hateBar:   hb.Clone(6, 6).Resize(-16, -10),
		phaseText: env.View.Clone(5, 5).Shift(0, -20),
	}
}

func (s *stage) currentPhase() obj.Phase {
	if s.phaseIndex >= len(s.phases) {
		return nil
	}
	if s.phaseIndex < 0 {
		return nil
	}
	return s.phases[s.phaseIndex]
}

func (s *stage) update() action.Action {
	s.timers.UpdateAll()

	st := s.timers["stage"]
	switch st.State {
	case "failed":
		s.updateFailed()
	case "cleared":
		s.updateCleared()
	default:
		s.updateMain()
	}
	return action.NoAction
}

func (s *stage) updateMain() {
	ph := s.currentPhase()
	act := ph.Update(s.objs)
	s.updateObjects()
	ph.Draw()
	s.updateUI()
	s.drawPhaseText()
	s.collision()
	if act == action.PhaseFinished {
		s.succPhase()
	}
}

func (s *stage) updateFailed() {
	s.updateObjects()
	s.updateUI()
	// s.drawFailed()
}

func (s *stage) updateCleared() {
	s.updateObjects()
	s.updateUI()
	// s.drawCleared()
}

func (s *stage) succPhase() {
	ut := s.timers["ui"]
	ut.Switch("result")

	// vanish all objects
	o := s.objs
	for _, sym := range o.Symbols {
		o.AppendEffect(obj.EffectRippleOnce, sym.Base().Pos)
	}
	o.Symbols = []obj.Symbol{}

	// check results
	st := s.timers["stage"]
	pb := s.currentPhase().Base()
	s.result.isLoveOk = pb.Love.IsOk(o.Player.Loves)
	s.result.isHateOk = pb.Hate.IsOk(o.Player.Hates)

	// 結果分岐
	if s.result.isLoveOk && s.result.isHateOk {
		s.phaseIndex++
		if s.phaseIndex >= len(s.phases) {
			st.Switch("cleared")
			return
		}
		st.Switch("") // reset count
		return
	}

	// failed
	st.Switch("failed")
	s.phaseIndex = -1
}

func (s *stage) updateObjects() {

	o := s.objs
	st := s.timers["stage"]
	dg := &draw.Group{}

	// play sound
	// if c.LastLenSymbols != len(c.Symbols) {
	// 	audio.PlaySe("paper")
	// }

	// update & draw symbols
	for _, sym := range o.Symbols {
		sym.Update()
		b := sym.Base()
		b.Timer.Limit = obj.BabyTime
		b.Timer.Update()

		i := images.Images["love"]
		if b.Type == obj.SymbolHate {
			i = images.Images["hate"]
		}

		alpha := b.Timer.Ratio()
		dg.DrawSprite(
			i,
			draw.Shift(sio.Ctof(b.Pos)),
			draw.Paint(1, 1, 1, alpha),
		)
	}

	// draw & update player
	s.updatePlayer()
	yellow := 0.8 * sio.UWave(st.RatioTo(20))
	dg.DrawSprite(
		images.Images["player"],
		draw.Shift(sio.Ctof(o.Player.Pos)),
		draw.Paint(1, 1, 1-yellow, 1),
	)

	// update & draw effects
	for _, e := range o.Effects {
		e.Update()
		b := e.Base()
		b.Timer.Update()
		b.Draw(dg)
	}

	// clean dead effects
	next := make([]obj.Effect, 0, len(o.Effects))
	for _, e := range o.Effects {
		if !e.Base().IsDead {
			next = append(next, e)
		}
	}
	o.Effects = next
}

func (s *stage) updatePlayer() {}

func (s *stage) updateUI() {}

func (s *stage) updateResultUI() {

	dg := &draw.Group{}
	ut := s.timers["ui"]

	then := ut.Do(0, 140, func(t sio.Timer) {
		if t.Count%30 < 20 {
			re := s.loveIcon.Clone(4, 6)
			if s.result.isLoveOk {
				dg.DrawText("OK ", re, obj.White)
			} else {
				dg.DrawText("NG ", re, obj.Red)
			}

			re = s.hateIcon.Clone(6, 4)
			if s.result.isHateOk {
				dg.DrawText(" OK", re, obj.White)
			} else {
				dg.DrawText(" NG", re, obj.Red)
			}
		}
	})

	then.Once(func() {
		phase := s.currentPhase()
		if phase != nil {
			ut.Switch("")
		}
	})

	// draw icons
	dg.DrawSprite(
		images.Images["love"],
		draw.Shift(s.loveIcon.Pos(5)),
	)
	dg.DrawSprite(
		images.Images["hate"],
		draw.Shift(s.hateIcon.Pos(5)),
	)
}

func (s *stage) collision() {

	o := s.objs
	for _, sym := range o.Symbols {
		b := sym.Base()
		if b.IsDead {
			continue
		}
		if !env.PlayArea.Contains(b.Pos) {
			b.IsDead = true
			continue
		}
		if b.Timer.Count < obj.BabyTime {
			continue
		}
	}

	// clean dead objects
	next := make([]obj.Symbol, 0, len(o.Symbols))
	for _, sym := range o.Symbols {
		if !sym.Base().IsDead {
			next = append(next, sym)
		}
	}
	o.Symbols = next
}

func (s *stage) drawPhaseText() {
	pb := s.currentPhase().Base()
	st := s.timers["stage"]
	dg := &draw.Group{}

	st.Do(0, 240, func(t sio.Timer) {
		y := 0.0
		alpha := 1.0
		then := st.Do(0, 80, func(t sio.Timer) {
			r := t.Ratio()
			y = 30 * (1 - ease.OutQuad(r))
			alpha = r
		})

		then.Do(80, 160, func(t sio.Timer) {
			r := t.Ratio()
			y = -30 * ease.InQuad(r)
			alpha = 1 - r
		})

		box := s.phaseText.Clone(5, 5).Shift(0, y)
		dg.DrawText(pb.Text, box, color.RGBA{255, 255, 255, uint8(255 * alpha)})
	})
}
