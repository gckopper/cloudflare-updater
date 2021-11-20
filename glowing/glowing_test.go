package glowing_test

import (
	"fmt"
	"testing"

	"github.com/gckopper/glowing-giggle/glowing"
)

func TestGetIP(t *testing.T) {
	if fmt.Sprint(glowing.GetIP()) != "2804:13d0:9928:9401:659c:8aa0:9388:3d9" {
		t.Fatal("Wrong IP", fmt.Sprint(glowing.GetIP()), "not", "2804:13d0:9928:9401:659c:8aa0:9388:3d9")
	}
}
