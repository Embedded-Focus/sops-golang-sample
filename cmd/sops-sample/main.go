package main

import (
	"fmt"
	"os"

	sops "github.com/getsops/sops/v3/decrypt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid number of arguments: %s <file-to-decrypt>", os.Args[0])
		os.Exit(1)
	}
	cleartext, err := sops.File(os.Args[1], "YAML")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decyrpting: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(cleartext))
}
