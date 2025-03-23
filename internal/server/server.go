package server

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/internal/worker"
	"github.com/koyo-os/nagent/pkg/logger"
	"github.com/koyo-os/nagent/pkg/proto/pb"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type Server struct{
	pb.UnimplementedNetworkAgentServer
	Logger *logger.Logger
	Cfg *config.Config
	TaskChan chan int
	StatusChan chan string
}

func (s *Server) SetTask(req *pb.WorkRequest,stream grpc.ServerStreamingServer[pb.StreamResp]) error {
	var cherr chan error

	workerCount := 4
	if s.Cfg.NotifyCfg.Send {
		workerCount++
	}

	worker := worker.Init(s.Cfg, req.RepoUrl, s.StatusChan, s.TaskChan, cherr)
	go worker.Run()

	for i := range workerCount{
		select {
		case err := <- cherr:
			s.Logger.Error("error run worker", zapcore.Field{
				Key: "err",
				String: err.Error(),
			})
		case status := <- s.StatusChan:
			stream.Send(&pb.StreamResp{
				Status: status,
				Task: int32(i),
			})
		}
	}

	return nil
}
