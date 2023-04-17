package wlog

import "go.uber.org/zap"

// echo 框架 xorm时会用的方法
var XLogger *XormLogger

type XormLogger struct{}

func (x *XormLogger) Write(p []byte) (n int, err error) {
	Wlogger.Info("数据库操作", zap.String("数据库", string(p)))
	return len(p), nil
}

var EchoLog *EchoLogger

type EchoLogger struct{}

func (e *EchoLogger) Write(p []byte) (n int, err error) {
	Wlogger.Info("ECHO", zap.String("请求", string(p)))
	return len(p), nil
}
