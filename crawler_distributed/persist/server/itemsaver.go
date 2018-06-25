package main

import (
	"log"

	"fmt"

	"flag"

	"gopkg.in/olivere/elastic.v5"
	"project/goproj1/crawler/config"
	"project/goproj1/crawler_distributed/persist"
	"project/goproj1/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		//ElasticSearch server address and port
		elastic.SetURL("http://192.168.1.188:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
