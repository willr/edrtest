package config

import (
	"flag"
	"os"
	_ "strings"
	"testing"
)

const (
	CmdName string = "edrtest"
)

func ResetFlagsForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = usage
	flag.String("test,v", "", "")
}

func TestProcessStartConfig(t *testing.T) {
	ResetFlagsForTesting(func() { t.Fatal("bad parse") })
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{CmdName, ExecActionCmdName, "-execPath", "/test/path", "exec_args1", "exec_args2"}

	env, err := Parse(nil)
	if err != nil {
		t.Fatalf("Failed to create ProcessStart config: %s", err)
	}

	if env.ConfigType != ExecActionConfigType {
		t.Errorf("failed to set correct ConfigType const")
	}

	if env.ExecActionConfig == nil {
		t.Errorf("Failed to create ProcessStartConfig struct for cmdline options: '%s'", os.Args)
	}

	execCfg := env.ExecActionConfig
	if execCfg.Path != nil && *execCfg.Path == "/test/path" {
		t.Logf("Found path: %s", *execCfg.Path)
	} else {
		t.Errorf("Failed to set path of file to exec")
	}
	if execCfg.Args != nil && len(execCfg.Args) == 2 {
		if execCfg.Args[0] == "exec_args1" && execCfg.Args[1] == "exec_args2" {
			t.Logf("found args: %s & %s", execCfg.Args[0], execCfg.Args[1])
		} else {
			t.Errorf("unexepected args to file for exec: %s & %s", execCfg.Args[0], execCfg.Args[1])
		}
	} else {
		t.Errorf("Failed to set args or unexpected args of file to exec, args: %s", execCfg.Args)
	}
}

func TestFileActionConfig(t *testing.T) {
	ResetFlagsForTesting(func() { t.Fatal("bad parse") })
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{CmdName, FileActionCmdName, FileSubCmdCreate, "-filePath", "/test/path"}

	env, err := Parse(nil)
	if err != nil {
		t.Fatalf("Failed to create FileAction config: %s", err)
	}

	if env.ConfigType != FileActionConfigType {
		t.Errorf("Failed to set correct ConfigType const")
	}

	if env.FileActionConfig == nil {
		t.Errorf("Failed to create FileActionConfig struct for cmdline options: '%s'", os.Args)
	}

	fileCfg := env.FileActionConfig
	if fileCfg.Action != CreateFile {
		t.Logf("cmdline: '%s'", os.Args)
		t.Errorf("Failed to set file subcmd to create, value was: %s", fileCfg.Action)
	}
	if fileCfg.Path != nil && *fileCfg.Path == "/test/path" {
		t.Logf("Found path: %s", *fileCfg.Path)
	} else {
		var xpath string
		if fileCfg.Path != nil {
			xpath = *fileCfg.Path
		} else {
			xpath = "nil"
		}
		t.Errorf("Failed to return path of file to create/modify/delete: '%s'", xpath)
	}
}

func TestFileActionConfigSubCmd(t *testing.T) {
	ResetFlagsForTesting(func() { t.Fatal("bad parse") })
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{CmdName, FileActionCmdName, FileSubCmdModify, "-filePath", "/test/path"}

	env, err := Parse(nil)
	if err != nil {
		t.Fatalf("Failed to create FileAction config: %s", err)
	}

	if env.ConfigType != FileActionConfigType {
		t.Errorf("Failed to set correct ConfigType const")
	}

	if env.FileActionConfig == nil {
		t.Errorf("Failed to create FileActionConfig struct for cmdline options: '%s'", os.Args)
	}

	fileCfg := env.FileActionConfig
	if fileCfg.Action != ModifyFile {
		t.Logf("cmdline: '%s'", os.Args)
		t.Errorf("Failed to set file subcmd to %s, value was: %s", ModifyFile, fileCfg.Action)
	}

	ResetFlagsForTesting(func() { t.Fatal("bad parse") })

	os.Args = []string{CmdName, FileActionCmdName, FileSubCmdDelete, "-filePath", "/test/path"}

	env, err = Parse(nil)
	if err != nil {
		t.Fatalf("Failed to create FileAction config: %s", err)
	}

	if env.ConfigType != FileActionConfigType {
		t.Errorf("Failed to set correct ConfigType const")
	}

	if env.FileActionConfig == nil {
		t.Errorf("Failed to create FileActionConfig struct for cmdline options: '%s'", os.Args)
	}

	fileCfg = env.FileActionConfig
	if fileCfg.Action != DeleteFile {
		t.Logf("cmdline: '%s'", os.Args)
		t.Errorf("Failed to set file subcmd to %s, value was: %s", DeleteFile, fileCfg.Action)
	}
}

func TestNetActionConfig(t *testing.T) {
	ResetFlagsForTesting(func() { t.Fatal("bad parse") })
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{CmdName, "net", "-url", "http://www.google.com"}
}
