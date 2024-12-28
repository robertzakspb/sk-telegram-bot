package main

import (
	"context"
	"encoding/json"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func main() {}

func YandexCFHandler(ctx context.Context, request []byte) ([]byte, error) {
	//Checking if the payload is provided by the cloud trigger
	requestBody := CloudTriggerPayload{}
	err := json.Unmarshal(request, &requestBody)
	if err == nil {
		fmt.Println("Payload: ", requestBody)
		cloudTriggerHandler(requestBody)
	} else {
		fmt.Println("Not a cloud trigger payload")
	}

	//Checking if the payload is provided by Telegram
	var update tele.Update
	err = json.Unmarshal(request, &update)
	if err == nil {
		fmt.Println("This is a Telegram update") //This is where Telegram upates will be processed
	} else {
		fmt.Println("Not a Telegram update")
	}

	body, err := json.Marshal(&ResponseBody{
		StatusCode: 200,
		Body:       "This response indicates that the cloud function successfully completed execution",
	})
	if err != nil {
		return nil, err
	}

	return body, nil
}

func cloudTriggerHandler(payload CloudTriggerPayload) {
	switch payload.Details.Payload {
	case "generateSKMealDistributionPolls":
		err := SendMealDistributionEnrolmentPoll()
		if err != nil {
			fmt.Println("an error has occurred while attempting to send the meal distribution poll: ", err)
		}
	case "generateDriversPolls":
		err := SendDriverEnrolmentPoll()
		if err != nil {
			fmt.Println("an error occurred while attempting to send the next drivers' poll: ", err)
		}
	default:
		fmt.Println("unknown cloud trigger payload: ", payload.Details.Payload)
	}
}

type ResponseBody struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}
