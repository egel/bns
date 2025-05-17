/*
Copyright © 2025 Maciej Sypien
*/
package cbn

import (
	"os"

	"github.com/egel/cbn/pkg/display"
	"github.com/egel/cbn/pkg/git"
	"github.com/egel/cbn/pkg/text"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

var flagOriginalCase bool
var flagForceAscii bool
var flagConnector string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cbn",
	Short: "Simple command tool to generate clean and friendly git branch name",
	Long: `
Tool that take the inputs and convert it into a clean and friendly name.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := git.CleanBranchName(args, flagConnector, flagOriginalCase, flagForceAscii)
		display.ShowResults(args, result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flagOriginalCase, "original-case", "o", false, "keep orginal text case")
	rootCmd.PersistentFlags().BoolVarP(&flagForceAscii, "force-ascii", "f", false, "force text convert to ASCII")
	rootCmd.PersistentFlags().StringVarP(&flagConnector, "connector", "c", text.DefaultConnectorChar, "change default strings connector")

}
