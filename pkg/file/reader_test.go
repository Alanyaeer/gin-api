package file 
import (
	"testing"
)

// 👇 函数名、参数必须严格这样写，少一个都不行
func TestAbsolutePath(t *testing.T) {
	path, err := GetAbsolutePath("./reader.go")
	// 用 t.Log 打印，确保测试函数有实际逻辑
	t.Log("绝对路径：", path)
	t.Log("错误：", err)
}

func TestFilePrivate(t *testing.T) {
	// 直接调用同包内的私有函数，确保它存在且可调用
	privateFunction()
}