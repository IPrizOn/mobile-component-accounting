package client

import (
	"log"
	"mobile/database"
	"mobile/domain"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	userRole  string
	usersList map[int]domain.Person
)

func openWindowAuth() {
	var err error

	usersList, err = database.GetPersonsList()
	if err != nil {
		log.Println(err)
	}

	windowMain.SetTitle("Авторизация")

	content := loadContent()

	windowMain.Resize(fyne.NewSize(300, 250))
	windowMain.CenterOnScreen()
	windowMain.SetContent(content)
}

func loadContent() *fyne.Container {
	labelLogin := widget.NewLabel("Логин")
	labelPassword := widget.NewLabel("Пароль")
	labelError := widget.NewLabel("")

	entryLogin := widget.NewEntry()
	entryPassword := widget.NewPasswordEntry()

	buttonAuth := widget.NewButton("Войти", func() {
		if entryLogin.Text != "" && entryPassword.Text != "" {
			isAuth := isAuth(entryLogin.Text, entryPassword.Text)
			if isAuth {
				openWindowMain()
			} else {
				labelError.SetText("Неправильный логин или пароль")
			}
		} else {
			labelError.SetText("Обнаружены пустые поля")
		}
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		labelLogin,
		entryLogin,
		labelPassword,
		entryPassword,
		container.NewHBox(layout.NewSpacer(), labelError, layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(), buttonAuth, layout.NewSpacer()),
		layout.NewSpacer(),
	)

	return content
}

func isAuth(login string, password string) bool {
	if login == usersList[0].Login && password == usersList[0].Password {
		userRole = usersList[0].Role
		return true
	} else if login == usersList[1].Login && password == usersList[1].Password {
		userRole = usersList[1].Role
		return true
	}

	return false
}
