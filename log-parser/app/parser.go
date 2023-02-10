package app

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	BufferSize = 4 * 1024 * 1024 // 4MB
)

type LogParser struct {
	scanner *bufio.Scanner
}

func NewLogParser(file *LogFile) *LogParser {
	scanner := bufio.NewScanner(file.file)
	scanner.Buffer(make([]byte, BufferSize), BufferSize)
	return &LogParser{scanner: scanner}
}

func (p *LogParser) Parse() {
	startTime := time.Now()
	p.parseChunk()
	endTime := time.Now()
	fmt.Println("Duration: ", endTime.Sub(startTime).String())
}

func (p *LogParser) parseChunk() {
	number := 1
	for p.scanner.Scan() {
		line := p.scanner.Text()
		// process line
		fmt.Printf("%v read line: %v\n", number, line)
		number++
	}
	fmt.Println("Finished: #", number)
	if err := p.scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
