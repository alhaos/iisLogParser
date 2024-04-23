package fsAdapter

import (
	"bufio"
	"os"
)

type FS struct{}

type Logfile struct {
	datafile *os.File
	Reader   *bufio.Reader
}
