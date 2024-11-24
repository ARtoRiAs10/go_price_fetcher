package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ARtoRiAs10/client"
)

func main() {

	client := client.New("http://localhost:3001")

	price, err := client.FetchPrice(context.Background(), "ET")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+V\n", price)

	
	listenAddr := flag.String("listenaddr", ":3000", "listen address server")
	flag.Parse()
	svc := NewLloggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

	// price, err := svc.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(price)

}
