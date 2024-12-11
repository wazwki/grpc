package app

import (
	"example/api/proto/examplepb"
	"example/internal/config"
	"example/internal/controllers/grpc"
	"example/internal/repository"
	"example/internal/service"
	"fmt"
	"net"
)

type App struct {
	serverHost string
	serverPort string
	server     *grpc.Server
}

func New(cfg *config.Config) (*App, error) {

	// logger

	// db

	repository := repository.NewRepository("inited db")
	service := service.NewService(repository)
	controllers := handlers.NewControllers(service)

	grpcServer := grpc.NewServer()

	examplepb.RegisterExampleServiceServer(grpcServer, controllers)

	return &App{server: grpcServer, serverHost: cfg.Host, serverPort: cfg.Port}, nil
}

func (a *App) Run() error {

	// migrate

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", a.serverHost, a.serverPort))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	go func() {
		fmt.Println(fmt.Sprintf("Server is running at %v", lis.Addr()))
		if err := a.server.Serve(lis); err != nil {
			fmt.Println(err.Error())
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.server.GracefulStop()

	//db

	return nil
}
