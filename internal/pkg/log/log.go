package log

import (
    "fmt"
    "github.com/LagrangeDev/LagrangeGo/utils"
    "github.com/mattn/go-colorable"
    "github.com/sirupsen/logrus"
    "strings"
    "time"
)




func Info(format string, arg ...any) {
    logger.Infof(format, arg...)
}

func Infof(format string) {
    logger.Infof(format)
}

func Warning(format string, arg ...any) {
    logger.Warnf(format, arg...)
}

func Warningf(format string) {
    logger.Warnf(format)
}

func Debug(format string, arg ...any) {
    logger.Debugf(format, arg...)
}

func Debugf(format string) {
    logger.Debugf(format)
}

func Error(format string, arg ...any) {
    logger.Errorf(format, arg...)
}

func Errorf(format string) {
    logger.Errorf(format)
}

func Fatal(format string, arg ...any) {
    logger.Fatalf(format, arg...)
}

func Fatalf(format string) {
    logger.Fatalf(format)
}

const (
    // 定义颜色代码
    colorReset  = "\x1b[0m"
    colorRed    = "\x1b[31m"
    colorYellow = "\x1b[33m"
    colorGreen  = "\x1b[32m"
    colorBlue   = "\x1b[34m"
    colorWhite  = "\x1b[37m"
)

var logger = logrus.New()

func init() {
    logger.SetLevel(logrus.TraceLevel)
    logger.SetFormatter(&ColoredFormatter{})
    logger.SetOutput(colorable.NewColorableStdout())
}

type ColoredFormatter struct{}

func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    // 获取当前时间戳
    timestamp := time.Now().Format("2006-01-02 15:04:05")

    // 根据日志级别设置相应的颜色
    var levelColor string
    switch entry.Level {
    case logrus.DebugLevel:
        levelColor = colorBlue
    case logrus.InfoLevel:
        levelColor = colorGreen
    case logrus.WarnLevel:
        levelColor = colorYellow
    case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
        levelColor = colorRed
    default:
        levelColor = colorWhite
    }

    return utils.S2B(fmt.Sprintf("[%s] [%s%s%s]: %s\n",
        timestamp, levelColor, strings.ToUpper(entry.Level.String()), colorReset, entry.Message)), nil
}