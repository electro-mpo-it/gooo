package logging

import (
	"log/slog"
	"os"
)

// Объект логгера. (Всё логирование через него.)
var Logger *slog.Logger

// Функция устанавливает в переменную Logger объект *slog.Logger
func SetLogger() {
	o := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	Logger = slog.New(slog.NewTextHandler(os.Stdout, o))
}
