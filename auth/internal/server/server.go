package server

import (
	"auth/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	cfgDB  *config.DBConfig
	router *gin.Engine
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%s", s.cfgDB.Host, s.cfgDB.Port)
	return s.router.Run(address)
}

func (s *Server) NewServer(cfg *config.Config, cfgDB *config.DBConfig) (*Server, error) {
	s.cfg = cfg
	s.cfgDB = cfgDB
	s.router = gin.Default()

	return s, nil
}
