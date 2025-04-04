package main

import (
	"flag"
	"fmt"
	"os"

	wphash "github.com/thundernet8/wordpress-hash-go"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hash":
		handleHash()
	case "verify":
		handleVerify()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: wp-hash-cli <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  hash   - Hash a password")
	fmt.Println("    Options:")
	fmt.Println("      -p string  Password to hash")
	fmt.Println()
	fmt.Println("  verify - Verify a password and hash")
	fmt.Println("    Options:")
	fmt.Println("      -p string  Password to verify")
	fmt.Println("      -h string  Hash to verify against")
}

func handleHash() {
	cmd := flag.NewFlagSet("hash", flag.ExitOnError)
	password := cmd.String("p", "", "Password to hash")
	cmd.Parse(os.Args[2:])

	if *password == "" {
		fmt.Println("Error: Please specify a password (-p option)")
		os.Exit(1)
	}

	hash := wphash.HashPassword(*password)
	fmt.Printf("Password <%s> hash result: <%s>\n", *password, hash)
}

func handleVerify() {
	cmd := flag.NewFlagSet("verify", flag.ExitOnError)
	password := cmd.String("p", "", "Password to verify")
	hash := cmd.String("h", "", "Hash to verify against")
	cmd.Parse(os.Args[2:])

	if *password == "" || *hash == "" {
		fmt.Println("Error: Please specify both password and hash (-p and -h options)")
		os.Exit(1)
	}

	match := wphash.CheckWordPressPasswordHash(*password, *hash)
	fmt.Printf("Password <%s> verify result with hash <%s>: %t\n", *password, *hash, match)
}
