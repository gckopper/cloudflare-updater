package cloudflare

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func UpdateAAAA(version int, zoneid string, recordid string, email string, authkey string, body string) int {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.cloudflare.com/client/v%d/zones/%s/dns_records/%s", version, zoneid, recordid), strings.NewReader(body))
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
