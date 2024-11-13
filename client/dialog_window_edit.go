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

func openDialogWindowEditComponent(id int, comp string, desc string, price int) {
	windowDialogEdit := AppFyne.NewWindow("Изменение компонента")

	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Компонент", entry1),
		widget.NewFormItem("Описание", entry2),
		widget.NewFormItem("Цена", entry3),
	)

	entry1.SetText(comp)
	entry2.SetText(desc)
	entry3.SetText(strconv.Itoa(price))

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Изменить", func() {
				thisPrice, _ := strconv.Atoi(entry3.Text)
				err := database.UpdateComponent(id, entry1.Text, entry2.Text, thisPrice)
				if err != nil {
					log.Println(err)
				}
				windowDialogEdit.Close()
				openWindowMain()
			}),
			widget.NewButton("  Отмена  ", func() {
				windowDialogEdit.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogEdit.Resize(fyne.NewSize(500, 300))
	windowDialogEdit.CenterOnScreen()
	windowDialogEdit.SetContent(content)
	windowDialogEdit.Show()
}

func openDialogWindowEditCustomer(id int, name string, phone string, email string, address string) {
	windowDialogEdit := AppFyne.NewWindow("Изменение клиента")

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

	entry1.SetText(name)
	entry2.SetText(phone)
	entry3.SetText(email)
	entry4.SetText(address)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Изменить", func() {
				err := database.UpdateCustomer(id, entry1.Text, entry2.Text, entry3.Text, entry4.Text)
				if err != nil {
					log.Println(err)
				}
				windowDialogEdit.Close()
				openWindowMain()
			}),
			widget.NewButton("  Отмена  ", func() {
				windowDialogEdit.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogEdit.Resize(fyne.NewSize(500, 300))
	windowDialogEdit.CenterOnScreen()
	windowDialogEdit.SetContent(content)
	windowDialogEdit.Show()
}

func openDialogWindowEditSale(id int, comp int, cust int, count int) {
	windowDialogEdit := AppFyne.NewWindow("Изменение продажи")

	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Компонент", entry1),
		widget.NewFormItem("Клиент", entry2),
		widget.NewFormItem("Количество", entry3),
	)

	entry1.SetText(strconv.Itoa(comp))
	entry2.SetText(strconv.Itoa(cust))
	entry3.SetText(strconv.Itoa(count))

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Изменить", func() {
				compID, _ := strconv.Atoi(entry1.Text)
				custID, _ := strconv.Atoi(entry2.Text)
				thisCount, _ := strconv.Atoi(entry3.Text)
				err := database.UpdateSale(id, compID, custID, thisCount)
				if err != nil {
					log.Println(err)
				}
				windowDialogEdit.Close()
				openWindowMain()
			}),
			widget.NewButton("  Отмена  ", func() {
				windowDialogEdit.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogEdit.Resize(fyne.NewSize(500, 300))
	windowDialogEdit.CenterOnScreen()
	windowDialogEdit.SetContent(content)
	windowDialogEdit.Show()
}
