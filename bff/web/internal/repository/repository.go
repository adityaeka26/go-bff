package repository

import (
	"context"

	pbUser "github.com/adityaeka26/go-bff/bff/web/internal/repository/user_proto"
)

type UserRepository interface {
	GetUserInfo(ctx context.Context, id int) (*pbUser.GetUserInfoResponse, error)
}
