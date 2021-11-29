package cloudflare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gckopper/cloudflare-updater/ip"
)

// Can update A and AAAA records. ipAddr and recordid can be set to auto
func UpdateRecord(version int, zoneid string, recordid string, email string, authkey string, recordType string, domain string, ipAddr string) int {
	if ipAddr == "auto" {
		if recordType == "A" || recordType == "AAAA" {
			ipAddr = fmt.Sprint(ip.GetIP(ip.ToVersion(recordType)))
		} else {
			log.Fatal("Unsuported record type for auto ip address")
		}
	}
	if recordid == "auto" {
		recordid = GetRecordId(version, zoneid, email, authkey, recordType, domain)
		if recordid == " " {
			log.Fatal("No record id")
		}
	}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.cloudflare.com/client/v%d/zones/%s/dns_records/%s", version, zoneid, recordid)
	body := fmt.Sprintf("{\"type\":\"%s\",\"name\":\"%s\",\"content\":\"%s\"}", recordType, domain, ipAddr)
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

// wrapper for the record updater with fixed type
func UpdateAAAA(version int, zoneid string, recordid string, email string, authkey string, domain string, ipAddr string) int {
	return UpdateRecord(version, zoneid, recordid, email, authkey, "AAAA", domain, ipAddr)
}

// Function to get the id of a dns record in cloudflare
func GetRecordId(version int, zoneid string, email string, authkey string, recordType string, domain string) string {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.cloudflare.com/client/v%d/zones/%s/dns_records?type=%s&name=%s&page=1&per_page=20&order=type&direction=desc&match=all", version, zoneid, recordType, domain)
	//fmt.Println(url)
	//fmt.Println(body)
	req, err := http.NewRequest("GET", url, strings.NewReader(" "))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	req.Header.Set("X-Auth-Email", email)
	req.Header.Set("X-Auth-Key", authkey)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
		return " "
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
		log.Fatal(err)
		return " "
	}
	//fmt.Println(body)
	var data = map[string][]map[string]string{}
	json.Unmarshal(body, &data)

	id, existance := data["result"][0]["id"]
	if !existance {
		log.Fatal("ID does not exist")
	}

	return fmt.Sprint(id)
}
