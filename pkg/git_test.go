package bumpversion

import (
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	git, err := newGit(filepath.Dir(sourcepath.MustAbsoluteDir()))
	assert.NoError(t, err)
	assert.NotEmpty(t, git)
	git.LatestTagInfo()
}
