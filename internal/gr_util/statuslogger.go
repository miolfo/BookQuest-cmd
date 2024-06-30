package gr_util

import (
	"log"
)

type status struct {
	Status string
	Done   int
}

var currentStatus *status = &status{
	Status: "Starting",
	Done:   0,
}

var enabled bool = true

func UpdateLoggerStatus(status string) {

	currentStatus.Status = status
	currentStatus.Done = 0
	flushLogger()
}

func DisableLogger() {
	enabled = false
}

func UpdateLoggerDoneCount(count int) {

	currentStatus.Done = count
	flushLogger()
}

func flushLogger() {
	if enabled {
		log.Printf("STATUS_LOG,STATUS:%s,DONECOUNT:%d", currentStatus.Status, currentStatus.Done)
	}
}
