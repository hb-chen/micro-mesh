package log

import (
	"io"

	"github.com/fatih/color"
)

const (
	DEBUG Lvl = iota
	INFO
	WARN
	ERROR
	OFF
	fatalLvl
	panicLvl
)

var (
	level       = DEBUG
	colorEnable = true
	global      = NewLogger()
)

func init() {
	global.SetCalldepth(4)
}

type (
	Lvl       uint
	colorFunc func(format string, a ...interface{}) string
)

func (lvl Lvl) String() string {
	switch lvl {
	case DEBUG:
		return lvl.colorString("DEBUG", color.WhiteString)
	case INFO:
		return lvl.colorString("INFO", color.GreenString)
	case WARN:
		return lvl.colorString("WARN", color.YellowString)
	case ERROR:
		return lvl.colorString("ERROR", color.RedString)
	case fatalLvl:
		return lvl.colorString("FATAL", color.HiRedString)
	case panicLvl:
		return lvl.colorString("PANIC", color.HiRedString)
	default:
		return lvl.colorString("-", color.WhiteString)
	}
}

func (lvl Lvl) colorString(str string, f colorFunc) string {
	if colorEnable {
		return f(str)
	} else {
		return str
	}
}

func SetLevel(lvl Lvl) {
	level = lvl
}

func SetColor(enable bool) {
	colorEnable = enable
}

func SetPrefix(prefix string) {
	global.SetPrefix(prefix)
}

func SetOutput(w io.Writer) {
	global.SetOutput(w)
}

func SetFlags(flag int) {
	global.SetFlags(flag)
}

func SetCalldepth(calldepth int) {
	global.SetCalldepth(calldepth)
}

func Debug(v ...interface{}) {
	global.Debug(v...)
}
func Debugf(format string, v ...interface{}) {
	global.Debugf(format, v...)
}

func Info(v ...interface{}) {
	global.Info(v...)
}
func Infof(format string, v ...interface{}) {
	global.Infof(format, v...)
}

func Warn(v ...interface{}) {
	global.Warn(v...)
}
func Warnf(format string, v ...interface{}) {
	global.Warnf(format, v...)
}

func Error(v ...interface{}) {
	global.Error(v...)
}
func Errorf(format string, v ...interface{}) {
	global.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	global.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	global.Fatalf(format, v...)
}

func Panic(v ...interface{}) {
	global.Panic(v...)
}
func Panicf(format string, v ...interface{}) {
	global.Panicf(format, v...)
}
