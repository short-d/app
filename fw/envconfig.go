package fw

type EnvConfig interface {
	ParseConfigFromEnv(config interface{}) error
}
