package details

import (
	"net"
	"os"
	"strings"
)

func GetHostName() (string, error) {
	hostName, err := os.Hostname()

	return hostName, err
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}
