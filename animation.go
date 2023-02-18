package main

import (
	"fmt"
	"strings"
	"time"
)

func animate(text string) {
	fmt.Printf("%s\r", strings.ToLower(text))
	time.Sleep(500 * time.Millisecond)

	for _, r := range text {
		fmt.Printf("%s", strings.ToUpper(string(r)))
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Printf("\r%s\r", strings.ToUpper(text))

	for _, r := range text {
		fmt.Printf("%s", strings.ToLower(string(r)))
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Printf("\r%s\r", strings.ToLower(text))

	for i:=0;i<len(text);i++ {
		fmt.Printf("%s", " ")
	}

	fmt.Printf("%s\r","")
}