package env

import (
	"runtime"
	"strings"
)

const Version = "v1.0.0"

func GetGoVersionNum() string{
	versionWithPrefix := runtime.Version()
	split := strings.Split(versionWithPrefix, "go")
	version := split[1]
	split = strings.Split(version, ".")
	return split[0]+"."+split[1]
}

func IsDebug() bool{
	return true
}
