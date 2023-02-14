package models

type Organizations struct {
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

// type ListOrganization []Organizations

type Organization_response struct {
	Code    int             `json:"code"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []Organizations `json:"data"`
}

type Organization_post struct {
	ID            string `json:"org_id" form:"org_id"`
	Name          string `json:"name" form:"name"`
	Website       string `json:"website" form:"website"`
	Country       string `json:"country" form:"country"`
	Description   string `json:"description" form:"description"`
	Founded       int    `json:"founded" form:"founded"`
	Industry      string `json:"industry" form:"industry"`
	NumOfEmployee int    `json:"noe" form:"noe"`
}

type Organization_response_single struct {
	Code    int           `json:"code"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    Organizations `json:"data"`
}
