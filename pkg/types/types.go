package types

type Miyuki struct {
	Id      int64  `json:"id"`
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}
