package main

import (
	"fmt"
	"os"
	"parse/iternal/config"
	"parse/pkg/database"
	"parse/pkg/entities"
	"parse/pkg/parse"

	"parse/iternal/logger"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		logger.ErrorLogger.Println("Can't load env variables!")
		os.Exit(1)
	}

	logger.InfoLogger.Println("Parser started working")

	positionPapersFiles, err := parse.GetPositionPapers()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing position papers! %s\n", err)
	}

	proceduresFiles, err := parse.GetProcedures()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing procedures! %s\n", err)
	}

	resolutionFiles, err := parse.GetResolutions()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing resolutions! %s\n", err)
		os.Exit(1)
	}

	allFiles := make([]entities.FileJSON, 0)
	for _, f := range positionPapersFiles {
		allFiles = append(allFiles, f)
	}

	for _, f := range proceduresFiles {
		allFiles = append(allFiles, f)
	}

	for _, f := range resolutionFiles {
		allFiles = append(allFiles, f)
	}
	messageChanges, err := database.FindAllChanges(allFiles)
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while working with database! %s\n", err)
		os.Exit(1)
	}
	for _, m := range messageChanges {
		fmt.Println(m.Name)
	}

	logger.InfoLogger.Println("Parser finished working")
}
