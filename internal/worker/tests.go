package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type Tester struct{
	cfg *config.Config
	logger *logger.Logger
}

func initTester(cfg *config.Config, logger *logger.Logger) *Tester {
	return &Tester{
		cfg: cfg,
		logger: logger,
	}
}

func (t *Tester) Run()