package rest

import (
	"goods/configs"
	"goods/internal/services"
	"goods/internal/transport/rest/handlers"
	"log"
	"net/http"
	"time"
)

func NewServer(conf configs.Config, logger *log.Logger, services *services.Services) *http.Server {
	return &http.Server{
		Addr:         conf.Addrs + conf.Port,
		Handler:      handlers.InitHandlers(logger, services),
		ReadTimeout:  15 * time.Microsecond,
		WriteTimeout: 15 * time.Microsecond,
	}
}
