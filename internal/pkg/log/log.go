package log

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func Error(s string, info interface{}) {
	color.Red(format("error", fmt.Sprintf("%s%s", s, info)))
}
func Errorf(s string) {
	color.Red(format("error", s))
}
func Info(s string, info interface{}) {
	color.Cyan(format("info", fmt.Sprintf("%s%s", s, info)))
}
func Infof(s string) {
	color.Cyan(format("info", s))
}
func Debug(s string, info interface{}) {
	color.Magenta(format("debug", fmt.Sprintf("%s%s", s, info)))
}
func Debugf(s string) {
	color.Cyan(format("debug", s))
}
func format(level string, format string) string {
	t := time.Now().Format("2006年01月02日15时04分05秒")
	return fmt.Sprintf("[%s] [%s] %s", level, t, format)
}
