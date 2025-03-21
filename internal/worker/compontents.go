package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type WorkerComponent interface{
	Run() error
}

type workerType string

func InitComponent(t workerType, cfg *config.Config, logger *logger.Logger, url string, statusChan chan string) WorkerComponent {
	switch t{
	case "cloner":
		return initCloner(cfg, logger, url, statusChan)
	default:
		return nil
	}
}