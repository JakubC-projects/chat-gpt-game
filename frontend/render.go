package frontend

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/JakubC-projects/chat-gpt-game/game"
)

var templates = template.Must(template.ParseGlob("templates/*.*"))

func Render(w io.Writer, s *game.State, lvl int) error {
	tmplData := gameToTmplData(s, lvl)
	return templates.ExecuteTemplate(w, "index.gohtml", tmplData)
}

func gameToTmplData(s *game.State, lvl int) templatePayload {
	var lvlData game.LevelState
	if len(s.Levels) > lvl {
		lvlData = s.Levels[lvl]
	}

	return templatePayload{
		TotalScore:            s.Score(),
		Levels:                getTmplLevels(s.MaxLevel()),
		OpenLevel:             lvl,
		LevelScore:            lvlData.BestScore,
		SanitizedSystemPrompt: getSanitizedPrompt(lvl),
		UserPrompt:            lvlData.UserPrompt,
		GptAnswer:             lvlData.Answer,
		NextLevel:             getNextLevel(lvl, s.MaxLevel()),
	}
}

func getTmplLevels(maxLvl int) []templateLevel {
	res := make([]templateLevel, len(game.Levels))

	for i := 0; i < len(game.Levels); i++ {
		res[i].Name = fmt.Sprintf("Level %d", i)
		res[i].Available = i <= maxLvl
	}
	return res
}

func getSanitizedPrompt(lvl int) string {
	levelData := game.Levels[lvl]

	return strings.ReplaceAll(
		levelData.SystemPrompt,
		levelData.Password,
		strings.Repeat("*", len(levelData.Password)))
}

func getNextLevel(lvl, maxLvl int) string {
	if maxLvl > lvl {
		return fmt.Sprintf("/%d", lvl+1)
	}
	return ""
}
