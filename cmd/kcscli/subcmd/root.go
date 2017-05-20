package subcmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is root object for cobra Command
var RootCmd = &cobra.Command{
	Use:   "kcscli",
	Short: "command line tool for kcs problem writer",
	Long:  "command line tool for kcs problem writer",
}
