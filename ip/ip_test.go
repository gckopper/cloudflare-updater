package ip_test

import (
	"testing"

	"github.com/gckopper/cloudflare-updater/ip"
)

func TestToVersion(t *testing.T) {
	if ip.ToVersion("A") != 4 {
		t.Fatal("Conversion from A to 4 is not working")
	}
	if ip.ToVersion("AAAA") != 6 {
		t.Fatal("Conversion from AAAA to 6 is not working")
	}
}
