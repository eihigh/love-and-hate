package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
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
		Target: 10,
		Shown:  15,
	}
	p5.Hate = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p5.Text = "窓から橙色の光が差した。\nそうだ、今日はオマツリだった。"
	return p5
}

func (p *phase05) Draw() {
	dg := &draw.Group{}
	for _, l := range p.lights {
		e := obj.EffectBase{
			Type:  obj.EffectRipple,
			Pos:   l.Pos,
			Timer: l.Timer,
		}
		e.Draw(dg)
	}
}

func (p *phase05) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 20 {
		return action.NoAction
	}
	if pt.Count > 60*17 {
		return action.PhaseFinished
	}

	// main process
	p.lights.Update()

	// spawn
	if pt.Count%40 == 0 {
		re := env.View.Clone(2, 2).Scale(1, 0.1)
		x, y := re.RandPos()

		p.lights = append(p.lights, &obj.Mover{
			Pos: complex(x, y),
			Dir: sio.UnitVector(rand.Float64()),
		})
	}

	for _, l := range p.lights {
		l.Timer.Loop(40, func(t sio.Timer) {
			if t.Count < 20 && t.Count%4 == 0 {
				rot := sio.Rot(7)
				dir := l.Dir
				for i := 0; i < 7; i++ {
					o.Spawn(obj.NewLinear(l.Pos, dir, obj.SymbolLove))
					dir *= rot
				}
			}
		})

		if l.Timer.Count > 200 {
			l.IsDead = true
		}
	}

	return action.NoAction
}
