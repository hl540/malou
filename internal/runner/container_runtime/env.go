package container_runtime

import "fmt"

type EnvValue struct {
	Key   string
	Value string
}

func (v *EnvValue) String() string {
	return fmt.Sprintf("%s=%s", v.Key, v.Value)
}

func EnvsToArray(envs []*EnvValue) []string {
	var values []string
	for _, item := range envs {
		values = append(values, item.String())
	}
	return values
}
