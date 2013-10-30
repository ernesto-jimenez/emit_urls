package url_extractor

import "net/url"
import "io"
import "code.google.com/p/go.net/html"

func ExtractURLs(url_str string, html_reader io.Reader) []url.URL {
  urls := []url.URL{}
  baseUrl, _ := url.Parse(url_str)

  d := html.NewTokenizer(html_reader)
  for {
    // token type
    tokenType := d.Next()
    if tokenType == html.ErrorToken {
      return urls
    }
    token := d.Token()
    switch tokenType {
      case html.StartTagToken: // <tag>
      if (token.Data == "a") {
        if href := href(token.Attr); href != "" {
          href_url, _ := url.Parse(href)
          link := baseUrl.ResolveReference(href_url)
          link.Fragment = ""
          if link.Scheme == "http" || link.Scheme == "https" {
            urls = append(urls, *link)
          }
        }
      }
    }
  }

  return urls
}

func href(attrs []html.Attribute) (href string) {
  for _, attr := range attrs {
    if attr.Key == "href" {
      return attr.Val
    }
  }
  return ""
}

