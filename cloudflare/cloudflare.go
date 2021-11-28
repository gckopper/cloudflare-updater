package cloudflare

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gckopper/cloudflare-updater/ip"
)

func UpdateRecord(version int, zoneid string, recordid string, email string, authkey string, recordType string, domain string, ipAddr string) int {
	if ipAddr == "auto" {
		if recordType == "A" || recordType == "AAAA" {
			ipAddr = fmt.Sprint(ip.GetIP(ip.ToVersion(recordType)))
		} else {
			log.Fatal("Unsuported record type for auto ip address")
		}
	}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.cloudflare.com/client/v%d/zones/%s/dns_records/%s", version, zoneid, recordid)
	body := fmt.Sprintf("{\"type\":\"%s\",\"name\":\"%s\",\"content\":\"%s\",\"ttl\":1,\"proxied\":false}", recordType, domain, ipAddr)
	//fmt.Println(url)
	//fmt.Println(body)
	req, err := http.NewRequest("PUT", url, strings.NewReader(body))
	req.Header.Set("X-Auth-Email", email)
	req.Header.Set("X-Auth-Key", authkey)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
		return 1
	}
	return 0
}

func UpdateAAAA(version int, zoneid string, recordid string, email string, authkey string, domain string, ipAddr string) int {
	return UpdateRecord(version, zoneid, recordid, email, authkey, "AAAA", domain, ipAddr)
}
