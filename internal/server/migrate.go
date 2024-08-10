package server

import "demo/internal/model/health"

func (s *Server) migrate() {
    health.AutoMigrate(s.DB.GetDB())
}
