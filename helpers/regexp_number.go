package helpers

import (
	"fmt"
	"os"
	"regexp"
)

func IsNumber() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("Number")
	} else {
		fmt.Println("Not number")
	}
}
