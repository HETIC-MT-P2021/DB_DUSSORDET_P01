package models

import "go_exercise/database"

//Customer is a customer DTO
type Customer struct {
	Number                 int                  `json:"customer_number,omitempy"`
	Name                   string               `json:"customer_name"`
	ContactLastName        string               `json:"contact_last_name"`
	ContactFirstName       string               `json:"contact_first_name"`
	Phone                  string               `json:"phone"`
	AddressLine            string               `json:"address_line"`
	AddressOptionalLine    database.NullString  `json:"address_optional_line,omitempty"`
	City                   string               `json:"city"`
	State                  database.NullString  `json:"state"`
	PostalCode             database.NullString  `json:"postal_code"`
	Country                string               `json:"country"`
	SalesRepEmployeeNumber database.NullInt64   `json:"sales_rep_employee_number"`
	CreditLimit            database.NullFloat64 `json:"credit_limit"`
}
