package mylog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var Log *logrus.Entry

type MyHook struct {
}

func getWriter() (io.Writer, error) {
	return rotatelogs.New("./logs/seckill"+".%Y%m%d%H%M",
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
}

func (m MyHook) Levels() []logrus.Level {
	//TODO implement me
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel}
}

func (m MyHook) Fire(entry *logrus.Entry) error {
	//TODO implement me
	writer, err := getWriter()
	if err != nil {
		return err
	}
	logrus.SetOutput(writer)
	return nil
}

func InitLog() {
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(&MyHook{})
	logrus.SetReportCaller(true)
	Log = logrus.WithField("version:", "1.0")
}
