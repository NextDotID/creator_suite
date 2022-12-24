package telnet

import (
	"fmt"
	"net"
	"time"
)

func isValidIp(ip string) bool {
	return net.ParseIP(ip) == nil
}

// TelnetIPPortTimeout telnet ip+port with timeout
func TelnetIPPortTimeout(host string, port int, timeout int) error {
	ip := host
	if !isValidIp(host) {
		ipaddr, err := net.LookupIP("host")
		if err != nil {
			return fmt.Errorf("lookup host err: %v", err.Error())
		}
		ip = string(ipaddr[0])
	}
	telnetParam := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", telnetParam, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return fmt.Errorf("create telnet connect err: %v", err.Error())
	}
	defer conn.Close()
	return nil
}
