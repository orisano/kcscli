package subcmd

import (
	"os"

	"fmt"

	"github.com/orisano/kcscli"
	"github.com/spf13/cobra"
)

func init() {
	subtaskAddCmd.Flags().IntVar(&score, "score", 100, "subtask's score")
	subtaskCmd.AddCommand(subtaskAddCmd)
	RootCmd.AddCommand(subtaskCmd)
}

var subtaskCmd = &cobra.Command{
	Use:   "subtask",
	Short: "manage subtasks",
	Long:  "manage subtasks",
}
var score int

var subtaskAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "add subtask",
	Long:  "add subtask",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		problem, err := kcscli.LoadProblem()
		if err != nil {
			fmt.Fprintln(os.Stderr, "LoadProblem: ", err)
			os.Exit(1)
		}

		name := args[0]
		_, exists := problem.Subtasks[name]
		if exists {
			fmt.Fprintln(os.Stderr, "Already exists subtask name.")
			os.Exit(1)
		}

		problem.Subtasks[name] =
			kcscli.Subtask{
				Files: []string{},
				Score: score,
			}

		problem.Save(false)
	},
}
