package handler

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/rtemb/srv-users/internal/config"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
	"github.com/sirupsen/logrus"
)

func StartService(cfg *config.Server, grpcHandler srvUsers.UsersServiceServer, logger *logrus.Entry) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	terminate := make(chan struct{}, 2)
	go startGRPC(terminate, cfg, logger, grpcHandler)

	<-interrupt
	terminate <- struct{}{}
	logger.Debug("Shutting down")

	time.Sleep(2 * time.Second)
	os.Exit(0)
}

func startGRPC(terminate chan struct{}, cfg *config.Server, logger *logrus.Entry, grpcHandler srvUsers.UsersServiceServer) {
	grpcAddr := cfg.GRPCHost + ":" + cfg.GRPCPort
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	srvUsers.RegisterUsersServiceServer(grpcServer, grpcHandler)
	var closed bool
	go func() {
		<-terminate
		closed = true
		logger.Debug("Shutting down gRPC")
		grpcServer.GracefulStop()
	}()

	logger.Info("Serving gRPC on http://", grpcAddr)
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Error(errors.Wrap(err, "unable to start gRPC server"))
	}

	if !closed {
		close(terminate)
	}
}
