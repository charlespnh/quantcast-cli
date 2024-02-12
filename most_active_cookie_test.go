package main

import (
	"testing"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/stretchr/testify/assert"
)

func TestOpenCSVFile(t *testing.T) {
	t.Parallel()
	validFilename := "logs/cookie_log.csv"
	invalidFilename := "cookie_log_invalid.csv"
	invalidFilename2 := "logs/cookie_log.json"

	file, err := os.Open(validFilename)
	assert.NoError(t, err)
	file.Close()

	file, err = os.Open(invalidFilename)
	assert.Error(t, err)
	file.Close()

	file, err = os.Open(invalidFilename2)
	assert.Error(t, err)
	file.Close()
}

func TestCommandLineFlags(t *testing.T) {
	t.Parallel()
	validArgs := []string{"logs/cookie_log.csv", 
						  "-d", "2018-12-07"}
	validArgs2 := []string{"logs/cookie_log.csv"}
	invalidArgs := []string{"logs/cookie_log.csv", "-d"}

	f := flag.NewFlagSet("test", flag.ContinueOnError)
	date := f.StringP("date", "d", "1970-01-01", "Date value 1")
	err := f.Parse(validArgs)
	assert.NoError(t, err)
	assert.Equal(t, "2018-12-07", *date)

	f = flag.NewFlagSet("test", flag.ContinueOnError)
	date = f.StringP("date", "d", "1970-01-01", "Date value 2")
	err = f.Parse(validArgs2)
	assert.NoError(t, err)
	assert.Equal(t, "1970-01-01", *date)

	f = flag.NewFlagSet("test", flag.ContinueOnError)
	err = f.Parse(invalidArgs)
	assert.Error(t, err)
}

func TestGetCookiesWithinDate(t *testing.T) {

}

func TestGetMostActiveCookies(t *testing.T) {
	
}