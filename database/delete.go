package database

func DeleteComponent(id int) error {
	delete(componentsList, id)

	return nil
}

func DeleteCustomer(id int) error {
	delete(customerList, id)

	return nil
}

func DeleteSale(id int) error {
	delete(salesList, id)

	return nil
}
