package server

import (
	"github.com/koyo-os/nagent/pkg/logger"
	"github.com/koyo-os/nagent/pkg/proto/pb"
)

type Server struct{
	pb.UnimplementedNetworkAgentServer
	logger *logger.Logger
}