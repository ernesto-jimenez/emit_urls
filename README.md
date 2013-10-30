# emit_urls commandline tool

This is an small tool that will crawl a website and print in STDOUT all crawled URLs.

# Possible uses

* See all the URLs crawlable in the local development server:
```emit_urls http://localhost:3000```
* Run some other command for all each crawlable url with GNU/Parallel
```emit_urls http://localhost:3000 | parallel "some_command {}"```

# Installation

```go install github.com/ernesto-jimenez/emit_urls```
