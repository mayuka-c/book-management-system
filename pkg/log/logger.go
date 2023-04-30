package log

import (
	"context"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	loggerTimeFormat = time.RFC3339Nano
)

var (
	logger    *logrus.Entry
	skipLevel = 3
	caller    = "context"
)

var logLevels = map[string]logrus.Level{
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

func init() {
	level, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		level = "info"
	}
	SetupLogger(logLevels[level])
}

func addCommonFields(ctx context.Context) *logrus.Entry {
	return logger.WithField(
		caller,
		fileInfo(),
	)
}

func fileInfo() string {
	_, file, line, ok := runtime.Caller(skipLevel)
	if !ok {
		file = "<???>"
		line = 1
	}
	return file + ":" + strconv.Itoa(line)
}

// Infof - logger for infof
func Infof(ctx context.Context, msg string, args ...interface{}) {
	if args == nil {
		addCommonFields(ctx).Info(msg)
		return
	}
	addCommonFields(ctx).Infof(msg, args...)
}

// Debugf - logger for debugf
func Debugf(ctx context.Context, msg string, args ...interface{}) {
	if args == nil {
		addCommonFields(ctx).Debug(msg)
		return
	}
	addCommonFields(ctx).Debugf(msg, args...)
}

// Errorf - logger for error
func Errorf(ctx context.Context, msg string, args ...interface{}) {
	if args == nil {
		addCommonFields(ctx).Error(msg)
		return
	}
	addCommonFields(ctx).Errorf(msg, args...)
}

// Warnf - logger for warning
func Warnf(ctx context.Context, msg string, args ...interface{}) {
	if args == nil {
		addCommonFields(ctx).Warn(msg)
		return
	}
	addCommonFields(ctx).Warnf(msg, args...)
}

// SetupLogger - setup logger
func SetupLogger(level logrus.Level) {
	logging := logrus.New()
	logging.SetLevel(level)
	formatter := &logrus.JSONFormatter{
		TimestampFormat: loggerTimeFormat,
	}
	logging.SetFormatter(formatter)
	logger = logging.WithContext(context.Background())
}
