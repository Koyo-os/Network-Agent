package worker

import (
	"os/exec"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type Cloner struct{
	cfg *config.Config
	logger *logger.Logger
	url string
}

func initCloner(cfg *config.Config, logger *logger.Logger, url string) *Cloner {
	return &Cloner{
		cfg: cfg,
		logger: logger,
	}
}

func (c *Cloner) Run() error {
	cmd := exec.Command("git", "clone", c.url, c.cfg.TempDirName)
}