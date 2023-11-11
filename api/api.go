package api

import (
	"encoding/json"
	"helperGPT/config"
	"helperGPT/gpt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type ReturnMessage struct {
	Message string
	Err     error
}

func GetResponse(conversation gpt.Conversation, ch chan<- ReturnMessage, wg *sync.WaitGroup) {

	defer wg.Done()
	url := "https://api.openai.com/v1/chat/completions"
	bearerToken := "Bearer " + config.GetConfig().ApiKey

	reqBody, err := json.Marshal(conversation)
	if err != nil {
		ch <- ReturnMessage{"", err}
	}

	log.Printf(string(reqBody))

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	if err != nil {
		ch <- ReturnMessage{"", err}
	}

	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		ch <- ReturnMessage{"", err}
	}

	defer resp.Body.Close()
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- ReturnMessage{"", err}
	}

	ch <- ReturnMessage{string(bodyResp), nil}
}
