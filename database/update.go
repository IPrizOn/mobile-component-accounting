package database

import (
	"mobile/domain"
)

func UpdateComponent(id int, comp string, desc string, price int) error {
	delete(componentsList, id)
	componentsList[len(componentsList)] = domain.Component{Type: comp, Description: desc, Price: price}

	return nil
}

func UpdateCustomer(id int, name string, phone string, email string, address string) error {
	delete(customerList, id)
	customerList[len(customerList)] = domain.Customer{Name: name, Phone: phone, Email: email, Address: address}

	return nil
}

func UpdateSale(id int, compID int, custID int, count int) error {
	delete(salesList, id)
	salesList[len(salesList)] = domain.Sale{CustomerID: custID, ComponentID: compID, Count: count}

	return nil
}
