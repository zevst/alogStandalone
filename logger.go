////////////////////////////////////////////////////////////////////////////////
// Author:   Nikita Koryabkin
// Email:    Nikita@Koryabk.in
// Telegram: https://t.me/Apologiz
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"io"
	"sync"
	"time"

	"github.com/mylockerteam/alog"
)

const (
	keyInfo = "ALOG_LOGGER_INFO"
	keyWrn  = "ALOG_LOGGER_WARNING"
	keyErr  = "ALOG_LOGGER_ERROR"
)

var logger struct {
	instance *alog.Log
	once     sync.Once
}

func GetLogger() *alog.Log {
	logger.once.Do(func() {
		logger.instance = alog.Create(&alog.Config{
			TimeFormat:  time.RFC3339Nano,
			LogFileLine: false,
			Loggers: alog.LoggerMap{
				alog.LoggerInfo: getInfoLogger(),
				alog.LoggerWrn:  getWarningLogger(),
				alog.LoggerErr:  getErrorLogger(),
			},
		})
	})
	return logger.instance
}

func getInfoLogger() *alog.Logger {
	return &alog.Logger{
		Channel: make(chan string, 100),
		Strategies: []io.Writer{
			alog.GetFileStrategy(alog.GetEnvStr(keyInfo)),
			alog.GetDefaultStrategy(),
		},
	}
}

func getWarningLogger() *alog.Logger {
	return &alog.Logger{
		Channel: make(chan string, 100),
		Strategies: []io.Writer{
			alog.GetFileStrategy(alog.GetEnvStr(keyWrn)),
			alog.GetDefaultStrategy(),
		},
	}
}

func getErrorLogger() *alog.Logger {
	return &alog.Logger{
		Channel: make(chan string, 100),
		Strategies: []io.Writer{
			alog.GetFileStrategy(alog.GetEnvStr(keyErr)),
			alog.GetDefaultStrategy(),
		},
	}
}
