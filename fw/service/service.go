package service

import "context"

// TODO(issue#67): support graceful shutdown.
type Service interface {
	Stop(ctx context.Context)
	StartAsync(port int)
	StartAndWait(port int)
}
