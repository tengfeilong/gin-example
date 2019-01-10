package mylog

import (
	"gin-example/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

//var Log = log.New(config.Dir + "/log.log","gin_",log.Ldate|log.Ltime|log.Lshortfile)
var Logger *zap.Logger

func init() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{config.Dir + "/log.log"}
	cfg.ErrorOutputPaths = []string{config.Dir + "/err.log"}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.StacktraceKey = ""

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
}
