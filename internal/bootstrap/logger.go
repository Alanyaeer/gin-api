package bootstrap

import (
	"log/slog"
	"os"
)

func init() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key != slog.SourceKey {
				return a
			}

			switch src := a.Value.Any().(type) {
			case *slog.Source:
				if src == nil {
					return a
				}
				return slog.Group(slog.SourceKey,
					slog.String("function", src.Function),
					slog.Int("line", src.Line),
				)
			case slog.Source:
				return slog.Group(slog.SourceKey,
					slog.String("function", src.Function),
					slog.Int("line", src.Line),
				)
			default:
				return a
			}
		},
	})
	slog.SetDefault(slog.New(handler))
}
