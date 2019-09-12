package action

import (
	"edrtest/config"
	_ "strings"
	"testing"
)

func buildFileOpEnv(fileOpType config.FileOpType) (env *config.Env) {

	env = new(config.Env)
	env.ConfigType = config.FileActionConfigType

	fileOpConfig := new(config.FileActionConfig)
	env.FileActionConfig = fileOpConfig
	env.FileActionConfig.Action = fileOpType
	path := "/home/will/Development/golang/src/edrtest/blah.txt"
	fileOpConfig.Path = &path

	return env
}

func TestRunFileOpCreateAction(t *testing.T) {

	actionAction := NewAction(config.FileActionConfigType)

	env := buildFileOpEnv(config.CreateFile)

	actionAction.Setup(env)
	generic := actionAction.Run(env)
	result := generic.(*FileOpResult)
	if result.Err != nil {
		t.Errorf("failed creating file : %s\n", result.Err)
	}
}

func TestRunFileOpModifyAction(t *testing.T) {

	actionAction := NewAction(config.FileActionConfigType)

	env := buildFileOpEnv(config.CreateFile)

	actionAction.Setup(env)
	generic := actionAction.Run(env)
	result := generic.(*FileOpResult)
	if result.Err != nil {
		t.Errorf("failed creating file : %s\n", result.Err)
	}
}

func TestRunFileOpDeleteAction(t *testing.T) {

	actionAction := NewAction(config.FileActionConfigType)

	env := buildFileOpEnv(config.DeleteFile)

	actionAction.Setup(env)
	generic := actionAction.Run(env)
	result := generic.(*FileOpResult)
	if result.Err != nil {
		t.Errorf("failed creating file : %s\n", result.Err)
	}
}
