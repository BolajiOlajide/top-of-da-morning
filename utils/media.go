package utils

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

// GetMediaIDFromTweet using the tweet ID get the media associated with the tweet
func GetMediaIDFromTweet(tweetID int64, client *twitter.Client) ([]twitter.MediaEntity, error) {
	tweet, _, err := client.Statuses.Show(tweetID, &twitter.StatusShowParams{
		IncludeEntities: twitter.Bool(true),
	})

	if err != nil {
		return []twitter.MediaEntity{}, fmt.Errorf("An error occurred while fetching the tweet with ID %d", tweetID)
	}

	return tweet.Entities.Media, nil
}
