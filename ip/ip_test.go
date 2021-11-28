package ip_test

import (
	"fmt"
	"testing"

	"github.com/gckopper/cloudflare-updater/ip"
)

func TestGetIP(t *testing.T) {
	if fmt.Sprint(ip.GetIP()) != "2804:13d0:9928:9401:659c:8aa0:9388:3d9" {
		t.Fatal("Wrong IP", fmt.Sprint(ip.GetIP()), "not", "2804:13d0:9928:9401:659c:8aa0:9388:3d9")
	}
}

func TestToVersion(t *testing.T) {
	if ip.ToVersion("A") != 4 {
		t.Fatal("Conversion from A to 4 is not working")
	}
	if ip.ToVersion("AAAA") != 6 {
		t.Fatal("Conversion from AAAA to 6 is not working")
	}
}
