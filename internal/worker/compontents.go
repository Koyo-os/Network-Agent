package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type WorkerComponent interface {
	Run() error
}

var SenderStatusOK bool = true

type workerType string

func InitComponent(t workerType, cfg *config.Config, logger *logger.Logger, url string, statusChan chan string) WorkerComponent {
	switch t {
	case CLONER:
		return initCloner(cfg, logger, url, statusChan)
	case BUILDER:
		return initBuiler(cfg, logger, url, statusChan)
	case LINTER:
		return initLinter(cfg, logger, url, statusChan)
	case TESTER:
		return initTester(cfg, logger, url, statusChan)
	default:
		return nil
	}
}
