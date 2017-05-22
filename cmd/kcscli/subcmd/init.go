package subcmd

import (
	"github.com/dustin/go-humanize"
	"github.com/orisano/kcscli"
	"github.com/spf13/cobra"
)

var (
	humanizedTimeLimit   string
	humanizedMemoryLimit string
	author               string
)

func init() {
	initCmd.Flags().StringVar(&humanizedMemoryLimit, "memory-limit", "512MB", "problem's memory limit")
	initCmd.Flags().StringVar(&humanizedTimeLimit, "time-limit", "3.0s", "problem's time limit")
	initCmd.Flags().StringVar(&author, "author", "None", "author name")
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init [problem_title]",
	Short:   "initialize kcs problem directory",
	Long:    "initialize kcs problem directory",
	PostRun: problemSaver(true),
	Run: func(cmd *cobra.Command, args []string) {
		title := take1Args(cmd, args)

		timeLimit, unit, err := humanize.ParseSI(humanizedTimeLimit)
		exitIf(unit != "s", "time-limit must be 'second'")
		exitIfError(err, "time-limit:")

		memoryLimit, err := humanize.ParseBytes(humanizedMemoryLimit)
		exitIfError(err, "memory-limit:")

		problem = kcscli.Problem{
			Title:       title,
			TimeLimit:   timeLimit,
			MemoryLimit: memoryLimit,
			Author:      author,
			Subtasks:    map[string]kcscli.Subtask{},
			Solvers:     []string{},
		}
	},
}
