package stage01

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/audio"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p3 *phase03
)

type phase03 struct {
	obj.PhaseBase

	timers sio.TimersMap
	grv    complex128
}

func newPhase03() *phase03 {
	p3 = &phase03{
		timers: sio.TimersMap{
			"phase": &sio.Timer{},
		},
		grv: complex(0, -0.02),
	}
	p3.Love = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p3.Hate = obj.Emo{
		Target: 5,
		Shown:  10,
	}
	p3.Text = "ごめんね、何もできずに頼ってばかりで。"

	audio.PlayBgm("Retrospect")
	return p3
}

func (p *phase03) Draw() {}

func (p *phase03) Update(o *obj.Objects) action.Action {
	p.timers.UpdateAll()
	pt := p.timers["phase"]
	if pt.Count < 5 {
		return action.NoAction
	}
	if pt.Count > 60*18 {
		return action.PhaseFinished
	}

	// main process
	// dg := &draw.Group{}

	// spawn
	if pt.Count%6 == 0 {

		orig := 80 + 280i
		t := *pt
		t.Count = t.Count % 300
		t.Do(0, 150, func(t sio.Timer) {
			orig += complex(160*t.Ratio(), 0)
		})
		t.Do(150, 300, func(t sio.Timer) {
			orig += complex(160*(1-t.Ratio()), 0)
		})

		dir := sio.UnitVector(rand.Float64() * 4)
		orig += dir * 8
		f := &fallen{}
		f.Pos = orig
		f.vec = dir * 0.8
		f.Type = obj.SymbolHate
		o.Symbols = append(o.Symbols, f)
	}

	return action.NoAction
}

// Symbols
type fallen struct {
	obj.SymbolBase
	vec complex128
}

func (f *fallen) Update() {
	f.vec += p3.grv
	f.Pos += f.vec
}
