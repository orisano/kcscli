package subcmd

import (
	"os"

	"log"

	"github.com/dustin/go-humanize"
	"github.com/orisano/kcscli"
	"github.com/spf13/cobra"
)

var (
	humanizedTimeLimit   string
	humanizedMemoryLimit string
	author string
)

func init() {
	initCmd.Flags().StringVar(&humanizedMemoryLimit, "memory-limit", "512MB", "problem's memory limit")
	initCmd.Flags().StringVar(&humanizedTimeLimit, "time-limit", "3.0s", "problem's time limit")
	initCmd.Flags().StringVar(&author, "author", "None", "author name")
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [problem_title]",
	Short: "initialize kcs problem directory",
	Long:  "initialize kcs problem directory",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		timeLimit, _, err := humanize.ParseSI(humanizedTimeLimit)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		memoryLimit, err := humanize.ParseBytes(humanizedMemoryLimit)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		problem := kcscli.Problem{
			Title:       args[0],
			TimeLimit:   timeLimit,
			MemoryLimit: memoryLimit,
			Author:      author,
			Subtasks:    []kcscli.Subtask{},
			Solvers:     []string{},
		}
		problem.Save(true)
	},
}
