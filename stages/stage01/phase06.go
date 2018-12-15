package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p6 *phase06
)

type phase06 struct {
	obj.PhaseBase

	timers sio.TimersMap
	lover  obj.Mover
}

func newPhase06() *phase06 {
	p6 = &phase06{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
	}

	p6.Love = obj.Emo{
		Target: 7,
		Shown:  20,
	}
	p6.Hate = obj.Emo{
		Target: 15,
		Shown:  20,
	}
	p6.Text = "友達と遊びに行けたら、きっと楽しかったんだろうな。"
	return p6
}

func (p *phase06) Draw() {}

func (p *phase06) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 1 {
		return action.NoAction
	}
	if pt.Count > 60*16 {
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
