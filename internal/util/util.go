package util

import (
	"io/fs"
	"log/slog"
)

func DeferFileClose(file fs.File) {
	err := file.Close()
	if err != nil {
		slog.Error("failed to close file", "cause", err)
	}
}
