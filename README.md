# Cowsay

A Go implementation of the classic Unix `cowsay` command. This program takes text input via stdin and displays it in a speech bubble above an ASCII art cow.

## Features
- Displays text in a formatted speech bubble
- Handles multi-line input with proper alignment
- Converts tabs to spaces for consistent formatting
- ASCII art cow display
- Designed to work with pipes and stdin

## Installation

### Prerequisites
- Go 1.24.5 or later

### Option 1: Run from the repository
```bash
go install github.com/davidalecrim1/cowsay@latest
```

### Option 2: Build it locally

Clone:
```bash
git clone github.com/davidalecrim1/cowsay
```

Build:
```bash
go build -o cowsay main.go
```

Run:
```bash
./cowsay
```

## Usage

The program is designed to work with pipes and stdin:

```bash
echo "Hello, world!" | cowsay
```

Or with multi-line input:

```bash
echo -e "Hello\nWorld" | cowsay
```

You can also use it with other commands:

```bash
fortune | cowsay
date | cowsay
```

## Examples

### Single line text
```bash
echo "Hello, world" | go run main.go
```

Output:
```bash
 ______________
< Hello, world >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## Development

### Running tests
```bash
go test -v ./...
```

### Using Makefile

```bash
# Run the program
make run
```

```bash
# Simulate input
make simulate
```

```bash
# Run tests
make test
```