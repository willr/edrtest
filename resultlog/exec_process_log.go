package resultlog

import (
	"edrtest/action"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func LogExecResult(result *action.ExecResult) {
	cfg := result.Config
	log := cfg.Logger

	log.WithFields(logrus.Fields{
		"timestamp":      result.ActitivityTimestamp.Format(time.RFC3339),
		"Action":         "ExecProcess",
		"UserName":       fmt.Sprintf("%s", result.UserName.Username),
		"ProcessName":    result.ProcessName,
		"ProcessCmdLine": result.ProcessCmdLine,
		"ProcessID":      result.ProcessID,
	}).Info()
}
