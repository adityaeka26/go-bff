package dto

type GetOrderHistoryRequest struct {
	UserId int `params:"user_id" validate:"required"`
}

type GetOrderHistoryResponse struct {
	UserId   int      `json:"user_id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Location string   `json:"location"`
	Orders   []string `json:"orders"`
}
