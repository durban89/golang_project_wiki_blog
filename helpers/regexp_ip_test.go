package helpers

import (
	"testing"
)

func Test_IsIP(t *testing.T) {
	is := IsIP("127.0.0.1")
	t.Fatal(is)
}
