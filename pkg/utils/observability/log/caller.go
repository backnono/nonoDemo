package log

import (
	kitlog "github.com/go-kit/log"
	"runtime"
	"strconv"
	"strings"
)

func Caller(depth int) kitlog.Valuer {
	return func() interface{} {
		for i := 4; i < depth; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				return "unknown"
			}
			if strings.Contains(file, strings.ReplaceAll(runtime.GOROOT(), "\\", "/")) ||
				strings.Contains(file, "logger") {
				continue
			}
			return file + ":" + strconv.Itoa(line)
		}

		return "unknown"
	}
}
