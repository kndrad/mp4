package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var ErrDateNotFound = errors.New("no date found in URL")

func parse(url string) (time.Time, error) {
	// Regular expression to match the date pattern (YYYY-MM-DD)
	dateRegex := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

	// Find the date in the URL
	dateMatch := dateRegex.FindString(url)
	if dateMatch == "" {
		return time.Time{}, ErrDateNotFound
	}

	// Parse the date
	date, err := time.Parse("2006-01-02", dateMatch)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}

	return date, nil
}

func MP4Name(date time.Time, name, part string) string {
	if name == "" {
		name = "video"
	}
	if part == "" {
		return fmt.Sprintf("%s_%s.mp4", date.Format("2006-01-02"), name)
	} else {
		return fmt.Sprintf("%s_%s_%s.mp4", date.Format("2006-01-02"), name, part)
	}
}
