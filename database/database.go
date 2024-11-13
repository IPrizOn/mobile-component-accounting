package database

import (
	"mobile/domain"
)

var (
	componentsList map[int]domain.Component
	customerList   map[int]domain.Customer
	salesList      map[int]domain.Sale
	personsList    map[int]domain.Person
)

func InitDatabase() {
	initComponents()
	initCustomers()
	initSales()
	initPersons()
}

func initComponents() {
	componentsList[1] = domain.Component{Type: "processor", Description: "amd ryzen 5 7600", Price: 15000}
	componentsList[2] = domain.Component{Type: "motherload", Description: "asus b250k", Price: 5000}
	componentsList[3] = domain.Component{Type: "video_card", Description: "amd rx 7700 xt", Price: 50000}
	componentsList[4] = domain.Component{Type: "ram", Description: "adata 16gb", Price: 7500}
	componentsList[5] = domain.Component{Type: "disk", Description: "adata ssd 1000gb", Price: 5000}
	componentsList[6] = domain.Component{Type: "power_unit", Description: "xilence 600w", Price: 5000}
	componentsList[7] = domain.Component{Type: "frame", Description: "ardor b15", Price: 2500}
}

func initCustomers() {
	customerList[1] = domain.Customer{Name: "Ivan", Phone: "+75440110980", Email: "ivan1@main.ru", Address: "Pushkina 3"}
	customerList[2] = domain.Customer{Name: "Andrey", Phone: "+79562675302", Email: "andrey5@main.ru", Address: "Proezdnoe 109"}
}

func initSales() {
	salesList[1] = domain.Sale{CustomerID: 1, ComponentID: 1, Count: 1}
	salesList[2] = domain.Sale{CustomerID: 2, ComponentID: 1, Count: 2}
	salesList[3] = domain.Sale{CustomerID: 1, ComponentID: 2, Count: 1}
	salesList[4] = domain.Sale{CustomerID: 2, ComponentID: 2, Count: 2}
	salesList[5] = domain.Sale{CustomerID: 1, ComponentID: 3, Count: 1}
	salesList[6] = domain.Sale{CustomerID: 2, ComponentID: 3, Count: 2}
	salesList[7] = domain.Sale{CustomerID: 1, ComponentID: 4, Count: 1}
	salesList[8] = domain.Sale{CustomerID: 2, ComponentID: 4, Count: 2}
	salesList[9] = domain.Sale{CustomerID: 1, ComponentID: 5, Count: 1}
	salesList[10] = domain.Sale{CustomerID: 2, ComponentID: 5, Count: 2}
	salesList[11] = domain.Sale{CustomerID: 1, ComponentID: 6, Count: 1}
	salesList[12] = domain.Sale{CustomerID: 2, ComponentID: 6, Count: 2}
	salesList[13] = domain.Sale{CustomerID: 1, ComponentID: 7, Count: 1}
	salesList[14] = domain.Sale{CustomerID: 2, ComponentID: 7, Count: 2}
}

func initPersons() {
	personsList[1] = domain.Person{Login: "admin", Password: "12345", Role: "admin"}
	personsList[2] = domain.Person{Login: "user", Password: "1111", Role: "common"}
}
