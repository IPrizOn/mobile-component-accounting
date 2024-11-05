package database

import (
	"mobile/domain"
)

func GetComponentsList() ([]domain.Component, error) {
	var (
		componentsList []domain.Component
	)

	componentsList = append(componentsList, domain.Component{ID: 1, Type: "processor", Description: "amd ryzen 5 7600", Price: 15000})
	componentsList = append(componentsList, domain.Component{ID: 2, Type: "motherload", Description: "asus b250k", Price: 5000})
	componentsList = append(componentsList, domain.Component{ID: 3, Type: "video_card", Description: "amd rx 7700 xt", Price: 50000})
	componentsList = append(componentsList, domain.Component{ID: 4, Type: "ram", Description: "adata 16gb", Price: 7500})
	componentsList = append(componentsList, domain.Component{ID: 5, Type: "disk", Description: "adata ssd 1000gb", Price: 5000})
	componentsList = append(componentsList, domain.Component{ID: 6, Type: "power_unit", Description: "xilence 600w", Price: 5000})
	componentsList = append(componentsList, domain.Component{ID: 7, Type: "frame", Description: "ardor b15", Price: 2500})

	return componentsList, nil
}

func GetCustomersList() ([]domain.Customer, error) {
	var (
		customerList []domain.Customer
	)

	customerList = append(customerList, domain.Customer{ID: 1, Name: "Ivan", Phone: "+75440110980", Email: "ivan1@main.ru", Address: "Pushkina 3"})
	customerList = append(customerList, domain.Customer{ID: 2, Name: "Andrey", Phone: "+79562675302", Email: "andrey5@main.ru", Address: "Proezdnoe 109"})

	return customerList, nil
}

func GetSalesList() ([]domain.Sale, error) {
	var (
		salesList []domain.Sale
	)

	salesList = append(salesList, domain.Sale{ID: 1, CustomerID: 1, ComponentID: 1, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 2, CustomerID: 2, ComponentID: 1, Count: 2})
	salesList = append(salesList, domain.Sale{ID: 3, CustomerID: 1, ComponentID: 2, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 4, CustomerID: 2, ComponentID: 2, Count: 2})
	salesList = append(salesList, domain.Sale{ID: 5, CustomerID: 1, ComponentID: 3, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 6, CustomerID: 2, ComponentID: 3, Count: 2})
	salesList = append(salesList, domain.Sale{ID: 7, CustomerID: 1, ComponentID: 4, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 8, CustomerID: 2, ComponentID: 4, Count: 2})
	salesList = append(salesList, domain.Sale{ID: 9, CustomerID: 1, ComponentID: 5, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 10, CustomerID: 2, ComponentID: 5, Count: 2})
	salesList = append(salesList, domain.Sale{ID: 11, CustomerID: 1, ComponentID: 6, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 12, CustomerID: 2, ComponentID: 6, Count: 2})
	salesList = append(salesList, domain.Sale{ID: 13, CustomerID: 1, ComponentID: 7, Count: 1})
	salesList = append(salesList, domain.Sale{ID: 14, CustomerID: 2, ComponentID: 7, Count: 2})

	return salesList, nil
}

func GetPersonsList() ([]domain.Person, error) {
	return nil, nil
}

func IsUserExist() error {
	return nil
}
