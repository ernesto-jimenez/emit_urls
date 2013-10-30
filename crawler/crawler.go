package crawler

import "log"
import "os"
import "strings"
import "net/url"
import "net/http"
import "github.com/ernesto-jimenez/emit_urls/url_extractor"

func Crawl(initial_url string, channel chan FoundURL, logfile string) {
  parsedUrl, _ := url.Parse(initial_url)
  var output *os.File
  switch logfile {
  case "/dev/stdout":
    output = os.Stdout
  case "/dev/stderr":
    output = os.Stderr
  default:
    output, _ = os.OpenFile(logfile, os.O_CREATE, 0666)
  }
  log.SetOutput(output)

  host := parsedUrl.Host

  crawl(initial_url, channel, host, make(map[string]bool), []string{})
}

func crawl(url string, channel chan FoundURL, host string,
            visitedURLs map[string]bool, queue []string) {
  resp, err := http.Get(url)
  if err != nil {
    log.Printf("%v", err)
  }
  defer resp.Body.Close()

  visitedURLs[url] = true
  contentType := resp.Header.Get("Content-Type")
  crawled := FoundURL{StatusCode: resp.StatusCode, Url: url, contentType: contentType}
  log.Printf("%v\n", crawled)

  if resp.StatusCode == 200 {
    channel <- crawled
    if strings.Contains(crawled.contentType, "text/html") {
      urls := url_extractor.ExtractURLs(url, resp.Body)
      for i := 0; i < len(urls); i++ {
        next_url := urls[i].String()
        if (!visitedURLs[next_url] && urls[i].Host == host) {
          visitedURLs[next_url] = true
          queue = append(queue, next_url)
        }
      }
      log.Printf("%v urls found\n", len(urls))
    }
  }

  if len(queue) > 0 {
    log.Printf("Queue: %v items\n", len(queue))
    crawl(queue[0], channel, host, visitedURLs, queue[1:])
  } else {
    close(channel)
  }
}

type FoundURL struct {
  Url string
  StatusCode int
  contentType string
}

