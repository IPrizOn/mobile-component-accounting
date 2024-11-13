package database

import (
	"mobile/domain"
)

func CreateComponent(comp string, desc string, price int) error {
	componentsList[len(componentsList)] = domain.Component{Type: comp, Description: desc, Price: price}

	return nil
}

func CreateCustomer(name string, phone string, email string, address string) error {
	customerList[len(customerList)] = domain.Customer{Name: name, Phone: phone, Email: email, Address: address}

	return nil
}

func CreateSale(compID int, custID int, count int) error {
	salesList[len(salesList)] = domain.Sale{CustomerID: custID, ComponentID: compID, Count: count}

	return nil
}
