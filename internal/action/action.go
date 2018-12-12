package action

type Action int

// Action definitions
const (
	NoAction Action = iota
	Cancel

	// from title
	ToStages
	ToHowTo
	NewGame
	HowTo

	// from play
	PlayContinue

	// from phases
	PhaseFinished
)
