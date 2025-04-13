# wp-hash-cli

A command-line tool for handling WordPress password hashes

## Installation

```bash
go install github.com/kinosuke01/wp-hash-cli@latest
```

## Usage

### Hash a password

```bash
wp-hash-cli hash -p "your-password"
```

### Verify a password and hash

```bash
wp-hash-cli verify -p "your-password" -h "$P$B..."
```

## Options

### hash command
- `-p`: Password to hash (required)

### verify command
- `-p`: Password to verify (required)
- `-h`: Hash to verify against (required)

## Examples

```bash
# Hash a password
$ wp-hash-cli hash -p "123456"
{"password":"123456","hash":"$P$B..."}

# Verify a password and hash
$ wp-hash-cli verify -p "123456" -h "$P$B..."
{"password":"123456","hash":"$P$B...","match":true}
```

