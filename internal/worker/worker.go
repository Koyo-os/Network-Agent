package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Worker struct {
	cfg        *config.Config
	logger     *logger.Logger
	statusChan chan string
	taskChan   chan int
	components []WorkerComponent
	url        string
}

func Init(cfg *config.Config, url string) *Worker {
	var statusChan chan string
	var taskChan chan int

	logger := logger.Init()

	var components []WorkerComponent

	components = append(components,
		InitComponent(
			CLONER,
			cfg,
			logger,
			url,
			statusChan,
		),
		InitComponent(
			TESTER,
			cfg,
			logger,
			url,
			statusChan,
		),
		InitComponent(
			BUILDER,
			cfg,
			logger,
			url,
			statusChan,
		),
		InitComponent(
			LINTER,
			cfg,
			logger,
			url,
			statusChan,
		))

	if cfg.NotifyCfg.Send {
		components = append(components, InitComponent(
			SENDER,
			cfg,
			logger,
			url,
			statusChan,
		))
	}

	return &Worker{
		cfg:        cfg,
		logger:     logger,
		components: components,
		statusChan: statusChan,
		taskChan:   taskChan,
	}
}

func (w *Worker) Run() error {
	for i, c := range w.components {

		w.taskChan <- i

		if err := c.Run(); err != nil {
			w.logger.Error("error run component", zapcore.Field{
				Key:    "err",
				String: err.Error(),
			})
		}
	}

	return nil
}
