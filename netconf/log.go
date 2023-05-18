package netconf

import (
	"io"
	"log"
	"os"
)

var (
	LogLevel *LogSeverity
)

type LogSeverity struct {
	Info    *log.Logger
	Fail    *log.Logger
	Message *log.Logger
}

func (logSeverity *LogSeverity) Default() {
	LogLevel = &LogSeverity{
		Info:    log.New(io.Discard, "[Verbose]: ", log.LstdFlags|log.Lshortfile),
		Fail:    log.New(os.Stderr, "[Error]: ", log.LstdFlags|log.Lshortfile),
		Message: log.New(io.Discard, "[Message]: ", log.LstdFlags|log.Lshortfile),
	}
}

func (logSeverity *LogSeverity) Verbose() {
	LogLevel = &LogSeverity{
		Info:    log.New(os.Stdout, "[Verbose]: ", log.LstdFlags|log.Lshortfile),
		Fail:    log.New(os.Stderr, "[Error]: ", log.LstdFlags|log.Lshortfile),
		Message: log.New(os.Stdout, "[Message]: ", log.LstdFlags|log.Lshortfile),
	}
}

func (logSeverity *LogSeverity) Messages() {
	LogLevel = &LogSeverity{
		Info:    log.New(io.Discard, "[Verbose]: ", log.LstdFlags|log.Lshortfile),
		Fail:    log.New(os.Stderr, "[Error]: ", log.LstdFlags|log.Lshortfile),
		Message: log.New(os.Stdout, "[Message]: ", log.LstdFlags|log.Lshortfile),
	}
}

func (logSeverity *LogSeverity) Silent() {
	LogLevel = &LogSeverity{
		Info:    log.New(io.Discard, "[Verbose]: ", log.LstdFlags|log.Lshortfile),
		Fail:    log.New(io.Discard, "[Error]: ", log.LstdFlags|log.Lshortfile),
		Message: log.New(io.Discard, "[Message]: ", log.LstdFlags|log.Lshortfile),
	}
}
