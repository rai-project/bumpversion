package bumpversion

import (
	configparser "github.com/bigkevmcd/go-configparser"
	"github.com/coreos/go-semver/semver"
	"github.com/pkg/errors"
)

var (
	invalidConfig = errors.New("invalid .bumpversion.cfg")
)

// Read ...
func Read(filePath string) (*BumpVersionConfig, error) {
	p, err := configparser.NewConfigParserFromFile(filePath)

	if !p.HasSection("bumpversion") {
		return nil, errors.Wrap(invalidConfig, "expecting a bumpversion section")
	}

	currentVersionString, err := p.Get("bumpversion", "current_version")
	if err != nil {
		return nil, errors.Wrap(invalidConfig, "expecting a current_version")
	}
	currentVersion, err := semver.NewVersion(currentVersionString)
	if err != nil {
		return nil, errors.Wrap(invalidConfig, "unable to parse current_version as a semantic version")
	}
	return &BumpVersionConfig{
		currentVersion: currentVersion,
	}, nil
}
