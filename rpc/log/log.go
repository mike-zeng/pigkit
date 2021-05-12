package log

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

type Log interface {
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	INFO(format string, v ...interface{})
	WARNING(format string, v ...interface{})
	ERROR(format string, v ...interface{})
	FATAL(format string, v ...interface{})
}

type logger struct{
	*log.Logger
	options *Options
}

func (log *logger) Trace(format string, v ...interface{}) {
	if log.options.level > TRACE {
		return
	}
	go logout("[TRACE]",format,v)
}

func (log logger) Debug(format string, v ...interface{}) {
	if log.options.level > DEBUG {
		return
	}
	go logout("[DEBUG]",format,v)
}

func (log logger) INFO(format string, v ...interface{}) {
	if log.options.level > INFO {
		return
	}
	go logout("[INFO]",format,v)
}

func (log logger) WARNING(format string, v ...interface{}) {
	if log.options.level > WARNING {
		return
	}
	go logout("[WARNING]",format,v)
}

func (log logger) ERROR(format string, v ...interface{}) {
	if log.options.level > ERROR {
		return
	}
	go logout("[ERROR]",format,v)
}

func (log logger) FATAL(format string, v ...interface{}) {
	if log.options.level > FATAL {
		return
	}
	logout("[FATAL]",format,v)
	data := log.Prefix() + fmt.Sprintf(format,v...)
	panic(errors.New(data))
}

func logout(leaveStr string,format string, v ...interface{}) {
	data := log.Prefix() + fmt.Sprintf(format,v...)
	var buffer bytes.Buffer
	buffer.WriteString(leaveStr+" ")
	buffer.WriteString(data)
	_ = log.Output(3, buffer.String())
}

var DefaultLog = &logger {
	Logger : log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	options : &Options {
		level : DEBUG,
	},
}
