package util

import (
	"encoding/json"
	"log"
	"os"
)

type StatusLogger struct {
	file string
}

type status struct {
	status string
	done   int
	total  int
}

var instance *StatusLogger
var currentStatus *status

func GetLogger() StatusLogger {

	if len(instance.file) == 0 {
		log.Fatal("Logger not initialized")
	}
	return *instance
}

func InitLogger(file string, initialCount int) {

	instance = &StatusLogger{file: file}
	currentStatus = &status{
		status: "Starting",
		done:   0,
		total:  initialCount,
	}
	flush()
}

func UpdateStatus(status string) {
	currentStatus.status = status
	flush()
}

func UpdateTotalCount(count int) {
	currentStatus.total = count
	flush()
}

func UpdateDoneCount(count int) {
	currentStatus.done = count
	flush()
}

func flush() {
	res, marshalErr := json.MarshalIndent(*currentStatus, "", "\t")
	if marshalErr != nil {
		log.Printf("Error writing status log")
		log.Fatal(marshalErr)
	}
	writeErr := os.WriteFile(instance.file, res, 0644)
	if writeErr != nil {
		log.Fatal("Unable to write results: %s", writeErr)
	}
}
