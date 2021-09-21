package actions

import (
	"github.com/BolajiOlajide/top-of-da-morning/logger"
	"github.com/dghubble/go-twitter/twitter"
)

const topOfDaMorningVideoID int64 = 1438926145966247938

// PublishTweet send top of the morning tweet haha
func PublishTweet(client *twitter.Client) error {
	logger.InfoLogger.Println("About to publish tweet.")

	randomVideo, err := pickRandomVideo()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while loading a random video. %s", err)
		return err
	}
	status := ComposeTweetText(randomVideo.Owner)
	params := twitter.StatusUpdateParams{
		MediaIds: []int64{randomVideo.MediaID},
	}

	_, _, err = client.Statuses.Update(status, &params)
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while sending the tweet. %s", err)
		return err
	}

	logger.InfoLogger.Println("Successfully published tweet.")
	return nil
}
