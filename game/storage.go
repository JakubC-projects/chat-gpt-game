package game

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

var stateCookie = "state"

func SaveState(w http.ResponseWriter, s *State) {

	state, err := json.Marshal(s)
	if err != nil {
		return
	}
	http.SetCookie(w, &http.Cookie{Name: stateCookie, Value: base64.StdEncoding.EncodeToString(state)})
}

func DeleteState(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: stateCookie, Value: ""})
}

func GetState(req *http.Request) *State {
	var s State

	c, err := req.Cookie(stateCookie)
	if err != nil {
		return &s
	}

	bytes, err := base64.StdEncoding.DecodeString(c.Value)
	if err != nil {
		return &s
	}

	err = json.Unmarshal(bytes, &s)
	if err != nil {
		return &s
	}

	return &s
}
