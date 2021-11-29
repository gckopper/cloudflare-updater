package main

import (
	"fmt"
	"os"

	"github.com/gckopper/cloudflare-updater/cloudflare"
)

// use $env:GOOS = "linux" to change the target os

func main() {
	//fmt.Println(glowing.GetIP())
	file, err := os.Open("secrets.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var version int
	var zoneid string
	var recordid string
	var email string
	var authkey string
	var recordType string
	var domain string
	var ttl string
	var proxied string
	fmt.Fscanf(file, "%d%s%s%s%s%s%s%s%s", &version, &zoneid, &recordid, &email, &authkey, &recordType, &domain, &ttl, &proxied)
	//fmt.Println("valores", version, zoneid, recordid, email, authkey, recordType, domain)
	cloudflare.UpdateRecord(version, zoneid, recordid, email, authkey, recordType, domain, "auto", ttl, proxied)
	//fmt.Println(result)
}
