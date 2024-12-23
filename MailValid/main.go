package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Print the CSV-style header
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	// Read input from stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domain := scanner.Text()
		checkDomain(domain)
	}

	// Check for any errors during input
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

// Validate the domain
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check for MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v\n", domain, err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Check for TXT records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v\n", domain, err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Check for DMARC records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for domain %s: %v\n", domain, err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Print the results in CSV format
	fmt.Printf("%s, %t, %t, %s, %t, %s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
