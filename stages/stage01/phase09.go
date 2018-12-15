package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p9 *phase09
)

type phase09 struct {
	obj.PhaseBase

	timers sio.TimersMap
	lover  obj.Mover
}

func newPhase09() *phase09 {
	p9 = &phase09{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
	}

	p9.Love = obj.Emo{
		Target: 4,
		Shown:  10,
	}
	p9.Hate = obj.Emo{
		Target: 3,
		Shown:  10,
	}
	p9.Text = "近くて遠い橙色の灯りが、今はただ辛かった。"
	return p9
}

func (p *phase09) Draw() {}

func (p *phase09) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 1 {
		return action.NoAction
	}
	if pt.Count > 60*15 {
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
