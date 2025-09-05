package root

import (
	"fmt"

	"github.com/docker/cagent/internal/telemetry"
	"github.com/spf13/cobra"
)

func NewLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate to agent registry",
		Long:  `Authenticate to a registry to push and pull agents`,
		Run: func(cmd *cobra.Command, args []string) {
			telemetry.TrackCommand("login", args)

			fmt.Println("Login functionality coming soon!")
		},
	}
}
