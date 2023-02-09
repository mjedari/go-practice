package app

import (
	"bufio"
	"fmt"
	"time"
)

const (
	ChunkSize = 100
	LogLines  = 100000000
)

type LogParser struct {
	reader *bufio.Reader
}

func NewLogParser(file *LogFile) *LogParser {
	reader := bufio.NewReader(file.file)
	return &LogParser{reader: reader}
}

func (p *LogParser) Parse() {
	iteration := LogLines / ChunkSize
	for i := 0; i < iteration; i++ {
		chunk := (i + 1) * 100
		p.parseChunk(chunk)
		// we should pars and delegate storing data to another function and async
		time.Sleep(time.Second * 3)
	}
}

func (p *LogParser) parseChunk(chunk int) {
	fmt.Printf("iteration chunk: #%v\n", chunk)
	for j := 0; j < ChunkSize; j++ {
		line, _, err := p.reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Printf("#%v======:%v\n", chunk+j, string(line))
	}
}
