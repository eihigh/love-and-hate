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
	GameClear
	GameOver

	// from stage
	FallbackToTitle
	StageClear
	StageFailed

	// from phases
	PhaseFinished

	// from ending
	EndingFinished
)
