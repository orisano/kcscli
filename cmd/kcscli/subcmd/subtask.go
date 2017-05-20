package subcmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	subtaskCmd.AddCommand(subtaskAddCmd)
	RootCmd.AddCommand(subtaskCmd)
}

var subtaskCmd = &cobra.Command{
	Use:   "subtask",
	Short: "manage subtasks",
	Long:  "manage subtasks",
}

var subtaskAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "add subtask",
	Long:  "add subtask",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}
		name := args[0]
	},
}
