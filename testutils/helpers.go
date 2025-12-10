package testutils

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func UseTempHomeDir(t testing.TB) func() {
	t.Helper()
	tmp := t.TempDir()
	oldHome, had := os.LookupEnv("HOME")
	_ = os.Setenv("HOME", tmp)
	return func() {
		if had {
			_ = os.Setenv("HOME", oldHome)
		} else {
			_ = os.Unsetenv("HOME")
		}
		viper.Reset()
	}
}
