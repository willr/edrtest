package action

import (
	"edrtest/config"
	_ "strings"
	"testing"
)

func buildExecProcessEnv() (env *config.Env) {

	env = new(config.Env)
	env.ConfigType = config.ExecActionConfigType

	execConfig := new(config.ExecActionConfig)
	env.ExecActionConfig = execConfig
	execConfig.Logger = config.NewLogger(nil)

	args := make([]string, 1)
	args[0] = "/home/will/Development/golang/src/edrtest/action"
	execConfig.Args = args
	cmd := "head"
	execConfig.Path = &cmd

	return env
}

func TestRunExecProcessAction(t *testing.T) {

	actionAction := NewAction(config.ExecActionConfigType)

	env := buildExecProcessEnv()

	actionAction.Setup(env)
	generic := actionAction.Run(env)
	result := generic.(*ExecResult)
	if result.Err != nil {
		t.Errorf("failed execing process : %s\n", result.Err)
	}
}
