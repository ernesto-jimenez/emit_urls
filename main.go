package main

import (
  "github.com/ernesto-jimenez/emit_urls/crawler"
  "os"
  "fmt"
  "regexp"
  "flag"
)

func main() {
  var logfile string
  var url string

  // Command line options
  flag.StringVar(&logfile, "logfile", os.DevNull, "File to output to")
  flag.Parse()
  if (len(flag.Args()) != 1 && !(flag.Lookup("help") != nil)) {
    os.Stderr.WriteString("Usage: emit_urls [url]\n")
    os.Exit(1)
  }

  // Process URL
  url = flag.Arg(0)
  regex, _ := regexp.Compile("https?://.*")

  if regex.MatchString(url) {
    channel := make(chan crawler.FoundURL)
    go crawler.Crawl(url, channel, logfile)
    printer(channel)
  } else {
    os.Stderr.WriteString("Enter a valid URL such as http://example.com\n")
    os.Exit(1)
  }
}

func printer(channel chan crawler.FoundURL) {
  var url crawler.FoundURL
  ok := true
  for ok {
    url, ok = <- channel
    if (ok) {
      fmt.Println(url.Url)
    }
  }
}

