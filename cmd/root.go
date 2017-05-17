package cmd

import (
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use: "kcscli",
    Short: "command line tool for kcs problem writer",
    Long: "command line tool for kcs problem writer",
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
    },
}
