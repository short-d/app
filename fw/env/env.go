package env

type Env interface {
	GetVar(key string, defaultValue string) string
	AutoLoadDotEnvFile()
}

type Runtime string
