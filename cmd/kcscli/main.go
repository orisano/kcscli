package main

import (
	"fmt"
	"os"

	"github.com/orisano/kcscli/cmd/kcscli/subcmd"
)

func main() {
	if err := subcmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
