package version

import "runtime"

var (
	Version   = "0.0.1"
	GoVersion = runtime.Version()
)
