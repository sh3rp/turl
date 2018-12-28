package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Log interface {
	Error(string, ...interface{})
	Info(string, ...interface{})
	Fatal(error)
}

type zeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger() Log {
	output := zerolog.NewConsoleWriter()
	output.TimeFormat = time.RFC3339
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return zeroLogger{
		logger: logger,
	}
}

func (z zeroLogger) Error(str string, vars ...interface{}) {
	if len(vars) == 0 {
		z.logger.Error().Msg(str)
	} else {
		z.logger.Error().Msg(fmt.Sprintf(str, vars))
	}
}

func (z zeroLogger) Info(str string, vars ...interface{}) {
	if len(vars) == 0 {
		z.logger.Info().Msg(str)
	} else {
		z.logger.Info().Msg(fmt.Sprintf(str, vars))
	}
}

func (z zeroLogger) Fatal(err error) {
	z.logger.Fatal().Err(err)
}
