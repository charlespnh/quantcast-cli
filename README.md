# quantcast-cli

## Binaries
The CLI application has been precompiled into the following binaries for each OS and architecture. These can also be found in Releases section on the right.

**Reminder**: Downloading the binary will require ```chmod u+x most_active_cookie``` to turn the binary into an executable.

### Linux x86-64
    https://github.com/charlespnh/quantcast-cli/releases/download/v0.1.0-x86-64/most_active_cookie

### Linux arm64
    https://github.com/charlespnh/quantcast-cli/releases/download/v0.1.0-arm64/most_active_cookie

### Darwin x86-64
    https://github.com/charlespnh/quantcast-cli/releases/download/v0.1.0-darwin-amd64/most_active_cookie

### Darwin arm64
    https://github.com/charlespnh/quantcast-cli/releases/download/v0.1.0-darwin-arm64/most_active_cookie


## Example

```
./most_active_cookie cookie_log.csv -d 2018-12-09
AtY0laUfhglK3lC7
```
```
./most_active_cookie cookie_log.csv -d 2018-12-08
SAZuXPGUrfbcn5UA
4sMM2LxV07bPJzwf
fbcn5UAVanZf6UtG
```
```
./most_active_cookie --help
Usage of ./most_active_cookie:
  -d, --date string   Date to filter by (default "1970-01-01")
```

Alternatively, if ```go``` is installed on your machine then within this project directory, run
```
go run . cookie_log.csv -d 2018-12-09
```

## Build
```
go build -o most_active_cookie
```

## Test
```
go test
```