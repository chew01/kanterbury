package proxy

import (
	"fmt"
	"github.com/chew01/kanterbury/utils"
	"github.com/elazarl/goproxy"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func (p *Proxy) Start() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Shutting down")
		os.Exit(0)
	}()

	ipstring := utils.GetOutboundIP()

	fmt.Printf("Proxy server listening on %s:8080\n", ipstring)
	utils.Must(http.ListenAndServe(":8080", p.Server))
}
