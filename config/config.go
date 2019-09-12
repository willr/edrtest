package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type ActionConfigType int

const (
	ExecActionConfigType ActionConfigType = iota
	FileActionConfigType
	NetActionConfigType
)

type Env struct {
	ConfigType       ActionConfigType
	ExecActionConfig *ExecActionConfig
	FileActionConfig *FileActionConfig
	NetActionConfig  *NetActionConfig
}

func NewLogger(f *os.File) (logger *logrus.Logger) {
	logger = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	logger.SetOutput(f)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.DebugLevel)

	return logger
}

func commonConfigSetup(env *Env) {

}

func Parse(f *os.File) (env *Env, err error) {

	if len(os.Args) < 2 {
		err = fmt.Errorf(fmt.Sprintf("expected one of sub-commands: (%s | %s | %s)",
			ExecActionCmdName, FileActionCmdName, NetActionCmdName))
		return nil, err
	}

	switch os.Args[1] {
	case ExecActionCmdName:
		env = NewExecActionConfig()
		env.ExecActionConfig.Logger = NewLogger(f)
	case FileActionCmdName:
		env, err = NewFileActionConfig()
		env.FileActionConfig.Logger = NewLogger(f)
		if err != nil {
			fmt.Printf("failed to setup FileAction config")
			return nil, err
		}
	case NetActionCmdName:
		env, err = NewNetworkActionConfig()
		env.NetActionConfig.Logger = NewLogger(f)
		if err != nil {
			fmt.Printf("failed to setup NetworkAction config")
			return nil, err
		}
	default:
		err = fmt.Errorf(fmt.Sprintf("expected one of sub-commands: (%s | %s | %s)",
			ExecActionCmdName, FileActionCmdName, NetActionCmdName))
	}

	return env, err
}
