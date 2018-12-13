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
	p1 *phase01
)

type phase01 struct {
	obj.PhaseBase

	timers sio.TimersMap
	lovers obj.Movers
}

func newPhase01() *phase01 {
	p1 = &phase01{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
		lovers: obj.Movers{},
	}

	p1.Love = obj.Emo{
		Target:     3,
		Shown:      10,
		IsPositive: true,
	}
	p1.Hate = obj.Emo{
		Target: 3,
		Shown:  10,
	}
	p1.Text = "襲いかかる吐き気に顔をしかめた。\n\"なんで自分だけ？\"って。"
	return p1
}

func (p *phase01) Draw() {
	// do nothing
}

func (p *phase01) Update(o *obj.Objects) action.Action {

	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 30 {
		return action.NoAction
	}
	if pt.Count > 60*10 {
		return action.PhaseFinished
	}

	dg := &draw.Group{}

	// update
	p.lovers.Update()

	for _, l := range p.lovers {
		e := obj.EffectBase{
			Type:  obj.EffectRipple,
			Pos:   l.Pos,
			Timer: l.Timer,
		}
		e.Draw(dg)

		if l.Timer.Count > 80 {
			// shoot
			dir := sio.UnitVector(rand.Float64()*4) * 1.3
			rot := sio.Rot(4)
			for i := 0; i < 4; i++ {
				o.Symbols = append(o.Symbols, obj.NewLinear(l.Pos, dir, obj.SymbolLove))
				dir *= rot
			}
			l.IsDead = true
		}
	}

	// spawn
	if pt.Count%140 == 0 {
		x, y := env.View.Clone(5, 5).Resize(-60, -40).RandPos()
		p.lovers = append(p.lovers, &obj.Mover{
			Pos: complex(x, y),
		})
	}

	return action.NoAction
}
