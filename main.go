package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// I guess io.Reader takes input in general
	// but to read it (linewise) you need another step
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wc := 0
	for scanner.Scan(){
		wc++
	}
	return wc
}
