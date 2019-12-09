package helpers

import (
	"fmt"
	"regexp"
)

// Expand byte
func Expand() {
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)

	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\S+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
