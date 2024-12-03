package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// main is the entry point of the program, where user input is taken and processed.
func main() {
	// Prompt the user to start the email checker or exit by entering 'q'
	fmt.Printf("Starting the Email-Checker....\n")
	fmt.Printf("If you want to exit the program enter 'q' else enter the email \n")

	// Initialize scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Loop to continually check emails until 'q' is entered
	for scanner.Scan() {
		// Exit if user inputs 'q'
		if scanner.Text() == "q" {
			break
		} else {
			email := scanner.Text()
			parts := strings.Split(email, "@")

			// Check if the email address is valid by ensuring it has exactly one '@' symbol
			if len(parts) != 2 {
				fmt.Printf("This is not a Valid email address!!\n")
				continue
			}

			// Check for domain's MX, SPF, and DMARC records
			checkDomain(parts[1])

			// Perform SMTP handshake to validate the email
			isValid := SmtpHandshake(email)

			// Output the result of email validity
			if isValid {
				fmt.Printf("Email %s is valid\n", email)
			} else {
				fmt.Printf("Email %s is invalid: \n", email)
			}
		}

		// Prompt for the next action
		fmt.Printf("If you want to exit the program enter 'q' else enter the email \n")
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning the input: %v\n", err)
	}
}

// checkDomain checks the DNS records (MX, SPF, DMARC) for the provided domain
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Checking MX Records
	getMXRecords(domain, &hasMX)

	// Checking SPF Records
	getSPFRecords(domain, &spfRecord, &hasSPF)

	// Checking DMARC Records
	getDMARCRecords(domain, &dmarcRecord, &hasDMARC)

	// Displaying the results of the checks for the domain
	fmt.Printf("\nResults for domain: %s\n", domain)
	fmt.Printf("Has MX Records: %t\n", hasMX)
	fmt.Printf("Has SPF Record: %t\n", hasSPF)
	if hasSPF {
		fmt.Printf("SPF Record: %s\n", spfRecord)
	}
	fmt.Printf("Has DMARC Record: %t\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("DMARC Record: %s\n", dmarcRecord)
	}
}

// getMXRecords fetches the MX (Mail Exchange) records for the given domain.
func getMXRecords(domain string, hasMX *bool) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error while checking MX Records: %v\n", err)
		return
	}

	// Loop through the MX records and print them
	fmt.Printf("\nMX Records for %s:\n", domain)
	for _, record := range mxRecords {
		*hasMX = true // Set the flag if MX records are found
		fmt.Printf("Host: %s, Preference: %d\n", record.Host, record.Pref)
	}
}

// getSPFRecords fetches the SPF (Sender Policy Framework) records for the domain
func getSPFRecords(domain string, spfRecord *string, hasSPF *bool) {
	// Look up TXT records for the domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error while fetching TXT Records: %v\n", err)
		return
	}

	// Loop through the TXT records to find an SPF record
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			*spfRecord = record // Set the SPF record
			*hasSPF = true      // Set the flag indicating SPF record is present
			return
		}
	}
}

// getDMARCRecords fetches the DMARC (Domain-based Message Authentication, Reporting & Conformance) records for the domain
func getDMARCRecords(domain string, dmarcRecords *string, hasDMARC *bool) {
	// Look up the TXT records for the "_dmarc" subdomain of the domain
	txtRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Fatalf("Error while fetching DMARC TXT Records %v\n", err)
	}

	// Loop through the TXT records to find a DMARC record
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			*dmarcRecords = record // Set the DMARC record
			*hasDMARC = true       // Set the flag indicating DMARC record is present
			return
		}
	}
}
