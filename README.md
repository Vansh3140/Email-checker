# Email Checker

**Email Checker** is a Go-based tool to verify the validity of email addresses. The tool checks several aspects of email validity, including domain records, MX records, SPF, and DMARC records to ensure that an email address is properly configured and functional. It also performs an SMTP handshake to check if the email is deliverable.

## Features

- **Domain Validation**: Checks if the domain has valid MX records, SPF records, and DMARC records.
- **SMTP Handshake**: Simulates an SMTP handshake to check whether the email address can be received by the mail server.
- **User-Friendly CLI**: Simple command-line interface to validate emails by entering them interactively.

## Installation

### Requirements

- Go version 1.23 or later.

### Steps to Install

1. Clone the repository:
   ```bash
   git clone https://github.com/Vansh3140/Email-checker.git
   cd email-checker
   ```

2. Initialize the Go module:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build
   ```

4. Run the program:
   ```bash
   ./email-checker
   ```

## Usage

### Start the Email Checker

Once the application is running, it will prompt you to enter an email address. You can check multiple emails in one session.

```bash
Starting the Email-Checker....
If you want to exit the program enter 'q' else enter the email
```

### Example:

```bash
Enter an email to check: example@domain.com

Results for domain: domain.com
Has MX Records: true
Has SPF Record: true
SPF Record: v=spf1 include:_spf.google.com ~all
Has DMARC Record: true
DMARC Record: v=DMARC1; p=none; rua=mailto:dmarc-reports@domain.com

Email example@domain.com is valid
```

- If the email address is valid and the domain's MX, SPF, and DMARC records are properly set up, it will print `Email is valid`.
- If any issues are detected (missing records or invalid SMTP response), the tool will notify you that the email is invalid.

## Code Overview

### main.go

- Handles the user input and validates email addresses.
- Checks the MX, SPF, and DMARC records for the email's domain.
- Calls the `SmtpHandshake` function to check email validity through an SMTP handshake.

### smtp.go

- Performs the SMTP handshake to check if the email address is deliverable.
- Interacts with the mail server to simulate sending an email to the specified address.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
