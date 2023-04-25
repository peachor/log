package log

import (
	"github.com/sirupsen/logrus"
	"io"
)

var g_logger *logrus.Logger

func init() {
	g_logger = logrus.New()
}

func GetLogger() *logrus.Logger {
	return g_logger
}

type Log struct {
	fields       map[string]interface{}
	level        logrus.Level
	output       io.Writer
	reportCaller bool
	formatter    logrus.Formatter
}

func NewLog() *Log {
	return &Log{
		fields: map[string]interface{}{},
	}
}

func (receiver *Log) Init(level logrus.Level, output io.Writer, reportCaller bool, formatter logrus.Formatter) *Log {
	g_logger.SetLevel(level)
	g_logger.SetFormatter(formatter)
	g_logger.SetOutput(output)
	g_logger.SetReportCaller(reportCaller)

	return &Log{
		fields:       map[string]interface{}{},
		level:        level,
		output:       output,
		reportCaller: reportCaller,
		formatter:    formatter,
	}
}

func (receiver *Log) SetField(key string, value interface{}) *Log {
	receiver.fields[key] = value
	return receiver
}

//TRACE
func (receiver *Log) Trace(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Trace(args...)
}

//DEBUG
func (receiver *Log) Debug(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Debug(args...)
}

//INFO
func (receiver *Log) Info(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Info(args...)
}

//WARN
func (receiver *Log) Warn(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Warn(args...)
}

//ERROR
func (receiver *Log) Error(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Error(args...)
}

//FATAL
func (receiver *Log) Fatal(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Fatal(args...)
}

//Panic
func (receiver *Log) Panic(args ...interface{}) {
	g_logger.WithFields(receiver.fields).Panic(args...)
}

//TRACEf
func (receiver *Log) Tracef(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Tracef(format, args...)
}

//DEBUGf
func (receiver *Log) Debugf(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Debugf(format, args...)
}

//INFOf
func (receiver *Log) Infof(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Infof(format, args...)
}

//WARNf
func (receiver *Log) Warnf(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Warnf(format, args...)
}

//ERRORf
func (receiver *Log) Errorf(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Errorf(format, args...)
}

//FATALf
func (receiver *Log) Fatalf(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Fatalf(format, args...)
}

//Panicf
func (receiver *Log) Panicf(format string, args ...interface{}) {
	g_logger.WithFields(receiver.fields).Panicf(format, args...)
}
