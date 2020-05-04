package service

type Service interface {
	Stop()
	StartAsync(port int)
	StartAndWait(port int)
}
