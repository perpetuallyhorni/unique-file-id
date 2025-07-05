package constants

const (
	// Startup messages
	LogBotStarting = "Bot is running! Send it any file, photo, or video."

	// Info messages
	LogMediaReceived = "Received media. FileID: '%s', UniqueID: '%s'"

	// Error messages
	LogErrTokenNotSet       = "Error: %s environment variable not set." // #nosec G101
	LogErrCouldNotCreateBot = "Could not create bot: %v"
)
