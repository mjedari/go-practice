package app

import "os"

type LogFile struct {
	file *os.File
}

func NewLogFile(name string) *LogFile {

	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		panic("can not read the file")
	}
	return &LogFile{file: file}
}

func (l *LogFile) Close() {
	l.file.Close()
}
