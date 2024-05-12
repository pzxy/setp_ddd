package log

import (
	"go.uber.org/zap"
)

var G = &logger{}

type logger struct {
	*zap.SugaredLogger
}

func init() {
	l, _ := zap.NewProduction()
	defer func(lg *zap.Logger) {
		_ = lg.Sync()
	}(l) // flushes buffer, if any
	G.SugaredLogger = l.Sugar()
}
