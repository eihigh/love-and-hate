package obj

import (
	"github.com/eihigh/love-and-hate/internal/action"
)

type Phase interface {
	Base() *PhaseBase
	Update(*Objects) action.Action
	Draw()
}

type PhaseBase struct {
	Love, Hate Emo
	Text       string
}

func (p *PhaseBase) Base() *PhaseBase {
	return p
}
