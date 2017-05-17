package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize kcs problem directory",
	Long:  "initialize kcs problem directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("success initialize")
	},
}
