package worker

import (
	"os/exec"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Linter struct {
	cfg        *config.Config
	logger     *logger.Logger
	url        string
	statusChan chan string
}

func initLinter(cfg *config.Config, logger *logger.Logger, url string, statusChan chan string) *Linter {
	return &Linter{
		cfg:        cfg,
		logger:     logger,
		statusChan: statusChan,
		url:        url,
	}
}

func (l *Linter) Run() error {
	cmd := exec.Command("golangci-lint", "run")
	if err := cmd.Run(); err != nil {
		l.logger.Error("error linter", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})
	}

	l.logger.Info("success lint")

	return nil
}
