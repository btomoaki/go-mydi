package config

import (
	"fmt"
	reflect "reflect"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Port int `envconfig:"PORT" default:"8080"`
}

func (e *Env) LoadEnv() map[string]string {
	var goenv Env
	envMap := map[string]string{}
	envconfig.Process("", &goenv)
	values := reflect.Indirect(reflect.ValueOf(goenv))
	envType := values.Type()

	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		iface := field.Interface()
		fieldName := envType.Field(i).Name
		fmt.Println(fieldName)
		if value, ok := iface.(int); ok {
			envMap[fieldName] = strconv.Itoa(value)
		} else {
			envMap[fieldName] = field.String()
		}
	}
	return envMap
}
