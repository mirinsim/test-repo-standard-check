package cmd

import (
	"fmt"

	"github.com/mirinsim/test-repo-standard-check/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test-repo-standard-check",
	Short: "A simple utility CLI tool",
	Long: `test-repo-standard-check is a simple utility CLI tool.
This is a demonstration of repository standards and bootstrap workflow.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommand is provided
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("test-repo-standard-check %s\n", version.Version)
		fmt.Printf("Build date: %s\n", version.BuildDate)
		fmt.Printf("Git commit: %s\n", version.GitCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
