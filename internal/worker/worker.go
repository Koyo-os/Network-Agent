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
	taskChan   chan string
	components []WorkerComponent
	url        string
}

func Init(cfg *config.Config, url string) *Worker {
	var taskChan chan string
	var statusChan chan string

	logger := logger.Init()

	components = append(components,
		InitComponent(
			CLONER,
			cfg,
			logger,
			url,
			taskChan,
		),
		InitComponent(
			TESTER,
			cfg,
			logger,
			url,
			taskChan,
		),
		InitComponent(
			BUILDER,
			cfg,
			logger,
			url,
			taskChan,
		),
		InitComponent(
			LINTER,
			cfg,
			logger,
			url,
			taskChan,
		))

	var components []WorkerComponent
	if cfg.NotifyCfg.Send {
		components = append(components, InitComponent(
			SENDER,
			cfg,
			logger,
			url,
			taskChan,
		))
	}

	return &Worker{
		cfg:        cfg,
		logger:     logger,
		components: components,
		taskChan:   taskChan,
	}
}

func (w *Worker) Run() error {
	for _, c := range w.components {
		if err := c.Run(); err != nil {
			w.logger.Error("error run component", zapcore.Field{
				Key:    "err",
				String: err.Error(),
			})
		}
	}

	return nil
}
