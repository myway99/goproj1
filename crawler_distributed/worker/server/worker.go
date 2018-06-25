package main

import (
	"fmt"

	"log"

	"flag"

	"project/goproj1/crawler/fetcher"
	"project/goproj1/crawler_distributed/rpcsupport"
	"project/goproj1/crawler_distributed/worker"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
