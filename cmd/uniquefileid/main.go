package main

import (
	"log"
	"os"

	"github.com/perpetuallyhorni/unique-file-id/internal/bot"
	"github.com/perpetuallyhorni/unique-file-id/internal/constants"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv(constants.EnvTelegramToken)
	if token == constants.EmptyString {
		log.Fatalf(constants.LogErrTokenNotSet, constants.EnvTelegramToken)
	}

	bot.Start(token)
}
