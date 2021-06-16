package main

import (
	"github.com/lpxxn/grpc_research/test/log_test/grpclog"
	internalgrpclog "github.com/lpxxn/grpc_research/test/log_test/internal/grpclog"
)

func main() {

	grpclog.Infoln("Info", "message.")

	grpclog.Infof("%v %v.", "Info", "message abcdef")
	grpclog.Warning("warning")
	grpclog.Errorln("error")

	grpclog.Fatalln("fatal")
	prefxLog := internalgrpclog.NewPrefixLogger(grpclog.Component("myApp"), "haha")
	prefxLog.Debugf("debug %v", "aaa")

}
