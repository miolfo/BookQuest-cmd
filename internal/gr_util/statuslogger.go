package gr_util

import (
	"log"
)

type StatusLogger struct {
	file string
}

type status struct {
	Status string
	Done   int
	Total  int
}

var instance *StatusLogger
var currentStatus *status

func GetLogger() StatusLogger {

	return *instance
}

func InitLogger(file string) {

	instance = &StatusLogger{file: file}
	currentStatus = &status{
		Status: "Starting",
		Done:   0,
		Total:  0,
	}
}

func UpdateLoggerStatus(status string) {

	if instance == nil {
		log.Printf("Logger not initialized, skipping update")
		return
	}
	currentStatus.Status = status
}

func UpdateLoggerTotalCount(count int) {
	if instance == nil {
		log.Printf("Logger not initialized, skipping update")
		return
	}
	currentStatus.Total = count
}

func UpdateLoggerDoneCount(count int) {
	if instance == nil {
		log.Printf("Logger not initialized, skipping update")
		return
	}
	currentStatus.Done = count
}

func FlushLogger() {

	log.Printf("STATUS LOG: FLUSH")
}
