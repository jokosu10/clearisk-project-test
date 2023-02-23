package models

type OrganizationsByCountry struct {
	Index         int
	ID            string
	Name          string
	Website       string
	Country       string
	Description   string
	Founded       int
	Industry      string
	NumOfEmployee int
}

type OrganizationsByCountry_response struct {
	Code    int                      `json:"code"`
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    []OrganizationsByCountry `json:"data"`
}
