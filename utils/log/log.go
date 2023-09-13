package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

const (
	TraceIdKey      = "traceID"
	SpanIdKey       = "spanID"
	SpanParentIdKey = "spanParentID"
	logErrFieldKey  = "err"
	logSrcFieldKey  = "src"
)

// init Set application log and beego log
func init() {
	// Set logrus configuration
	level, err := logrus.ParseLevel("trace") // to be change from .env
	if err != nil {
		log.Fatalf("Invalid log level '%s', %s", level, err)
	}

	// Mkdir log folder
	t := time.Now().Format("2006-01-02")
	filePath := fmt.Sprintf("./logs/%s.log", t)
	dir := "./logs"
	if dir != "" {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Make log dir '%s' failed, %s", dir, err)
		}
	}

	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   filePath,
			MaxSize:    100,
			MaxBackups: 100,
			MaxAge:     3,
			Compress:   false,
			LocalTime:  true,
		},
		level,
		&logrus.JSONFormatter{},
		&lumberjackrus.LogFileOpts{},
	)

	if err != nil {
		log.Fatal("Log hook creation failed")
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetLevel(level)
	logrus.AddHook(hook)
	if err != nil {
		log.Fatal(err)
	}
}

func stdEntries(ctx context.Context) *logrus.Entry {
	sc := trace.SpanContextFromContext(ctx)
	return logrus.
		WithField(TraceIdKey, sc.TraceID().String()).
		WithField(SpanIdKey, sc.SpanID().String()).
		WithField(SpanParentIdKey, ctx.Value(SpanParentIdKey))
}

func stdErrEntries(ctx context.Context, err error) *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(2); ok {
		file = file[strings.LastIndex(file, "/")+1:]
		funcName := runtime.FuncForPC(pc).Name()
		return stdEntries(ctx).
			WithField(logErrFieldKey, err).
			WithField(logSrcFieldKey, fmt.Sprintf("%s:%s:%d", file, funcName, line))
	}
	return stdEntries(ctx)
}

func Trace(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Trace(args)
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx).Tracef(format, args)
}

func Traceln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Traceln(args)
}

func Debug(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Debug(args)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx).Debugf(format, args)
}

func Debugln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Debugln(args)
}

func Print(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Print(args)
}

func Printf(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx).Printf(format, args)
}

func Println(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Println(args)
}

func Info(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Info(args)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx).Infof(format, args)
}

func Infoln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Infoln(args)
}

func Warn(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Warn(args)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx).Warnf(format, args)
}

func Warnln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx).Warnln(args)
}

func Error(ctx context.Context, err error, args ...interface{}) {
	stdErrEntries(ctx, err).Error(args)
}

func Errorf(ctx context.Context, err error, format string, args ...interface{}) {
	stdErrEntries(ctx, err).Errorf(format, args)
}

func Errorln(ctx context.Context, err error, args ...interface{}) {
	stdErrEntries(ctx, err).Errorln(args)
}

func Fatal(ctx context.Context, err error, args ...interface{}) {
	stdErrEntries(ctx, err).Fatal(args)
}

func Fatalf(ctx context.Context, err error, format string, args ...interface{}) {
	stdErrEntries(ctx, err).Fatalf(format, args)
}

func Fatalln(ctx context.Context, err error, args ...interface{}) {
	stdErrEntries(ctx, err).Fatalln(args)
}

func Panic(ctx context.Context, err error, args ...interface{}) {
	stdErrEntries(ctx, err).Panic(args)
}

func Panicf(ctx context.Context, err error, format string, args ...interface{}) {
	stdErrEntries(ctx, err).Panicf(format, args)
}

func Panicln(ctx context.Context, err error, args ...interface{}) {
	stdErrEntries(ctx, err).Panicln(args)
}
