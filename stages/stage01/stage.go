package stage01

import "github.com/eihigh/love-and-hate/internal/obj"

func NewPhases() []obj.Phase {
	ps := []obj.Phase{}
	ps = append(ps, newPhase01())
	ps = append(ps, newPhase02())
	ps = append(ps, newPhase03())
	ps = append(ps, newPhase04())
	ps = append(ps, newPhase05())
	ps = append(ps, newPhase06())
	ps = append(ps, newPhase07())
	ps = append(ps, newPhase08())
	ps = append(ps, newPhase09())
	ps = append(ps, newPhase10())
	return ps
}
