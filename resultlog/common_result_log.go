package resultlog

import (
	"edrtest/action"
	"edrtest/config"
	"fmt"
)

func LogResult(rtype config.ActionConfigType, r interface{}) {

	switch rtype {
	case config.NetActionConfigType:
		result := r.(*action.NetConResult)
		if result.Err != nil {
			fmt.Printf("Action failed: %s\n", result.Err)
		}
		LogNetConResult(result)
	case config.FileActionConfigType:
		result := r.(*action.FileOpResult)
		if result.Err != nil {
			fmt.Printf("Action failed: %s\n", result.Err)
		}
		LogFileOpResult(result)
	case config.ExecActionConfigType:
		result := r.(*action.ExecResult)
		if result.Err != nil {
			fmt.Printf("Action failed: %s\n", result.Err)
		}
		LogExecResult(result)
	default:
		fmt.Printf("Failed to write resultlog\n")
	}
}
