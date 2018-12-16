package stage01

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
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
	pt.Loop(30, func(t sio.Timer) {
		t.Do(0, 20, func(t sio.Timer) {
			if t.Count%6 == 0 {
				x := float64(t.Count) / 30 * 16
				for i := 0; i < 8; i++ {
					pos := complex(float64(i)*40+x, 240)
					o.Spawn(obj.NewLinear(pos, complex(0.2, -1.5), obj.SymbolLove))
				}
			}
		})
	})

	return action.NoAction
}
