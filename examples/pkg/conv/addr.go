package conv

import (
	"strings"
)

func ServiceTargetParse(serviceName, addr string) string {
	if len(addr) > 0 && addr[:1] == ":" {
		return strings.Replace(serviceName, "_", "-", -1) + addr
	}

	return addr
}
