package golwip

import (
	"net"
)

// TCPConnHandler handles TCP connections comming from TUN.
type TCPConnHandler interface {
	// Handle handles the conn for target.
	Handle(conn net.Conn, target *net.TCPAddr) error
}

// UDPConnHandler handles UDP connections comming from TUN.
type UDPConnHandler interface {
	// Connect connects the proxy server. Note that target can be nil.
	Connect(conn UDPConn, target *net.UDPAddr) error

	// ReceiveTo will be called when data arrives from TUN.
	ReceiveTo(conn UDPConn, data []byte, addr *net.UDPAddr) error
}

// DnsHandler handles DNS
type DnsHandler interface {
	ResolveIP(host string) (net.IP, error)
}

var tcpConnHandler TCPConnHandler
var udpConnHandler UDPConnHandler
var dnsHandler DnsHandler

func RegisterTCPConnHandler(h TCPConnHandler) {
	tcpConnHandler = h
}

func RegisterUDPConnHandler(h UDPConnHandler) {
	udpConnHandler = h
}

func RegisterDnsHandler(h DnsHandler) {
	dnsHandler = h
}
