package main

import (
	"os"
	"parse/pkg/email"

	"parse/iternal/config"
	"parse/iternal/logger"
	"parse/pkg/database"
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

	allFiles[40].ACF.FutureName = "edite 1.0"
	allFiles[40].ACF.FutureLink = "http:/localhost:80180"
	allFiles[41].ACF.Description = "was deleted"
	allFiles[41].ACF.Notes = "would be added new soon 2.0"
	messages, err := database.FindAllChanges(allFiles)
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while working with database! %s\n", err)
		os.Exit(1)
	}

	if len(messages) != 0 {
		logger.InfoLogger.Printf("Changes were found: %d", len(messages))
		email.SendNotificationEmail(messages)
	} else {
		logger.InfoLogger.Println("Changes weren't found")
	}

	logger.InfoLogger.Println("Parser finished working")
}
