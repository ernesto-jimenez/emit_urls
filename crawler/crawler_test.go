package crawler

import "fmt"
import "testing"
import "net"
import "net/http"
import "github.com/bmizerany/pat"
import "github.com/bmizerany/assert"
import "io"
import "os"
import "path"

const (
  FileFixturesBasePath = "./test_fixtures/"
)

type testexpectation struct {
  url string
  expected []string
}

var tests = []testexpectation{
  {"http://localhost:8080/basictest/index.html", []string{
    "http://localhost:8080/basictest/index.html",
    "http://localhost:8080/basictest/second.html",
    "http://localhost:8080/basictest/third.html",
    "http://localhost:8080/basictest/fourth.html",
  }},
  {"http://localhost:8080/non-existent/index.html", []string{
  }},
}


func Handler() http.Handler {
  m := pat.New()

  m.Get("/:dir/:file", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    dir := req.URL.Query().Get(":dir")
    file := req.URL.Query().Get(":file")
    f, e := os.Open(path.Join(FileFixturesBasePath, dir, file))
    if e != nil {
      // Treat errors as 404s - file not found
      w.WriteHeader(http.StatusNotFound)
    } else {
      w.WriteHeader(http.StatusOK)
      io.Copy(w, f)
    }
  }))

  return m
}

func TestCrawl(t *testing.T) {
  l, err := net.Listen("tcp", ":8080")
  if err != nil {
    t.Fatal(err)
  }
  http.Handle("/", Handler())
  go func() {
    http.Serve(l, nil)
  }()

  for _, test := range tests {
    channel := make(chan FoundURL)
    go Crawl(test.url, channel, os.DevNull)

    result := make([]string, 0)
    var url FoundURL
    ok := true
    for ok {
      url, ok = <- channel
      if (ok) {
        result = append(result, url.Url)
        fmt.Println(url.Url)
      } else {
        fmt.Println("End")
      }
    }

    assert.Equal(t, result, test.expected)
  }

  // Close listener
  if err = l.Close(); err != nil {
    panic(err)
  }
}

