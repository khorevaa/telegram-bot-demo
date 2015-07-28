# telegram-bot-demo
Example of a Telegram bot using Syfaro's telegram-bot-api [AppEngine]
This is a simple Telegram bot that replies to every user message with the same message. It uses [Go](https://golang.org/) and deploys to [AppEngine](https://appengine.google.com/).

## How to use this
1. Download this repo
2. Change *HERE_GOES_YOUR_APP_ID* for your AppEngine app's ID in *app.yaml*
3. Change *HERE_GOES_YOUR_TELEGRAM_BOT_TOKEN* for your Telegram Bot Token
4. Deploy your app: ```goapp deploy -oauth app.yaml```
5. Go to http://APP_ID.appspot.com/start to create the Webhook (you can go to */stop* to remove the Webhook)
6. You should deploy again without the */start* and */stop* handlers (comment them in case you need them in the future) so that nobody can start or stop your bot.

## How does this work?
The app has three handlers:
* **/start**  
Creates the Webhook for Telegram to send you the messages.
* **/stop**  
Removes the Webhook so Telegram will stop sending you messages.
* **/YOUR_TOKEN**  
Telegram will send the messages to this URL, this handles those requests. Only Telegram and you know the token, so nobody else has access to this URL.

For more information about the Telegram Bot API go to https://core.telegram.org/bots/api

## NOTE
This is currently using my fork of Syfaro's telegram-bot-api which can be found in https://github.com/rausntos/telegram-bot-api.  
The reason for creating that fork is that Syfaro's version uses http.DefaultClient which is not available in AppEngine. I made a pull request: https://github.com/Syfaro/telegram-bot-api/pull/6.
