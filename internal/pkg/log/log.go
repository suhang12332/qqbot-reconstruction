package log

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "time"
)

// 定义颜色常量
const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    White  = "\033[37m"
)

// 定义日志级别常量
const (
    INFO = iota
    WARN
    ERROR
    DEBUG
)

var base *ColorLogger

// ColorLogger 包装器，用于为不同日志级别添加颜色和前缀
type ColorLogger struct {
    logger  *log.Logger
    level   int
}

// NewColorLogger 创建一个新的 ColorLogger
func NewColorLogger(level int) *ColorLogger {
    return &ColorLogger{
        logger:  log.New(os.Stdout,"", log.Llongfile),
        level:   level,
    }
}

// Log 方法根据级别和颜色打印日志
func (cl *ColorLogger) Log(level int, color string, prefix string, v ...interface{}) {
    if level >= cl.level {
        now := time.Now().Format("2006年01月02日15时04分05秒")
        coloredPrefix := fmt.Sprintf("%s[%s]%s%s", color, now, prefix, Reset)
        cl.logger.SetPrefix(coloredPrefix)
        cl.logger.Output(2, fmt.Sprintf("%s%s%s", color, fmt.Sprint(v...), Reset))
    }
}
func (cl *ColorLogger) getCallerFileName() string {
	_, file, line, ok := runtime.Caller(2) // Adjust the number based on your call stack depth
	if !ok {
		file = "???"
	}
	return fmt.Sprintf("%s%s%s:%d", Purple, file, Reset, line)
}

// Fatal 方法用于记录致命错误并终止程序执行
func (cl *ColorLogger) Fatal(color string, format string, v ...interface{}) {
	now := time.Now().Format("2006年01月02日15时04分05秒")
	coloredPrefix := fmt.Sprintf("%s[%s][FATAL]: %s", color, now, Reset)
	cl.logger.SetPrefix(coloredPrefix)
	cl.logger.Fatalf(fmt.Sprintf("%s%s%s", color, format, Reset), v...)
}

func init() {

    base = NewColorLogger(INFO)
}

func Error(s string, info interface{}) {
    base.Log(ERROR, Red, "[ERROR]: ", fmt.Sprintf(s, info))
}
func Errorf(s string) {
    base.Log(ERROR, Red, "[ERROR]: ", s)
}
func Info(s string, info interface{}) {
    base.Log(INFO, Green, "[INFO]: ", fmt.Sprintf(s, info))
}
func Infof(s string) {
    base.Log(INFO, Green, "[INFO]: ", s)
}
func Debug(s string, info interface{}) {
    base.Log(DEBUG, Cyan, "[DEBUG]: ", fmt.Sprintf( s, info))
}
func Debugf(s string) {
    base.Log(DEBUG, Cyan, "[DEBUG]: ", s)
}

func Fatalf(s string) {
    base.Fatal(Red, "%s", s)
}

func Fatal(s string, info interface{}) {
    base.Fatal(Red, "%s", fmt.Sprintf(s, info))
}
