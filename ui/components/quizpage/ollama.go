package quizpage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Message            Message   `json:"message"`
	Done               bool      `json:"done"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

type ResponseChunk struct {
	Model   string  `json:"model"`
	Created int64   `json:"created_at"`
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}

const DefaultOllamaURL = "http://localhost:11434/api/chat"

func StreamChunks(url string, req Request) (<-chan string, <-chan error) {
	out := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		req.Stream = true
		js, err := json.Marshal(&req)
		if err != nil {
			errc <- err
			return
		}

		httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(js))
		if err != nil {
			errc <- err
			return
		}

		client := http.Client{}
		httpResp, err := client.Do(httpReq)
		if err != nil {
			errc <- err
			return
		}
		defer httpResp.Body.Close()

		reader := bufio.NewReader(httpResp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				errc <- err
				return
			}

			var chunk ResponseChunk

			if err := json.Unmarshal(line, &chunk); err != nil {
				errc <- err
				return
			}

			if chunk.Done {
				break
			}

			out <- chunk.Message.Content
		}
	}()

	return out, errc
}
