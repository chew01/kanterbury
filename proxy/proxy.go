package proxy

import (
	"fmt"
	"github.com/chew01/kanterbury/utils"
	"github.com/elazarl/goproxy"
	"os"
	"time"
)

const (
	certPath = "cert.pem"
	keyPath  = "key.pem"
)

type Proxy struct {
	Server *goproxy.ProxyHttpServer
	State  *GameState
}

func NewProxy() *Proxy {
	server := goproxy.NewProxyHttpServer()
	state := &GameState{
		Player:    &PlayerData{StartTime: time.Now().Unix()},
		Character: &CharacterData{},
		Activity:  &ActivityData{},
	}

	_, certStatErr := os.Stat(certPath)
	_, keyStatErr := os.Stat(keyPath)
	proxy := &Proxy{Server: server, State: state}

	if os.IsNotExist(certStatErr) || os.IsNotExist(keyStatErr) {
		fmt.Println("Generating CA...")
		if err := utils.GenerateCA(certPath, keyPath); err != nil {
			panic(err)
		}
		println("Cert and key saved")
	} else {
		utils.Must(certStatErr)
		utils.Must(keyStatErr)
	}

	server.OnRequest().DoFunc(proxy.handleReq)
	server.OnResponse().DoFunc(proxy.handleRes)

	server.OnRequest().HandleConnect(goproxy.FuncHttpsHandler(proxy.handleHttps))

	return proxy
}
