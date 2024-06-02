package global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type MyFormatter struct {
}

var TPLogger *logrus.Logger

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取日志级别
	level := entry.Level
	var levelColor int
	switch level {
	case logrus.TraceLevel, logrus.DebugLevel:
		// 设置为灰色
		levelColor = gray
	case logrus.InfoLevel:
		// 设置为蓝色
		levelColor = blue
	case logrus.WarnLevel:
		// 设置为黄色
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		// 设置为红色
		levelColor = red
	default:
		// 默认为蓝色
		levelColor = blue
	}
	// 获取调用者函数信息
	funcVal := entry.Caller.Function //获取调用者函数信息
	// 格式化调用者的文件名和行号信息
	//式化调用者的文件名和行号信息
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	// 获取日志消息
	msg := entry.Message
	// 格式化时间
	time := entry.Time.Format("2006-01-02 15:04:05")
	// 返回格式化后的日志条目
	return []byte(fmt.Sprintf("%s [ \033[%dm%s\033[0m ] [ %s ] [%s] %s\n", time, levelColor, level.String(), funcVal, fileVal, msg)), nil
}

func InitLog() {
	//创建Logrus实例
	TPLogger = logrus.New()
	TPLogger.SetReportCaller(true)
	TPLogger.SetOutput(os.Stdout)
	//使用自定义的日志格式
	TPLogger.SetFormatter(&MyFormatter{})
}
