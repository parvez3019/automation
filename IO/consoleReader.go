package IO

import (
	"bufio"
	. "os"
)

type Reader interface {
	Read() string
}

type ConsoleReader struct {
}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (c *ConsoleReader) Read() string {
	reader := bufio.NewReader(Stdin)
	inputStringFromConsole, _ := reader.ReadString('\n')
	return inputStringFromConsole
}
