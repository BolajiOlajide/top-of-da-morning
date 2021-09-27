package main

import (
	"fmt"
	"os"

	"github.com/BolajiOlajide/top-of-da-morning/actions"
	"github.com/BolajiOlajide/top-of-da-morning/logger"
	"github.com/BolajiOlajide/top-of-da-morning/utils"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	logger.InitializeLogger()
	logger.InfoLogger.Println("Initialized Top of Da Morning Script")
	client := actions.InitializeBot()

	switch command := os.Args[1]; command {
	case "bot":
		runBot(client)
	case "fetchMediaId":
		fetchMediaID(1442485251117248513, client)
	default:
		fmt.Printf("Unknown command %s", command)
		os.Exit(1)
	}
}

func runBot(client *twitter.Client) {
	var operationCode int = 0

	if actions.IsWeekDay() {
		err := actions.PublishTweet(client)

		if err != nil {
			operationCode = 1
		}
	} else {
		logger.InfoLogger.Println("Today is not a weekday. Take your time out to flex.")
	}

	os.Exit(operationCode)
}

func fetchMediaID(tweetID int64, client *twitter.Client) {
	media, _ := utils.GetMediaIDFromTweet(tweetID, client)

	if len(media) == 1 {
		fmt.Printf("Media ID: %d\n", media[0].ID)
	} else {
		fmt.Printf("Tweet contains more than one media or no media.\nLength: %d", len(media))
	}
}
