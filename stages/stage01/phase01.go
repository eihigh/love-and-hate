package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p1 *phase01
)

type phase01 struct {
	obj.PhaseBase

	timers   sio.TimersMap
	cyclones obj.Movers
}

type cyclone struct {
	pos, vec, dir complex128
	timer         sio.Timer
}

func newPhase01() *phase01 {
	p1 = &phase01{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
		cyclones: obj.Movers{},
	}
	p1.Love = obj.Emo{
		Target:     2,
		Shown:      10,
		IsPositive: true,
	}
	p1.Hate = obj.Emo{
		Target: 2,
		Shown:  10,
	}
	p1.Text = ""
	return p1
}

func (p *phase01) Update(o *obj.Objects) action.Action {
	p.timers.UpdateAll()

	pt := p.timers["phase"]
	if pt.Count > 30 {
		p.updateMain(o)
	}
	if pt.Count > 800 {
		return action.PhaseFinished
	}

	return action.NoAction
}

func (p *phase01) updateMain(o *obj.Objects) {

	dg := &draw.Group{}
	p.cyclones.Update()
	pt := p.timers["phase"]

	if pt.Count%90 == 0 {
		p.cyclones = append(p.cyclones, &obj.Mover{
			Pos: complex(30+300*rand.Float64(), 250),
			Vec: complex(0, -1),
			Dir: sio.UnitVector(rand.Float64()) * 1.2,
		})
	}

	for _, c := range p.cyclones {
		e := obj.EffectBase{
			Type:  obj.EffectRipple,
			Pos:   c.Pos,
			Timer: c.Timer,
		}
		e.Draw(dg)

		c.Dir *= sio.Rot(200)

		then := c.Timer.Do(20, 120, func(t sio.Timer) {
			if t.Count%2 != 0 {
				return
			}

			dir := c.Dir
			rot := sio.Rot(8)
			for i := 0; i < 8; i++ {
				t.Do(0, 25, func(sio.Timer) {
					o.Symbols = append(o.Symbols, obj.NewLinear(c.Pos, dir, obj.SymbolLove))
				})
				t.Do(50, 75, func(sio.Timer) {
					o.Symbols = append(o.Symbols, obj.NewLinear(c.Pos, dir, obj.SymbolHate))
				})
				dir *= rot
			}
		})

		then.Once(func() {
			c.IsDead = true
		})
	}

}

func (p *phase01) Draw() {
	// do nothing
}
