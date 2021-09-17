package actions

import "github.com/dghubble/oauth1"

func InitializeTwitterClient() {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
}
