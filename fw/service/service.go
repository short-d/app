package service

// TODO(issue#67): support graceful shutdown.
type Service interface {
	Stop()
	StartAsync(port int)
	StartAndWait(port int)
}
