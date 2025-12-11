package version

import (
	"fmt"
	"runtime"
)

var (
	Version = "v0.0.0-dev"

	Commit = "none"

	Date = "unknown"
)

func Info() string {
	return fmt.Sprintf("gitgen version %s\ncommit: %s\nbuilt: %s\ngoos: %s\ngoarch: %s",
		Version, Commit, Date, runtime.GOOS, runtime.GOARCH)
}

func ShortInfo() string {
	return Version
}
