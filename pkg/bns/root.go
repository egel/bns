/*
Copyright © 2025 Maciej Sypien
*/
package bns

import (
	"os"

	"github.com/egel/bns/pkg/display"
	"github.com/egel/bns/pkg/git"
	"github.com/egel/bns/pkg/text"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

var flagOriginalCase bool
var flagForceAscii bool
var flagConnector string
var flagVerbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bns",
	Short: "Simple command tool to generate clean and friendly git branch name",
	Long: `
Simple command tool which savely cleans text inputs focusing on comply with Git
repository naming conventions.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := git.CleanBranchName(args, flagConnector, flagOriginalCase, flagForceAscii)
		os.Stdout.WriteString(result)
		if flagVerbose {
			display.ShowResults(args, result)
		}
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
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "print original and modified texts")
}
