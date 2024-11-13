package client

import (
	"log"
	"mobile/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func openDialogWindowError() {
	windowDialogInfo := AppFyne.NewWindow("Ошибка")

	content := container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("У вас недостаточно прав"),
			layout.NewSpacer(),
		),
		widget.NewButton("ОК", func() {
			windowDialogInfo.Close()
		}),
	)

	windowDialogInfo.Resize(fyne.NewSize(250, 75))
	windowDialogInfo.CenterOnScreen()
	windowDialogInfo.SetContent(content)
	windowDialogInfo.Show()
}

func openDialogWindowConfirm(selectedTab string, id int) {
	windowDialogInfo := AppFyne.NewWindow("Подтверждение")

	content := container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Вы точно хотите удалить эту запись?"),
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("  Да  ", func() {
				switch selectedTab {
				case "component":
					err := database.DeleteComponent(id)
					if err != nil {
						log.Println(err)
					}
				case "customer":
					err := database.DeleteCustomer(id)
					if err != nil {
						log.Println(err)
					}
				case "sale":
					err := database.DeleteSale(id)
					if err != nil {
						log.Println(err)
					}
				}
				windowDialogInfo.Close()
				openWindowMain()
			}),
			widget.NewButton(" Нет ", func() {
				windowDialogInfo.Close()
			}),
			layout.NewSpacer(),
		),
	)

	windowDialogInfo.Resize(fyne.NewSize(250, 75))
	windowDialogInfo.CenterOnScreen()
	windowDialogInfo.SetContent(content)
	windowDialogInfo.Show()
}
