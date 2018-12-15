package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p5 *phase05
)

type phase05 struct {
	obj.PhaseBase

	timers sio.TimersMap
	lights obj.Movers
}

func newPhase05() *phase05 {
	p5 = &phase05{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
		lights: obj.Movers{},
	}

	p5.Love = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p5.Hate = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p5.Text = "窓から橙色の光が差した。\nそうだ、今日はオマツリだった。"
	return p5
}

func (p *phase05) Draw() {}

func (p *phase05) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 1 {
		return action.NoAction
	}
	if pt.Count > 60*17 {
		return action.PhaseFinished
	}

	// main process
	p.lights.Update()

	// spawn
	if pt.Count%15 == 0 {
		p.lights = append(p.lights, &obj.Mover{
			Pos: complex(30+300*rand.Float64(), 250),
			Dir: sio.UnitVector(rand.Float64()*0.2 - 0.1),
		})
	}

	for _, l := range p.lights {
		if l.Timer.Count%6 == 0 {
			// spawn
			o.Spawn(obj.NewLinear(l.Pos, complex(0, -2.1)*l.Dir, obj.SymbolLove))
		}
		if l.Timer.Count > 24 {
			l.IsDead = true
		}
	}

	return action.NoAction
}
