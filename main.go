package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/akamensky/argparse"
	"github.com/nadhirxz/cryptr/utils"
	"golang.org/x/term"
)

func main() {
	parser := argparse.NewParser("cryptr", "Encrypt/Decrypt files")

	filename := parser.StringPositional(&argparse.Options{Required: true, Help: "File to encrypt/decrypt"})
	password := parser.String("p", "password", &argparse.Options{Required: false, Help: "Password to encrypt/decrypt file"})
	encrypt := parser.Flag("e", "encrypt", &argparse.Options{Required: false, Help: "Encrypt file"})
	decrypt := parser.Flag("d", "decrypt", &argparse.Options{Required: false, Help: "Decrypt file"})
	output := parser.String("o", "output", &argparse.Options{Required: false, Help: "Output file"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if *filename == "" {
		fmt.Println("Please specify file to encrypt/decrypt")
		return
	}

	if !*encrypt && !*decrypt {
		*encrypt = true
	}

	if *encrypt == *decrypt {
		fmt.Println("Please choose either encrypt or decrypt")
		return
	}

	if *password == "" {
		fmt.Print("Password: ")
		bytepw, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println()
		*password = string(bytepw)
	}

	key := utils.GenerateKey(*password)

	if *encrypt {
		err := utils.EncryptFile(*filename, key, *output)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("File encrypted")
		return
	}

	if *decrypt {
		err := utils.DecryptFile(*filename, key, *output)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("File decrypted")
		return
	}
}
