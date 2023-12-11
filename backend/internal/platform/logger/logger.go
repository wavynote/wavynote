package logger

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LOG_LEVEL_DEBUG = 1
	LOG_LEVEL_INFO  = 2
	LOG_LEVEL_WARN  = 3
	LOG_LEVEL_ERROR = 4
	LOG_LEVEL_FATAL = 5
)

type Logger struct {
	logPath      string
	maxAge       int // 단위는 시간
	rotationTime int // 단위는 시간
	logLevel     zapcore.Level
	atom         zap.AtomicLevel

	logger *zap.Logger
}

func NewLogger(logPath string, maxAge int, rotationTime int, logLevel int) *Logger {
	var level zapcore.Level
	switch logLevel {
	case LOG_LEVEL_DEBUG:
		level = zap.DebugLevel

	case LOG_LEVEL_INFO:
		level = zap.InfoLevel

	case LOG_LEVEL_WARN:
		level = zap.WarnLevel

	case LOG_LEVEL_ERROR:
		level = zap.ErrorLevel

	case LOG_LEVEL_FATAL:
		level = zap.FatalLevel
	}

	l := Logger{
		logPath:      logPath,
		maxAge:       maxAge,
		rotationTime: rotationTime,
		logLevel:     level,
		atom:         zap.NewAtomicLevel(),
	}
	l.Initialize()
	return &l
}

func (l *Logger) Initialize() {
	rotator, err := rotatelogs.New(
		l.logPath,
		rotatelogs.WithMaxAge(time.Duration(l.maxAge)*time.Hour),             // 파일 자동 삭제 주기(기본값: 7일)
		rotatelogs.WithRotationTime(time.Duration(l.rotationTime)*time.Hour), // 파일 로테이션 주기(기본값: 86400초)
	)
	if err != nil {
		fmt.Println(err)
		l.logger = nil
		return
	}

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "message",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(rotator),
		l.atom,
	)

	l.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	l.atom.SetLevel(l.logLevel)
}

func (l *Logger) SetLogLevel(logLevel int) {
	var level zapcore.Level

	switch logLevel {
	case LOG_LEVEL_DEBUG:
		level = zap.DebugLevel

	case LOG_LEVEL_INFO:
		level = zap.InfoLevel

	case LOG_LEVEL_WARN:
		level = zap.WarnLevel

	case LOG_LEVEL_ERROR:
		level = zap.ErrorLevel

	case LOG_LEVEL_FATAL:
		level = zap.FatalLevel
	}

	l.atom.SetLevel(level)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Fatal(fmt.Sprintf(format, args...))
}
