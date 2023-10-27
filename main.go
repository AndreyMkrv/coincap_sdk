package main

import (
	"coincap_sdk/coincap"
	"fmt"
	"log"
	"time"
)

func main() {

	coincapClient, error := coincap.NewClient(time.Second * 10)
	if error != nil {
		log.Fatal(error)
	}

	assets, error := coincapClient.GetAssets()
	if error != nil {
		log.Fatal(error)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

}
