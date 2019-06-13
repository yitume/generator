package main

import (
	"log"
	"os"

	"github.com/yitume/generator/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
