package kata

import (
	"strconv"
	"strings"
)

func IsValidIp(ip string) bool {
	parts := strings.Split(ip, ".")

	if len(parts) < 4 {
		return false
	}
	for _, part := range parts {
		if len(part) > 1 && part[0] == '0' {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
		if num > 255 {
			return false
		}
		if num < 0 {
			return false
		}
	}
	return true
}
