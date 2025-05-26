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

func Error(data string) {
	logger(zerolog.ErrorLevel, data)
}

func Fatal(data string) {
	logger(zerolog.FatalLevel, data)
}

func Panic(data string) {
	logger(zerolog.PanicLevel, data)
}

func Warn(data string) {
	logger(zerolog.WarnLevel, data)
}

func logger(level zerolog.Level, data string) {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.WithLevel(level).Msg(data)
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
