package main

import (
	"encoding/json"
	cf "github.com/cloudflare/cloudflare-go"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func dumpZones() (err error) {
	api, err := cf.New(cfCreds.ID, cfCreds.Email)
	if err != nil {
		log.Printf("Error auth. Msg: %s", err)
	}

	zones, err := api.ListZones()
	if err != nil {
		log.Printf("Can't get zones list. Msg: %s", err)
	}

	for _, zone := range zones {
		zoneName := zone.Name

		dumpPath := filepath.Join(".", zoneName)
		if _, err := os.Stat(dumpPath); os.IsNotExist(err) {
			os.Mkdir(dumpPath, os.ModePerm)
		}

		z, err := json.MarshalIndent(zone, "  ", "  ")
		if err != nil {
			log.Printf("Error to marshal zone data. Msg: %s", err)
		}

		zoneFileSettings := filepath.Join(dumpPath, "zone_settings.json")
		err = ioutil.WriteFile(zoneFileSettings, z, 0644)
		if err != nil {
			log.Printf("Error dump zome settings. Msg: %s", err)
		}

		log.Printf("Name: %s\n %s", zoneName, string(z))
	}

	return err
}
