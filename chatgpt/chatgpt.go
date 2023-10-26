package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey}
}

func (c *Client) Prompt(systemPrompt, userPrompt string) (string, error) {

	completionRequest := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		MaxTokens: 100,
	}

	response, err := c.createChatCompletion(completionRequest)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}

func (c *Client) createChatCompletion(request ChatCompletionRequest) (ChatCompletionResponse, error) {
	var result ChatCompletionResponse

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return result, fmt.Errorf("cannot marshal json: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(jsonBody))
	if err != nil {
		return result, fmt.Errorf("cannot create request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, fmt.Errorf("cannot get response: %w", err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("cannot read response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("received invalid response")
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return result, fmt.Errorf("cannot unmarshal response: %w", err)
	}

	return result, nil
}
