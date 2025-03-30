package config

import (
	"errors"
	"runtime/debug"
)

// GetVersion returns the current version based on the module's version which is dictated by the tag version.
func GetVersion() (string, error) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return "", errors.New("unable to read version from the build info")
	}

	if len(buildInfo.Main.Version) == 0 {
		return "unknown", nil
	}

	return buildInfo.Main.Version, nil
}
