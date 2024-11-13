package database

import (
	"mobile/domain"
)

func GetComponentsList() (map[int]domain.Component, error) {
	return componentsList, nil
}

func GetCustomersList() (map[int]domain.Customer, error) {
	return customerList, nil
}

func GetSalesList() (map[int]domain.Sale, error) {
	return salesList, nil
}

func GetPersonsList() (map[int]domain.Person, error) {
	return personsList, nil
}
