package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p10 *phase10
)

type phase10 struct {
	obj.PhaseBase

	timers sio.TimersMap
	lover  obj.Mover
}

func newPhase10() *phase10 {
	p10 = &phase10{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
	}

	p10.Love = obj.Emo{
		Target: 4,
		Shown:  10,
	}
	p10.Hate = obj.Emo{
		Target: 3,
		Shown:  10,
	}
	p10.Text = "……目を閉じたそのとき、灯りは消えた。"
	return p10
}

func (p *phase10) Draw() {}

func (p *phase10) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 1 {
		return action.NoAction
	}
	if pt.Count > 400 {
		return action.PhaseFinished
	}

	// main process
	dg := &draw.Group{}
	p.lover.Update()
	p.lover.Pos = o.Player.Pos

	// draw ripple
	e := obj.EffectBase{
		Type:  obj.EffectRipple,
		Pos:   p.lover.Pos,
		Timer: *pt,
	}
	e.Draw(dg)

	// spawn
	if pt.Count%80 == 0 {
		dir := sio.UnitVector(rand.Float64())
		rot := sio.Rot(9)
		for i := 0; i < 9; i++ {
			pos := p.lover.Pos + dir*10
			o.Symbols = append(o.Symbols, obj.NewLinear(pos, dir*1.3, obj.SymbolLove))
			dir *= rot
		}
	}

	return action.NoAction
}
