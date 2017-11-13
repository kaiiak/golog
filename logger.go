package golog

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 可以看到被调用者的logger
type Logger struct {
	isDebug  bool
	callDeep int
	Logger   *zap.Logger
	field    []zapcore.Field
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

func httpStatusCodeEncoder(code int, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt(code)
}

func ios8601Encoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

// New new logger
func New(debugLevel bool) (*Logger, error) {
	productConfig := zap.NewProductionConfig()
	productConfig.EncoderConfig.EncodeTime = ios8601Encoder
	productConfig.EncoderConfig.EncodeDuration = milliSecondsDurationEncoder
	productConfig.EncoderConfig.EncodeCaller = callerEncoder
	if debugLevel {
		productConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	l, err := productConfig.Build()
	if err != nil {
		return nil, err
	}
	return &Logger{
		isDebug: debugLevel,
		Logger:  l,
	}, nil
}

// Caller Add caller field
func (l Logger) Caller(i int) Logger {
	newLogger := l
	newLogger.Logger.WithOptions(zap.AddCallerSkip(i + 1))
	return newLogger
}

// Read output log message
func (l Logger) Read(p []byte) (int, error) {
	return 0, nil
}

// func (l Logger) Write(p []byte) (int, error) {
// 	return 0, nil
// }

// AddField Add new filed to log message
func (l Logger) AddField(key string, value interface{}) {
	l.field = append(l.field, zap.Any(key, value))
}

// Print Print
func (l Logger) Print(args ...interface{}) {
	l.Logger.Info(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Println Println
func (l Logger) Println(args ...interface{}) {
	l.Logger.Info(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Printf Printf
func (l Logger) Printf(format string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}

// Debug Debug
func (l Logger) Debug(args ...interface{}) {
	l.Logger.Debug(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Debugln Debugln
func (l Logger) Debugln(args ...interface{}) {
	l.Logger.Debug(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Debugf Debugf
func (l Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}

// Info Info
func (l Logger) Info(args ...interface{}) {
	l.Logger.Info(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Infoln Infoln
func (l Logger) Infoln(args ...interface{}) {
	l.Logger.Info(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Infof Infof
func (l Logger) Infof(format string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}

// Warn Warn
func (l Logger) Warn(args ...interface{}) {
	l.Logger.Warn(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Warnln Warnln
func (l Logger) Warnln(args ...interface{}) {
	l.Logger.Warn(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Warnf Warnf
func (l Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}

// Error Error
func (l Logger) Error(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Errorln Errorln
func (l Logger) Errorln(args ...interface{}) {
	l.Logger.Error(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Errorf Errorf
func (l Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}

// Panic Panic
func (l Logger) Panic(args ...interface{}) {
	l.Logger.Panic(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Panicln Panicln
func (l Logger) Panicln(args ...interface{}) {
	l.Logger.Panic(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Panicf Panicf
func (l Logger) Panicf(format string, args ...interface{}) {
	l.Logger.Panic(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}

// Fatal Fatal
func (l Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(fmt.Sprint(args...), l.field...)
	l.Logger.Sync()
}

// Fatalln Fatalln
func (l Logger) Fatalln(args ...interface{}) {
	l.Logger.Fatal(fmt.Sprintln(args...), l.field...)
	l.Logger.Sync()
}

// Fatalf Fatalf
func (l Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, args...), l.field...)
	l.Logger.Sync()
}
