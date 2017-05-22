package subcmd

import (
	"github.com/orisano/kcscli"
	"github.com/spf13/cobra"
)

func init() {
	subtaskAddCmd.Flags().IntVar(&subtaskScore, "score", 100, "subtask's score")
	subtaskCmd.AddCommand(subtaskAddCmd)
	RootCmd.AddCommand(subtaskCmd)
}

var subtaskCmd = &cobra.Command{
	Use:   "subtask",
	Short: "manage subtasks",
	Long:  "manage subtasks",
}
var subtaskScore int

var subtaskAddCmd = &cobra.Command{
	Use:     "add [name]",
	Short:   "add subtask",
	Long:    "add subtask",
	PreRun:  problemLoader,
	PostRun: problemSaver(true),
	Run: func(cmd *cobra.Command, args []string) {
		name := take1Args(cmd, args)

		_, exists := problem.Subtasks[name]
		exitIf(exists, "Already exists subtask name.")

		problem.Subtasks[name] = kcscli.Subtask{
			Files: []string{},
			Score: subtaskScore,
		}
	},
}
