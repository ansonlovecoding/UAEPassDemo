package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Printf("Config.Env: %s \n", LocalConfig.Env)
	fmt.Printf("Config.Staging.authorization: %s \n", LocalConfig.Endpoints.Staging.Authorization)
}
