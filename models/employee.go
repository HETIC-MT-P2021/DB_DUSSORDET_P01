package models

import (
	"go_exercise/database"
)

//Employee is an employee DTO
type Employee struct {
	Number        int    `json:"employee_number"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	Extension     string `json:"extension,omitempty"`
	Email         string `json:"email,omitempty"`
	OfficeCountry string `json:"office_country,omitempty"`
	OfficeCity    string `json:"office_city,omitempty"`
	JobTitle      string `json:"job_title,omitempty"`
}

//Office is an office DTO
type Office struct {
	Code                string              `json:"office_code"`
	City                string              `json:"city"`
	Phone               string              `json:"phone"`
	AddressLine         string              `json:"address_line"`
	AddressOptionalLine database.NullString `json:"address_optional_line"`
	State               database.NullString `json:"state"`
	Country             string              `json:"country"`
	PostalCode          string              `json:"postal_code"`
	Territory           string              `json:"territory"`
	Employees           []*Employee         `json:"employees"`
}
