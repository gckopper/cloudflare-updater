package ip

import (
	"log"
	"net"
)

// Func that returns the new ip (version is based on the int passed to the func)
func GetIP(version int) net.IP {
	dnsserver := " "
	switch version {
	case 4:
		dnsserver = "8.8.8.8:80"
	case 6:
		dnsserver = "[2620:119:35::35]:80"
	}
	if dnsserver == " " {
		log.Fatal("Unsuported ip version: ", version)
	}
	conn, err := net.Dial("udp", dnsserver)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Converts a dns type record to ip version
func ToVersion(iptype string) int {
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
