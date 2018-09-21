package main

import (
	"log"
	"os"
)

//type dumpSettings struct {
//	pathPrefix string
//}

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
