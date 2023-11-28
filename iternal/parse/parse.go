package parse

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"

	"parse/iternal/entities"
	"parse/iternal/httpreq"
	"parse/iternal/logger"
)

const ResolutionsURL = "https://iacs.flumeserver.co.za/wp-json/wp/v2/publications?sections=?&per_page=10&page=?&status=publish&orderby=menu_order&order=asc&acf_format=standard"

const ProceduresURL = "https://iacs.org.uk/membership/procedures"

const PositionPapersURL = "https://iacs.org.uk/about-us/position-papers"

func GetResolutions() ([]entities.File, error) {
	logger.InfoLogger.Println("Started collecting resolutions")
	currentURL, err := url.Parse(ResolutionsURL)
	if err != nil {
		return nil, err
	}
	values := currentURL.Query()
	files := make([]entities.File, 0)
	for section := 10; section <= 11; section++ {
		values.Set("sections", strconv.Itoa(section))
		size := len(files)
		for page := 1; ; page++ {
			values.Set("page", strconv.Itoa(page))
			currentURL.RawQuery = values.Encode()
			filesURL, err := httpreq.GetFilesRequest(currentURL.String())
			for _, f := range filesURL {
				files = append(files, f)
			}
			if len(filesURL) == 0 || err != nil {
				break
			}

		}
		logger.InfoLogger.Printf("Number of files collected from section %d : %d\n", section, len(files)-size)
	}
	return files, nil
}

func GetProcedures() ([]entities.File, error) {
	logger.InfoLogger.Println("Started collecting procedures")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(ProceduresURL),
		chromedp.Nodes(`//a[contains(., 'DOWNLOAD FILE')]`, &nodes),
	)
	if err != nil {
		return nil, err
	}
	logger.InfoLogger.Printf("Files collected in procedures: %d\n", len(nodes))
	return nodesToFiles(nodes, 1), nil
}

func GetPositionPapers() ([]entities.File, error) {
	logger.InfoLogger.Println("Started collecting position papers")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(PositionPapersURL),
		chromedp.Nodes(`//a[contains(@title, 'IACS')]`, &nodes),
	)
	if err != nil {
		return nil, err
	}
	logger.InfoLogger.Printf("Files collected in position papers: %d\n", len(nodes))
	return nodesToFiles(nodes, 3), nil
}

func nodesToFiles(nodes []*cdp.Node, attributePosition int) []entities.File {
	var files []entities.File
	for _, node := range nodes {
		var builder strings.Builder
		builder.Grow(100)
		link := node.Attributes[attributePosition]
		for i := len(link) - 1; i >= 0; i-- {
			if link[i] == '/' {
				break
			}
			builder.WriteByte(link[i])
		}
		name := reverse(builder.String())
		files = append(files, *entities.NewFile(name, link))
	}
	return files
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
