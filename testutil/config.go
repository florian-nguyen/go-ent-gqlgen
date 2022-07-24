package testutil

import (
	"go-ent-gqlgen/config"
	"go-ent-gqlgen/pkg/util/environment"
)

func ReadConfig() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.Test,
	})
}

func ReadConfigE2E() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.E2E,
	})
}
