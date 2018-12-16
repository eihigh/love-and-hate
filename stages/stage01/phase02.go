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
	re := env.View.Clone(8, 8).Resize(-120, +20)
	var pos complex128
	if pt.Count%120 == 0 {
		pos = re.CPos(1)
	}
	if pt.Count%120 == 60 {
		pos = re.CPos(3)
	}

	if pt.Count%60 == 0 {
		dir := o.AimFrom(pos) * 1.2
		rot := sio.Rot(22)
		o.Spawn(obj.NewLinear(pos, dir/rot/rot, obj.SymbolHate))
		o.Spawn(obj.NewLinear(pos, dir/rot, obj.SymbolLove))
		o.Spawn(obj.NewLinear(pos, dir, obj.SymbolHate))
		o.Spawn(obj.NewLinear(pos, dir*rot, obj.SymbolLove))
		o.Spawn(obj.NewLinear(pos, dir*rot*rot, obj.SymbolHate))
	}

	return action.NoAction
}
