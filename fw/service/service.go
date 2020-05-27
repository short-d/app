package service

import "context"

type Service interface {
	Stop(ctx context.Context)
	StartAsync(port int)
	StartAndWait(port int)
}
