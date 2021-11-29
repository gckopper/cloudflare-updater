package cloudflare_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gckopper/cloudflare-updater/cloudflare"
)

func getCrefentials(version *int, zoneid *string, recordid *string, email *string, authkey *string, recordType *string, domain *string, ttl *string, proxied *string) {
	file, err := os.Open("secrets.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//var version int
	fmt.Fscanf(file, "%d%s%s%s%s%s%s%s%s", version, zoneid, recordid, email, authkey, recordType, domain, ttl, proxied)
}

func TestGetRecordId(t *testing.T) {
	var version int
	var zoneid string
	var recordid string
	var email string
	var authkey string
	var recordType string
	var domain string
	var ttl string
	var proxied string
	getCrefentials(&version, &zoneid, &recordid, &email, &authkey, &recordType, &domain, &ttl, &proxied)
	id := cloudflare.GetRecordId(version, zoneid, email, authkey, recordType, domain)
	if id != recordid {
		log.Fatal("ERROR: ", id)
	}
}
