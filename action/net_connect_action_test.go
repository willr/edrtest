package action

import (
	"edrtest/config"
	"fmt"
	"net/url"
	_ "strings"
	"testing"
)

func buildNetConEnv(t *testing.T) (env *config.Env) {

	env = new(config.Env)
	env.ConfigType = config.NetActionConfigType

	netConnConfig := new(config.NetActionConfig)
	env.NetActionConfig = netConnConfig
	urlStr := "http://192.168.25.1"
	u, err := url.Parse(urlStr)
	netConnConfig.Host = u.Host
	netConnConfig.Port = u.Port()
	netConnConfig.Scheme = u.Scheme
	netConnConfig.Logger = config.NewLogger(nil)
	if err != nil {
		fmt.Printf("error parsing: %s", err)
	}

	switch netConnConfig.Scheme {
	case "http":
		if netConnConfig.Port == "" {
			netConnConfig.Port = "80"
		}
	case "https":
		if netConnConfig.Port == "" {
			netConnConfig.Port = "443"
		}
	}

	return env
}

func TestRunNetConnectAction(t *testing.T) {

	actionAction := NewAction(config.NetActionConfigType)

	env := buildNetConEnv(t)

	actionAction.Setup(env)
	t.Logf("trying connection to: %s:%s", env.NetActionConfig.Host, env.NetActionConfig.Port)
	generic := actionAction.Run(env)
	result := generic.(*NetConResult)
	if result.Err != nil {
		t.Errorf("failed creating file : %s\n", result.Err)
	}
}
