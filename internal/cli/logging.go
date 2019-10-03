package cli

import (
	log "github.com/sirupsen/logrus"
)

// Logging - handles CLI log level option
type Logging struct {
	Level level `cli:"log-level" usage:"set log level, possible values: error, warn, info, debug, trace" dft:"info"`
}

type level struct {
	value log.Level
}

func (l *level) Decode(input string) error {
	var err error = nil
	l.value, err = log.ParseLevel(input)
	return err
}
