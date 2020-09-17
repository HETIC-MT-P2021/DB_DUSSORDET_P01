package models

import (
	"database/sql"
	"go_exercise/database"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

//GetCustomerByNumber repository method to get a customer by number from db
func (repository *Repository) GetCustomerByNumber(customerNumber int) (*Customer, error) {
	row := repository.Conn.QueryRow(`SELECT
	c.customerName,
	c.contactFirstName,
	c.contactLastName,
	c.phone,
	c.addressLine1,
	c.addressLine2,
	c.city,
	c.state,
	c.postalCode,
	c.country,
	c.salesRepEmployeeNumber,
	c.creditLimit
FROM
	customers c
WHERE
	c.customerNumber = (?);`, customerNumber)
	var salesRepEmployeeNumber database.NullInt64
	var creditLimit database.NullFloat64
	var name, firstName, lastName, phone, addressLine, city, country string
	var addressOptionalLine, state, postalCode database.NullString
	switch err := row.Scan(&name, &firstName, &lastName, &phone, &addressLine, &addressOptionalLine,
		&city, &state, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		customer := &Customer{
			Name:                   name,
			ContactLastName:        firstName,
			ContactFirstName:       lastName,
			Phone:                  phone,
			AddressLine:            addressLine,
			AddressOptionalLine:    addressOptionalLine,
			City:                   city,
			State:                  state,
			PostalCode:             postalCode,
			Country:                country,
			SalesRepEmployeeNumber: salesRepEmployeeNumber,
			CreditLimit:            creditLimit,
		}
		return customer, nil
	default:
		return nil, err
	}
}
