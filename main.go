package main

import (
	"os"

	"github.com/BolajiOlajide/top-of-da-morning/actions"
	"github.com/BolajiOlajide/top-of-da-morning/logger"
)

func main() {
	logger.InitializeLogger()
	logger.InfoLogger.Println("Initialized Top of Da Morning Script")

	client := actions.InitializeBot()

	var operationCode int = 0
	err := actions.PublishTweet(client)

	if err != nil {
		operationCode = 1
	}

	os.Exit(operationCode)
}
