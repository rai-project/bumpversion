package bumpversion

import "github.com/coreos/go-semver/semver"

type BumpVersionConfig struct {
	currentVersion *semver.Version
}
