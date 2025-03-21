package worker

import (
	"os/exec"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Tester struct {
	cfg    *config.Config
	logger *logger.Logger
}

func initTester(cfg *config.Config, logger *logger.Logger) *Tester {
	return &Tester{
		cfg:    cfg,
		logger: logger,
	}
}

func (t *Tester) Run() error {
	cmd := exec.Command("go", "test", "...")
	cmd.Dir = t.cfg.TempDirName

	if err := cmd.Run(); err != nil {
		t.logger.Error("cant run tests", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})
		return err
	}

	return nil
}

