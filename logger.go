package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Priority values based on log/syslog package & unix values
type Priority int

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERROR
	LOG_WARN
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

type Logger struct {
	log      *log.Logger
	priority Priority
}

func New(out io.Writer, prefix string, priority Priority, flag int) *Logger {
	return &Logger{
		log:      log.New(out, prefix, flag),
		priority: priority,
	}
}

func (l Logger) priorityPrefix(p Priority) string {
	switch p {
	case LOG_EMERG:
		return "EMERGENCY: "
	case LOG_ALERT:
		return "ALERT: "
	case LOG_CRIT:
		return "CRITICAL: "
	case LOG_ERROR:
		return "ERROR: "
	case LOG_WARN:
		return "WARNING: "
	case LOG_NOTICE:
		return "NOTICE: "
	case LOG_INFO:
		return "INFO: "
	case LOG_DEBUG:
		return "DEBUG: "
	}
	return ""
}

func (l Logger) withPriorityAbove(p Priority, m string) {
	if l.priority < p {
		return
	}
	m = fmt.Sprint(l.priorityPrefix(p), m)
	// Set the correct depth value to '3' so we get the file that actually made
	// the log call not this library.
	l.log.Output(3, m)
}

func (l Logger) Emerg(v ...interface{})  { l.withPriorityAbove(LOG_EMERG, fmt.Sprint(v...)) }
func (l Logger) Crit(v ...interface{})   { l.withPriorityAbove(LOG_CRIT, fmt.Sprint(v...)) }
func (l Logger) Error(v ...interface{})  { l.withPriorityAbove(LOG_ERROR, fmt.Sprint(v...)) }
func (l Logger) Warn(v ...interface{})   { l.withPriorityAbove(LOG_WARN, fmt.Sprint(v...)) }
func (l Logger) Notice(v ...interface{}) { l.withPriorityAbove(LOG_NOTICE, fmt.Sprint(v...)) }
func (l Logger) Info(v ...interface{})   { l.withPriorityAbove(LOG_INFO, fmt.Sprint(v...)) }
func (l Logger) Debug(v ...interface{})  { l.withPriorityAbove(LOG_DEBUG, fmt.Sprint(v...)) }

func (l Logger) Emergf(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_EMERG, fmt.Sprintf(format, v...))
}
func (l Logger) Critf(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_CRIT, fmt.Sprintf(format, v...))
}
func (l Logger) Errorf(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_ERROR, fmt.Sprintf(format, v...))
}
func (l Logger) Warnf(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_WARN, fmt.Sprintf(format, v...))
}
func (l Logger) Noticef(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_NOTICE, fmt.Sprintf(format, v...))
}
func (l Logger) Infof(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_INFO, fmt.Sprintf(format, v...))
}
func (l Logger) Debugf(format string, v ...interface{}) {
	l.withPriorityAbove(LOG_DEBUG, fmt.Sprintf(format, v...))
}

func (l Logger) Fatal(v ...interface{}) {
	l.log.Output(3, fmt.Sprint(v...))
	os.Exit(1)
}

func (l Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log.Output(3, s)
	panic(s)
}
