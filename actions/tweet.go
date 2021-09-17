package actions

import (
	"github.com/BolajiOlajide/top-of-da-morning/logger"
	"github.com/dghubble/go-twitter/twitter"
)

const topOfDaMorningVideoID int64 = 1438926338660999171

// PublishTweet send top of the morning tweet haha
func PublishTweet(client *twitter.Client) error {
	logger.InfoLogger.Println("About to publish tweet.")
	status := ComposeTweetText()
	params := twitter.StatusUpdateParams{
		MediaIds: []int64{topOfDaMorningVideoID},
	}

	_, _, err := client.Statuses.Update(status, &params)
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while sending the tweet. %s", err)
		return err
	}

	return nil
}
