package bot

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"github.com/raulsntos/telegram-bot-api"
	"io/ioutil"
	"net/http"
)

var token string = "HERE_GOES_YOUR_TELEGRAM_BOT_TOKEN"

func init() {
	http.HandleFunc("/start", startBot)
	http.HandleFunc("/stop", stopBot)
	http.HandleFunc("/"+token, handler)
}

func startBot(w http.ResponseWriter, r *http.Request) {
	setWebhook("https://"+r.URL.Host+"/"+token, r)
}

func stopBot(w http.ResponseWriter, r *http.Request) {
	setWebhook(""+token, r)
}

func getBot(client *http.Client) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPIwithClient(token, client)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func setWebhook(link string, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	bot, err := getBot(client)
	if err != nil {
		c.Errorf("Error getting the bot: " + err.Error())
		return
	}
	bot.SetWebhook(tgbotapi.NewWebhook("https://" + r.URL.Host + "/" + token))
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.Errorf(err.Error())
		return
	}

	var update tgbotapi.Update
	json.Unmarshal(bytes, &update)

	bot, err := getBot(client)
	if err != nil {
		c.Errorf(err.Error())
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.SendMessage(msg)
}
