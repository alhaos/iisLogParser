package parsing

import (
	"fmt"
	"iisLogParser/internal/model"
	"strings"
)

// newLineParser constructor lineParser struct
func newLineParser(fieldLine string) *lineParser {

	fields := strings.Fields(fieldLine)[1:]

	lp := &lineParser{fields: make(map[string]int)}

	for i, field := range fields {
		lp.fields[field] = i
	}

	return lp
}

// parseLine parse data line to *model.LogEntry
func (p *lineParser) parseLine(line string) (*model.LogEntry, error) {

	fields := strings.Fields(line)

	if len(fields) < len(p.fields) {
		return nil, fmt.Errorf("line has %d fields, expected %d", len(fields), len(p.fields))
	}

	le := &model.LogEntry{
		Datetime: fields[p.fields["date"]] + fields[p.fields["time"]],
		URL:      fields[p.fields["cs-uri-stem"]],
		Username: fields[p.fields["cs-username"]],
		IP:       fields[p.fields[",cs-uri-query"]],
	}

	return le, nil
}
