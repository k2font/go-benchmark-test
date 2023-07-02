package main

import (
	"fmt"
	"strconv"
)

func main() {
	ItoaByFmt(1000000)
	ItoaByStrconv1(1000000)
	ItoaByStrconv2(1000000)
}

func ItoaByFmt(n int) {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, fmt.Sprint(i))
	}
}

func ItoaByStrconv1(n int) {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(i))
	}
}

func ItoaByStrconv2(n int) {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(i))
	}
}
