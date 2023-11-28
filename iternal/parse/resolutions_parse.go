package parse

import (
	"context"
	"fmt"
	"parse/iternal/logger"
	"strings"
	"time"
	"unicode"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const WebsiteURL = "https://iacs.org.uk"

//const ResolutionsURL = "https://iacs.org.uk/resolutions"

func ResolutionsParse(ctx context.Context) ([]string, error) {
	subsections, err := getSideMenuLinks(ctx, ResolutionsURL, ResolutionsURL)
	if err != nil {
		return nil, err
	}
	logger.InfoLogger.Printf("Number of subsections received: %d \n", len(subsections))

	fileNames := make([]string, 0)
	for _, subsection := range subsections {
		pages, err := getSideMenuLinks(ctx, subsection, WebsiteURL)
		if err != nil {
			return nil, err
		}
		logger.InfoLogger.Printf("Subsection - %s is being parsed\n", subsection)
		logger.InfoLogger.Printf("Number of pages received: %d\n", len(pages))
		for _, page := range pages {
			logger.InfoLogger.Printf("Page - %s is being parsed\n", page)
			pageNames, err := getFileNames(ctx, page)
			if err != nil {
				logger.ErrorLogger.Printf("An error occurred while parsing page %s: %s\n", page, err)
				continue
			}
			for _, s := range pageNames {
				fileNames = append(fileNames, s)
			}

		}

	}
	return fileNames, nil
}

func getSideMenuLinks(ctx context.Context, currentURL string, parentURL string) ([]string, error) {
	var sideMenuNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(currentURL),
		chromedp.Nodes(".side-menu", &sideMenuNodes),
	)
	if err != nil {
		return nil, err
	}
	sideMenuLinks := make([]string, 0)
	var builder strings.Builder
	for _, node := range sideMenuNodes {
		builder.Reset()
		builder.Grow(100)
		builder.WriteString(parentURL)
		builder.WriteString(node.Attributes[1])
		sideMenuLinks = append(sideMenuLinks, builder.String())
	}
	return sideMenuLinks, nil
}

func getFileNames(ctx context.Context, currentURL string) ([]string, error) {
	wasAdded := make(map[string]bool, 500)
	/*err := chromedp.Run(ctx,
		chromedp.Navigate(currentURL),
	)
	if err != nil {
		return nil, err
	}*/
	fileNames := make([]string, 0)

Loop:
	for i := 0; ; i++ {
		fmt.Println(i)
		var fileNodes []*cdp.Node
		err := chromedp.Run(ctx,
			chromedp.Navigate(currentURL),
		)
		if err != nil {
			return nil, err
		}
		for j := 0; j < i; j++ {
			err := chromedp.Run(ctx,
				chromedp.Click(`//button[@type='button' and contains(., 'Next')]`),
				chromedp.Sleep(time.Duration(1)*time.Second),
			)
			if err != nil {
				return nil, err
			}
		}
		err = chromedp.Run(ctx,
			chromedp.Nodes(".leading-snug", &fileNodes),
		)
		if err != nil {
			return nil, err
		}
		for _, node := range fileNodes {
			name := parseName(node.Children[0].NodeValue)
			if !wasAdded[name] {
				wasAdded[name] = true
				fileNames = append(fileNames, name)
			} else {
				break Loop
			}

		}
		/*err = chromedp.Run(ctx,
			chromedp.Click(`//button[@type='button' and contains(., 'Next')]`),
			chromedp.Sleep(time.Duration(1)*time.Second),
			chromedp.NavigateForward(),
		)
		if err != nil {
			return nil, err
		}*/
	}
	logger.InfoLogger.Printf("%d files collected", len(fileNames))
	return fileNames, nil
}

func parseName(s string) string {
	firstLetter := false
	l := 0
	r := 0
	for i := range s {
		if unicode.IsLetter(rune(s[i])) {
			if !firstLetter {
				l = i
				firstLetter = true
			} else {
				r = i
			}
		}
	}
	return s[l : r+1]
}
