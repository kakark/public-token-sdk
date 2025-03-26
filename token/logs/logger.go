package logs

import (
	"context"

	"github.com/sirupsen/logrus"
)

const loggerKey string = "public_token_sdk_logger"

func Logger(ctx context.Context) ILogger {
	l, _ := GetLogger(ctx)
	return l
}

// NewLogger always create a new log instance with fields in ctx and cache into the ctx
// fields have default values, "request_id" and "K_LOGID"
func NewLogger(ctx context.Context, fields ...string) (ILogger, context.Context) {

	fields = append(fields, "request_id")
	fieldMap := make(map[string]interface{}, len(fields)+1)
	for _, k := range fields {
		if v := ctx.Value(k); v != nil {
			fieldMap[k] = v
		}
	}
	if _, ok := fieldMap["request_id"]; !ok {
		if v := ctx.Value("K_LOGID"); v != nil {
			fieldMap["request_id"] = v
		}
	}
	logs := logrus.WithFields(fieldMap)
	logger := &logrusWrapper{Entry: logs}
	return logger, SetLogger(ctx, logger)
}

type ctxWithSetValue interface {
	Set(key string, value interface{})
}

// GetLogger from ctx, create a new logger not exist and return the context with the log entry
func GetLogger(ctx context.Context) (ILogger, context.Context) {
	if v := ctx.Value(string(loggerKey)); v != nil {
		vv, ok := v.(ILogger)
		if ok {
			return vv, ctx
		}
	}
	return NewLogger(ctx)
}

// SetLogger set ILogger into context
func SetLogger(ctx context.Context, log ILogger) context.Context {
	if v, ok := ctx.(ctxWithSetValue); ok {
		v.Set(string(loggerKey), log)
		return ctx
	}
	return context.WithValue(ctx, string(loggerKey), log)
}

// AddField add kv into current context's logger,
// return logger with kv and new context with the new logger
func AddField(ctx context.Context, key string, value interface{}) (ILogger, context.Context) {
	logger, ctx := GetLogger(ctx)
	logger.AddField(key, value)
	return logger, ctx
}

// AddFields add kvs into current context's logger,
// return logger with kvs and new context with the new logger
func AddFields(ctx context.Context, fields map[string]interface{}) (ILogger, context.Context) {
	logger, ctx := GetLogger(ctx)
	logger.AddFields(fields)
	return logger, ctx
}
