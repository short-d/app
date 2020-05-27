package env

var _ Env = (*Stub)(nil)

type Stub struct {
	envs map[string]string
}

func (s Stub) GetVar(key string, defaultValue string) string {
	val, ok := s.envs[key]
	if !ok {
		return defaultValue
	}
	return val
}

func (s Stub) AutoLoadDotEnvFile() {
}

func NewStub(envs map[string]string) Stub {
	return Stub{envs: envs}
}
