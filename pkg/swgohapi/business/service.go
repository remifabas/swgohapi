package business

import (
	"net/http"
	"time"

	guild "github.com/remifabas/swgohapi/pkg/guild"
)

type SWgohConfig struct {
	URL string // http://api.swgoh.gg/player/
}

type Service struct {
	config    SWgohConfig
	svcLogray guild.Service

	httpClient http.Client
}

func NewService(cfg SWgohConfig, g guild.Service) *Service {
	return &Service{
		config:     cfg,
		svcLogray:  g,
		httpClient: http.Client{Timeout: time.Duration(1) * time.Second},
	}
}
