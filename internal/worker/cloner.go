package worker

import (
	"fmt"
	"os/exec"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
)

type Cloner struct{
	cfg *config.Config
	logger *logger.Logger
	url string
	statusChan chan string

}

func initCloner(cfg *config.Config, logger *logger.Logger, url string, statusChan chan string) *Cloner {
	return &Cloner{
		cfg: cfg,
		logger: logger,
		url: url,
		statusChan: statusChan,
	}
}

func (c *Cloner) Run() error {
	c.statusChan <- "ok"

	cmd := exec.Command("git", "clone", c.url, c.cfg.TempDirName)
	err := cmd.Run()
	if err != nil{
		c.statusChan <- fmt.Sprintf("error cloner: %v", err)
	}

	return nil
}