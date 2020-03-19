package logging

import (
	"fmt"
	"go.uber.org/zap"
)

// Log is our logger for use elsewhere
var Log = loggerInit()

func loggerInit() *zap.SugaredLogger {
	Logger := zap.NewExample().Sugar()
	defer Logger.Sync()
	if Logger == nil {
		fmt.Print("expected a non-nil logger")
	}
	Logger.Info("constructed a logger")
	return Logger
}
