package logging

import (
	"os"

	"github.com/aliblue2/khodro45/configs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zerologLogger zerolog.Logger
var zeroLogLevel = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"warn":  zerolog.WarnLevel,
}

func getZerologLevel(l string) zerolog.Level {
	level, exist := zeroLogLevel[l]
	if !exist {
		return zerolog.DebugLevel
	}
	return level
}

type zeroLogger struct {
	cfg    *configs.Config
	logger *zerolog.Logger
}

func NewZeroLogger(cfg *configs.Config) *zeroLogger {
	logger := &zeroLogger{
		cfg: cfg,
	}
	logger.Init()
	return logger
}

func (l *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			panic(err)
		}

		zerologLogger = zerolog.New(file).
			With().
			Str("appName", "khodro45").
			Str("level", "logger").
			Timestamp().
			Logger()
		zerolog.SetGlobalLevel(getZerologLevel(l.cfg.Logger.Level))
	})
	l.logger = &zerologLogger
}

func (l *zeroLogger) Info(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := mapToZeroLogParam(extra)
	l.logger.Info().Str("category", string(cat)).Str("subCategory", string(subCat)).Fields(params).Msg(msg)
}
func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}
func (l *zeroLogger) Debug(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := mapToZeroLogParam(extra)
	l.logger.Debug().Str("category", string(cat)).Str("subCategory", string(subCat)).Fields(params).Msg(msg)
}
func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}
func (l *zeroLogger) Warn(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := mapToZeroLogParam(extra)
	l.logger.Warn().Str("category", string(cat)).Str("subCategory", string(subCat)).Fields(params).Msg(msg)
}
func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}
func (l *zeroLogger) Error(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := mapToZeroLogParam(extra)
	l.logger.Error().Str("category", string(cat)).Str("subCategory", string(subCat)).Fields(params).Msg(msg)
}
func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}
func (l *zeroLogger) Fatal(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := mapToZeroLogParam(extra)
	l.logger.Fatal().Str("category", string(cat)).Str("subCategory", string(subCat)).Fields(params).Msg(msg)
}
func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}
