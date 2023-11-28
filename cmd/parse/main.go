package main

import (
	"fmt"
	"os"

	"parse/iternal/logger"
	"parse/iternal/parse"
)

func main() {

	logger.InfoLogger.Println("Parser started working")

	positionPapersFiles, err := parse.GetPositionPapers()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing position papers! %s\n", err)
	}
	_ = positionPapersFiles
	proceduresFiles, err := parse.GetProcedures()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing procedures! %s\n", err)
	}
	fmt.Println(proceduresFiles)

	resolutionFiles, err := parse.GetResolutions()
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing resolutions! %s\n", err)
		os.Exit(1)
	}
	for _, rf := range resolutionFiles {
		fmt.Println(rf.RenderedName.Name, rf.ACFLink.Link)
	}
	logger.InfoLogger.Println("Parser finished working")
}
