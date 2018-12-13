package action

type Action int

// Action definitions
const (
	NoAction Action = iota
	Cancel

	// from title
	StartPlay

	// from play
	PlayContinue
	BackToTitle

	// from stage
	FallbackToTitle

	// from phases
	PhaseFinished
)
