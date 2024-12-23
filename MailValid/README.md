# MailValid

**MailValid** is a command-line tool written in Go that validates email domains by checking their DNS records for various configurations, such as MX, SPF, and DMARC. This tool helps ensure that an email domain is configured properly for sending and receiving emails securely.

## Features
- Validates if a domain has:
  - **MX (Mail Exchange) Records**: Ensures the domain can receive emails.
  - **SPF (Sender Policy Framework) Records**: Checks if the domain has defined servers authorized to send emails on its behalf.
  - **DMARC (Domain-based Message Authentication, Reporting, and Conformance) Records**: Verifies if the domain has implemented DMARC policies for email authentication.
- Outputs results in a CSV-like format for easy parsing and analysis.

---

## How It Works

The tool takes a list of domains as input and validates the following DNS records:

### 1. MX Records
MX records specify the mail servers responsible for accepting email messages on behalf of a domain.  
- If a domain has MX records, it can receive emails.
- The tool uses `net.LookupMX` to fetch the MX records.

### 2. SPF Records
SPF is a DNS record used to prevent email spoofing by specifying the mail servers authorized to send emails for a domain.  
- SPF records start with `v=spf1`.  
- The tool uses `net.LookupTXT` to find the TXT records and checks for SPF entries.

### 3. DMARC Records
DMARC builds on SPF and DKIM to provide a mechanism for domain owners to protect their domain from unauthorized use.  
- DMARC records are stored under the `_dmarc` subdomain (e.g., `_dmarc.example.com`).  
- DMARC records start with `v=DMARC1`.  
- The tool uses `net.LookupTXT` to locate DMARC records.

---

## Input and Output

### Input
The tool reads domain names line by line from the standard input (stdin). For example:
```plaintext
example.com
gmail.com
nonexistentdomain.xyz
```

### Output
The results are printed in CSV format with the following columns:

- domain: The domain name being validated.
- hasMX: Indicates whether the domain has MX records (true/false).
- hasSPF: Indicates whether the domain has SPF records (true/false).
- spfRecord: The actual SPF record, if found.
- hasDMARC: Indicates whether the domain has DMARC records (true/false).
- dmarcRecord: The actual DMARC record, if found.

Example Output
```plaintext
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
example.com, true, true, v=spf1 include:_spf.example.com ~all, true, v=DMARC1; p=none; rua=mailto:dmarc@example.com
gmail.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=reject; rua=mailto:dmarc-reports@google.com
nonexistentdomain.xyz, false, false, , false, 
```

Example Usage
```bash
example.com
gmail.com
nonexistentdomain.xyz
```

## Explanation of DNS Records Used
### 1. MX Records
Purpose: MX records ensure that a domain can receive emails by specifying the mail servers responsible for handling emails.

Example:
```plaintext
example.com mail exchanger = 10 mail.example.com.
```
If no MX records are found, the domain cannot receive emails.

### 2. SPF Records
Purpose: SPF prevents email spoofing by specifying the mail servers allowed to send emails on behalf of the domain.

Example:
```plaintext
v=spf1 include:_spf.example.com ~all
```
If no SPF record is found, the domain lacks protection against email spoofing.

### 3. DMARC Records
Purpose: DMARC enhances email authentication by providing domain owners with the ability to specify how emails failing SPF/DKIM checks should be handled.
Example:

```plaintext
v=DMARC1; p=reject; rua=mailto:dmarc-reports@example.com
```

If no DMARC record is found, the domain is not enforcing any email authentication policies.
Example Usage

**Input**
```plaintext
example.com
gmail.com
```

**Output**
```plaintext
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
example.com, true, true, v=spf1 include:_spf.example.com ~all, true, v=DMARC1; p=none; rua=mailto:dmarc@example.com
gmail.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=reject; rua=mailto:dmarc-reports@google.com
```

## Potential Use Cases
1. **Email Security Analysis**: Quickly verify if a domain has proper email configurations to prevent spoofing and phishing attacks.
2. **Email Server Setup**: Use this tool to check if your domain is ready to send and receive emails.
2. **Domain Validation**: Validate multiple domains in bulk for their email capabilities.

## Contributing
If you'd like to contribute to this project, feel free to open an issue or submit a pull request.