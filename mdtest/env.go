package mdtest

import "github.com/short-d/app/fw"

var _ fw.Environment = (*EnvironmentFake)(nil)

type EnvironmentFake struct {
	envs map[string]string
}

func (e EnvironmentFake) GetEnv(key string, defaultValue string) string {
	val, ok := e.envs[key]
	if !ok {
		return defaultValue
	}
	return val
}

func (e EnvironmentFake) AutoLoadDotEnvFile() {
}

func NewEnvironmentFake(envs map[string]string) EnvironmentFake {
	return EnvironmentFake{envs: envs}
}
