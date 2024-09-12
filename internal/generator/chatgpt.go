package generator

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mehmetymw/grpc-testgen/internal/config"
)

func getChatGPTResponses(cfg *config.Config, prompt string) (string, error) {
	url := "https://api.openai.com/v1/engines/" + cfg.ChatGPTModel + "/completions"
	reqBody := map[string]interface{}{
		"prompt":     prompt,
		"max_tokens": 150,
	}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+cfg.ChatGPTKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["choices"].([]interface{})[0].(map[string]interface{})["text"].(string), nil
}
