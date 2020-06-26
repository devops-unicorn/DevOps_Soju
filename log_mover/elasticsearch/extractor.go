package main

import (
  "log"
  elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

func main() {
  es, _ := elasticsearch7.NewDefaultClient()
  log.Println(es.Info())
}