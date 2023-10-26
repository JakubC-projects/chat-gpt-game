package frontend

type templatePayload struct {
	TotalScore            int
	Levels                []templateLevel
	OpenLevel             int
	LevelScore            int
	SanitizedSystemPrompt string
	UserPrompt            string
	GptAnswer             string
	NextLevel             string
}

type templateLevel struct {
	Name      string
	Available bool
}
