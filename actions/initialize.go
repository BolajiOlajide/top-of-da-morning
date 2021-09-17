package actions

import (
	"github.com/BolajiOlajide/top-of-da-morning/config"
	"github.com/BolajiOlajide/top-of-da-morning/constants"
	"github.com/BolajiOlajide/top-of-da-morning/models"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func initializeTwitterClient(credentials models.TwitterCredentials) *twitter.Client {
	config := oauth1.NewConfig(credentials.ConsumerKey, credentials.ConsumerSecret)
	token := oauth1.NewToken(credentials.AccessToken, credentials.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}

func generateTwitterCreds() models.TwitterCredentials {
	variables := config.GetEnvVars()

	apiKey := variables[constants.TwitterAPIKey]
	apiSecret := variables[constants.TwitterAPISecret]
	accessToken := variables[constants.TwitterAccessToken]
	accessSecret := variables[constants.TwitterAccessTokenSecret]

	return models.TwitterCredentials{
		AccessSecret:   accessSecret,
		AccessToken:    accessToken,
		ConsumerKey:    apiKey,
		ConsumerSecret: apiSecret,
	}
}

// InitializeBot starts up all the relevant services and prerequisites for this app
func InitializeBot() *twitter.Client {
	credentials := generateTwitterCreds()

	client := initializeTwitterClient(credentials)
	return client
}
