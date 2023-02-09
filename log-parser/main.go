package main

import (
	"log-parser/app"
)

func main() {

	logFile := app.NewLogFile("../log.log")
	defer logFile.Close()

	parser := app.NewLogParser(logFile)

	parser.Parse()

}
