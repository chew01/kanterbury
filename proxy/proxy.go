package proxy

import (
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

type Proxy struct {
	Server *goproxy.ProxyHttpServer
	State  *GameState
	log.Logger
}

// NewProxy returns a proxy instance and generates a key pair if not present in executable directory
func NewProxy() *Proxy {
	server := goproxy.NewProxyHttpServer()
	state := newGameState()

	_, certStatErr := os.Stat(certPath)
	_, keyStatErr := os.Stat(keyPath)
	proxy := &Proxy{
		Server: server,
		State:  state,
		Logger: log.New(true, "/dev/null", 0),
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

	p.Printf("Proxy server listening on %s:8080\n", ipstring)
	utils.Must(http.ListenAndServe(":8080", p.Server))
}
