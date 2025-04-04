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
		fmt.Printf("不明なコマンド: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("使用方法: wp-hash-cli <command> [options]")
	fmt.Println("コマンド:")
	fmt.Println("  hash   - パスワードをハッシュ化")
	fmt.Println("    オプション:")
	fmt.Println("      -p string  ハッシュ化するパスワード")
	fmt.Println()
	fmt.Println("  verify - パスワードとハッシュを検証")
	fmt.Println("    オプション:")
	fmt.Println("      -p string  検証するパスワード")
	fmt.Println("      -h string  検証するハッシュ")
}

func handleHash() {
	cmd := flag.NewFlagSet("hash", flag.ExitOnError)
	password := cmd.String("p", "", "ハッシュ化するパスワード")
	cmd.Parse(os.Args[2:])

	if *password == "" {
		fmt.Println("エラー: パスワードを指定してください (-p オプション)")
		os.Exit(1)
	}

	hash := wphash.HashPassword(*password)
	fmt.Printf("パスワード <%s> のハッシュ結果: <%s>\n", *password, hash)
}

func handleVerify() {
	cmd := flag.NewFlagSet("verify", flag.ExitOnError)
	password := cmd.String("p", "", "検証するパスワード")
	hash := cmd.String("h", "", "検証するハッシュ")
	cmd.Parse(os.Args[2:])

	if *password == "" || *hash == "" {
		fmt.Println("エラー: パスワードとハッシュの両方を指定してください (-p と -h オプション)")
		os.Exit(1)
	}

	match := wphash.CheckWordPressPasswordHash(*password, *hash)
	fmt.Printf("パスワード <%s> とハッシュ <%s> の検証結果: %t\n", *password, *hash, match)
}
