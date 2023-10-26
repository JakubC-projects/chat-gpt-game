package game

type State struct {
	Levels []LevelState
}

type LevelState struct {
	UserPrompt string
	BestScore  int
	Answer     string
}

func (s State) MaxLevel() int {
	for i, l := range s.Levels {
		if !l.IsSolved() {
			return i
		}
	}
	return len(s.Levels)
}

func (s State) Score() int {
	total := 0

	for _, l := range s.Levels {
		total += l.BestScore
	}

	return total
}

func (l LevelState) IsSolved() bool {
	return l.BestScore > 0
}
