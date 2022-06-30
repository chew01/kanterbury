package main

import (
	"fmt"
	"github.com/chew01/kanterbury/proxy"
	"github.com/chew01/kanterbury/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	p := proxy.NewProxy()

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
