package models

import "database/sql"

//Customer is a customer DTO
type Customer struct {
	Number                 int             `json:"customer_number"`
	Name                   string          `json:"customer_name"`
	ContactLastName        string          `json:"contact_last_name"`
	ContactFirstName       string          `json:"contact_first_name"`
	Phone                  string          `json:"phone"`
	Address                string          `json:"address"` //@toDO decide if one or two strings
	City                   string          `json:"city"`
	State                  string          `json:"state"`
	PostalCode             string          `json:"postal_code"`
	Country                string          `json:"country"`
	SalesRepEmployeeNumber int             `json:"sales_rep_employee_number"`
	CreditLimit            sql.NullFloat64 `json:"credit_limit"`
}
