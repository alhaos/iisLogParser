package main

import (
	"flag"
	"iisLogParser/internal/configuration"
	"iisLogParser/internal/fsAdapter"
	"iisLogParser/internal/model"
	"iisLogParser/internal/parsing"
)

func main() {

	// parse flag
	configFilePointer := flag.String("config", "config.json", "Path to config file")
	flag.Parse()
	configFile := *configFilePointer

	// init config
	config, err := configuration.NewConfig(configFile)
	if err != nil {
		panic(err)
	}

	// init fs
	fs := fsAdapter.FS{}

	// init parser
	parser := parsing.Parser{}

	// get files from config.Source directory
	files, err := fs.GetFiles(config.Source)
	if err != nil {
		panic(err)
	}

	var commonLogEntries []*model.LogEntry

	// process files
	for _, file := range files {

		// extract log entries from file
		logEntries, err := parser.Parse(file)
		if err != nil {
			panic(err)
		}

		// append current logfile entries to common entries
		commonLogEntries = append(commonLogEntries, logEntries...)
	}

	// export common log entries to config.OutFileName file
	err = fs.ExportToJSON(commonLogEntries, config.OutFileName)
	if err != nil {
		panic(err)
	}
}
