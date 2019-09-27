package kata

// parseip solution does not work for example test 123.045.067.089
import (
	"net"
)

func IsValidIp(ip string) bool {
	res := net.ParseIP(ip)
	return res != nil
}
