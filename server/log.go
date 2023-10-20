package server

import (
	"time"

	"github.com/rs/zerolog"
)

func newLog(debug bool) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
