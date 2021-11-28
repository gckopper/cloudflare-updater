package ip

import (
	"log"
	"net"
)

// Func that returns the new ip
func GetIP() net.IP {
	conn, err := net.Dial("udp", "[2620:119:35::35]:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func ToVersion(iptype string) int8 {
	switch iptype {
	case "A":
		return 4
	case "AAAA":
		return 6
	default:
		log.Fatal("invalid or unsuported record type")
	}
	return 0
}
