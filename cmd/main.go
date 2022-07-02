package main

import (
	"flag"
	"github.com/chew01/kanterbury/proxy"
	"log"
)

var logPath = flag.String("l", "logs/proxy.log", "file to output log to")
var port = flag.Int("p", 8080, "port for proxy server")

func main() {
	flag.Parse()
	logFlags := log.Lshortfile | log.Ltime

	options := &proxy.Options{
		LogPath:  *logPath,
		LogFlags: logFlags,
		Port:     *port,
	}

	p := proxy.NewProxy(options)
	p.Start()
}
