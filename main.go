package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	hostname, err := os.Hostname()

	user := os.Getenv("USER")

	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
	} else {
		fmt.Printf("%s@%s\n", user, hostname)

	}

	fmt.Println("------------------------------")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
}
