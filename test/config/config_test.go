package config

import (
	"github/linfengOu/blog-backend/config"
	"testing"
)

func TestConfig(t *testing.T) {
	configuration := config.Config
	if configuration.HTTPServer.Port != ":1234" {
		t.Errorf("HTTPServer.Port expect %s but %s present", ":1234", configuration.HTTPServer.Port)
	}
}
