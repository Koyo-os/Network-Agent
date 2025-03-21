package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type Worker struct{
	cfg *config.Config
	logger *logger.Logger
	statusChan chan string
	taskChan chan string
	components []WorkerComponent
}

func Init(cfg *config.Config) *Worker {
	return &Worker{
		cfg: cfg,
		logger: logger.Init(),
	}
}

func (w *Worker) Run(url string) error {

}