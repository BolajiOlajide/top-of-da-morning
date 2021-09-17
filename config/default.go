package config

import (
	"github.com/BolajiOlajide/go-env-utils"
	"github.com/BolajiOlajide/top-of-da-morning/constants"
)

var envVariables map[string]string
var err error

// GetEnvVars fetch all environment variables in .env
func GetEnvVars() map[string]string {
	_ = env.InitializeEnv()

	if len(envVariables) == 0 {
		envVariables = map[string]string{
			constants.TwitterAPIKey:            env.GetEnvVar("TWITTER_API_KEY", env.Options{}),
			constants.TwitterAPISecret:         env.GetEnvVar("TWITTER_API_SECRET", env.Options{}),
			constants.TwitterAccessToken:       env.GetEnvVar("TWITTER_ACCESS_TOKEN", env.Options{}),
			constants.TwitterAccessTokenSecret: env.GetEnvVar("TWITTER_ACCESS_TOKEN_SECRET", env.Options{}),
		}
	}

	return envVariables
}
