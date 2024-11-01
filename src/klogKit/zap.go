package klogKit

import (
	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// UseZap
/*
PS: 会设置 kratos 的全局logger.

@param zapLogger 可以为nil
*/
func UseZap(zapLogger *zap.Logger, id, name, version string) (logger log.Logger) {
	if zapLogger == nil {
		writeSyncer := zapcore.AddSync(os.Stderr)
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder := zapcore.NewConsoleEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
		zapLogger = zap.New(core)
	}

	kLogger := kratoszap.NewLogger(zapLogger)
	logger = log.With(kLogger,
		//"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
		"service.id", id,
		"service.name", name,
		"service.version", version,
	)
	log.SetLogger(logger)
	return
}
