package log

import "fmt"


// LogRecord mock的日志函数，用于展示BUILD之间的关联
func LogRecord(port string) {
	fmt.Printf("🚀 Server is running on http://localhost:%s", port)
}