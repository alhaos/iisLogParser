package parsing

import (
	"errors"
	"fmt"
	"iisLogParser/internal/fsAdapter"
	"iisLogParser/internal/model"
	"io"
	"strings"
)

func (p Parser) Parse(filename string) ([]*model.LogEntry, error) {

	var LogEntries []*model.LogEntry

	logfile, err := fsAdapter.NewLogfile(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer logfile.Close()

	var line string
	var lp *lineParser
	for {
		line, err = logfile.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading line: %w", err)
		}

		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "#Fields") {
			continue
		}

		if strings.HasPrefix(line, "#Fields") {
			lp = newLineParser(line)
			continue
		}

		le, err := lp.parseLine(line)
		if err != nil {
			return nil, err
		}

		LogEntries = append(LogEntries, le)

	}

	return LogEntries, nil
}
