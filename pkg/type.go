package bumpversion

import "github.com/coreos/go-semver/semver"

// BumpVersionConfig ...
type BumpVersionConfig struct {
	currentVersion *semver.Version
}
