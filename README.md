# Lambda function to relay notifications to Telegram
The idea is to have a Lambda function hosted on AWS for which a Monzo user can register a webhook. Then everytime a transaction is created, this Lambda function runs and forwards the data to a Telegram chat. Why use this? To build social pressure to spend wisely, what if your frugal friend knew everytime you bought a pint?

## To test
```
$ cd functions/monzo_fwd
$ telegram_bot=BOT_ID telegram_recipient=RECIPIENT_ID go run main.go < transaction.json
```

or use something awesome like [direnv](https://direnv.net/) to manage your env variables.
