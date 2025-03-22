package worker

import (
	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/internal/notify"
	"github.com/koyo-os/nagent/pkg/logger"
)

type Sender struct {
	cfg        *config.Config
	logger     *logger.Logger
	statusChan chan string
	url        string
	notify     *notify.Notifyler
}

func initSender(cfg *config.Config, logger *logger.Logger, url string, statusChan chan string, notify *notify.Notifyler) *Sender {
	return &Sender{
		cfg:        cfg,
		logger:     logger,
		url:        url,
		statusChan: statusChan,
		notify:     notify,
	}
}

func (s *Sender) Run() error {
	if SenderStatusOK {
		return s.notify.Send(notify.OK_AGENT)
	} else {
		return s.notify.Send(notify.ERROR_AGENT)
	}
}
