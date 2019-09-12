package action

import (
	"edrtest/config"
	_ "fmt"
	"os/user"
	"time"
)

type Actioner interface {
	Setup(*config.Env)
	Run(*config.Env) interface{}
}

type CommonResult struct {
	ActitivityTimestamp time.Time
	UserName            *user.User
	Err                 error
	ProcessName         *string
	ProcessCmdLine      []string
	ProcessID           int
}

func (result *CommonResult) SetCurrentUser() (err error) {

	user, err := user.Current()
	if err == nil {
		result.UserName = user
	}

	return err
}

func (result *CommonResult) SetTimestamp() {

	tstamp := time.Now()

	result.ActitivityTimestamp = tstamp
}
