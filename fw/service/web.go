package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/logger"

	"github.com/short-d/app/modern/mdio"
)

type WebServer struct {
	mux    *http.ServeMux
	server *http.Server
	logger logger.Logger
}

func (s *WebServer) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	s.server = &http.Server{Addr: addr, Handler: s.mux}
	err := s.server.ListenAndServe()

	if err == nil || err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (s WebServer) Shutdown() error {
	return s.server.Shutdown(context.Background())
}

func (s WebServer) HandleFunc(pattern string, handler http.Handler) {
	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		w = setupPreFlight(w)
		if (*r).Method == "OPTIONS" {
			return
		}

		w = enableCors(w)
		r.Body = mdio.Tap(r.Body, func(body string) {
			s.logger.Info(fmt.Sprintf("HTTP: url=%s host=%s method=%s body=%s", r.URL, r.Host, r.Method, body))
		})
		handler.ServeHTTP(w, r)
	})
}

func setupPreFlight(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	return w
}

func enableCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func NewWebServer(logger logger.Logger) WebServer {
	mux := http.NewServeMux()
	return WebServer{
		mux:    mux,
		logger: logger,
	}
}
