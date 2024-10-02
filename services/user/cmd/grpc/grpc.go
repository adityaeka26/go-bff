package grpc

import (
	"fmt"
	"net"

	"github.com/adityaeka26/go-bff/services/user/config"
	"github.com/adityaeka26/go-bff/services/user/internal/handler"
	"github.com/adityaeka26/go-bff/services/user/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServeGRPC(config *config.EnvConfig, userUsecase usecase.UserUsecase) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", config.GrpcPort))
	if err != nil {
		return err
	}

	app := grpc.NewServer()

	handler.InitGrpcHandler(
		app,
		userUsecase,
	)
	reflection.Register(app)

	fmt.Println("running grpc on port", config.GrpcPort)
	err = app.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
