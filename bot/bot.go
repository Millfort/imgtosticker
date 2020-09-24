package bot

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/millfort/imgfit/sticker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func New(tgToken string) *Bot {
	api, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal(err)
	}
	u := tgbotapi.NewUpdate(0)
	updates, err := api.GetUpdatesChan(u)
	return &Bot{
		api:     api,
		updates: updates,
	}
}

func (b *Bot) Start() {
	for update := range b.updates {
		msg := update.Message
		log.Printf("got message")
		if msg == nil {
			log.Printf("msg is nil")
			continue
		}
		doc := msg.Document
		if doc == nil {
			log.Printf("document is nil")
			continue
		}
		b.handlePhoto(msg.Chat.ID, *doc)
	}
}

func (b *Bot) handlePhoto(chatID int64, photo tgbotapi.Document) error {
	fileID := photo.FileID
	fileURL, err := b.api.GetFileDirectURL(fileID)
	if err != nil {
		return err
	}
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	srcFile := resp.Body
	defer srcFile.Close()
	srcImg, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}
	dstImg := sticker.New(srcImg, sticker.HalfSize)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, dstImg)
	if err != nil {
		return err
	}
	dstFile := tgbotapi.FileBytes{
		Name:  "img.png",
		Bytes: buf.Bytes(),
	}
	doc := tgbotapi.NewDocumentUpload(chatID, dstFile)
	b.api.Send(doc)
	return nil
}
