package util

import (
	"github.com/microcosm-cc/bluemonday"
	"parse/pkg/entities"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CleanFromTags(files *[]entities.FileJSON) {
	p := bluemonday.StripTagsPolicy()
	for _, f := range *files {
		f.ACF.Description = p.Sanitize(f.ACF.Description)
		f.ACF.FutureDescription = p.Sanitize(f.ACF.FutureDescription)
	}
}
