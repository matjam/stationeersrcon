package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chzyer/readline"
	"github.com/howeyc/gopass"
)

func loginPrompt() string {
	if !readline.IsTerminal(int(os.Stdin.Fd())) {
		log.Fatal("need a working terminal to prompt for RCON password")
	}

	fmt.Print("RCON Password: ")

	password, err := gopass.GetPasswdMasked()
	if err != nil {
		log.Fatal(err)
	}

	return string(password)
}
