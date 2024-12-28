package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v4"
)

func SendDriverEnrolmentPoll() error {
	pref := tele.Settings{
		Token:   os.Getenv("SK_TELEGRAM_BOT_TOKEN"),
		Verbose: true,
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		fmt.Println(err)
		return err
	}

	poll := tele.Poll{
		Question:        generateDriverPollTitle(),
		MultipleAnswers: true,
		Anonymous:       false,
		Type:            tele.PollRegular,
		Options: []tele.PollOption{
			{
				Text: "Уторак 19 ч",
			},
			{
				Text: "Среда 14.30 ч. - 17 ч.",
			},
			{
				Text: "Петак 19 ч.",
			},
			{
				Text: "Субота 14.30 ч. - 17 ч.",
			},
			{
				Text: "Недеља 15 ч. - 17 ч.",
			}, {
				Text: "Увид",
			},
		}}
	//TODO: Replace IDs with the real value
	solidarDriversChat := tele.Chat{ID: -1002301646854} //ID of the SK Vozaci chat
	pollTopicID := 2                                    //ID of the topic where drivers' polls are posted
	message, err := poll.Send(bot, &solidarDriversChat, &tele.SendOptions{ThreadID: pollTopicID})
	if err != nil {
		fmt.Println(err)
	}
	if message != nil {
		fmt.Println(message)
	}

	return nil
}

func generateDriverPollTitle() string {
	currentDate := time.Now() //Must be Sunday
	nextMonday := currentDate.AddDate(0, 0, 1)
	nextSunday := currentDate.AddDate(0, 0, 7)

	pollTitle := strconv.Itoa(nextMonday.Day()) + "." + strconv.Itoa(int(nextMonday.Month())) + " – " + strconv.Itoa(nextSunday.Day()) + "." + strconv.Itoa(int(nextSunday.Month()))
	return pollTitle
}
