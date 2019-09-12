package action

import (
	"edrtest/config"
	_ "fmt"
)

func NewAction(actionType config.ActionConfigType) (action Actioner) {

	switch actionType {
	case config.ExecActionConfigType:
		action = NewExecProcessAction()
	case config.FileActionConfigType:
		action = NewFileOpAction()
	case config.NetActionConfigType:
		action = NewNetConAction()
	}

	return action
}
