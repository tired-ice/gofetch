package main

import (
	"fmt"
	"os"
	"os/exec"
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

	cmd := exec.Command("uname", "-r")
	kernel, err := cmd.Output()

	if err != nil {
		fmt.Println("Error retrieving kernel version:", err)
	} else {
		fmt.Printf("Kernel: %s", kernel)
	}

	fmt.Println("Architecture:", runtime.GOARCH)
}
