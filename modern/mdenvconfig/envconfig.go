package envconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/short-d/app/fw"
)

var _ fw.EnvConfig = (*EnvConfig)(nil)

// EnvConfig parses configuration from environmental variables.
type EnvConfig struct {
	parseBool   func(newValue string, typeOfValue reflect.Type) (bool, error)
	parseInt    func(newValue string, typeOfValue reflect.Type) (int64, error)
	parseString func(newValue string, typeOfValue reflect.Type) (string, error)
	environment fw.Environment
}

// ParseConfigFromEnv retrieves configurations from environmental variables and
// parse them into the given struct.
func (e EnvConfig) ParseConfigFromEnv(config interface{}) error {
	configVal := reflect.ValueOf(config)
	if configVal.Kind() != reflect.Ptr {
		return errors.New("config must be a pointer")
	}

	if configVal.IsNil() {
		return errors.New("config can't be nil")
	}

	elem := configVal.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("config must be a struct")
	}

	numFields := elem.NumField()
	configType := elem.Type()

	for idx := 0; idx < numFields; idx++ {
		field := configType.Field(idx)
		envName, ok := field.Tag.Lookup("env")
		if !ok {
			continue
		}
		defaultVal := field.Tag.Get("default")
		envVal := e.environment.GetEnv(envName, defaultVal)
		err := e.setFieldValue(field, elem.Field(idx), envVal)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e EnvConfig) setFieldValue(field reflect.StructField, fieldValue reflect.Value, newValue string) error {
	kind := field.Type.Kind()
	switch kind {
	case reflect.Bool:
		boolean, err := strconv.ParseBool(newValue)
		if err != nil {
			return err
		}
		fieldValue.SetBool(boolean)
		return nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		num, err := e.parseInt(newValue, field.Type)
		if err != nil {
			return err
		}
		fieldValue.SetInt(num)
		return nil
	case reflect.String:
		str, err := e.parseString(newValue, field.Type)
		if err != nil {
			return err
		}
		fieldValue.SetString(str)
		return nil
	default:
		return fmt.Errorf("unexpected field type: %s", kind)
	}
}

// NewEnvConfig creates EnvConfig.
func NewEnvConfig(
	parseBool func(newValue string, typeOfValue reflect.Type) (bool, error),
	parseInt func(newValue string, typeOfValue reflect.Type) (int64, error),
	parseString func(newValue string, typeOfValue reflect.Type) (string, error),
	environment fw.Environment,
) EnvConfig {
	return EnvConfig{parseBool: parseBool, parseInt: parseInt, parseString: parseString, environment: environment}
}
