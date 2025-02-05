package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AdviceResponse struct {
	Slip struct {
		ID     int    `json:"id"`
		Advice string `json:"advice"`
	} `json:""slip`
}

func FetchAdviceAndPublish() {
	resp, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		log.Println("Error when seeking advice:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading API response")
		log.Println("Error reading API response:", err)
		return
	}

	var advice AdviceResponse
	if err := json.Unmarshal(body, &advice); err != nil {
		fmt.Println("Error decoding JSON")
		log.Println("Error decoding JSON:", err)
		return
	}

	messageBody, _ := json.Marshal(advice.Slip)
	PublishMessage(messageBody)
}
