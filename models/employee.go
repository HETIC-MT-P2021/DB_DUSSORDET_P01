package models

//Employee is an employee DTO
type Employee struct {
	Number     int    `json:"employee_number"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Extension  string `json:"extension"`
	Email      string `json:"email"`
	OfficeCode string `json:"office_code"`
	ReportsTo  string `json:"reports_to"`
	JobTitle   string `json:"job_title"`
}

//Office is an office DTO
type Office struct {
	Code       string `json:"office_code"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	Territory  string `json:"territory"`
}
