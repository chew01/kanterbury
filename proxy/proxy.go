package proxy

import (
	"fmt"
	"github.com/chew01/kanterbury/log"
	"github.com/chew01/kanterbury/utils"
	"github.com/elazarl/goproxy"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	certPath = "cert.pem"
	keyPath  = "key.pem"
)

type Options struct {
	LogPath  string
	LogFlags int
	Port     int
}

type Proxy struct {
	Options *Options
	Server  *goproxy.ProxyHttpServer
	State   *GameState
	log.Logger
}

// NewProxy returns a proxy instance and generates a key pair if not present in executable directory
func NewProxy(options *Options) *Proxy {
	server := goproxy.NewProxyHttpServer()
	state := newGameState()

	_, certStatErr := os.Stat(certPath)
	_, keyStatErr := os.Stat(keyPath)
	proxy := &Proxy{
		Options: options,
		Server:  server,
		State:   state,
		Logger:  log.New(true, options.LogPath, options.LogFlags),
	}

	if os.IsNotExist(certStatErr) || os.IsNotExist(keyStatErr) {
		proxy.Println("Certificate file was not found. Generating CA...")
		if err := utils.GenerateCA(certPath, keyPath); err != nil {
			panic(err)
		}
		proxy.Printf("Cert and key saved in %s\n", utils.BinDir)
	} else {
		utils.Must(certStatErr)
		utils.Must(keyStatErr)
	}

	server.OnRequest().DoFunc(proxy.handleReq)
	server.OnResponse().DoFunc(proxy.handleRes)

	server.OnRequest().HandleConnect(goproxy.FuncHttpsHandler(proxy.handleHttps))

	return proxy
}

// Start the proxy.
func (p *Proxy) Start() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		p.Println("Shutting down")
		os.Exit(0)
	}()

	ipstring := utils.GetOutboundIP()
	addr := fmt.Sprintf(":%d", p.Options.Port)

	p.Printf("Proxy server listening on %s%s\n", ipstring, addr)
	utils.Must(http.ListenAndServe(addr, p.Server))
}
