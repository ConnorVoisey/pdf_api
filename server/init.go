package server

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Options struct {
	LogLevel string `doc:"Set log level, valid options: trace, debug, info, warn, error" default:"info"`
	LogPath  string `doc:"Set log file output path defaults to stdout"`
	Port     int    `doc:"Port to listen on." short:"p" default:"3000"`
}

func Init(options *Options) error {
	// Set global log level
	level, err := zerolog.ParseLevel(strings.ToLower(options.LogLevel))
	if err != nil {
		log.Warn().Msgf("Invalid log level '%s', falling back to default: 'warn'", options.LogLevel)
		level = zerolog.WarnLevel
	}
	zerolog.SetGlobalLevel(level)

	// Configure console logging
	if options.LogPath == "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	} else {
		// File logging
		logFile, err := os.OpenFile(options.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal().Err(err).Str("logPath", options.LogPath).Msg("Failed to open log file")
			return err
		}

		log.Logger = zerolog.New(logFile).With().Timestamp().Logger()
	}
	return nil
}
