package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	writerFile, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	l := NewLog(logrus.TraceLevel, io.MultiWriter(os.Stdout, writerFile), true, &Formatter{PrintKey: true}).SetField("requestId", 123445)

	l.Trace("trace msg")
	l.Debug("debug msg")
	l.Info("info msg")
	l.Warn("warn msg")
	l.Error("error msg")
	l.Fatal("fatal msg")
	l.Panic("panic msg")
}
