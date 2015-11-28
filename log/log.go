package log

import (
	log15 "gopkg.in/inconshreveable/log15.v2"
)

type LogLevel log15.Lvl

const (
	DEBUG    = LogLevel(log15.LvlDebug)
	INFO     = LogLevel(log15.LvlInfo)
	WARN     = LogLevel(log15.LvlWarn)
	ERROR    = LogLevel(log15.LvlError)
	CRITICAL = LogLevel(log15.LvlCrit)
)

func New(lvl LogLevel) log15.Logger {
	logger := log15.New()
	logger.SetHandler(log15.CallerFileHandler(log15.StdoutHandler))
	logger.SetHandler(log15.LvlFilterHandler(log15.Lvl(lvl), log15.StdoutHandler))

	return logger
}
