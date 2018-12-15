package stage01

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p2 *phase02
)

type phase02 struct {
	obj.PhaseBase

	timers sio.TimersMap
}

func newPhase02() *phase02 {
	p2 = &phase02{
		timers: sio.TimersMap{
			"phase": sio.NewTimer(""),
		},
	}

	p2.Love = obj.Emo{
		Target:     3,
		Shown:      10,
		IsPositive: true,
	}
	p2.Hate = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p2.Text = "病気の身体では、水を飲むのさえ一苦労だ。"
	return p2
}

func (p *phase02) Draw() {}

func (p *phase02) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 30 {
		return action.NoAction
	}
	if pt.Count > 60*17 {
		return action.PhaseFinished
	}
	// dg := &draw.Group{}

	// spawn
	re := env.View.Clone(8, 8).Resize(-120, +40)
	var pos1, pos2 complex128

	if pt.Count%60 == 0 {
		pos1 = re.CPos(1)
		pos2 = re.CPos(3)
	}
	if pt.Count%60 == 30 {
		pos1 = re.CPos(3)
		pos2 = re.CPos(1)
	}

	if pt.Count%30 == 0 {
		aim := sio.Normalize(o.Player.Pos-pos1) * 2
		o.Symbols = append(o.Symbols, obj.NewLinear(pos1, aim, obj.SymbolLove))
		aim = sio.Normalize(o.Player.Pos-pos2) * 2
		o.Symbols = append(o.Symbols, obj.NewLinear(pos2, aim, obj.SymbolHate))
	}

	return action.NoAction
}
