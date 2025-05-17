/*
Copyright © 2025 Maciej Sypien
*/
package cbn

import (
	"os"

	"github.com/egel/cbn/v2/pkg/display"
	"github.com/egel/cbn/v2/pkg/git"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cbn",
	Short: "Simple command tool to generate clean and friendly git branch name",
	Long: `
Tool that take the inputs and convert it into a clean and friendly name.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := git.CleanBranchName(args)
		display.ShowResults(args, result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
