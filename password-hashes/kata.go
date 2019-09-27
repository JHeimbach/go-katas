package kata

import (
	"crypto/md5"
	"fmt"
)

func PassHash(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
