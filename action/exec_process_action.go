package action

import (
	"edrtest/config"
	"os"
	"os/exec"
)

type ExecProcessAction struct {
	Actioner
}

type ExecResult struct {
	CommonResult
	Config         *config.ExecActionConfig
	ExecProcessPid int
}

func start(args ...string) (p *os.Process, err error, procAttr os.ProcAttr) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)

		if err == nil {
			return p, nil, procAttr
		}
	}

	return nil, err, procAttr
}

func NewExecProcessAction() (action Actioner) {

	action = new(ExecProcessAction)
	return action
}

func newExecResult(runcfg *config.Env) (result *ExecResult) {
	result = new(ExecResult)

	return result
}

func (execAction *ExecProcessAction) Setup(runcfg *config.Env) {

}

func (execAction *ExecProcessAction) Run(env *config.Env) interface{} {

	execCmd := exec.Command(*env.ExecActionConfig.Path, env.ExecActionConfig.Args...)

	result := newExecResult(env)
	// execIn, err := execCmd.StdinPipe()
	_, err := execCmd.StdinPipe()
	if err != nil {
		result.Err = err
		return result
	}
	// execOut, err := execCmd.StdoutPipe()
	_, err = execCmd.StdoutPipe()
	if err != nil {
		result.Err = err
		return result
	}
	err = execCmd.Start()
	if err != nil {
		result.Err = err
		return result
	}
	result.SetTimestamp()
	result.SetCurrentUser()

	// execOut, err = ioutil.ReadAll(execOut)
	execCmd.Wait()

	result.ExecProcessPid = execCmd.Process.Pid
	result.updateResult(execCmd, env.ExecActionConfig)

	return result
}

func (result *ExecResult) updateResult(execCmd *exec.Cmd, execCfg *config.ExecActionConfig) {

	result.ProcessID = os.Getpid()
	result.ProcessCmdLine = os.Args
	result.ProcessName = &os.Args[0]
	result.Config = execCfg
}
