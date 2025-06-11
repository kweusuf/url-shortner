package log

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Using a simple library here. Can be changed as we are interfacing this in other classes.
func Info(data string) {
	logger(zerolog.InfoLevel, data)
}

func Debug(data string) {
	logger(zerolog.DebugLevel, data)
}

func Error(data string, error ...error) {
	logger(zerolog.ErrorLevel, data, error...)
}

func Fatal(data string, error ...error) {
	logger(zerolog.FatalLevel, data, error...)
}

func Panic(data string, error ...error) {
	logger(zerolog.PanicLevel, data, error...)
}

func Warn(data string) {
	logger(zerolog.WarnLevel, data)
}

func logger(level zerolog.Level, data string, error ...error) {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.WithLevel(level).Msg(data)
	if len(error) > 0 {
		log.WithLevel(level).Err(error[0]).Msg("An error occurred")
	}
}

func InitializeLogging() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
