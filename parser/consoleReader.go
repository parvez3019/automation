package parser

import (
	"bufio"
	. "os"
)

type Reader interface {
	ReadLine() string
}

type ConsoleReader struct {
}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (c *ConsoleReader) ReadLine() string {
	reader := bufio.NewReader(Stdin)
	inputStringFromConsole, _ := reader.ReadString('\n')
	return inputStringFromConsole
}
