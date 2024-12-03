package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// smtpHandshake validates an email address via SMTP handshake
func SmtpHandshake(email string) bool {
	// Split the email into user and domain parts
	parts := strings.Split(email, "@")
	domain := parts[1]

	// Get the MX records for the domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v\n", domain, err)
		return false
	}

	if len(mxRecords) == 0 {
		log.Printf("No MX records found for domain %s\n", domain)
		return false
	}

	// Connect to the first available MX server
	server := mxRecords[0].Host
	conn, err := net.Dial("tcp", server+":25")
	if err != nil {
		log.Printf("Error connecting to mail server %s: %v\n", server, err)
		return false
	}
	defer conn.Close()

	// Read initial server response
	ReadResponse(conn)

	// Send EHLO command
	WriteCommand(conn, "EHLO example.com")
	ReadResponse(conn)

	// Send MAIL FROM command
	WriteCommand(conn, "MAIL FROM:<test@example.com>")
	ReadResponse(conn)

	// Send RCPT TO command
	WriteCommand(conn, fmt.Sprintf("RCPT TO:<%s>", email))
	response := ReadResponse(conn)

	// Check if the RCPT TO command was accepted
	if strings.HasPrefix(response, "250") {
		return true
	} else {
		return false
	}
}

// WriteCommand writes a command to the SMTP connection
func WriteCommand(conn net.Conn, cmd string) {
	fmt.Fprintf(conn, "%s\r\n", cmd)
}

// ReadResponse reads a response from the SMTP connection
func ReadResponse(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading server response: %v\n", err)
		return ""
	}
	return strings.TrimSpace(response)
}
