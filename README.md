# emit_urls commandline tool

This is an small tool that will crawl a website and print in STDOUT all crawled URLs.

# Possible uses

* See all the URLs crawlable in the local development server:
```emit_urls http://localhost:3000```
* Run some other command for all each crawlable url with GNU/Parallel:
```emit_urls http://localhost:3000 | parallel "some_command {}"```
* Find out about your links producing 404:
```emit_urls --logfile=/dev/stdout http://0.0.0.0:8000 | grep 404```

# Installation

```
go get github.com/ernesto-jimenez/emit_urls
go install github.com/ernesto-jimenez/emit_urls
```

