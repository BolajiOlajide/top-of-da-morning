package main

import (
	"github.com/BolajiOlajide/top-of-da-morning/actions"
	"github.com/BolajiOlajide/top-of-da-morning/logger"
)

func main() {
	logger.InitializeLogger()
	logger.InfoLogger.Println("Initialized Top of Da Morning Script")

	actions.InitializeBot()
}
