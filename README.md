
# GhoZZ
Ghost Fuzzer

## Description
**GhoZZ** is a high-performance directory fuzzing tool built with Go, designed for security researchers and penetration testers. This tool allows users to discover hidden directories and files on web servers by making rapid HTTP requests based on a wordlist. 

A key feature of GhoZZ is its support for the TOR network, enabling anonymous scanning and bypassing certain IP-based restrictions. Additionally, GhoZZ includes customizable filtering options to hide specific HTTP status codes from results using the `--hide-code` flag.

## Usage

**GhoZZ** is using TOR network by default. To run ghozz, ensure you have Go installed and the TOR service running.

### Basic Example
```bash
./ghozz --target http://example.com --wordlist wordlist.txt
```
- `--target`: URL to fuzz
- `--wordlist`: Path to the wordlist file
- Requires the TOR service running on `localhost:9050`

### Hiding Specific HTTP Status Codes
Exclude results with specific HTTP status codes:
```bash
./ghozz --target http://example.com --wordlist wordlist.txt --hide-code 403,404,500
```
- `--hide-code`: Comma-separated status codes to filter out

### Saving Output to a File
```bash
./ghozz --target http://example.com --wordlist wordlist.txt --output results.txt
```
- `--output`: Save results to the specified file

### Combining Options
```bash
./ghozz --target http://example.com --wordlist wordlist.txt --hide-code 403,404 --output results.txt
```
This example runs a fuzz through the TOR network, excludes 403 and 404 responses, and saves results to `results.txt`.

## Disclaimer

This tool is developed for educational purposes and cybersecurity research only. It is intended to help security professionals and enthusiasts understand web security concepts and identify vulnerabilities in authorized environments. 

The author strictly prohibits any illegal activity, including but not limited to unauthorized testing, hacking, or disrupting systems and networks. Users are solely responsible for their actions, and the author assumes no liability for any misuse of this tool. 

By using this tool, you agree that you are fully responsible for complying with local laws and regulations. This project is created purely out of curiosity and for ethical learning purposes. Use it responsibly and only in environments where you have explicit permission.

## Authors

This project was created and maintained by:

- **Fadila Ahmad S**  
  - GitHub: [https://github.com/fadilaahmads](https://github.com/fadilaahmads)  
  - GitLab: [https://gitlab.com/fadilaahmads](https://gitlab.com/fadilaahmads)  
  - LinkedIn: [https://linkedin.com/in/fadilaahmads](https://linkedin.com/in/fadilaahmads)

