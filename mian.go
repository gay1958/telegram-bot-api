
import (
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("7661595964:AAEKiia_Q9cMkhhFCkyq5hWS_EIZqQZ_ez"))
    if err != nil {
        panic(err)
    }

    bot.Debug = true
}
    // Create a new UpdateConfig struct with an offset of 0. Offsets are used
    // to make sure Telegram knows we've handled previous values and we don't
    // need them repeated.
    updateConfig := tgbotapi.NewUpdate(0)

    // Tell Telegram we should wait up to 30 seconds on each request for an
    // update. This way we can get information just as quickly as making many
    // frequent requests without having to send nearly as many.
    updateConfig.Timeout = 30

    // Start polling Telegram for updates.
    updates := bot.GetUpdatesChan(updateConfig)

    // Let's go through each update that we're getting from Telegram.
    for update := range updates {
        // Telegram can send many types of updates depending on what your Bot
        // is up to. We only want to look at messages for now, so we can
        // discard any other updates.
        if update.Message == nil {
            continue
        }

        // Now that we know we've gotten a new message, we can construct a
        // reply! We'll take the Chat ID and Text from the incoming message
        // and use it to create a new message.
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
        // We'll also say that this message is a reply to the previous message.
        // For any other specifications than Chat ID or Text, you'll need to
        // set fields on the `MessageConfig`.
        msg.ReplyToMessageID = update.Message.MessageID

        // Okay, we're sending our message off! We don't care about the message
        // we just sent, so we'll discard it.
        if _, err := bot.Send(msg); err != nil {
            // Note that panics are a bad way to handle errors. Telegram can
            // have service outages or network errors, you should retry sending
            // messages or more gracefully handle failures.
            panic(err)
        }
    }

‏openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes


‏package main

‏import (
‏	"log"

‏	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

‏func main() {
‏	bot, err := tgbotapi.NewBotAPI("7661595964:AAEKiia_Q9cMkhhFCkyq5hWS_EIZqQZ_ez")
‏	if err != nil {
‏		log.Panic(err)
	}

‏	bot.Debug = true

‏	log.Printf("Authorized on account %s", bot.Self.UserName)

‏	u := tgbotapi.NewUpdate(0)
‏	u.Timeout = 60

‏	updates := bot.GetUpdatesChan(u)

‏	for update := range updates {
‏		if update.Message != nil { // If we got a message
‏			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

‏			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
‏			msg.ReplyToMessageID = update.Message.MessageID

‏			bot.Send(msg)
		}
	}
}

‏package main

‏import (
‏	"log"
‏	"net/http"

‏	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

‏func main() {
‏	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
‏	if err != nil {
‏		log.Fatal(err)
	}

‏	bot.Debug = true

‏	log.Printf("Authorized on account %s", bot.Self.UserName)

‏	wh, _ := tgbotapi.NewWebhookWithCert("https://www.example.com:8443/"+bot.Token, "cert.pem")

‏	_, err = bot.Request(wh)
‏	if err != nil {
‏		log.Fatal(err)
	}

‏	info, err := bot.GetWebhookInfo()
‏	if err != nil {
‏		log.Fatal(err)
	}

‏	if info.LastErrorDate != 0 {
‏		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

‏	updates := bot.ListenForWebhook("/" + bot.Token)
‏	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

‏	for update := range updates {
‏		log.Printf("%+v\n", update)
	}
}