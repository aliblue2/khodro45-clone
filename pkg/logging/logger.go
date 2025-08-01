package logging

import "github.com/aliblue2/khodro45/configs"

type Logger interface {
	Init()

	Info(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, subCat SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *configs.Config) Logger {
	switch cfg.Logger.Logger {
	case "zap":
		return NewZapLogger(cfg)
	case "zerolog":
		return NewZeroLogger(cfg)
	}

	panic("logger is not supported")

}

// file <- fileBeat -> elasticSearch -> kibana
