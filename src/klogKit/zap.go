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

@param zapLogger 	可以为nil（将采用默认值）
@param globalArgs	是否设置为kratos的全局logger？（默认true）
*/
func UseZap(zapLogger *zap.Logger, id, name, version string, globalArgs ...bool) (logger log.Logger) {
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
		//"ts", log.DefaultTimestamp, /* Richelieu: zap的日志输出已经有时间戳了，所以此处不再需要时间戳 */
		"caller", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
		"service.id", id,
		"service.name", name,
		"service.version", version,
	)

	global := true
	if globalArgs != nil {
		global = globalArgs[0]
	}
	if global {
		log.SetLogger(logger)
	}

	return
}
