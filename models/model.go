package models

type Users struct {
	ID   int    `json: "id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Users  `json:"data"`
}

type UsersResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Users `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
