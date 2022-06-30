package utils

import (
	"net"
)

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

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
