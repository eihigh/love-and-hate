package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p4 *phase04
)

type phase04 struct {
	obj.PhaseBase

	timers sio.TimersMap
	lover  obj.Mover
}

func newPhase04() *phase04 {
	p4 = &phase04{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
	}

	p4.Love = obj.Emo{
		Target: 4,
		Shown:  10,
	}
	p4.Hate = obj.Emo{
		Target: 3,
		Shown:  10,
	}
	p4.Text = "こんな私はとんだ親不孝者だ。"
	return p4
}

func (p *phase04) Draw() {}

func (p *phase04) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 1 {
		return action.NoAction
	}
	if pt.Count > 60*18 {
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
	if pt.Count%110 == 0 {
		dir := sio.UnitVector(rand.Float64())
		rot := sio.Rot(9)
		for i := 0; i < 9; i++ {
			pos := o.Player.Pos + dir*120
			o.Spawn(obj.NewLinear(pos, -dir*0.8, obj.SymbolHate))
			dir *= rot
		}
	}

	return action.NoAction
}
