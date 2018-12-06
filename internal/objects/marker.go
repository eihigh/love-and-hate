package objects

type Marker struct {
	Pos, Vec complex128
	Count    int
	IsDead   bool
}

func (m *Marker) Update() {
	m.Pos += m.Vec
	m.Count++
}
