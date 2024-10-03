package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	initOnce sync.Once
	logger   *zap.Logger
)

type CustomeLogger struct {
	Logger *zap.Logger
}

func NewCustomLogger(defaultLogLevel zapcore.Level) (cl *CustomeLogger) {
	initOnce.Do(func() {
		consoleDebuging := zapcore.Lock(os.Stdout)
		consoleError := zapcore.Lock(os.Stderr)

		encoderConfig, fileSyncWriter := initiateLoggerConfig()

		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		prodEncoder := zapcore.NewJSONEncoder(encoderConfig)

		core := zapcore.NewTee(

			zapcore.NewCore(prodEncoder, fileSyncWriter, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				if l >= defaultLogLevel {
					return l == zap.ErrorLevel
				}
				return false
			})),
			zapcore.NewCore(prodEncoder, fileSyncWriter, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				if l >= defaultLogLevel {
					return l == zap.FatalLevel
				}
				return false
			})),
			zapcore.NewCore(prodEncoder, fileSyncWriter, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				if l >= defaultLogLevel {
					return l == zap.DebugLevel
				}
				return false
			})),
			zapcore.NewCore(consoleEncoder, consoleDebuging, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				if l >= defaultLogLevel {
					return l >= zap.DebugLevel && l < zap.ErrorLevel
				}
				return false
			})),
			zapcore.NewCore(consoleEncoder, consoleError, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				if l >= defaultLogLevel {
					return l >= zap.ErrorLevel
				}
				return false
			})),
		)

		logger = zap.New(core)
		defer logger.Sync()
	})

	return &CustomeLogger{Logger: logger}
}

func initiateLoggerConfig() (zapcore.EncoderConfig, zapcore.WriteSyncer) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	lumberJack := &lumberjack.Logger{
		Filename:   fmt.Sprintf("/tmp/payment-%v.log", time.Now()),
		MaxSize:    100,
		MaxAge:     2,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true,
	}

	file := zapcore.AddSync(lumberJack)

	return encoderConfig, file
}

func (cl *CustomeLogger) ChangeTheDefaultLogLevel(lvl zapcore.Level) *zap.Logger {
	c := NewCustomLogger(lvl)
	return c.Logger
}
