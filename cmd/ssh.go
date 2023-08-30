package cmd

import (
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH key pairs configurations",
	Long: `An API for managing ssh configurations, it is useful for changing the
active private SSH key from my personal Github to my organization's
github.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
