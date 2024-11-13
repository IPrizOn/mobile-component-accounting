package client

import (
	"log"
	"mobile/database"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func openDialogWindowAddComponent() {
	windowDialogAdd := AppFyne.NewWindow("Добавление компонента")

	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	selecter := widget.NewSelect(
		[]string{
			"processor",
			"motherboard",
			"video_card",
			"ram",
			"disk",
			"power_unit",
			"frame",
		},

		func(s string) {

		},
	)

	form := widget.NewForm(
		widget.NewFormItem("Компонент", selecter),
		widget.NewFormItem("Описание", entry1),
		widget.NewFormItem("Цена", entry2),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Добавить", func() {
				price, _ := strconv.Atoi(entry2.Text)
				err := database.CreateComponent(selecter.Selected, entry1.Text, price)
				if err != nil {
					log.Println(err)
				}
				windowDialogAdd.Close()
				openWindowMain()
			}),
			widget.NewButton("  Отмена  ", func() {
				windowDialogAdd.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogAdd.Resize(fyne.NewSize(500, 300))
	windowDialogAdd.CenterOnScreen()
	windowDialogAdd.SetContent(content)
	windowDialogAdd.Show()
}

func openDialogWindowAddCustomer() {
	windowDialogAdd := AppFyne.NewWindow("Добавление клиента")

	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()
	entry4 := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Имя", entry1),
		widget.NewFormItem("Телефон", entry2),
		widget.NewFormItem("Почта", entry3),
		widget.NewFormItem("Адрес", entry4),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Добавить", func() {
				err := database.CreateCustomer(entry1.Text, entry2.Text, entry3.Text, entry4.Text)
				if err != nil {
					log.Println(err)
				}
				windowDialogAdd.Close()
				openWindowMain()
			}),
			widget.NewButton("  Отмена  ", func() {
				windowDialogAdd.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogAdd.Resize(fyne.NewSize(500, 300))
	windowDialogAdd.CenterOnScreen()
	windowDialogAdd.SetContent(content)
	windowDialogAdd.Show()
}

func openDialogWindowAddSale() {
	windowDialogAdd := AppFyne.NewWindow("Добавление продажи")

	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Клиент", entry1),
		widget.NewFormItem("Компонент", entry2),
		widget.NewFormItem("Количество", entry3),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Добавить", func() {
				compID, _ := strconv.Atoi(entry1.Text)
				custID, _ := strconv.Atoi(entry2.Text)
				count, _ := strconv.Atoi(entry3.Text)
				err := database.CreateSale(compID, custID, count)
				if err != nil {
					log.Println(err)
				}
				windowDialogAdd.Close()
				openWindowMain()
			}),
			widget.NewButton("  Отмена  ", func() {
				windowDialogAdd.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogAdd.Resize(fyne.NewSize(500, 300))
	windowDialogAdd.CenterOnScreen()
	windowDialogAdd.SetContent(content)
	windowDialogAdd.Show()
}
