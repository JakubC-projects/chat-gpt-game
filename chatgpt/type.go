package chatgpt

type ChatCompletionRequest struct {
	Model     string    `json:"model,omitempty"`
	Messages  []Message `json:"messages,omitempty"`
	MaxTokens int       `json:"max_tokens,omitempty"`
}

type ChatCompletionResponse struct {
	Choices []Choice `json:"choices,omitempty"`
}

type Choice struct {
	Index   int     `json:"index,omitempty"`
	Message Message `json:"message,omitempty"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
