package url_extractor

import "testing"
import "net/url"
import "strings"

type testexpectation struct {
  url string
  html string
  expected []url.URL
}

var tests = []testexpectation{
  {"http://example.com", "", []url.URL{}},
  {"http://example.com", "<a href=\"/this/works\">", []url.URL{
    {Scheme: "http", Host: "example.com", Path: "/this/works"},
  }},
  {"http://example.com", "<a href=\"http://whatever.es/\">", []url.URL{
    {Scheme: "http", Host: "whatever.es", Path: "/"},
  }},
  {"https://example.com", "<a href=\"http://whatever.es/\">", []url.URL{
    {Scheme: "http", Host: "whatever.es", Path: "/"},
  }},
  {"http://example.com", "<a href=\"https://whatever.es/\">", []url.URL{
    {Scheme: "https", Host: "whatever.es", Path: "/"},
  }},
  {"http://example.com", "<a href=\"https://whatever.es/whatever#remove\">", []url.URL{
    {Scheme: "https", Host: "whatever.es", Path: "/whatever"},
  }},
}

func TestExtractURLs(t *testing.T) {
  var urls []url.URL
  for _, test := range tests {
    urls = ExtractURLs(test.url, strings.NewReader(test.html))
    if !equalResults(test.expected, urls) {
      t.Error("Expected", test.expected, "Got", urls)
    }
  }
}

func equalResults(expected []url.URL, returned []url.URL) bool {
  exp_len := len(expected)
  ret_len := len(returned)
  if exp_len != ret_len {
    return false
  }
  for i := 0; i < exp_len; i++ {
    if expected[i] != returned[i] {
      return false
    }
  }
  return true
}

