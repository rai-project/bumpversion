package bumpversion

import (
	"github.com/sirupsen/logrus"

	logger "github.com/rai-project/logger"
)

var (
	log *logrus.Entry
)

func init() {
	log = logger.New().WithField("pkg", "bumpversion")
}
