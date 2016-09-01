package common

import (
	"fmt"
	"os"
)

func MyError(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(-1)
	}
}
