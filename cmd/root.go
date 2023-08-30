package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "matheus",
	Short: "Matheus environment setup kit",
	Long: `An application that helps matheus change between differents
environments. It is helpful for changing ssh key pairs.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
