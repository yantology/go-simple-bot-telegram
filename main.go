package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Inisialisasi bot dengan token
	bot, err := tgbotapi.NewBotAPI("8013191447:AAGg9AEbC6s2gLv8DP9g2_tK5BUlVqZjcXk")
	if err != nil {
		log.Panic(err)
	}

	// Mengatur mode debug (opsional)
	bot.Debug = true

	// Konfigurasi update
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// Mendapatkan channel updates
	updates := bot.GetUpdatesChan(updateConfig)

	// Handle updates
	for update := range updates {
		// Handle commands
		if update.Message != nil && update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "start":
				msg.Text = "Selamat datang! Gunakan /help untuk melihat perintah yang tersedia."
			case "help":
				msg.Text = "Perintah yang tersedia:\n" +
					"/start - Mulai bot\n" +
					"/help - Tampilkan bantuan\n" +
					"/photo - Kirim foto contoh\n" +
					"/keyboard - Tampilkan keyboard"
			case "photo":
				// Kirim foto
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("path/to/photo.jpg"))
				bot.Send(photo)
				continue
			case "keyboard":
				// Membuat keyboard
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Button 1"),
						tgbotapi.NewKeyboardButton("Button 2"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Button 3"),
						tgbotapi.NewKeyboardButton("Button 4"),
					),
				)
				msg.ReplyMarkup = keyboard
				msg.Text = "Keyboard ditampilkan!"
			default:
				msg.Text = "Perintah tidak dikenal"
			}

			bot.Send(msg)
		}

		// Handle text biasa (bukan command)
		if update.Message != nil && !update.Message.IsCommand() {
			// Echo pesan user
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}

		// Handle callback dari inline keyboard
		if update.CallbackQuery != nil {
			// Respond to button press
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Tombol ditekan!")
			bot.Send(callback)

			// Send message
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Anda menekan: "+update.CallbackQuery.Data)
			bot.Send(msg)
		}
	}
}

// // Fungsi untuk membuat inline keyboard
// func createInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
// 	return tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("Button 1", "btn1"),
// 			tgbotapi.NewInlineKeyboardButtonData("Button 2", "btn2"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonURL("Google", "https://google.com"),
// 		),
// 	)
// }
