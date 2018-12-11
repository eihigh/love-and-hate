package obj

import "github.com/eihigh/sio"

// Mover is just a useful struct.
type Mover struct {
	Pos, Vec, Dir complex128
	Timer         sio.Timer
	IsDead        bool
}

func (m *Mover) Update() {
	m.Pos += m.Vec
	m.Timer.Update()
}

// Movers is useful type.
type Movers []*Mover

func (ms *Movers) Update() {
	ms.Cleanup()
	for _, m := range *ms {
		m.Update()
	}
}

func (ms *Movers) Cleanup() {
	next := make(Movers, 0, len(*ms))
	for _, m := range *ms {
		if !m.IsDead {
			next = append(next, m)
		}
	}
	*ms = next
}
