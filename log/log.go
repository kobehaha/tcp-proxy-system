package log

import (
	"log"
	"os"
	"path/filepath"
)

// description
// init log file
func Init(name string, DefaultLogFileLocation string) {
	filename := filepath.Join(DefaultLogFileLocation, name)
	os.MkdirAll(DefaultLogFileLocation, os.ModePerm)
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println("cannot create log file:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("success create log file:", logFile.Name())
}
