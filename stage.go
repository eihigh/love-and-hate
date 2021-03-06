package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/audio"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/images"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/love-and-hate/stages/stage01"
	"github.com/eihigh/sio"
	"github.com/fogleman/ease"
)

var (
	phaseLists = map[int]func() []obj.Phase{
		1: stage01.NewPhases,
	}
)

type stage struct {
	timers     sio.TimersMap
	phases     []obj.Phase
	phaseIndex int

	result struct {
		isLoveOk, isHateOk bool
	}

	loveIcon, loveBar *sio.Rect
	hateIcon, hateBar *sio.Rect
	phaseText         *sio.Rect
	resultText        *sio.Rect
}

func newStage(level int) *stage {

	// play bgm
	audio.PlayBgm(bgms[level])

	bars := env.View.Clone(8, 8).SetSize(-1, 16).Shift(0, 4)
	// love
	lb := bars.Clone(4, 4).Scale(0.5, 1)
	hb := bars.Clone(6, 6).Scale(0.5, 1)

	return &stage{
		timers: sio.TimersMap{
			"ui":    &sio.Timer{},
			"stage": &sio.Timer{},
		},
		phases:     phaseLists[level](),
		loveIcon:   lb.Clone(6, 6).SetSize(16, 16),
		hateIcon:   hb.Clone(4, 4).SetSize(16, 16),
		loveBar:    lb.Clone(4, 4).Resize(-16, -10),
		hateBar:    hb.Clone(6, 6).Resize(-16, -10),
		phaseText:  env.View.Clone(5, 5).Shift(0, -20),
		resultText: env.View.Clone(5, 5).Shift(0, -20),
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
		if st.Count > 240 {
			return action.StageFailed
		}

	case "cleared":
		s.updateCleared()
		if st.Count > 120 && input.JustDecided() {
			return action.StageClear
		}

	default:
		s.updateMain()
	}
	return action.NoAction
}

func (s *stage) updateMain() {
	ph := s.currentPhase()
	act := ph.Update(o)
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

	s.timers["stage"].Do(0, 180, func(sio.Timer) {
		dg := &draw.Group{}
		dg.DrawText("LOST", s.resultText, obj.White)
		s.updateObjects()
		s.updateUI()
	})
}

func (s *stage) updateCleared() {
	s.updateObjects()
	s.updateUI()

	dg := &draw.Group{}
	dg.DrawText("STAGE CLEARED", s.resultText, obj.White)
}

func (s *stage) succPhase() {
	ut := s.timers["ui"]
	ut.Switch("result")

	// vanish all objects
	for _, sym := range o.Symbols {
		o.AppendEffect(obj.EffectRippleOnce, sym.Base().Pos)
	}
	o.Symbols = []obj.Symbol{}

	// check results
	st := s.timers["stage"]
	pb := s.currentPhase().Base()
	s.result.isLoveOk = pb.Love.IsOk(o.Player.Loves)
	s.result.isHateOk = pb.Hate.IsOk(o.Player.Hates)

	// reset count
	o.Player.Loves = 0
	o.Player.Hates = 0

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

	st := s.timers["stage"]
	dg := &draw.Group{}

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

func (s *stage) updateUI() {

	ut := s.timers["ui"]
	if ut.State == "result" {
		s.updateResultUI()
		return
	}

	dg := &draw.Group{}
	phase := s.currentPhase()
	pb := phase.Base()
	pl := &o.Player

	// draw bars
	scale := 1.0
	ut.Do(0, 60, func(t sio.Timer) {
		scale = ease.OutQuad(t.Ratio())
	})

	// love
	bk, fr := pb.Love.Ratios(pl.Loves)
	bc, fc := pb.Love.Colors()
	bar := s.loveBar.Clone(6, 6).Scale(bk*scale, 1)
	dg.DrawRect(bar, bc)
	bar = s.loveBar.Clone(6, 6).Scale(fr, 1)
	dg.DrawRect(bar, fc)

	// hate
	bk, fr = pb.Hate.Ratios(pl.Hates)
	bc, fc = pb.Hate.Colors()
	bar = s.hateBar.Clone(4, 4).Scale(bk*scale, 1)
	dg.DrawRect(bar, bc)
	bar = s.hateBar.Clone(4, 4).Scale(fr, 1)
	dg.DrawRect(bar, fc)

	// love icon
	alpha := 1.0
	if pb.Love.IsPoor(pl.Loves) {
		alpha = 1 - 0.5*sio.UWave(ut.RatioTo(40))
	}
	dg.DrawSprite(
		images.Images["love"],
		draw.Shift(s.loveIcon.Pos(5)),
		draw.Paint(1, 1, 1, alpha),
	)

	// hate icon
	alpha = 1.0
	if pb.Hate.IsPoor(pl.Hates) {
		alpha = 1 - 0.5*sio.UWave(ut.RatioTo(40))
	}
	dg.DrawSprite(
		images.Images["hate"],
		draw.Shift(s.hateIcon.Pos(5)),
		draw.Paint(1, 1, 1, alpha),
	)
}

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
		if s.currentPhase() != nil {
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

	pb := s.currentPhase().Base()
	pl := &o.Player
	pl.LastLoves = pl.Loves
	pl.LastHates = pl.Hates

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

		positive := false
		if pb.Love.IsPositive && b.Type == obj.SymbolLove {
			positive = true
		}
		if pb.Hate.IsPositive && b.Type == obj.SymbolHate {
			positive = true
		}

		limit := 6.0
		if positive {
			limit = 12.0
		}

		diff := sio.AbsSq(pl.Pos - b.Pos)
		if diff < limit*limit {
			b.IsDead = true
			if b.Type == obj.SymbolLove {
				pl.Loves++
			}
			if b.Type == obj.SymbolHate {
				pl.Hates++
			}
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

	// play sounds
	if (pl.Loves != pl.LastLoves && !pb.Love.IsPositive) || (pl.Hates != pl.LastHates && !pb.Hate.IsPositive) {
		audio.PlaySe("buzzer")
	}
	if (pl.Loves != pl.LastLoves && pb.Love.IsPositive) || (pl.Hates != pl.LastHates && pb.Hate.IsPositive) {
		audio.PlaySe("drop")
	}
}

func (s *stage) drawPhaseText() {
	pb := s.currentPhase().Base()
	st := s.timers["stage"]
	dg := &draw.Group{}

	st.Do(0, 240, func(t sio.Timer) {
		y := 0.0
		alpha := 1.0
		t.Do(0, 80, func(t sio.Timer) {
			r := t.Ratio()
			y = 30 * (1 - ease.OutQuad(r))
			alpha = r
		})

		t.Do(-80, 0, func(t sio.Timer) {
			r := t.Ratio()
			y = -30 * ease.InQuad(r)
			alpha = 1 - r
		})

		box := s.phaseText.Clone(5, 5).Shift(0, y)
		dg.DrawText(pb.Text, box, color.RGBA{255, 255, 255, uint8(255 * alpha)})
	})
}
