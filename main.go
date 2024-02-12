package main

import (
	"encoding/csv"
	"os"
	// "flag"
	"log"
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	// Parse command line arguments
	var date string
	flag.StringVarP(&date, "date", "d", "1970-01-01", "Date to filter by")
	flag.Parse()

	// Open the file
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the file
	records, err := GetCookiesWithinTimestamp(csv.NewReader(file), date)
	if err != nil {
		log.Fatal(err)
	}
	// Get the most active cookies
	mostActiveCookies := GetMostActiveCookies(records)
	for _, cookie := range mostActiveCookies {
		fmt.Println(cookie)
	}
}