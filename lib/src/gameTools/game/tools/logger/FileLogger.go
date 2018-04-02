package logger

import (
	"time"
	"os"
)

//	"log"

const (
	//go-logger version
	_VER string = "1.0.3"
)

type LEVEL int32
type UNIT int64
type _ROLLTYPE int //dailyRolling ,rollingFile

const _DATEFORMAT = "2006-01-02"

var logLevel LEVEL = 1

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

const (
	_DAILY _ROLLTYPE = iota
	_ROLLFILE
	LOG_ROOT = "logs"
)


type FileLogger struct {
	name string
}


func NewFileLogger(name string) * FileLogger {
	this := new(FileLogger)
	this.name = name

	this.createRollingFile(this.name ,0, 300, MB)
	this.setConsole(false)

	return this
}


func (this * FileLogger) Writer(content string){
	this.Debug(content)
}

func (this * FileLogger) setConsole(isConsole bool) {
	defaultlog.setConsole(isConsole)
}

func (this * FileLogger) setLevel(_level LEVEL) {
	defaultlog.setLevel(_level)
}

func (this * FileLogger) setFormat(logFormat string) {
	defaultlog.setFormat(logFormat)
}


func (this * FileLogger) createRollingFile( fileName string, maxNumber int32, maxSize int64, _unit UNIT){
	path := this.getPath()+"/"+LOG_ROOT+"/"+fileName+"/"+time.Now().Format(_DATEFORMAT)
	defaultlog.setRollingFile(path, fileName, maxNumber, maxSize, _unit)
}

func (this * FileLogger) setRollingFile(fileDir, fileName string, maxNumber int32, maxSize int64, _unit UNIT) {
	defaultlog.setRollingFile(fileDir, fileName, maxNumber, maxSize, _unit)
}

func (this * FileLogger) setRollingDaily(fileDir, fileName string) {
	defaultlog.setRollingDaily(fileDir+"/"+time.Now().Format(_DATEFORMAT), fileName)
}

func (this * FileLogger) Debug(v ...interface{}) {
	defaultlog.debug(v...)
}
func (this * FileLogger) Info(v ...interface{}) {
	defaultlog.info(v...)
}
func (this * FileLogger) Warn(v ...interface{}) {
	defaultlog.warn(v...)
}
func (this * FileLogger) Error(v ...interface{}) {
	defaultlog.error(v...)
}
func (this * FileLogger) Fatal(v ...interface{}) {
	defaultlog.fatal(v...)
}

func (this * FileLogger) setLevelFile(level LEVEL, dir, fileName string) {
	defaultlog.setLevelFile(level, dir, fileName)
}

func (this * FileLogger)  getPath() (string){
	dir, _ := os.Getwd()
	return dir
}
