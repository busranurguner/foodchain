package logger

import (
	"go.uber.org/zap"
)

var _ Logger = ZapLoggerStruct{}

type ZapLoggerStruct struct {
	l *zap.Logger
}

func NewZapLogger() (Logger, error) {

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return ZapLoggerStruct{l: zapLogger}, nil
}

func (z ZapLoggerStruct) Fatal(data ...interface{}) {
	z.l.Sugar().Fatal(data)
}

func (z ZapLoggerStruct) Error(data ...interface{}) {
	z.l.Sugar().Error(data)
}

func (z ZapLoggerStruct) Warning(data ...interface{}) {
	z.l.Sugar().Warn(data)
}

func (z ZapLoggerStruct) Info(data ...interface{}) {
	z.l.Sugar().Info(data)
}

func (z ZapLoggerStruct) Debug(data ...interface{}) {
	z.l.Sugar().Debug(data)
}
