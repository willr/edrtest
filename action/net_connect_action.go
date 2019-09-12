package action

import (
	"edrtest/config"
	"fmt"
	"net"
	"os"
)

type NetConAction struct {
	Actioner
}

type NetConResult struct {
	CommonResult
	Config        *config.NetActionConfig
	RemoteAddress *string
	LocalAddress  *string
	CountDataSent uint64
	ProtocolSent  *string
}

func NewNetConAction() (action Actioner) {

	action = new(NetConAction)
	return action
}

func newNetConResult(runcfg *config.Env) (result *NetConResult) {
	result = new(NetConResult)

	return result
}

func (netConAction *NetConAction) Setup(runcfg *config.Env) {

}

func (netConAction *NetConAction) Run(env *config.Env) interface{} {

	result := newNetConResult(env)

	result.sendRequest(env)

	result.updateResult(env.NetActionConfig)

	return result
}

func (result *NetConResult) sendRequest(env *config.Env) (err error) {

	// GET /docs/index.html HTTP/1.1
	// Host: www.nowhere123.com
	// Accept: image/gif, image/jpeg, */*
	// Accept-Language: en-us
	// User-Agent: Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)

	netConn := env.NetActionConfig
	remoteAddress := fmt.Sprintf("%s:%s", netConn.Host, netConn.Port)
	result.ProtocolSent = &netConn.Scheme
	result.RemoteAddress = &remoteAddress
	conn, err := net.Dial("tcp", remoteAddress)
	if err != nil {
		result.Err = err
		return err
	}
	result.SetTimestamp()
	localAddress := conn.LocalAddr().String()
	result.LocalAddress = &(localAddress)

	counter := NewCountWriter(conn)
	fmt.Fprintf(counter, "GET / HTTP/1.1\nHost: %s\nUser-Agent: curl/7.58.0\nAccept: */*\n\n", netConn.Host)
	// fmt.Printf("size: %d", counter.Count())
	result.CountDataSent = counter.Count()

	return nil
}

func (result *NetConResult) updateResult(netConCfg *config.NetActionConfig) {

	result.SetCurrentUser()
	result.Config = netConCfg
	result.ProcessID = os.Getpid()
	result.ProcessCmdLine = os.Args
	result.ProcessName = &os.Args[0]
}
