package main

import (
	"fmt"
	"os"

	"github.com/gckopper/glowing-giggle/cloudflare"
	"github.com/gckopper/glowing-giggle/glowing"
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
	fmt.Fscanf(file, "%d%s%s%s%s%s%s", &version, &zoneid, &recordid, &email, &authkey, &recordType, &domain)
	//fmt.Println("valores", version, zoneid, recordid, email, authkey, recordType, domain)
	cloudflare.UpdateAAAA(version, zoneid, recordid, email, authkey, recordType, domain, fmt.Sprint(glowing.GetIP()))
	//fmt.Println(result)
}
