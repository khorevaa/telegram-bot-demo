# telegram-bot-demo
Example of a Telegram bot using [Syfaro's telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) [AppEngine]
This is a simple Telegram bot that replies to every user message with the same message. It uses [Go](https://golang.org) and deploys to [AppEngine](https://appengine.google.com/).

## Prerequisites
* Have **Go** installed  
https://golang.org
* Have **telegram-bot-api** installed  
```go get github.com/go-telegram-bot-api/telegram-bot-api```

## How to use this
1. Download this repo
2. Change *HERE_GOES_YOUR_APP_ID* for your AppEngine app's ID in *app.yaml*
3. Change *HERE_GOES_YOUR_TELEGRAM_BOT_TOKEN* for your Telegram Bot Token
4. Deploy your app: ```goapp deploy app.yaml```
5. Go to http://APP_ID.appspot.com/YOUR_TOKEN/start to create the Webhook (you can go to */YOUR_TOKEN/stop* to remove the Webhook)

## How does this work?
The app has three handlers:
* **/YOUR_TOKEN/start**  
Creates the Webhook for Telegram to send you the messages.
* **/YOUR_TOKEN/stop**  
Removes the Webhook so Telegram will stop sending you messages.
* **/YOUR_TOKEN**  
Telegram will send the messages to this URL, this handles those requests. Only Telegram and you know the token, so nobody else has access to this URL.

For more information about the Telegram Bot API go to https://core.telegram.org/bots/api
