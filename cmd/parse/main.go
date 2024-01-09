package main

import (
	"os"
	"parse/iternal/config"
	"parse/iternal/logger"
	"parse/iternal/util"
	"parse/pkg/csv"
	"parse/pkg/database"
	"parse/pkg/email"
	"parse/pkg/entities"
	"parse/pkg/parse"
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

	util.CleanFromTags(&allFiles)

	changedFiles, err := database.FindAllChanges(allFiles)
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while working with database! %s\n", err)
		os.Exit(1)
	}

	if len(changedFiles) != 0 {
		err := csv.ConvertToCSV(changedFiles)
		if err != nil {
			logger.ErrorLogger.Printf("An error occurred while saving changes to csv file! %s\n", err)
			os.Exit(1)
		}
		logger.InfoLogger.Printf("Changes were found: %d", len(changedFiles))
		err = email.SendNotificationEmail()
		if err != nil {
			logger.ErrorLogger.Printf("An error occurred while trying to send notification messages! %s\n", err)
			os.Exit(1)
		}
	} else {
		logger.InfoLogger.Println("Changes weren't found")
	}

	logger.InfoLogger.Println("Parser finished working")
}
