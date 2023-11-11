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

func GetResponse(conversation gpt.Conversation, ch chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()
	url := "https://api.openai.com/v1/chat/completions"
	bearerToken := "Bearer " + config.GetConfig().ApiKey

	reqBody, _ := json.Marshal(conversation)
	// if err != nil {
	// 	return err
	// }

	log.Printf(string(reqBody))

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	// if err != nil {
	// 	return err
	// }

	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-type", "application/json")

	client := http.Client{}

	resp, _ := client.Do(req)
	// if err != nil {
	// 	return err
	// }

	defer resp.Body.Close()
	bodyResp, _ := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	ch <- string(bodyResp)
	// return nil
}
