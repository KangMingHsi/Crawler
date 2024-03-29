package main

import (
	"flag"
	"fmt"
	"log"
	"spider_distributed/rpcsupport"
	"spider_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRPC(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
