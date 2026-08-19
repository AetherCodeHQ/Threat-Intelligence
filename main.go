package main

import (
	"fmt"
	"os"
)

// threat_intelligence - Analyze threat feeds
func threat_intelligence(path string) {
	fmt.Println("========================================")
	fmt.Println("  Threat-Intelligence")
	fmt.Println("  Analyze threat feeds")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	threat_intelligence(path)
}
