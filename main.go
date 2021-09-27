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

	runBot(client)
	// fetchMediaID(1442399972641628162, client)
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
