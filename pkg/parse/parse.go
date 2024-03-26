package parse

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"

	"parse/internal/config"
	"parse/internal/logger"
	"parse/internal/util"
	"parse/pkg/entities"
	"parse/pkg/httpreq"
)

func GetResolutions() ([]entities.FileJSON, error) {
	logger.InfoLogger.Println("Started collecting resolutions")
	currentURL, err := url.Parse(config.ResolutionsURL)
	if err != nil {
		return nil, err
	}
	values := currentURL.Query()
	files := make([]entities.FileJSON, 0)
	for section := config.SectionNumberStart; section <= config.SectionNumberEnd; section++ {
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

func GetProcedures() ([]entities.FileJSON, error) {
	logger.InfoLogger.Println("Started collecting procedures")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(config.ProceduresURL),
		chromedp.Nodes(`//a[contains(., 'DOWNLOAD FILE')]`, &nodes),
	)
	if err != nil {
		return nil, err
	}
	logger.InfoLogger.Printf("Files collected in procedures: %d\n", len(nodes))
	return nodesToFiles(nodes, 1), nil
}

func GetPositionPapers() ([]entities.FileJSON, error) {
	logger.InfoLogger.Println("Started collecting position papers")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(config.PositionPapersURL),
		chromedp.Nodes(`//a[contains(@title, 'IACS')]`, &nodes),
	)
	if err != nil {
		return nil, err
	}
	logger.InfoLogger.Printf("Files collected in position papers: %d\n", len(nodes))
	return nodesToFiles(nodes, 3), nil
}

func nodesToFiles(nodes []*cdp.Node, attributePosition int) []entities.FileJSON {
	var files []entities.FileJSON
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
		name := util.Reverse(builder.String())
		files = append(files, *entities.NewFileJSON(name, link))
	}
	return files
}
