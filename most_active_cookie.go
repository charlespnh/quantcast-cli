package main

import (
	"encoding/csv"
	"io"
	"strings"
	"log"
)

func GetCookiesWithinTimestamp(reader *csv.Reader, date string) ([]string, error) {
	// Read the header
	_, err := reader.Read()
	if err != nil {
		return "", err
	}

	// Read the records
	var cookies []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		// Parse the record
		recordParsed = strings.Split(record, " ")
		cookie := recordParsed[0]
		timestamp := recordParsed[1]

		// Append the cookie if the timestamp contains the date of interest
		if strings.Contains(timestamp, date){
			cookies = append(cookies, cookie)
		}
	}

	return records, nil
}

func GetMostActiveCookies(records []string) ([]string, error) {

	return records, nil
}