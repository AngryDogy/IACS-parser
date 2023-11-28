package main

import (
	"fmt"
	"parse/iternal/logger"
	"parse/iternal/parse"
)

func main() {

	logger.InfoLogger.Println("Parser started working")
	resolutionFiles, err := parse.GetResolutions()
	if err != nil {
		panic(err)
	}
	for _, rf := range resolutionFiles {
		fmt.Println(rf.RenderedName.Name, rf.ACFLink.Link)
	}

	/*ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	resolutionsNames, err := parse.ResolutionsParse(ctx)
	if err != nil {
		logger.ErrorLogger.Printf("An error occurred while parsing resolutions: %s", err)
		os.Exit(1)
	}
	fmt.Println(resolutionsNames)*/
}
