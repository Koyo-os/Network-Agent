package notify

import (
	"context"

	"github.com/koyo-os/nagent/internal/config"
	"github.com/koyo-os/nagent/pkg/logger"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

type NotifyTheme string

var OK_AGENT NotifyTheme = "ok_agent"
var ERROR_AGENT NotifyTheme = "error_agent"

const ERR_MESSAGE = `
Hi, i am Murocami CI

your service cant build, check it please`

const SUCCESS_MESSAGE = `
Hi, i am Murocami CI

your service successfully get test and build
Enjoy!<3`

type Notifyler struct {
	cfg    *config.Config
	tg     *telegram.Telegram
	logger *logger.Logger
}

func Init(cfg *config.Config) (*Notifyler, error) {
	service, err := telegram.New(cfg.NotifyCfg.Token)
	if err != nil {
		return nil, err
	}

	return &Notifyler{
		tg:     service,
		logger: logger.Init(),
	}, nil
}

func (n *Notifyler) Send(themeid NotifyTheme) error {
	n.tg.AddReceivers(n.cfg.NotifyCfg.ChatId)

	notify.UseServices(n.tg)

	switch themeid {
	case OK_AGENT:
		return notify.Send(
			context.Background(),
			"Murocami CI",
			SUCCESS_MESSAGE,
		)
	case ERROR_AGENT:
		return notify.Send(
			context.Background(),
			"Murocami CI",
			ERR_MESSAGE,
		)
	default:
		return nil
	}
}
