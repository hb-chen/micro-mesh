package log

import (
	"fmt"
	"log"
	"os"
)

type defaultLogger struct {
	*log.Logger
	calldepth int
}

func NewLogger() *defaultLogger {
	return &defaultLogger{
		Logger:    log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile),
		calldepth: 3,
	}
}

func (l *defaultLogger) SetCalldepth(calldepth int) {
	l.calldepth = calldepth
}

func (l *defaultLogger) Debug(v ...interface{}) {
	l.output(DEBUG, v...)
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	l.outputf(DEBUG, format, v...)
}

func (l *defaultLogger) Info(v ...interface{}) {
	l.output(INFO, v...)
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	l.outputf(INFO, format, v...)
}

func (l *defaultLogger) Warn(v ...interface{}) {
	l.output(WARN, v...)
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.outputf(WARN, format, v...)
}

func (l *defaultLogger) Error(v ...interface{}) {
	l.output(ERROR, v...)
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.outputf(ERROR, format, v...)
}

func (l *defaultLogger) Fatal(v ...interface{}) {
	l.output(fatalLvl, v...)
	os.Exit(1)
}

func (l *defaultLogger) Fatalf(format string, v ...interface{}) {
	l.outputf(fatalLvl, format, v...)
	os.Exit(1)
}

func (l *defaultLogger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.output(panicLvl, s)
	panic(s)
}

func (l *defaultLogger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(panicLvl, s)
	panic(s)
}

func (l *defaultLogger) output(lvl Lvl, v ...interface{}) {
	if lvl < level {
		return
	}
	l.Output(l.calldepth, header(lvl, fmt.Sprint(v...)))
}

func (l *defaultLogger) outputf(lvl Lvl, format string, v ...interface{}) {
	if lvl < level {
		return
	}
	l.Output(l.calldepth, header(lvl, fmt.Sprintf(format, v...)))
}

func header(lvl Lvl, msg string) string {
	return fmt.Sprintf("[%s] %s", lvl.String(), msg)
}
