package models

import (
	"database/sql"
	"go_exercise/database"
	"time"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

//GetCustomerInfo repository method to get a customer by number from db
func (repository *Repository) GetCustomerInfo(customerNumber int) (*Customer, error) {
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

//GetCustomerOrders repository method to get a customer's orders
func (repository *Repository) GetCustomerOrders(customerNumber int) ([]*Order, error) {
	var orders []*Order
	rows, err := repository.Conn.Query(`SELECT
	o.orderNumber,
	o.orderDate,
	o.requiredDate,
	o.shippedDate,
	o.status,
	o.comments,
	SUM(od.quantityOrdered) AS number_of_items,
	SUM(od.priceEach * od.quantityOrdered) AS total_cost
FROM
	orders o
	INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber
WHERE
	o.customerNumber = (?)
GROUP BY
	o.orderNumber;`, customerNumber)
	if err != nil {
		return nil, err
	}
	var orderNumber, numberOfItems int
	var totalCost float32
	var orderDate, requiredDate time.Time
	var shippedDate database.NullTime
	var status string
	var comments database.NullString

	for rows.Next() {
		if err := rows.Scan(&orderNumber, &orderDate, &requiredDate, &shippedDate, &status,
			&comments, &numberOfItems, &totalCost); err != nil {
			return nil, err
		}

		order := &Order{
			Number:         orderNumber,
			CreationDate:   orderDate,
			RequiredDate:   requiredDate,
			ShippedDate:    shippedDate,
			Status:         status,
			Comments:       comments,
			CustomerNumber: customerNumber,
			NumberOfItems:  numberOfItems,
			TotalCost:      totalCost,
		}

		orders = append(orders, order)
	}

	return orders, nil
}

//GetOrderProductList repository method to get an order details
func (repository *Repository) GetOrderProductList(orderNumber int) ([]*Product, error) {
	var poducts []*Product
	rows, err := repository.Conn.Query(`SELECT
	p.productCode,
	p.productName,
	p.productLine,
	p.productScale,
	p.productVendor,
	p.productDescription,
	p.quantityInStock,
	p.buyPrice,
	p.MSRP
FROM
	orderdetails od
	INNER JOIN products p ON od.productCode = p.productCode
WHERE
	od.orderNumber = (?);`, orderNumber)
	if err != nil {
		return nil, err
	}
	var buyPrice, MSRP database.NullFloat64
	var code, name, line, scale, vendor, description string
	var quantityInStock int64

	for rows.Next() {
		if err := rows.Scan(&code, &name, &line, &scale, &vendor, &description, &quantityInStock,
			&buyPrice, &MSRP); err != nil {
			return nil, err
		}

		product := &Product{
			Code:            code,
			Name:            name,
			Line:            line,
			Scale:           scale,
			Vendor:          vendor,
			Description:     description,
			QuantityInStock: quantityInStock,
			BuyPrice:        buyPrice,
			MSRP:            MSRP,
		}

		poducts = append(poducts, product)
	}

	return poducts, nil
}

//GetProductInfo repository method to get a product by its code from db
func (repository *Repository) GetProductInfo(productCode string) (*Product, error) {
	row := repository.Conn.QueryRow(`SELECT
	p.productCode,
	p.productName,
	p.productLine,
	p.productScale,
	p.productVendor,
	p.productDescription,
	p.quantityInStock,
	p.buyPrice,
	p.MSRP
FROM
	products p
WHERE
	p.productCode = (?);`, productCode)
	var buyPrice, MSRP database.NullFloat64
	var code, name, line, scale, vendor, description string
	var quantityInStock int64
	switch err := row.Scan(&code, &name, &line, &scale, &vendor, &description, &quantityInStock,
		&buyPrice, &MSRP); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		product := &Product{
			Code:            productCode,
			Name:            name,
			Line:            line,
			Scale:           scale,
			Vendor:          vendor,
			Description:     description,
			QuantityInStock: quantityInStock,
			BuyPrice:        buyPrice,
			MSRP:            MSRP,
		}
		return product, nil
	default:
		return nil, err
	}
}

//GetEmployees repository method to get all employees
func (repository *Repository) GetEmployees() ([]*Employee, error) {
	var employees []*Employee
	rows, err := repository.Conn.Query(`SELECT
	e.employeeNumber,
	e.lastName,
	e.firstName,
	e.extension,
	e.email,
	o.country AS office_country,
	o.city AS office_city,
	e.jobTitle
FROM
	employees e
	INNER JOIN offices o ON e.officeCode = o.officeCode;
`)
	if err != nil {
		return nil, err
	}
	var employeeNumber int
	var lastName, firstName, extension, email, country, city, jobTitle string

	for rows.Next() {
		if err := rows.Scan(&employeeNumber, &lastName, &firstName, &extension, &email, &country,
			&city, &jobTitle); err != nil {
			return nil, err
		}

		employee := &Employee{
			Number:        employeeNumber,
			LastName:      lastName,
			FirstName:     firstName,
			Extension:     extension,
			Email:         email,
			OfficeCountry: country,
			OfficeCity:    city,
			JobTitle:      jobTitle,
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
