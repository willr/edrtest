package config

import (
	"flag"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	NetActionCmdName string = "net"

	UrlFlagName string = "url"
)

type NetActionConfig struct {
	Scheme string
	Host   string
	Port   string
	Logger *logrus.Logger
}

func NewNetworkActionConfig() (env *Env, err error) {
	netCmd := flag.NewFlagSet("net", flag.ExitOnError)
	remoteRequest := netCmd.String(UrlFlagName, "", "")

	netCmd.Parse(os.Args[2:])

	env = new(Env)
	env.ConfigType = NetActionConfigType
	commonConfigSetup(env)

	netConfig := new(NetActionConfig)
	env.NetActionConfig = netConfig
	url, err := url.Parse(*remoteRequest)
	if err != nil {
		return nil, err
	}
	netConfig.Host = url.Hostname()
	netConfig.Port = url.Port()
	netConfig.Scheme = url.Scheme

	switch netConfig.Scheme {
	case "http":
		if netConfig.Port == "" {
			netConfig.Port = "80"
		}
	case "https":
		if netConfig.Port == "" {
			netConfig.Port = "443"
		}
	}

	return env, nil
}
