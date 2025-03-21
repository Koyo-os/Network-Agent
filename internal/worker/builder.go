package worker

import (
	"os/exec"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type Builder struct {
	cfg    *config.Config
	logger *logger.Logger
}

func initBuiler(cfg *config.Config, logger *logger.Logger) *Builder {
	return &Builder{
		cfg:    cfg,
		logger: logger,
	}
}

func (b *Builder) Run() error {
	cmd := exec.Command("go", "build", "-o", b.cfg.BuildCfg.OutputPoint, b.cfg.BuildCfg.InputPoint)

}
