package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	windowDialog fyne.Window
)

func openDialogWindow() {
	windowDialog = AppFyne.NewWindow("Ошибка")

	content := container.NewVBox(
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("У вас недостаточно прав"), layout.NewSpacer()),
		widget.NewButton("Ок", func() {
			windowDialog.Close()
		}),
	)

	windowDialog.Resize(fyne.NewSize(250, 75))
	windowDialog.CenterOnScreen()
	windowDialog.SetContent(content)
	windowDialog.Show()
}
