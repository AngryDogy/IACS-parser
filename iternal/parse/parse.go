package parse

import (
	"net/url"
	"parse/iternal/entities"
	"parse/iternal/httpreq"
	"parse/iternal/logger"
	"strconv"
)

const GetURL = "https://iacs.flumeserver.co.za/wp-json/wp/v2/publications?sections=?&per_page=10&page=?&status=publish&orderby=menu_order&order=asc&acf_format=standard"

func GetResolutions() ([]entities.File, error) {
	currentURL, err := url.Parse(GetURL)
	if err != nil {
		return nil, err
	}
	values := currentURL.Query()
	files := make([]entities.File, 0)
	for section := 1; section <= 100; section++ {
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
		logger.InfoLogger.Println("Number of files collected from section %d : %d", section, len(files)-size)
	}
	return files, nil
}
