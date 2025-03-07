package model

// mapping for request, response

type EmployeeResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
