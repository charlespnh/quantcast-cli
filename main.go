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

	// Get the most active cookies
	cookie, err := mac.GetMostActiveCookies(records)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cookie)
}