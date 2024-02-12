package main

import (
	"encoding/csv"
	"os"
	"io"
	"strings"
)

// Reads the CSV file and returns the cookies that were active on the given date
func GetCookiesWithinDate(filename string, date string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Read the records
	cookies := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Parse for values
		cookie := record[0]
		timestamp := record[1]
		// Append the cookie if the timestamp contains the date of interest
		if strings.Contains(timestamp, date) {
			cookies = append(cookies, cookie)
		}
	}

	return cookies, nil
}

// Returns the most active cookies
func GetMostActiveCookies(records []string) []string {
	// Count the occurrences of each cookie
	cookieFreq := make(map[string]int)
	maxCount := 0
	for _, cookie := range records {
		cookieFreq[cookie]++
		maxCount = max(maxCount, cookieFreq[cookie])
	}

	// Find the most active cookie
	var mostActiveCookies []string
	for cookie, freq := range cookieFreq {
		if freq == maxCount {
			mostActiveCookies = append(mostActiveCookies, cookie)
		}
	}

	return mostActiveCookies
}

// Find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}