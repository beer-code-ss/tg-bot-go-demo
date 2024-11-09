package main

import (
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    botToken := "Ваш_Токен"
    
    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message != nil {
            handleMessage(bot, update.Message)
        } else if update.CallbackQuery != nil {
            handleCallbackQuery(bot, update.CallbackQuery)
        }
    }
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    if message.Contact != nil {
        phoneNumber := message.Contact.PhoneNumber
        reply := tgbotapi.NewMessage(message.Chat.ID, "Спасибо! Ваш номер телефона: "+phoneNumber)
        bot.Send(reply)
        return
    }

    switch message.Text {
    case "/start":
        startCommand(bot, message)
    case "/help":
        helpCommand(bot, message)
    case "/buttons":
        buttonsExample(bot, message)
    case "/photo":
        sendPhoto(bot, message)
    case "/document":
        sendDocument(bot, message)
    case "/contact":
        requestPhoneNumber(bot, message)
    case "/poll":
        sendCustomPoll(bot, message)
    case "/format":
        sendFormattedText(bot, message)
    case "/quickbutton":
        quickReplyButton(bot, message)
    default:
        msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Введите /help для списка команд.")
        bot.Send(msg)
    }
}

func startCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    startText := `
Добро пожаловать! Этот бот демонстрирует основные функции Telegram API.
Вот список доступных команд:
/start - начало работы
/help - помощь
/buttons - пример инлайн-кнопок
/photo - отправить фото
/document - отправить документ
/contact - запросить ваш контактный номер
/poll - создать кастомный опрос
/format - показать форматированный текст
/quickbutton - быстрая кнопка
`
    msg := tgbotapi.NewMessage(message.Chat.ID, startText)
    bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    helpText := `
Список команд:
/start - начало работы
/help - помощь
/buttons - пример инлайн-кнопок
/photo - отправить фото
/document - отправить документ
/contact - запросить ваш контактный номер
/poll - создать кастомный опрос
/format - показать форматированный текст
/quickbutton - быстрая кнопка
`
    msg := tgbotapi.NewMessage(message.Chat.ID, helpText)
    bot.Send(msg)
}

func buttonsExample(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    button1 := tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "button1")
    button2 := tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "button2")

    keyboard := tgbotapi.NewInlineKeyboardMarkup(
        tgbotapi.NewInlineKeyboardRow(button1, button2),
    )

    msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите одну из кнопок:")
    msg.ReplyMarkup = keyboard

    bot.Send(msg)
}

func sendPhoto(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath("files/pivocode.jpg"))
    photo.Caption = "Вот фото!"
    bot.Send(photo)
}

func sendDocument(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    doc := tgbotapi.NewDocument(message.Chat.ID, tgbotapi.FilePath("files/pivocode.txt"))
    doc.Caption = "Вот документ!"
    bot.Send(doc)
}

func requestPhoneNumber(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    button := tgbotapi.NewKeyboardButton("Отправить свой номер телефона")
    button.RequestContact = true
    keyboard := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(button),
    )

    msg := tgbotapi.NewMessage(message.Chat.ID, "Нажмите на кнопку ниже, чтобы отправить свой номер телефона.")
    msg.ReplyMarkup = keyboard
    bot.Send(msg)
}

func sendCustomPoll(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    poll := tgbotapi.NewPoll(message.Chat.ID, "Любишь пиво?", "Да", "Нет")
    poll.IsAnonymous = true
    poll.AllowsMultipleAnswers = false
    bot.Send(poll)
}

func sendFormattedText(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    formattedText := "*Жирный текст*\n_Курсив_\n~Зачеркнутый~\n__Подчёркнутый__\n\n" +
        "`Моноширинный текст`\n\n> Цитата текста\n\n||Скрытый текст||"

    msg := tgbotapi.NewMessage(message.Chat.ID, formattedText)
    msg.ParseMode = "MarkdownV2"
    bot.Send(msg)
}

func quickReplyButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    button := tgbotapi.NewKeyboardButton("Отправить свою локацию")
    button.RequestLocation = true
    keyboard := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(button),
    )

    msg := tgbotapi.NewMessage(message.Chat.ID, "Нажмите на кнопку, чтобы отправить вашу локацию.")
    msg.ReplyMarkup = keyboard
    bot.Send(msg)
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
    var responseText string
    switch callback.Data {
    case "button1":
        responseText = "Вы нажали на Кнопка 1"
    case "button2":
        responseText = "Вы нажали на Кнопка 2"
    default:
        responseText = "Неизвестная кнопка"
    }

    msg := tgbotapi.NewMessage(callback.Message.Chat.ID, responseText)
    bot.Send(msg)

    callbackResp := tgbotapi.NewCallback(callback.ID, responseText)
    bot.Request(callbackResp)
}
