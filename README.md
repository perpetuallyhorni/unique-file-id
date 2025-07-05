# unique-file-id

This is a Telegram bot that retrieves and displays the `file_id` and `unique_id` of media files sent to it.

## Usage

1. **Set up the bot:**
    * Create a couple of Telegram bots using BotFather.
    * Obtain both bot tokens.
    * Set the `BOT_TOKEN` variable in the `.env` file to your first bot's token.

2. **Run the bot:**
    * Execute the compiled program.
3. **Interact with the bot:**
    * Send any media file (photo, video, etc.) to the bot.
    * The bot will respond with the `file_id` and `unique_id` of the media.

## Example

* Send a photo to the first bot.
* The bot responds with:

> The `file_id` for this media is `BAA...gACAgIAAxkBAAIBa2X`</br>
> The `unique_id` for this media is `AgACAgIAAxkBAAIBa2X`

* Stop the bot.
* Replace the Token in the `.env` file with your second bot token and run the program again.
* Send the same photo to the new bot.

* The bot responds with:

> The `file_id` for this media is `AAC...CAAAQEAAwIAA20AAR4`</br>
> The `unique_id` for this media is `AgACAgIAAxkBAAIBa2X`

* Notice that the `unique_id` remains the same, while the `file_id` changes.
