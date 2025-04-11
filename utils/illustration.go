package utils

import "regexp"

func FindIllustrations(text string) ([]string, error) {
	regex := `!\[([^\]]*)\]\(([^)]+)\)`

	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(text, -1)

	var illustrations []string

	for _, match := range matches {
		if len(match) > 2 {
			illustrations = append(illustrations, match[2])
		}
	}

	return illustrations, nil
}
