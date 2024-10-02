package dto

type GetOrdersRequest struct {
	UserId int
}

type GetOrdersResponse struct {
	OrderId  int
	UserId   int
	Item     string
	Quantity int
}
