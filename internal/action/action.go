package action

type Action int

// Action definitions
const (
	NoAction Action = iota

	// from title
	NewGame
	HowTo

	// from play
	PlayContinue

	// from phases
	PhaseFinished
)
