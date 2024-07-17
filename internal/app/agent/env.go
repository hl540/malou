package agent

import "fmt"

type EnvValue struct {
	Key   string
	Value string
}

func (v *EnvValue) Strings() string {
	return fmt.Sprintf("%s=%s", v.Key, v.Value)
}

var envs = make(map[string]*EnvValue)

func GetEnvs() []string {
	var values []string
	for _, item := range envs {
		values = append(values, item.Strings())
	}
	return values
}
