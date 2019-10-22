package fw

type Environment interface {
	GetEnv(key string, defaultValue string) string
	AutoLoadDotEnvFile()
}
