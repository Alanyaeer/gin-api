package demo

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestDemoRedisOps(t *testing.T) {
	DemoRedisOps()
	result := DemoRedisGetVal()
	slog.Info(fmt.Sprintf("Retrieved value: %s", result))
	t.Logf("Retrieved value: %s", result)
}
/**
	当超时的参数调整为1s时，可以正常输出 INFO Successfully set value in Redis
	当超时的参数调整为1μs时，会输出 ERROR Failed to set value in Redis error context deadline exceeded
*/
func TestDemoTimeOutRedisOps(t *testing.T) {
	DemoTimeOutRedisOps()
}
