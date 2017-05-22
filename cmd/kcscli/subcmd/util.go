package subcmd

import (
	"fmt"
	"os"

	"github.com/orisano/kcscli"
	"github.com/spf13/cobra"
)

func exitIf(cond bool, msg string) {
	if cond {
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}

func exitIfError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg, err)
		os.Exit(1)
	}
}

func usageIf(cond bool, cmd *cobra.Command) {
	if cond {
		cmd.Usage()
		os.Exit(1)
	}
}

func hasPositionalArgs(cmd *cobra.Command, args []string, expected int) {
	usageIf(len(args) != expected, cmd)
}

func take1Args(cmd *cobra.Command, args []string) string {
	hasPositionalArgs(cmd, args, 1)
	return args[0]
}

func take2Args(cmd *cobra.Command, args []string) (string, string) {
	hasPositionalArgs(cmd, args, 2)
	return args[0], args[1]
}

func take3Args(cmd *cobra.Command, args []string) (string, string, string) {
	hasPositionalArgs(cmd, args, 3)
	return args[0], args[1], args[2]
}

func take4Args(cmd *cobra.Command, args []string) (string, string, string, string) {
	hasPositionalArgs(cmd, args, 4)
	return args[0], args[1], args[2], args[3]
}

func problemLoader(cmd *cobra.Command, args []string) {
	var err error
	problem, err = kcscli.LoadProblem()
	exitIfError(err, "LoadProblem:")
}

func problemSaver(withStdout bool) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		problem.Save(withStdout)
	}
}
