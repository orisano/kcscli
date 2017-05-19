package cmd

import (
	"encoding/json"
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
		problem := Problem{
			Title:       args[0],
			TimeLimit:   3.0,
			MemoryLimit: 512,
			Author:      "None",
			Subtasks:    []Subtask{},
			Solvers:     []string{},
		}
		b, _ := json.MarshalIndent(&problem, "", "  ")
		fmt.Println(string(b))
	},
}
