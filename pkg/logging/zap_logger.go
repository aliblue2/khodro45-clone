package logging

import (
	"sync"

	"github.com/aliblue2/khodro45/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogLevel = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"error": zap.ErrorLevel,
	"fatal": zap.FatalLevel,
	"warn":  zap.WarnLevel,
}

var once sync.Once
var zapSingLogger *zap.SugaredLogger

func (l *zapLogger) getLogeLevel() zapcore.Level {
	level, exists := zapLogLevel[l.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}

	return level

}

type zapLogger struct {
	cfg    *configs.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *configs.Config) *zapLogger {
	logger := &zapLogger{
		cfg: cfg,
	}

	logger.Init()
	return logger
}

func (l *zapLogger) Init() {
	once.Do(func() {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.cfg.Logger.FilePath,
			MaxSize:    10,
			MaxAge:     5,
			MaxBackups: 10,
			LocalTime:  true,
			Compress:   true,
		})

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(zapcore.NewJSONEncoder(config), w, l.getLogeLevel())

		zapSingLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

	})

	l.logger = zapSingLogger

}

func (l *zapLogger) Info(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, subCat)
	l.logger.Info(msg, params)
}
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Debug(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, subCat)
	l.logger.Debug(msg, params)
}
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Warn(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, subCat)
	l.logger.Warn(msg, params)
}
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, subCat)
	l.logger.Error(msg, params)
}
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Fatal(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, subCat)
	l.logger.Fatal(msg, params)
}
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, subCat SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["cat"] = cat
	extra["sub"] = subCat
	params := mapToZapParam(extra)
	return params
}
