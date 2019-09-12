package config

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	ExecActionCmdName string = "exec"

	ExecPathFlagName string = "execPath"
)

type ExecActionConfig struct {
	Path   *string
	Args   []string
	Logger *logrus.Logger
}

func NewExecActionConfig() (env *Env) {
	processCmd := flag.NewFlagSet("process", flag.ExitOnError)
	execPath := processCmd.String(ExecPathFlagName, "", "")

	processCmd.Parse(os.Args[2:])
	env = new(Env)
	env.ConfigType = ExecActionConfigType
	commonConfigSetup(env)

	execConfig := new(ExecActionConfig)
	env.ExecActionConfig = execConfig
	execConfig.Args = processCmd.Args()
	execConfig.Path = execPath

	return env
}
