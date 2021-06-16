package glogger

import (
	"fmt"

	"github.com/lpxxn/grpc_research/test/log_test/grpclog"
)

const d = 2

func init() {
	grpclog.SetLoggerV2(&glogger{})
}

type glogger struct{}

func (g *glogger) Info(args ...interface{}) {
	InfoDepth(d, args...)
}

func (g *glogger) Infoln(args ...interface{}) {
	InfoDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Infof(format string, args ...interface{}) {
	InfoDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) InfoDepth(depth int, args ...interface{}) {
	InfoDepth(depth+d, args...)
}

func (g *glogger) Warning(args ...interface{}) {
	WarningDepth(d, args...)
}

func (g *glogger) Warningln(args ...interface{}) {
	WarningDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Warningf(format string, args ...interface{}) {
	WarningDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) WarningDepth(depth int, args ...interface{}) {
	WarningDepth(depth+d, args...)
}

func (g *glogger) Error(args ...interface{}) {
	ErrorDepth(d, args...)
}

func (g *glogger) Errorln(args ...interface{}) {
	ErrorDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Errorf(format string, args ...interface{}) {
	ErrorDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) ErrorDepth(depth int, args ...interface{}) {
	ErrorDepth(depth+d, args...)
}

func (g *glogger) Fatal(args ...interface{}) {
	FatalDepth(d, args...)
}

func (g *glogger) Fatalln(args ...interface{}) {
	FatalDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Fatalf(format string, args ...interface{}) {
	FatalDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) FatalDepth(depth int, args ...interface{}) {
	FatalDepth(depth+d, args...)
}

func (g *glogger) V(l int) bool {
	return bool(V(Level(l)))
}
