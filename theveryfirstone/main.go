package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("im runnin")
	s, sep := "", "\r\n"
	for i, v := range os.Args[:] {
		s += sep + "index:" + strconv.Itoa(i) + " val:" + v
		sep = "  "
	}
	fmt.Println(s)
}
