package utils

import (
	"net"
	"os"
	"path/filepath"
)

// BinDir returns the path of the executable
var BinDir = func() string {
	ex, err := os.Executable()
	Must(err)
	return filepath.Clean(filepath.Dir(ex) + "/")
}()

// GetOutboundIP returns preferred outbound IP of the machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer Must(conn.Close())

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// Must is a panic check for must-fail errors
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
