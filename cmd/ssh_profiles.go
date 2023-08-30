/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nimbo1999/environment-setup/internal/services"
	"github.com/spf13/cobra"
)

func NewProfilesCmd(service services.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "profiles",
		Short: "List all SSH profile availables",
		Long: `This command is usefull to understand whitch profiles does the user
	has available in this OS. It returns a list of all profiles in the CLI.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			entries, err := service.List()
			if err != nil {
				return err
			}
			color.White("Profiles:")
			for _, entry := range entries {
				if entry.IsDir() {
					color.Blue(fmt.Sprintf(" - %s\n", entry.Name()))
				}
			}
			return nil
		},
	}
}

func init() {
	sshService := services.NewSSHService("installed-ssh")
	profilesCmd := NewProfilesCmd(sshService)
	sshCmd.AddCommand(profilesCmd)
}
