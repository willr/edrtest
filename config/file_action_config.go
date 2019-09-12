package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	FileActionCmdName string = "file"

	FileSubCmdCreate string = "create"
	FileSubCmdModify string = "modify"
	FileSubCmdDelete string = "delete"

	FilePathFlagName string = "filePath"
)

type FileOpType int

const (
	Unknown FileOpType = iota
	CreateFile
	ModifyFile
	DeleteFile
)

func (e FileOpType) String() string {
	switch e {
	case CreateFile:
		return "create"
	case ModifyFile:
		return "modify"
	case DeleteFile:
		return "delete"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type FileActionConfig struct {
	Path   *string
	Args   []string
	Action FileOpType
	Logger *logrus.Logger
}

func parseFileActionType(arg string) (actionType FileOpType, err error) {
	lowerArg := strings.ToLower(arg)

	switch lowerArg {
	case FileSubCmdCreate:
		actionType = CreateFile
	case FileSubCmdModify:
		actionType = ModifyFile
	case FileSubCmdDelete:
		actionType = DeleteFile
	default:
		err = fmt.Errorf(fmt.Sprintf("Invalid file action: '%s'", lowerArg))
	}
	return actionType, err
}

func NewFileActionConfig() (env *Env, err error) {
	fileCmd := flag.NewFlagSet("file", flag.ExitOnError)
	filePath := fileCmd.String(FilePathFlagName, "", "")

	fileCmd.Parse(os.Args[3:])

	env = new(Env)
	env.ConfigType = FileActionConfigType
	env.FileActionConfig = new(FileActionConfig)
	commonConfigSetup(env)

	action, err := parseFileActionType(os.Args[2])
	if err != nil {
		return nil, err
	}
	fileConfig := new(FileActionConfig)
	env.FileActionConfig = fileConfig
	fileConfig.Action = action
	fileConfig.Path = filePath

	return env, nil
}
