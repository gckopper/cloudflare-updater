package cloudflare

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func UpdateAAAA(version int, zoneid string, recordid string, email string, authkey string, recordType string, domain string, ipAddr string) int {
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
