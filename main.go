package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	fmt.Println("Gofetch")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("Hostname: unknown")
	} else {
		fmt.Println("Hostname:", hostname)
	}
}
