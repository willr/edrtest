package action

import (
	"edrtest/config"
	"fmt"
	"os"
)

type FileOpAction struct {
	Actioner
}

type FileOpResult struct {
	CommonResult
	Config *config.FileActionConfig
	Op     config.FileOpType
	Path   *string
}

func NewFileOpAction() (action Actioner) {

	action = new(FileOpAction)
	return action
}

func newFileOpResult(runcfg *config.Env) (result *FileOpResult) {
	result = new(FileOpResult)

	return result
}

func (fileOpAction *FileOpAction) Setup(runcfg *config.Env) {

}

func (fileOpAction *FileOpAction) Run(env *config.Env) interface{} {

	result := newFileOpResult(env)

	result.runSubCmd(env)

	return result
}

func (result *FileOpResult) createFile(env *config.Env) {
	fileOpCfg := env.FileActionConfig

	dataToWrite := []byte("hello\n")
	f, err := os.Create(*fileOpCfg.Path)
	defer f.Close()
	result.SetTimestamp()
	if err != nil {
		result.Err = err
		return
	}

	f.Write(dataToWrite)
	if err != nil {
		result.Err = err
		return
	}

	f.Sync()
}

func (result *FileOpResult) modifyFile(env *config.Env) {
	fileOpCfg := env.FileActionConfig

	dataToWrite := []byte("hello - modified\n")
	f, err := os.OpenFile(*fileOpCfg.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	result.SetTimestamp()
	if err != nil {
		result.Err = err
	}

	_, err = f.Write(dataToWrite)
	if err != nil {
		result.Err = err
		return
	}

	f.Sync()
}

func (result *FileOpResult) deleteFile(env *config.Env) {
	fileOpCfg := env.FileActionConfig

	err := os.Remove(*fileOpCfg.Path)
	result.SetTimestamp()
	if err != nil {
		result.Err = err
	}
}

func (result *FileOpResult) runSubCmd(env *config.Env) {

	fileOpSubCmd := env.FileActionConfig.Action

	switch fileOpSubCmd {
	case config.CreateFile:
		result.Op = config.CreateFile
		result.createFile(env)
	case config.ModifyFile:
		result.Op = config.ModifyFile
		result.modifyFile(env)
	case config.DeleteFile:
		result.Op = config.DeleteFile
		result.deleteFile(env)
	default:
		result.Err = fmt.Errorf("File Operation not handled: %d", fileOpSubCmd)
	}
	result.updateResult(env.FileActionConfig)
}

func (result *FileOpResult) updateResult(fileOpCfg *config.FileActionConfig) {

	result.SetCurrentUser()
	result.ProcessID = os.Getpid()
	result.ProcessCmdLine = os.Args
	result.ProcessName = &os.Args[0]
	result.Config = fileOpCfg
}
