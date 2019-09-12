package resultlog

import (
	"edrtest/action"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func LogFileOpResult(result *action.FileOpResult) {
	cfg := result.Config
	log := cfg.Logger

	log.WithFields(logrus.Fields{
		"timestamp":          result.ActitivityTimestamp.Format(time.RFC3339),
		"Action":             "FileOpProcess",
		"ActivityDescriptor": cfg.Action.String(),
		"UserName":           fmt.Sprintf("%s", result.UserName.Username),
		"ProcessCmdLine":     result.ProcessCmdLine,
		"ProcessID":          result.ProcessID,
	}).Info()
}
