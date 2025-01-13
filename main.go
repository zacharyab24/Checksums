/*
 * File: sums.go
 * Author: Zachary Bower
 * Last updated: 13/01/2025
 * About: Small CLI application used for verifying check sums and comparing computed sums with given values
 */

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Flags
	var algorithmFlag = flag.String("a", "sha256", "Algorithm used to calculate the checksum") // Used to specify the algorithm (default: sha256)
	var fileName = flag.String("f", "", "Path to file") // Used to specify the file path
	var inputChecksum  = flag.String("s", "", "Provided checksum to check against") // Used to specify if the user wants to compare the calculated value to a given input and provide said input
	flag.Parse()

	// Attempt to open file
	if *fileName == "" {
		log.Fatal("File path not provided. Please set it with the -f flag")
	}
	data, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate checksum of file using specified algorithm
	var checksum []byte
	switch *algorithmFlag {
	case "sha256":
		s := sha256.Sum256(data)
		checksum = s[:]
	case "md5":
		s := md5.Sum(data)
		checksum = s[:]
	default:
		log.Fatal("Invalid algorithm given")
	}

	// Compare the two checksums
	if *inputChecksum != "" {
		if *inputChecksum == fmt.Sprintf("%x", checksum) {
			fmt.Println("Checksums match")
		} else {
			fmt.Println("Warning: checksums do not match")
		}
	}
	fmt.Printf("%x", checksum)
}
