package common

import "os"
import "log"

type LogSys struct {
	FileHandle *os.File
	LogHandle  *log.Logger
}

func (this *LogSys) InitLogFile(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		this.FileHandle, err = os.Create(filePath)
	} else {
		this.FileHandle, err = os.OpenFile(filePath, os.O_APPEND, 0666)
	}

	if err != nil {
		os.Exit(1)
	}
	this.LogHandle = log.New(this.FileHandle, "[runtime]", log.Ldate|log.Ltime)
	this.LogHandle.Println("created")
}

func (this *LogSys) SaveLog(message string) {
	this.LogHandle.Println(message)
}
