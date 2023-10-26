package game

import (
	"strings"

	"github.com/JakubC-projects/chat-gpt-game/chatgpt"
)

const defaultLvlScore = 1000

func (g *State) Answer(client *chatgpt.Client, lvl int, prompt string) error {
	g.ensureLevelSize()

	lvlDef := Levels[lvl]

	gptAnswer, err := client.Prompt(lvlDef.SystemPrompt, prompt)
	if err != nil {
		return err
	}

	if strings.Contains(gptAnswer, lvlDef.Password) {
		score := defaultLvlScore - len(prompt)
		if score > g.Levels[lvl].BestScore {
			g.Levels[lvl].BestScore = score
		}
	}

	g.Levels[lvl].Answer = gptAnswer
	g.Levels[lvl].UserPrompt = prompt

	return nil
}

func (g *State) ensureLevelSize() {
	sizeDiff := len(Levels) - len(g.Levels)
	for i := 0; i < sizeDiff; i++ {
		g.Levels = append(g.Levels, LevelState{})
	}
}
