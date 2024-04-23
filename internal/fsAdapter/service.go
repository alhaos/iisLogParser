package fsAdapter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"iisLogParser/internal/model"
	"os"
	"path/filepath"
	"strings"
)

// GetFiles get files list from directory
func (f *FS) GetFiles(directory string) ([]string, error) {

	var files []string

	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("ReadDir Error: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		files = append(files, filepath.Join(directory, entry.Name()))
	}

	return files, nil
}

// NewLogfile constructor for Logfile struct
func NewLogfile(filename string) (*Logfile, error) {

	datafile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open error: %w", err)
	}

	reader := bufio.NewReader(datafile)

	logfile := &Logfile{
		datafile: datafile,
		Reader:   reader,
	}

	return logfile, nil
}

// Close logfile
func (f Logfile) Close() error {
	return f.datafile.Close()
}

// ReadLine from file
func (f Logfile) ReadLine() (string, error) {
	read, err := f.Reader.ReadString('\r')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(read), nil
}

// ExportToJSON export []*model.LogEntry to json file
func (f *FS) ExportToJSON(logfileEntries []*model.LogEntry, filename string) error {

	data, err := json.Marshal(logfileEntries)
	if err != nil {
		return err
	}

	b := bytes.Buffer{}

	err = json.Indent(&b, data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
