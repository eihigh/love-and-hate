package stage01

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p1 *phase01
)

type phase01 struct {
	obj.PhaseBase

	timers sio.TimersMap
}

func newPhase01() *phase01 {
	p1 = &phase01{
		timers: sio.TimersMap{
			"phase": sio.NewTimer(""),
		},
	}

	p1.Love = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p1.Hate = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p1.Text = "目が覚めて、重たい身体をゆっくり起こす。"
	return p1
}

func (p *phase01) Draw() {
	// do nothing
}

func (p *phase01) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 60 {
		return action.NoAction
	}
	if pt.Count > 60*16 {
		return action.PhaseFinished
	}
	// dg := &draw.Group{}

	// spawn
	re := env.View.Clone(8, 8).Resize(-120, +40)
	pos1 := re.CPos(1)
	pos2 := re.CPos(3)

	rot := sio.Rot(20)
	if pt.Count%120 == 0 {
		// spawn from pos1
		aim := sio.Normalize(o.Player.Pos - pos1)
		o.Symbols = append(o.Symbols, obj.NewLinear(pos1, aim, obj.SymbolHate))
		o.Symbols = append(o.Symbols, obj.NewLinear(pos1, aim*rot, obj.SymbolHate))
		o.Symbols = append(o.Symbols, obj.NewLinear(pos1, aim/rot, obj.SymbolHate))
	}
	if pt.Count%120 == 60 {
		// spawn from pos2
		aim := sio.Normalize(o.Player.Pos - pos2)
		o.Symbols = append(o.Symbols, obj.NewLinear(pos2, aim, obj.SymbolHate))
		o.Symbols = append(o.Symbols, obj.NewLinear(pos2, aim*rot, obj.SymbolHate))
		o.Symbols = append(o.Symbols, obj.NewLinear(pos2, aim/rot, obj.SymbolHate))
	}

	return action.NoAction
}
