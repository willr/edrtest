package main

import (
	"edrtest/action"
	"edrtest/config"
	"edrtest/resultlog"
	"fmt"
	"os"
)

func setupLogfile() (f *os.File) {
	// open a file
	f, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	return f
}

func main() {

	// open a file
	f := setupLogfile()
	// don't forget to close it
	defer f.Close()

	env, err := config.Parse(f)
	if err != nil {
		panic(err)
	}
	actionAction := action.NewAction(env.ConfigType)
	actionAction.Setup(env)
	generic := actionAction.Run(env)

	resultlog.LogResult(env.ConfigType, generic)

}
