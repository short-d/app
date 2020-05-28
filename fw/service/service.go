package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Service interface {
	Stop(ctx context.Context, cancel context.CancelFunc)
	StartAsync(port int)
	StartAndWait(port int)
}

func listenForSignals(s Service) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	s.Stop(ctx, cancel)
}
