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
		Question:        generateDriverPollTitle(time.Now()),
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

	solidarDriversChat := tele.Chat{ID: -1001721324233} //ID of the SK Vozaci chat
	message, err := poll.Send(bot, &solidarDriversChat, &tele.SendOptions{})
	if err != nil {
		fmt.Println(err)
	}
	if message != nil {
		fmt.Println(message)
	}

	return nil
}

func generateDriverPollTitle(currentDate time.Time) string {
	//Polls must be generated on Sundays; however, we added a fallback in case the poll was generated during the week
	startDate := getThisOrNextMonday(currentDate)
	endDate := startDate.AddDate(0, 0, 6)

	pollTitle := strconv.Itoa(startDate.Day()) + "." + strconv.Itoa(int(startDate.Month())) + " – " + strconv.Itoa(endDate.Day()) + "." + strconv.Itoa(int(endDate.Month()))
	return pollTitle
}

func getThisOrNextMonday(fromDate time.Time) time.Time {
	return fromDate.AddDate(0, 0, -int(fromDate.Weekday())+1)
}
