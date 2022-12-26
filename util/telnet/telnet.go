package telnet

import (
	"fmt"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func isValidIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

func parseHttps(host string) string {
	h1 := strings.TrimPrefix(host, "http://")
	h2 := strings.TrimPrefix(h1, "https://")
	return h2
}

func parseIP(s string) (net.IP, int) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return ip, 4
		case ':':
			return ip, 6
		}
	}
	return nil, 0
}

// TelnetIPPortTimeout telnet ip+port with timeout
func TelnetIPPortTimeout(host string, port int, timeout int) error {
	parseHost := parseHttps(host)
	ip := parseHost
	if !isValidIp(parseHost) {
		ipaddrs, err := net.LookupIP(parseHost)
		if err != nil {
			return fmt.Errorf("lookup host err: %v", err.Error())
		}
		for _, ipaddr := range ipaddrs {
			// if isValidIp(ipaddr.String()) {
			// 	ip = string(ipaddr.String())
			// 	break
			// }
			_ip, flag := parseIP(ipaddr.String())
			if _ip != nil && flag == 4 {
				ip = _ip.String()
				break
			}
		}
	}
	telnetParam := fmt.Sprintf("%s:%d", ip, port)
	log.Infof("telnet %s", telnetParam)
	conn, err := net.DialTimeout("tcp", telnetParam, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return fmt.Errorf("create telnet connect err: %v", err.Error())
	}
	defer conn.Close()
	return nil
}
