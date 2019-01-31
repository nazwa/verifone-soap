package verifone

import (
	"testing"

	"github.com/spf13/viper"
)

// This test relies on a valid config file in the bin directory.
// I know it could be done with mocks, but for now this was easier
// Feel free to improve!

func createClientFromConfig(t *testing.T) *Client {
	viper.SetConfigFile("./../../bin/config.toml")
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("Config file not loaded: %s", err)
	}

	return NewClient(Config{
		SystemID:   viper.GetString("verifone.SystemID"),
		SystemGUID: viper.GetString("verifone.SystemGUID"),
		Passcode:   viper.GetString("verifone.Passcode"),
		Url:        viper.GetString("verifone.Url"),
	})
}

func TestSum(t *testing.T) {
	client := createClientFromConfig(t)

	token, err := client.BeginSession()
	if err != nil {
		t.Error("Session not started", err)
	} else {
		if token == "" {
			t.Error("Session token empty")
		}
	}
}
