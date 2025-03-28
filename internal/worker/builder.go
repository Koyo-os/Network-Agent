package worker

import (
	"os/exec"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Builder struct {
	cfg        *config.Config
	logger     *logger.Logger
	statusChan chan string
	url        string
}

func initBuiler(cfg *config.Config, logger *logger.Logger, url string, statusChan chan string) *Builder {
	return &Builder{
		cfg:        cfg,
		logger:     logger,
		statusChan: statusChan,
		url:        url,
	}
}

func (b *Builder) Run() error {
	b.statusChan <- "ok"

	cmd := exec.Command("go", "build", "-o", b.cfg.BuildCfg.OutputPoint, b.cfg.BuildCfg.InputPoint)
	if err := cmd.Run(); err != nil {
		b.logger.Error("cant build", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		b.statusChan <- "err build"

		return err
	}

	return nil
}
