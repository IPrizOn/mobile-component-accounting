package client

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"mobile/database"
)

var (
	buttonBack *widget.Button
)

func openWindowMain() {
	windowMain.SetTitle("Учёт комплектующих")

	buttonBack = widget.NewButton("Выход", func() {
		openWindowAuth()
	})

	contentComponents := loadTabComponents()
	contentCustomers := loadTabCustomers()
	contentSales := loadTabSales()

	appTabs := container.NewAppTabs(
		container.NewTabItem("Комплектующие", contentComponents),
		container.NewTabItem("Клиенты", contentCustomers),
		container.NewTabItem("Продажи", contentSales),
	)
	appTabs.SetTabLocation(container.TabLocationTop)

	windowMain.Resize(fyne.NewSize(1200, 650))
	windowMain.CenterOnScreen()
	windowMain.SetContent(appTabs)
}

func loadTabComponents() *fyne.Container {
	componentsList, err := database.GetComponentsList()
	if err != nil {
		log.Println(err)
	}

	dataTable := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			return len(componentsList), 5
		},

		func() fyne.CanvasObject {
			return container.NewStack(
				widget.NewLabel("data"),
				container.NewHBox(),
			)
		},

		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*fyne.Container).Objects[0].(*widget.Label)
			buttonsContainer := co.(*fyne.Container).Objects[1].(*fyne.Container)
			label.Show()
			buttonsContainer.Hide()
			if tci.Col == 0 {
				label.SetText(strconv.Itoa(tci.Row + 1))
			} else if tci.Col == 1 {
				label.SetText(componentsList[tci.Row].Type)
			} else if tci.Col == 2 {
				label.SetText(componentsList[tci.Row].Description)
			} else if tci.Col == 3 {
				label.SetText(strconv.Itoa(componentsList[tci.Row].Price))
			} else if tci.Col == 4 {
				label.Hide()
				buttonsContainer.Show()

				co.(*fyne.Container).Objects[1] = container.NewHBox(
					widget.NewButton("Изменить", func() {
						if userRole != "common" {
							openDialogWindowEditComponent(tci.Row, componentsList[tci.Row].Type, componentsList[tci.Row].Description, componentsList[tci.Row].Price)
						} else {
							openDialogWindowError()
						}
					}),
					widget.NewButton("Удалить", func() {
						if userRole != "common" {
							openDialogWindowConfirm("component", tci.Row)
						} else {
							openDialogWindowError()
						}
					}),
				)
			}
		},
	)

	dataTable.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		label := template.(*widget.Label)
		if id.Col == -1 {
			label.SetText(strconv.Itoa(id.Row + 1))
		} else if id.Col == 0 {
			label.SetText("Код")
		} else if id.Col == 1 {
			label.SetText("Компонент")
		} else if id.Col == 2 {
			label.SetText("Описание")
		} else if id.Col == 3 {
			label.SetText("Цена")
		} else if id.Col == 4 {
			label.SetText("Действия")
		}
	}
	dataTable.SetColumnWidth(0, 100)
	dataTable.SetColumnWidth(1, 200)
	dataTable.SetColumnWidth(2, 400)
	dataTable.SetColumnWidth(3, 200)
	dataTable.SetColumnWidth(4, 200)

	box := container.NewHBox(
		widget.NewLabel("СПИСОК КОМПОНЕНТОВ"),
		layout.NewSpacer(),
		widget.NewButton("Добавить", func() {
			if userRole != "common" {
				openDialogWindowAddComponent()
			} else {
				openDialogWindowError()
			}
		}),
	)

	container := layout.NewBorderLayout(box, buttonBack, nil, nil)

	content := fyne.NewContainerWithLayout(
		container,
		box,
		dataTable,
		buttonBack,
	)

	return content
}

