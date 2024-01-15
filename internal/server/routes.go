package server

import (
	"net/http"
	"webpDAV/internal/webdav"
)

func (s *Server) webDavHandler() http.Handler {
	handler := webdav.CreateHandler("./data")
	return handler
}
