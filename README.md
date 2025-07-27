# IpLookup

A simple command-line tool written in Go that performs DNS lookups using DNS-over-HTTPS (DoH) to resolve domain names to IP addresses.

## Features

- 🔍 DNS lookup using Cloudflare's secure DoH service (1.1.1.1)
- 🛡️ HTTPS-encrypted DNS queries for privacy
- 📝 Interactive command-line interface
- ⚡ Fast and lightweight
- 🎯 Simple A record resolution

## Project Structure
```
IpLookup/
├── go.mod # Go module definition
├── dns.go # DNS response type definitions
├── http.go # DoH client implementation
└── main.go # CLI interface
```
text

## Installation

1. Clone this repository:
```bash
git clone https://github.com/yourusername/IpLookup.git
cd IpLookup
```

2. Build the application:
```bash
go build -o iplookup
```

## Usage

### Interactive Mode

Run the application and follow the prompts:
```bash
./iplookup
```

Example session:
```bash
Enter domain name (e.g., boot.dev): google.com
The IP address for google.com is: 142.250.191.14
```

### Direct Execution

You can also run it directly with Go:
```bash
go run .
```

## How It Works

1. **DNS-over-HTTPS**: Uses Cloudflare's DoH service for secure DNS resolution
2. **JSON Response Parsing**: Parses DNS responses in JSON format
3. **A Record Lookup**: Resolves domain names to IPv4 addresses
4. **Error Handling**: Provides clear error messages for failed lookups

## Technical Details

- **Go Version**: 1.24.1
- **DNS Provider**: Cloudflare (1.1.1.1)
- **Protocol**: DNS-over-HTTPS (DoH)
- **Record Type**: A records (IPv4 addresses)

## Dependencies

This project uses only Go standard library packages:
- `encoding/json` - JSON parsing
- `fmt` - Formatted I/O
- `io` - I/O utilities  
- `net/http` - HTTP client
- `bufio` - Buffered I/O
- `os` - Operating system interface
- `strings` - String utilities

## Example API Response

The application parses DNS responses in this format:
```bash
{
  "Status": 0,
  "TC": false,
  "RD": true,  
  "RA": true,
  "AD": false,
  "CD": false,
  "Question": [
    {
      "name": "example.com",
      "type": 1
    }
  ],
  "Answer": [
    {
      "name": "example.com",
      "type": 1,
      "TTL": 300,
      "data": "93.184.216.34"
    }
  ]
}
```

## Error Handling

The application handles various error scenarios:
- Invalid HTTP requests
- Network connectivity issues
- JSON parsing errors
- Empty DNS responses
- Invalid domain names

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- Uses Cloudflare's public DoH service
- Inspired by the need for simple, secure DNS lookups
