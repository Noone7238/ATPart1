package main

import (
	network "gitstuff/AT/internal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	err := network.Start()
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
