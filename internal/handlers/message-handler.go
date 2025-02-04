package handlers

import (
	"context"
	"errors"
	"files-bot-service/domain"
	"files-bot-service/file"
	"files-bot-service/user"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createKeyboard(fs []domain.File) tgbotapi.ReplyKeyboardMarkup {
	var rows [][]tgbotapi.KeyboardButton

	// Проходим по всем элементам массива fs
	for _, f := range fs {
		// Создаем кнопку с текстом из f.Name
		button := tgbotapi.KeyboardButton{
			Text: f.Name,
		}
		// Добавляем кнопку в строку
		rows = append(rows, []tgbotapi.KeyboardButton{button})
	}

	// Возвращаем клавиатуру
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       rows,
		ResizeKeyboard: true, // Чтобы клавиатура была адаптирована под размер экрана
	}
}

type MessageHandler struct {
	UserService *user.Service
	FileService *file.Service
}

func NewMessageHandler(userService *user.Service, fileService *file.Service) *MessageHandler {
	return &MessageHandler{
		UserService: userService,
		FileService: fileService,
	}
}

func (handler *MessageHandler) HandleMessage(message tgbotapi.Message) (msg tgbotapi.Chattable, err error) {
	var model domain.User
	model, err = handler.UserService.Repository.GetByID(context.Background(), uint(message.Chat.ID)) //TODO поправить на сервис
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	switch message.Text {
	case "/start":
		{
			if model.ID == 0 {
				model, err = handler.UserService.CreateUser(message.Chat.ID)
				if err != nil {
					return tgbotapi.MessageConfig{}, err
				}
				return tgbotapi.NewMessage(message.Chat.ID, "Вы успешно зарегистрированы!"), nil
			}

			return tgbotapi.NewMessage(message.Chat.ID, "Чем могу помочь?"), nil
		}
	case "/add":
		{
			if model.ID == 0 {
				return nil, errors.New("пользователь не найден")
			}
			model.Terminator = domain.WAIT_INPUT
			err = handler.UserService.Update(&model)
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}

			return tgbotapi.NewMessage(int64(model.ID), "Пришлите файл, который вы хотите сохранить"), nil
		}
	case "/getID": //TODO сделать автосетап по команде
		{
			return tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Chat ID: %d", message.Chat.ID)), nil
		}
	case "/getFilesList":
		{
			var fs []domain.File
			fs, err = handler.FileService.FileRepository.All(context.Background())
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			keyboard := createKeyboard(fs)
			m := tgbotapi.NewMessage(message.Chat.ID, "Доступные файлы")
			m.ReplyMarkup = keyboard
			model.Terminator = domain.GET_FILE
			err = handler.UserService.Update(&model)
			return m, nil
		}
	default:
		{
			if message.Text == "" && model.Terminator == "" {
				return nil, errors.New("некорректный запрос")
			}
			switch model.Terminator {
			case domain.WAIT_INPUT:
				{
					var fileName string
					if message.Document.FileName != "" {
						fileName = message.Document.FileName
					}
					f := domain.NewFile(message.MessageID, fileName)
					err = handler.FileService.FileRepository.Store(context.Background(), &f)
					if err != nil {
						return tgbotapi.MessageConfig{}, err
					}
					model.Terminator = ""
					err = handler.UserService.Update(&model)
					return tgbotapi.NewForward(-4655598408, int64(model.ID), int(f.ID)), nil
				}
			case domain.GET_FILE:
				{
					var f domain.File
					handler.FileService.FileRepository.Conn.Where("name = ?", message.Text).First(&f)
					if f.ID == 0 {
						return nil, errors.New("искомый вами файл не найден")
					}

					m := tgbotapi.NewCopyMessage(message.Chat.ID, 980196074, int(f.ID))
					m.AllowSendingWithoutReply = false
					m.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{
						RemoveKeyboard: true, // Это указывает на удаление клавиатуры
					}
					model.Terminator = ""
					err = handler.UserService.Update(&model)
					return m, nil
				}
			}
		}
	}
	return tgbotapi.MessageConfig{}, errors.New("отсутствует сообщение")
}
