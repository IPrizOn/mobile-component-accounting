package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	windowMain fyne.Window
	AppFyne    fyne.App
)

func Start() {
	AppFyne = app.New()

	windowMain = AppFyne.NewWindow("")

	openWindowAuth()

	windowMain.Show()
	windowMain.SetMaster()
	AppFyne.Run()
}
