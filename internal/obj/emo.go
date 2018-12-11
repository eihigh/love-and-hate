package obj

import "image/color"

type Emo struct {
	Target     int
	Shown      int
	IsPositive bool
}

func (e *Emo) IsOver(current int) bool {
	if e.IsPositive {
		return false
	}
	return e.Target <= current
}

func (e *Emo) IsPoor(current int) bool {
	if !e.IsPositive {
		return false
	}
	return e.Target > current
}

func (e *Emo) IsOk(current int) bool {
	if e.IsPositive {
		return current >= e.Target
	}
	return current < e.Target
}

func (e *Emo) Colors() (back, front color.Color) {
	if e.IsPositive {
		return Red, White
	}
	return White, Red
}

func (e *Emo) Ratios(current int) (back, front float64) {
	s := float64(e.Shown)
	back = float64(e.Target) / s
	front = float64(current) / s
	return back, front
}
