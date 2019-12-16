# PPSearch (pinpoint search)
A small and simple search engine created as an interview project for pinpoint)

## Use

build binary and run
```
go build -o ppsearch

# Index a webpage and all linked pages up to 3 pages deep `ppsearch construct <url>`
./ppsearch construct austinpetsalive.com

# Get results `ppsearch search <term>`
./ppsearch search dog

# Clear results `ppearch reset`
./ppsearch reset
```

### Caveats
- Once a url has been indexed it will not update it's word count if crawled again. This can be changed at a later point but has been set up this way to work with the simple append only DB