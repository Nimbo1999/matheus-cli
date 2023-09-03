package cmd

import (
	"github.com/fatih/color"
	"github.com/nimbo1999/environment-setup/internal/entities"
	"github.com/nimbo1999/environment-setup/internal/services"
	"github.com/spf13/cobra"
)

func NewChangeCmd(service services.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "change",
		Short: "Command to change the ssh configuration folder",
		Long: `It changes the default SSH private and public keys to the one specified
		by parameters`,
		RunE: func(cmd *cobra.Command, args []string) error {
			profile, err := cmd.Flags().GetString("profile")
			if err != nil {
				return err
			}
			p, err := entities.GetProfile(profile)
			if err = service.Update(*p); err != nil {
				return err
			}
			color.Green("Profile updated successfully!")
			return nil
		},
	}
}

func init() {
	sshService := services.NewSSHService("installed-ssh")
	changeCmd := NewChangeCmd(sshService)
	sshCmd.AddCommand(changeCmd)
	changeCmd.Flags().StringP("profile", "p", "", "Profile to be used")
	changeCmd.MarkFlagRequired("profile")
}
