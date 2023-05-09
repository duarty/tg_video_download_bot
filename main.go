package main

import (
	"log"

	"os"

	"dev.azure/duarty/tg_bot/third_party/twitter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Environment variable is not set.", err)
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_API_KEY"))
	if err != nil {
		log.Panic("Env var errors", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			url := update.Message.Text

			videoBytes := twitter.TwitterDownloader(url)

			uuidVideoName := uuid.New().String()

			video := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FileBytes{
				Name:  uuidVideoName + ".mp4",
				Bytes: videoBytes,
			})

			bot.Send(video)

		}
	}
}
