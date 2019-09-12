package resultlog

import (
	"edrtest/action"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func LogNetConResult(result *action.NetConResult) {
	cfg := result.Config
	log := cfg.Logger

	log.WithFields(logrus.Fields{
		"timestamp":      result.ActitivityTimestamp.Format(time.RFC3339),
		"NetDestination": result.RemoteAddress,
		"NetSource":      result.LocalAddress,
		"DataSize":       result.CountDataSent,
		"Protocol":       result.ProtocolSent,
		"UserName":       fmt.Sprintf("%s", result.UserName.Username),
		"ProcessName":    result.ProcessName,
		"ProcessCmdLine": result.ProcessCmdLine,
		"ProcessID":      result.ProcessID,
	}).Info()
}
