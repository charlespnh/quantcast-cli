package main

import (
	"log"
	"os"
	"flag"

	mac "most_active_cookie"
)

func main() {
	// Parse command line arguments
	timestamp := flag.String("d", "", "Timestamp to filter by")
	flag.Parse()
	
	// Open the file
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the file
	records, err := mac.GetCookiesWithinTimestamp(csv.NewReader(file), timestamp)
	if err != nil {
		log.Fatal(err)
	}

	// Get the most active cookies
	mostActiveCookies := mac.GetMostActiveCookies(records)
	log.Println(mostActiveCookies)
}