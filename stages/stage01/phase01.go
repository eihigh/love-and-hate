package stage01

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	p1 *phase01
)

type phase01 struct {
	obj.PhaseBase

	dir, rot complex128
}

func newPhase01() *phase01 {
	p1 = &phase01{
		dir: 0 + 1.2i,
		rot: sio.Rot(300),
	}
	p1.Love = obj.Emo{
		Target: 2,
		Shown:  10,
	}
	p1.Hate = obj.Emo{
		Target: 2,
		Shown:  10,
	}
	p1.Text = "あしたはオマツリだよ！\nねえねえ、ママみてる？"
	return p1
}

func (p *phase01) Update(o *obj.Objects) action.Action {
	// switch pw.State {
	// case "main":
	// 	p.updateMain(o)
	// 	if pw.Count > 400 {
	// 		return action.PhaseFinished
	// 	}
	//
	// default:
	// 	if pw.Count > 30 {
	// 		pw.Switch("main")
	// 	}
	// }

	return action.NoAction
}

func (p *phase01) updateMain(o *obj.Objects) {
	p.dir *= p.rot

	// pw := p.workers["phase"]
	// if pw.Count%50 < 25 {
	// 	if pw.Count%5 == 0 {
	// 		dir := p.dir
	// 		n := 8
	// 		rot := sio.Rot(float64(n))
	// 		for i := 0; i < n; i++ {
	// 			o.Symbols = append(o.Symbols, obj.NewLinear(200+200i, dir, obj.SymbolLove))
	// 			dir *= rot
	// 		}
	// 	}
	// }
	//
	// if pw.Count%30 == 0 {
	// 	aim := c.Player.Pos - (100 + 150i)
	// 	aim = sio.Normalize(aim) * 2.2
	// 	c.Symbols = append(c.Symbols, NewLinear(100+150i, aim, SymbolHate))
	// 	c.Effects = append(c.Effects, NewEffectRippleOnce(100+150i))
	// }
}

func (p *phase01) Draw() {
	// do nothing
}
