package main

import (
	"log"
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	// Parse command line arguments
	var date string
	flag.StringVarP(&date, "date", "d", "1970-01-01", "Date to filter by")
	flag.Parse()

	filename := flag.Arg(0)

	// Parse the file
	records, err := GetCookiesWithinDate(filename, date)
	if err != nil {
		log.Fatal(err)
	}

	// Get the most active cookies
	mostActiveCookies := GetMostActiveCookies(records)
	for _, cookie := range mostActiveCookies {
		fmt.Println(cookie)
	}
}