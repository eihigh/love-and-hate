package stage01

import "github.com/eihigh/love-and-hate/internal/obj"

func NewPhases() []obj.Phase {
	ps := []obj.Phase{}
	ps = append(ps, newPhase01())
	// ps = append(ps, newPhase02())
	// ps = append(ps, newPhase03())
	// ps = append(ps, newPhase04())
	// ps = append(ps, newPhase05())
	// ps = append(ps, newPhase06())
	// ps = append(ps, newPhase07())
	return ps
}
