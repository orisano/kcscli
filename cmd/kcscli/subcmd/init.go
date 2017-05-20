package subcmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/orisano/kcscli"
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
		if len(args) == 0 {
			cmd.Usage()
			os.Exit(1)
		}
		problem := kcscli.Problem{
			Title:       args[0],
			TimeLimit:   3.0,
			MemoryLimit: 512,
			Author:      "None",
			Subtasks:    []kcscli.Subtask{},
			Solvers:     []string{},
		}
		b, _ := json.MarshalIndent(&problem, "", "  ")
		fmt.Println(string(b))
	},
}
