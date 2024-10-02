package dto

type GetUserInfoRequest struct {
	UserId int
}

type GetUserInfoResponse struct {
	UserId   int
	Name     string
	Email    string
	Location string
}
