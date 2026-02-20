package log

import (
	"log/slog"
	"os"
	"io"
	"context"
)
// 自定义Handler：重写日志输出格式，去掉键名只保留值的拼接
type customTextHandler struct {
	level slog.Leveler // 日志级别控制
}

// NewCustomTextHandler 创建自定义Handler实例
func NewCustomTextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &customTextHandler{
		level: opts.Level,
	}
}

// Enabled 判断当前级别是否允许输出
func (h *customTextHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

// Handle 核心：自定义日志输出格式
func (h *customTextHandler) Handle(_ context.Context, r slog.Record) error {
	// 1. 提取级别（转大写）、时间（自定义格式）、消息内容
	levelStr := r.Level.String() // DEBUG/INFO/WARN/ERROR
	timeStr := r.Time.Format("2006-01-02 15:04:05.000")
	msg := r.Message

	// 2. 拼接成目标格式: [级别] 时间  消息
	logLine := "[" + levelStr + "] " + timeStr + "  " + msg + "\n"

	// 3. 输出到标准输出
	_, err := os.Stdout.WriteString(logLine)
	return err
}

// WithAttrs 实现Handler接口（自定义场景无需额外字段，直接返回自身）
func (h *customTextHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup 实现Handler接口（无需分组，直接返回自身）
func (h *customTextHandler) WithGroup(_ string) slog.Handler {
	return h
}
func init() {
	slogHandler := NewCustomTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	loggerPtr := slog.New(slogHandler)
	slog.SetDefault(loggerPtr)
}
