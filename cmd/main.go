package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/JakubC-projects/chat-gpt-game/chatgpt"
	"github.com/JakubC-projects/chat-gpt-game/frontend"
	"github.com/JakubC-projects/chat-gpt-game/game"
)

var ChatGptApiKey = os.Getenv("CHAT_GPT_API_KEY")

func main() {
	mux := http.NewServeMux()

	client := chatgpt.New(ChatGptApiKey)

	mux.HandleFunc("/submit", func(w http.ResponseWriter, req *http.Request) {
		g := game.GetState(req)
		lvl := game.CurrentLevel(req, g.MaxLevel())

		if err := frontend.Render(w, g, lvl); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Fatalf("cannot render page: %s", err)
			return
		}
		w.WriteHeader(http.StatusOK)

	})

	mux.HandleFunc("/reset", func(w http.ResponseWriter, req *http.Request) {
		game.DeleteState(w)
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		g := game.GetState(req)
		lvl := game.CurrentLevel(req, g.MaxLevel())

		prompt, err := getPrompt(req)
		if err == nil {
			err := g.Answer(client, lvl, prompt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalf("cannot get answer: %s", err)
				return
			}
			game.SaveState(w, g)
		}

		if err := frontend.Render(w, g, lvl); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Fatalf("cannot render page: %s", err)
			return
		}

	})

	http.ListenAndServe(":8080", mux)
}

func getPrompt(req *http.Request) (string, error) {
	err := req.ParseForm()
	if err != nil {
		return "", err

	}

	prompt := req.FormValue("prompt")

	if prompt == "" {
		return "", errors.New("missing prompt")
	}

	return prompt, nil
}
