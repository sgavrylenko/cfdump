package cfdump

import (
	"log"
	"os"
)

type CFAuth struct {
	ID    string
	Email string
}

func (CFA *CFAuth) init() {
	var ok bool
	CFA.ID, ok = os.LookupEnv("CLOUDFLARE_API_KEY")
	if !ok {
		log.Fatalln("Set env variable CLOUDFLARE_API_KEY")
	}

	CFA.Email, ok = os.LookupEnv("CLOUDFLARE_EMAIL")
	if !ok {
		log.Fatalln("Set env variable CLOUDFLARE_EMAIL")
	}
}
