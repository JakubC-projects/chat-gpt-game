package game

import (
	"net/http"
	"strconv"
	"strings"
)

func CurrentLevel(r *http.Request, maxLvl int) int {
	path, _ := strings.CutPrefix(r.URL.Path, "/")

	lvl, err := strconv.Atoi(path)
	if err != nil {
		return 0
	}
	if lvl > maxLvl {
		return maxLvl
	}
	return lvl
}
