package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Worker struct{
	cfg *config.Config
	logger *logger.Logger
	statusChan chan string
	taskChan chan string
	components []WorkerComponent
}

func Init(cfg *config.Config) *Worker {
	var components []WorkerComponent


	return &Worker{
		cfg: cfg,
		logger: logger.Init(),
	}
}

func (w *Worker) Run(url string) error {
	for _, c := range w.components{
		if err := c.Run();err != nil{
			w.logger.Error("error run component", zapcore.Field{
				Key: "err",
				String: err.Error(),
			})
		}
	}

	return nil
}