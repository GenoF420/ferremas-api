package server

import (
	"github.com/charmbracelet/log"
	"github.com/genof420/ferremas-api/internal/config"
	"github.com/genof420/ferremas-api/internal/database"
	"github.com/labstack/echo/v4"
	"net"
	"strconv"
)

type Server struct {
	Application *echo.Echo
	Config      *config.Config
	Database    database.Database
}

func New(cfg *config.Config) (*Server, error) {
	db, err := database.New(cfg)
	if err != nil {
		return nil, err
	}

	echoEngine := echo.New()
	echoEngine.Debug = cfg.General.Debug

	server := &Server{
		Application: echoEngine,
		Config:      cfg,
		Database:    db,
	}

	return server, nil
}

func (s *Server) Start() error {
	addr := net.JoinHostPort(s.Config.HTTP.Address, strconv.Itoa(s.Config.HTTP.Port))

	if !s.Config.HTTP.SSL {
		log.Warn("SSL is disabled. This is not recommended for production use.")
		return s.Application.Start(addr)
	}

	return s.Application.StartTLS(addr, s.Config.HTTP.SSLCert, s.Config.HTTP.SSLKey)
}