func loadTabCustomers() *fyne.Container {
	customersList, err := database.GetCustomersList()
	if err != nil {
		log.Println(err)
	}

	dataTable := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			return len(customersList), 6
		},

		func() fyne.CanvasObject {
			return container.NewStack(
				widget.NewLabel("data"),
				container.NewHBox(),
			)
		},

		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*fyne.Container).Objects[0].(*widget.Label)
			buttonsContainer := co.(*fyne.Container).Objects[1].(*fyne.Container)
			label.Show()
			buttonsContainer.Hide()
			if tci.Col == 0 {
				label.SetText(strconv.Itoa(tci.Row + 1))
			} else if tci.Col == 1 {
				label.SetText(customersList[tci.Row].Name)
			} else if tci.Col == 2 {
				label.SetText(customersList[tci.Row].Phone)
			} else if tci.Col == 3 {
				label.SetText(customersList[tci.Row].Email)
			} else if tci.Col == 4 {
				label.SetText(customersList[tci.Row].Address)
			} else if tci.Col == 5 {
				label.Hide()
				buttonsContainer.Show()

				co.(*fyne.Container).Objects[1] = container.NewHBox(
					widget.NewButton("Изменить", func() {
						if userRole != "common" {
							openDialogWindowEditCustomer(tci.Row, customersList[tci.Row].Name, customersList[tci.Row].Phone, customersList[tci.Row].Email, customersList[tci.Row].Address)
						} else {
							openDialogWindowError()
						}
					}),
					widget.NewButton("Удалить", func() {
						if userRole != "common" {
							openDialogWindowConfirm("customer", tci.Row)
						} else {
							openDialogWindowError()
						}
					}),
				)
			}
		},
	)

	dataTable.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		label := template.(*widget.Label)
		if id.Col == -1 {
			label.SetText(strconv.Itoa(id.Row + 1))
		} else if id.Col == 0 {
			label.SetText("Код")
		} else if id.Col == 1 {
			label.SetText("Имя")
		} else if id.Col == 2 {
			label.SetText("Телефон")
		} else if id.Col == 3 {
			label.SetText("Почта")
		} else if id.Col == 4 {
			label.SetText("Адрес")
		} else if id.Col == 5 {
			label.SetText("Действия")
		}
	}
	dataTable.SetColumnWidth(0, 100)
	dataTable.SetColumnWidth(1, 200)
	dataTable.SetColumnWidth(2, 200)
	dataTable.SetColumnWidth(3, 200)
	dataTable.SetColumnWidth(4, 200)
	dataTable.SetColumnWidth(5, 200)

	box := container.NewHBox(
		widget.NewLabel("СПИСОК КЛИЕНТОВ"),
		layout.NewSpacer(),
		widget.NewButton("Добавить", func() {
			if userRole != "common" {
				openDialogWindowAddCustomer()
			} else {
				openDialogWindowError()
			}
		}),
	)

	container := layout.NewBorderLayout(box, buttonBack, nil, nil)

	content := fyne.NewContainerWithLayout(
		container,
		box,
		dataTable,
		buttonBack,
	)

	return content
}

func loadTabSales() *fyne.Container {
	salesList, err := database.GetSalesList()
	if err != nil {
		log.Println(err)
	}

	dataTable := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			return len(salesList), 5
		},

		func() fyne.CanvasObject {
			return container.NewStack(
				widget.NewLabel("data"),
				container.NewHBox(),
			)
		},

		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*fyne.Container).Objects[0].(*widget.Label)
			buttonsContainer := co.(*fyne.Container).Objects[1].(*fyne.Container)
			label.Show()
			buttonsContainer.Hide()
			if tci.Col == 0 {
				label.SetText(strconv.Itoa(tci.Row + 1))
			} else if tci.Col == 1 {
				label.SetText(strconv.Itoa(salesList[tci.Row].ComponentID))
			} else if tci.Col == 2 {
				label.SetText(strconv.Itoa(salesList[tci.Row].CustomerID))
			} else if tci.Col == 3 {
				label.SetText(strconv.Itoa(salesList[tci.Row].Count))
			} else if tci.Col == 4 {
				label.Hide()
				buttonsContainer.Show()

				co.(*fyne.Container).Objects[1] = container.NewHBox(
					widget.NewButton("Изменить", func() {
						if userRole != "common" {
							openDialogWindowEditSale(tci.Row, salesList[tci.Row].ComponentID, salesList[tci.Row].CustomerID, salesList[tci.Row].Count)
						} else {
							openDialogWindowError()
						}
					}),
					widget.NewButton("Удалить", func() {
						if userRole != "common" {
							openDialogWindowConfirm("sale", tci.Row)
						} else {
							openDialogWindowError()
						}
					}),
				)
			}
		},
	)

	dataTable.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		label := template.(*widget.Label)
		if id.Col == -1 {
			label.SetText(strconv.Itoa(id.Row + 1))
		} else if id.Col == 0 {
			label.SetText("Код")
		} else if id.Col == 1 {
			label.SetText("Компонент")
		} else if id.Col == 2 {
			label.SetText("Клиент")
		} else if id.Col == 3 {
			label.SetText("Количество")
		} else if id.Col == 4 {
			label.SetText("Действия")
		}
	}
	dataTable.SetColumnWidth(0, 100)
	dataTable.SetColumnWidth(1, 200)
	dataTable.SetColumnWidth(2, 200)
	dataTable.SetColumnWidth(3, 200)
	dataTable.SetColumnWidth(4, 200)

	box := container.NewHBox(
		widget.NewLabel("СПИСОК ПРОДАЖ"),
		layout.NewSpacer(),
		widget.NewButton("Добавить", func() {
			if userRole != "common" {
				openDialogWindowAddSale()
			} else {
				openDialogWindowError()
			}
		}),
	)

	container := layout.NewBorderLayout(box, buttonBack, nil, nil)

	content := fyne.NewContainerWithLayout(
		container,
		box,
		dataTable,
		buttonBack,
	)

	return content
}
