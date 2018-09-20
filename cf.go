package main

import (
	"fmt"
	cf "github.com/cloudflare/cloudflare-go"
	"log"
	"os"
)

type cfAuth struct {
	ID    string
	Email string
}

func (cfA *cfAuth) init() {
	var ok bool
	cfA.ID, ok = os.LookupEnv("CLOUDFLARE_API_KEY")
	if !ok {
		log.Fatalln("Set env variable CLOUDFLARE_API_KEY")
	}

	cfA.Email, ok = os.LookupEnv("CLOUDFLARE_EMAIL")
	if !ok {
		log.Fatalln("Set env variable CLOUDFLARE_EMAIL")
	}
}

func getZones() {
	api, err := cf.New(cfCreds.ID, cfCreds.Email)
	if err != nil {
		log.Printf("Error auth. Msg: %s", err)
	}

	zones, err := api.ListZones()

	for key, zone := range zones {
		fmt.Printf("key: %v\t value:%s", key, zone)

	}

}
