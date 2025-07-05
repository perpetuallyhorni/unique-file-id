package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/perpetuallyhorni/unique-file-id/internal/constants" // Import the constants package

	"gopkg.in/telebot.v3"
)

// Start initializes and starts the Telegram bot.
func Start(token string) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf(constants.LogErrCouldNotCreateBot, err)
	}

	// Register handlers
	registerHandlers(b)

	log.Println(constants.LogBotStarting)
	b.Start()
}

// registerHandlers sets up all the bot's message handlers.
func registerHandlers(b *telebot.Bot) {
	b.Handle(telebot.OnMedia, func(c telebot.Context) error {
		return handleMedia(c)
	})
}

// handleMedia processes any incoming media message to find its file_unique_id.
func handleMedia(c telebot.Context) error {
	file := c.Message().Media().MediaFile()
	fileID := file.FileID
	uniqueID := file.UniqueID

	log.Printf(constants.LogMediaReceived, fileID, uniqueID)

	replyText := fmt.Sprintf(constants.ReplyFileIDAndUniqueID, fileID, uniqueID)

	return c.Reply(replyText, telebot.ModeMarkdownV2)
}
