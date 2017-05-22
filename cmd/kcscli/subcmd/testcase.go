package subcmd

import (
	_ "github.com/orisano/kcscli"
	"github.com/spf13/cobra"
)

func init() {
	testcaseAddCommand.Flags().StringVar(&subtaskName, "subtask", "all", "subtask name")
	testcaseCommand.AddCommand(testcaseAddCommand)
	RootCmd.AddCommand(testcaseCommand)
}

var subtaskName string

var testcaseCommand = &cobra.Command{
	Use:   "testcase",
	Short: "manage testcase",
	Long:  "manage testcase",
}

var testcaseAddCommand = &cobra.Command{
	Use:     "add [input-file]",
	Short:   "add testcase",
	Long:    "add testcase",
	PreRun:  problemLoader,
	PostRun: problemSaver(true),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := take1Args(cmd, args)
		subtask, exists := problem.Subtasks[subtaskName]
		exitIf(!exists, "subtask name '"+subtaskName+"' is not found")
		subtask.Files = append(subtask.Files, inputFile)
		problem.Subtasks[subtaskName] = subtask
	},
}
