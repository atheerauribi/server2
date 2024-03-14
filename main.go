package main

import "log"

// run server
func main() {
	if err := RunServer(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
