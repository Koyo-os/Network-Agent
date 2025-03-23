package app

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/internal/server"
	"github.com/koyo-os/nagent/pkg/logger"
	"google.golang.org/grpc"
)

type App struct{
	logger *logger.Logger
	cfg *config.Config
	server *server.Server
	grpc *grpc.Server
}

func Init() {
	
}