package api

import (
	"encoding/json"
	"helperGPT/config"
	"helperGPT/gpt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetResponse(conversation gpt.Conversation) (string, error) {

	url := "https://api.openai.com/v1/chat/completions"
	bearerToken := "Bearer " + config.GetConfig().ApiKey

	reqBody, err := json.Marshal(conversation)
	if err != nil {
		return "", err
	}

	log.Printf(string(reqBody))

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyResp), nil
}
