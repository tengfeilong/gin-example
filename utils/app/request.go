package app

import (
	"gin-example/mylog"
	"github.com/astaxie/beego/validation"
	"go.uber.org/zap"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		mylog.Logger.Info("app err", zap.String("key", err.Key), zap.String("err", err.Message))
	}

	return
}
