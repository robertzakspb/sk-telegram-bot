package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v4"
)

// Provided by the Yandex Cloud trigger that calls the cloud function
type CloudTriggerPayload struct {
	EventMetadata struct {
		EventID   string    `json:"event_id"`
		EventType string    `json:"event_type"`
		CreatedAt time.Time `json:"created_at"`
		CloudID   string    `json:"cloud_id"`
		FolderID  string    `json:"folder_id"`
	} `json:"event_metadata"`
	Details struct {
		TriggerID string `json:"trigger_id"`
		Payload   string `json:"payload"`
	} `json:"details"`
}

func SendMealDistributionEnrolmentPoll() error {
	pref := tele.Settings{
		Token:   os.Getenv("SK_TELEGRAM_BOT_TOKEN"),
		Verbose: true,
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		fmt.Println(err)
		return err
	}

	weeklyPolls, err := generateMealDistributionTelegramPoll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, poll := range weeklyPolls {
		solidarKitchenChat := tele.Chat{ID: -1002011462579} //ID of the SK Terenci chat
		pollTopicID := 2                                    //ID of the topic where meal distribution polls are posted
		message, err := poll.Send(bot, &solidarKitchenChat, &tele.SendOptions{ThreadID: pollTopicID})
		if err != nil {
			fmt.Println(err)
		}
		if message != nil {
			fmt.Println(message)
		}

	}

	return nil
}

func generateMealDistributionTelegramPoll() ([]tele.Poll, error) {

	weeklyPolls := []tele.Poll{}
	for _, date := range nextDistributionDates() {
		poll := tele.Poll{
			Question:        generateMealDistributionPollTitle(date),
			MultipleAnswers: true,
			Anonymous:       false,
			Type:            tele.PollRegular,
			Options: []tele.PollOption{
				{
					Text: "°ДОЛАЗИМ " + requiredNumberOfPeople(date) + "°: -клик ако долазиш!-",
				}, {
					Text: "°РЕД (1-2)°: водим рачуна о реду од 15.30 ч. (током поделе треба две особе, пре једна)",
				},
				{
					Text: breadPollOptionTitle(date),
				},
				{
					Text: "°СТО И РУКСАК (1-2)°: из \"ДЦ Кров\"",
				},
				{
					Text: "°ИЗВЕШТАЈ (1)°: пишем после поделе",
				}, {
					Text: "°ФОТКЕ (1)°: -не фоткати лица-",
				},
				{
					Text: "°ПОДЕЛА (3)°: код стола сам",
				},
				{
					Text: "°СМЕЋЕ°(1)",
				},
			},
		}
		if shouldDistributeDrink(date) {
			poll.Options = append(poll.Options, tele.PollOption{Text: "°НАПИТАК° (1-2): делим напитак"})
		}
		poll.Options = append(poll.Options, tele.PollOption{Text: "°Увид у стање°: види пријављене"})

		weeklyPolls = append(weeklyPolls, poll)
	}

	return weeklyPolls, nil

}

func nextDistributionDates() []time.Time {
	mealDistibutionDays := [...]time.Weekday{3, 6, 0}

	nextDistributionDates := []time.Time{}
	today := time.Now()
	//This logic assumes that the poll is being scheduled on Sundays
	for _, distributionDay := range mealDistibutionDays {
		if distributionDay.String() == "Sunday" { //As Sunday is 0 in Go, we need to add 7
			nextSunday := today.AddDate(0, 0, 7)
			nextDistributionDates = append(nextDistributionDates, nextSunday)
		} else {
			nextDate := today.AddDate(0, 0, int(distributionDay))
			nextDistributionDates = append(nextDistributionDates, nextDate)
		}
	}
	return nextDistributionDates
}

func generateMealDistributionPollTitle(date time.Time) string {
	title := ""
	switch date.Weekday().String() {
	case "Sunday":
		title += "НЕДЕЉА"
	case "Monday":
		title += "ПОНЕДЕЉАК"
	case "Tuesday":
		title += "УТОРАК"
	case "Wednesday":
		title += "СРЕДА"
	case "Thursday":
		title += "ЧЕТВРТАК"
	case "Friday":
		title += "ПЕТАК"
	case "Saturday":
		title += "СУБОТА"
	}

	title += " "
	title += strconv.Itoa(date.Day()) + "."
	title += strconv.Itoa(int(date.Month())) + "."
	title += strconv.Itoa(date.Year())

	return title
}

func shouldDistributeDrink(date time.Time) bool {
	drinkDistributionDay := "Sunday"
	return date.Weekday().String() == drinkDistributionDay
}

func requiredNumberOfPeople(date time.Time) string {
	if date.Weekday().String() == "Sunday" {
		return "(6-7)"
	} else {
		return "(5)"
	}
}

func breadPollOptionTitle(date time.Time) string {
	title := "°ХЛЕБ (1-2)°" 

	switch date.Weekday().String() {
	case "Sunday": 
		title += ": из ДЦ \"Кров\""
	case "Saturday", "Wednesday": 
		title += ": из продавнице \"Расина\""
	}

	return title
}
